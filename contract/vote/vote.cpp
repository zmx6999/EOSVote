#include <eosiolib/eosio.hpp>

class youvote: eosio::contract {
private:
    // @abi table poll
    struct poll {
        uint64_t key;
        uint64_t poll_id;
        std::string poll_name;
        uint64_t poll_status = 0;
        std::string option;
        uint64_t count = 0;

        uint64_t primary_key() const { return key; }
        uint64_t by_poll_id() const { return poll_id; }
    };
    typedef eosio::multi_index<N(poll), poll, eosio::indexed_by<N(poll_id), eosio::const_mem_fun<poll, uint64_t, &poll::by_poll_id>>> t_poll;

    // @abi table pollvote
    struct pollvote {
        uint64_t key;
        uint64_t poll_id;
        std::string poll_name;
        account_name account;

        uint64_t primary_key() const { return key; }
        uint64_t by_poll_id() const { return poll_id; }
    };
    typedef eosio::multi_index<N(pollvote), pollvote, eosio::indexed_by<N(poll_id), eosio::const_mem_fun<pollvote, uint64_t, &pollvote::by_poll_id>>> t_pollvote;

    t_poll _poll;
    t_pollvote _pollvote;

public:
    youvote(account_name self): eosio::contract(self), _poll(self, self), _pollvote(self, self) {}

    // @abi action
    void addpoll(std::string poll_name) {
        require_auth(_self);
        _poll.emplace(_self, [&](auto& p) {
            p.key = _poll.available_primary_key();
            p.poll_id = _poll.available_primary_key();
            p.poll_name = poll_name;
            p.poll_status = 0;
            p.option = "";
            p.count = 0;
        });
    }

    // @abi action
    void rmpoll(std::string poll_name) {
        require_auth(_self);

        std::vector<uint64_t> keyForDeletion;
        for (auto& item : _poll) {
            if (item.poll_name == poll_name) {
                keyForDeletion.push_back(item.key);
            }
        }

        for (uint64_t key : keyForDeletion) {
            auto itr = _poll.find(key);
            if (itr != _poll.end()) {
                _poll.erase(itr);
            }
        }

        std::vector<uint64_t> keyForDeletionFromVote;
        for (auto& item : _pollvote) {
            if (item.poll_name == poll_name) {
                keyForDeletionFromVote.push_back(item.key);
            }
        }

        for (uint64_t key : keyForDeletionFromVote) {
            auto itr = _pollvote.find(key);
            if (itr != _pollvote.end()) {
                _pollvote.erase(itr);
            }
        }
    }

    // @abi action
    void addpollopt(std::string poll_name, std::string option) {
        require_auth(_self);

        for (auto& item : _poll) {
            if (item.poll_name == poll_name) {
                if (item.poll_status == 0) {
                    _poll.emplace(_self, [&](auto& p) {
                        p.key = _poll.available_primary_key();
                        p.poll_id = item.poll_id;
                        p.poll_name = poll_name;
                        p.poll_status = 0;
                        p.option = option;
                        p.count = 0;
                    });
                }
                break;
            }
        }
    }

    // @abi action
    void rmpollopt(std::string poll_name, std::string option) {
        require_auth(_self);

        std::vector<uint64_t> keyForDeletion;
        for (auto& item : _poll) {
            if (item.poll_name == poll_name && item.option == option) {
                keyForDeletion.push_back(item.key);
            }
        }

        for (uint64_t key : keyForDeletion) {
            auto itr = _poll.find(key);
            if (itr != _poll.end()) {
                _poll.erase(itr);
            }
        }
    }

    // @abi action
    void vote(std::string poll_name, std::string option, account_name account) {
        require_auth(account);

        std::vector<uint64_t> keyForModify;
        for (auto& item : _poll) {
            if (item.poll_name == poll_name && item.option == option) {
                if (item.poll_status != 1) return;
                keyForModify.push_back(item.key);
                break;
            }
        }

        for (auto& v : _pollvote) {
            if (v.account == account) return;
        }

        uint64_t poll_id;

        for (uint64_t key : keyForModify) {
            auto itr = _poll.find(key);
            if (itr != _poll.end()) {
                _poll.modify(itr, _self, [&](auto& p) {
                    p.count++;
                });
                poll_id = itr->poll_id;
            }
        }

        _pollvote.emplace(_self, [&](auto& pv) {
            pv.key = _pollvote.available_primary_key();
            pv.poll_id = poll_id;
            pv.poll_name = poll_name;
            pv.account = account;
        });
    }

    // @abi action
    void status(std::string poll_name) {
        require_auth(_self);

        std::vector<uint64_t> keyForModify;
        for (auto& item : _poll) {
            if (item.poll_name == poll_name) {
                keyForModify.push_back(item.key);
            }
        }

        for (uint64_t key : keyForModify) {
            auto itr = _poll.find(key);
            if (itr != _poll.end()) {
                _poll.modify(itr, _self, [&](auto& p) {
                    p.poll_status++;
                });
            }
        }
    }

    // @abi action
    void statusreset(std::string poll_name) {
        require_auth(_self);

        std::vector<uint64_t> keyForModify;
        for (auto& item : _poll) {
            if (item.poll_name == poll_name) {
                keyForModify.push_back(item.key);
            }
        }

        for (uint64_t key : keyForModify) {
            auto itr = _poll.find(key);
            if (itr != _poll.end()) {
                _poll.modify(itr, _self, [&](auto& p) {
                    p.poll_status = 0;
                });
            }
        }
    }
};

EOSIO_ABI(youvote, (addpoll)(rmpoll)(addpollopt)(rmpollopt)(vote)(status)(statusreset))