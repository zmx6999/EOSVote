package youvote

import "github.com/eoscanada/eos-go"

type VoteData struct {
	PolloptId uint64 `json:"pollopt_id"`
	Account eos.AccountName `json:"account"`
}

func NewVote(polloptId uint64, account string) *eos.Action {
	return NewAction("vote", account, VoteData{
		PolloptId:polloptId,
		Account:eos.AN(account),
	})
}

func Vote(polloptId uint64, account string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewVote(polloptId, account), privateKey)
	return
}
