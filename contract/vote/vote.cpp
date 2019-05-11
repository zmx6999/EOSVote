#include <eosiolib/eosio.hpp>

class youvote: eosio::contract {
private:
    // @abi table poll
    struct poll {
        uint64_t poll_id;
        std::string poll_name;
        uint64_t poll_status;
        uint64_t winner_id;

        uint64_t primary_key() const { return poll_id; }
    };
    typedef eosio::multi_index<N(poll), poll> t_poll;

    // @abi table pollopt
    struct pollopt {
        uint64_t pollopt_id;
        uint64_t poll_id;
        std::string option;
        uint64_t num;

        uint64_t primary_key() const { return pollopt_id; }
        uint64_t by_poll_id() const { return poll_id; }
    };
    typedef eosio::multi_index<N(pollopt), pollopt, eosio::indexed_by<N(poll_id), eosio::const_mem_fun<pollopt, uint64_t, &pollopt::by_poll_id>>> t_pollopt;

    // @abi table pollvote
    struct pollvote {
        uint64_t pollvote_id;
        uint64_t poll_id;
        uint64_t pollopt_id;
        account_name account;

        uint64_t primary_key() const { return pollvote_id; }
        uint64_t by_pollopt_id() const { return pollopt_id; }
    };
    typedef eosio::multi_index<N(pollvote), pollvote, eosio::indexed_by<N(pollopt_id), eosio::const_mem_fun<pollvote, uint64_t, &pollvote::by_pollopt_id>>> t_pollvote;

    t_poll _poll;
    t_pollopt _pollopt;
    t_pollvote _pollvote;

public:
    youvote(account_name self): eosio::contract(self), _poll(self, self), _pollopt(self, self), _pollvote(self, self) {}

    // @abi action
    void addpoll(std::string poll_name) {
        require_auth(_self);
        _poll.emplace(_self, [&](auto& p) {
            p.poll_id = _poll.available_primary_key();
            p.poll_name = poll_name;
            p.poll_status = 0;
            p.winner_id = 0;
        });
    }

    // @abi action
    void rmpoll(uint64_t poll_id) {
        require_auth(_self);

        auto itr = _poll.find(poll_id);
        if (itr == _poll.end() || itr->poll_status != 0) return;

        _poll.erase(itr);

        std::vector<uint64_t> polloptKeyListForDeletion;
        for (auto& item : _pollopt) {
            if (item.poll_id == poll_id) polloptKeyListForDeletion.push_back(item.pollopt_id);
        }

        for (uint64_t key : polloptKeyListForDeletion) {
            rmpollopt(key);
        }
    }

    // @abi action
    void status(uint64_t poll_id) {
        require_auth(_self);

        auto itr = _poll.find(poll_id);
        if (itr == _poll.end()) return;

        if (itr->poll_status >= 2) return;

        if (itr->poll_status == 1) {
            uint64_t winner_id = 0;
            uint64_t max = 0;
            for (auto& item : _pollopt) {
                if (item.poll_id == poll_id && item.num > max) {
                    winner_id = item.pollopt_id;
                    max = item.num;
                }
            }

            _poll.modify(itr, _self, [&](auto& p) {
                p.poll_status++;
                p.winner_id = winner_id;
            });

            return;
        }

        _poll.modify(itr, _self, [&](auto& p) {
            p.poll_status++;
        });
    }

    // @abi action
    void reset(uint64_t poll_id) {
        require_auth(_self);

        auto itr = _poll.find(poll_id);
        if (itr == _poll.end()) return;

        _poll.modify(itr, _self, [&](auto& p) {
            p.poll_status = 0;
        });
    }

    // @abi action
    void addpollopt(uint64_t poll_id, std::string option) {
        require_auth(_self);

        auto itr = _poll.find(poll_id);
        if (itr == _poll.end()) return;

        if (itr->poll_status != 0) return;

        _pollopt.emplace(_self, [&](auto& p) {
            p.pollopt_id = _pollopt.available_primary_key();
            p.poll_id = poll_id;
            p.option = option;
            p.num = 0;
        });
    }

    // @abi action
    void rmpollopt(uint64_t pollopt_id) {
        require_auth(_self);

        auto itr = _pollopt.find(pollopt_id);
        if (itr == _pollopt.end()) return;

        uint64_t poll_id = itr->poll_id;
        auto poll_itr = _poll.find(poll_id);
        if (poll_itr != _poll.end() && poll_itr->poll_status != 0) return;

        _pollopt.erase(itr);

        std::vector<uint64_t> pollvoteKeyListForDeletion;
        for (auto& item : _pollvote) {
            if (item.pollopt_id == pollopt_id) pollvoteKeyListForDeletion.push_back(item.pollvote_id);
        }

        for (uint64_t key : pollvoteKeyListForDeletion) {
            auto itr = _pollvote.find(key);
            if (itr != _pollvote.end()) _pollvote.erase(itr);
        }
    }

    // @abi action
    void vote(uint64_t pollopt_id, account_name account) {
        require_auth(account);

        auto itr = _pollopt.find(pollopt_id);
        if (itr == _pollopt.end()) return;

        uint64_t poll_id = itr->poll_id;
        auto poll_itr = _poll.find(poll_id);
        if (poll_itr == _poll.end()) return;

        if (poll_itr->poll_status != 1) return;

        for (auto& item : _pollvote) {
            if (item.poll_id == poll_id && item.account == account) return;
        }

        _pollvote.emplace(_self, [&](auto& p) {
            p.pollvote_id = _pollvote.available_primary_key();
            p.poll_id = poll_id;
            p.pollopt_id = pollopt_id;
            p.account = account;
        });

        _pollopt.modify(itr, _self, [&](auto& p) {
            p.num++;
        });
    }
};

EOSIO_ABI(youvote, (addpoll)(rmpoll)(status)(reset)(addpollopt)(rmpollopt)(vote))