package youvote

import "github.com/eoscanada/eos-go"

type VoteData struct {
	PollName string `json:"poll_name"`
	Option string `json:"option"`
	Account eos.AccountName `json:"account"`
}

func NewVote(pollName string, option string, account string) *eos.Action {
	return NewAction("vote", account, VoteData{
		PollName:pollName,
		Option:option,
		Account:eos.AN(account),
	})
}

func Vote(pollName string, option string, account string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewVote(pollName, option, account), privateKey)
	return
}
