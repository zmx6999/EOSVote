package youvote

import "github.com/eoscanada/eos-go"

type AddPollOptData struct {
	PollId uint64 `json:"poll_id"`
	Option string `json:"option"`
}

func NewAddPollOpt(pollId uint64, option string) *eos.Action {
	return NewAction("addpollopt", CONTRACT_NAME, AddPollOptData{
		PollId:pollId,
		Option:option,
	})
}

func AddPollOpt(pollId uint64, option string, privateKey string) (signedTx *eos.SignedTransaction, r *eos.PushTransactionFullResp, err error) {
	signedTx, r, err = PushAction(NewAddPollOpt(pollId, option), privateKey)
	return
}
