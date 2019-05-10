package youvote

import "github.com/eoscanada/eos-go"

type StatusData struct {
	PollName string `json:"poll_name"`
}

func NewStatus(pollName string) *eos.Action {
	return NewAction("status", CONTRACT_NAME, StatusData{
		PollName:pollName,
	})
}

func Status(pollName string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewStatus(pollName), privateKey)
	return
}
