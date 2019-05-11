package youvote

import "github.com/eoscanada/eos-go"

type RmPollData struct {
	PollId uint64 `json:"poll_id"`
}

func NewRmPoll(pollId uint64) *eos.Action {
	return NewAction("rmpoll", CONTRACT_NAME, RmPollData{
		PollId:pollId,
	})
}

func RmPoll(pollId uint64, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewRmPoll(pollId), privateKey)
	return
}
