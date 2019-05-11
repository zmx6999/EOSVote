package youvote

import "github.com/eoscanada/eos-go"

type AddPollData struct {
	PollName string `json:"poll_name"`
}

func NewAddPoll(pollName string) *eos.Action {
	return NewAction("addpoll", CONTRACT_NAME, AddPollData{
		PollName:pollName,
	})
}

func AddPoll(pollName string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewAddPoll(pollName), privateKey)
	return
}
