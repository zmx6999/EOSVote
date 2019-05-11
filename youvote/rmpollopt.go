package youvote

import "github.com/eoscanada/eos-go"

type RmPollOptData struct {
	PolloptId uint64 `json:"pollopt_id"`
}

func NewRmPollOpt(polloptId uint64) *eos.Action {
	return NewAction("rmpollopt", CONTRACT_NAME, RmPollOptData{
		PolloptId:polloptId,
	})
}

func RmPollOpt(polloptId uint64, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewRmPollOpt(polloptId), privateKey)
	return
}
