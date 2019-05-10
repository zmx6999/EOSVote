package youvote

import "github.com/eoscanada/eos-go"

type AddPollOptData struct {
	PollName string `json:"poll_name"`
	Option string `json:"option"`
}

func NewAddPollOpt(pollName string, option string) *eos.Action {
	return NewAction("addpollopt", CONTRACT_NAME, AddPollOptData{
		PollName:pollName,
		Option:option,
	})
}

func AddPollOpt(pollName string, option string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewAddPollOpt(pollName, option), privateKey)
	return
}
