package youvote

import "github.com/eoscanada/eos-go"

type RmPollData struct {
	PollName string `json:"poll_name"`
}

func NewRmPoll(pollName string) *eos.Action {
	return NewAction("rmpoll", CONTRACT_NAME, RmPollData{
		PollName:pollName,
	})
}

func RmPoll(pollName string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewRmPoll(pollName), privateKey)
	return
}
