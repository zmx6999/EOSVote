package controller

import (
	"190510/youvote"
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
	pollName, ok := data["poll_name"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

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

	_, r, err := youvote.AddPollOpt(pollName, option, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) Vote()  {
	data := this.getPost()
	pollName, ok := data["poll_name"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

	option, ok := data["option"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

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

	_, r, err := youvote.Vote(pollName, option, account, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) RmPoll()  {
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

	_, r, err := youvote.RmPoll(pollName, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) RmPollOpt()  {
	data := this.getPost()
	pollName, ok := data["poll_name"].(string)
	if !ok {
		this.error(1012, "Invalid Request")
		return
	}

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

	_, r, err := youvote.RmPollOpt(pollName, option, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) Status()  {
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

	_, r, err := youvote.Status(pollName, privateKey)
	if err != nil {
		this.error(1012, err.Error())
		return
	}

	this.success(map[string]interface{}{"transaction_id": hex.EncodeToString(r.Processed.ID)})
}

func (this *YouVoteController) StatusReset()  {
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

	_, r, err := youvote.StatusReset(pollName, privateKey)
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
