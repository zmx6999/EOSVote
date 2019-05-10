package youvote

import "github.com/eoscanada/eos-go"

type RmPollOptData struct {
	PollName string `json:"poll_name"`
	Option string `json:"option"`
}

func NewRmPollOpt(pollName string, option string) *eos.Action {
	return NewAction("rmpollopt", CONTRACT_NAME, RmPollOptData{
		PollName:pollName,
		Option:option,
	})
}

func RmPollOpt(pollName string, option string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewRmPollOpt(pollName, option), privateKey)
	return
}
