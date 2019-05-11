package youvote

import "github.com/eoscanada/eos-go"

type ResetData struct {
	PollId uint64 `json:"poll_id"`
}

func NewReset(pollId uint64) *eos.Action {
	return NewAction("reset", CONTRACT_NAME, ResetData{
		PollId:pollId,
	})
}

func Reset(pollId uint64, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewReset(pollId), privateKey)
	return
}
