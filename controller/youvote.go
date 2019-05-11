package controller

import (
	"190511/youvote"
	"encoding/hex"
)

type YouVoteController struct {
	BaseController
}

func (this *YouVoteController) AddPoll()  {
	data := this.getPost()
	pollName, ok := data["poll_name"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	privateKey, ok := data["private_key"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	_, r, err := youvote.AddPoll(pollName, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) AddPollOpt()  {
	data := this.getPost()
	_pollId, ok := data["poll_id"].(float64)
	if !ok || _pollId < 0 {
		this.error(1012, "Invalid Request")
		return
	}
	pollId := uint64(_pollId)

	option, ok := data["option"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	privateKey, ok := data["private_key"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	_, r, err := youvote.AddPollOpt(pollId, option, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) Vote()  {
	data := this.getPost()
	_polloptId, ok := data["pollopt_id"].(float64)
	if !ok || _polloptId < 0 {
		this.error(1012, "Invalid Request")
		return
	}
	polloptId := uint64(_polloptId)

	account, ok := data["account"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	privateKey, ok := data["private_key"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	_, r, err := youvote.Vote(polloptId, account, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) RmPoll()  {
	data := this.getPost()
	_pollId, ok := data["poll_id"].(float64)
	if !ok || _pollId < 0 {
		this.error(1012, "Invalid Request")
		return
	}
	pollId := uint64(_pollId)

	privateKey, ok := data["private_key"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	_, r, err := youvote.RmPoll(pollId, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) RmPollOpt()  {
	data := this.getPost()
	_polloptId, ok := data["pollopt_id"].(float64)
	if !ok || _polloptId < 0 {
		this.error(1012, "Invalid Request")
		return
	}
	polloptId := uint64(_polloptId)

	privateKey, ok := data["private_key"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	_, r, err := youvote.RmPollOpt(polloptId, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) Status()  {
	data := this.getPost()
	_pollId, ok := data["poll_id"].(float64)
	if !ok || _pollId < 0 {
		this.error(1012, "Invalid Request")
		return
	}
	pollId := uint64(_pollId)

	privateKey, ok := data["private_key"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	_, r, err := youvote.Status(pollId, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) Reset()  {
	data := this.getPost()
	_pollId, ok := data["poll_id"].(float64)
	if !ok || _pollId < 0 {
		this.error(1012, "Invalid Request")
		return
	}
	pollId := uint64(_pollId)

	privateKey, ok := data["private_key"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	_, r, err := youvote.Reset(pollId, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) GetList()  {
	table := this.GetString("t")
	data, err := youvote.GetList(table)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(data)
}
