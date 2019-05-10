package youvote

import "github.com/eoscanada/eos-go"

type StatusResetData struct {
	PollName string `json:"poll_name"`
}

func NewStatusReset(pollName string) *eos.Action {
	return NewAction("statusreset", CONTRACT_NAME, StatusResetData{
		PollName:pollName,
	})
}

func StatusReset(pollName string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewStatusReset(pollName), privateKey)
	return
}
