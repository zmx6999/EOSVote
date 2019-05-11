package youvote

import "github.com/eoscanada/eos-go"

type StatusData struct {
	PollId uint64 `json:"poll_id"`
}

func NewStatus(pollId uint64) *eos.Action {
	return NewAction("status", CONTRACT_NAME, StatusData{
		PollId:pollId,
	})
}

func Status(pollId uint64, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewStatus(pollId), privateKey)
	return
}
