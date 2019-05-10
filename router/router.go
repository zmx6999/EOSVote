package router

import (
	"github.com/astaxie/beego"
	"190510/controller"
)

func init()  {
	beego.Router("/poll/add", &controller.YouVoteController{}, "POST:AddPoll")
	beego.Router("/poll/rm", &controller.YouVoteController{}, "POST:RmPoll")
	beego.Router("/poll/status", &controller.YouVoteController{}, "POST:Status")
	beego.Router("/poll/reset", &controller.YouVoteController{}, "POST:StatusReset")
	beego.Router("/pollopt/add", &controller.YouVoteController{}, "POST:AddPollOpt")
	beego.Router("/pollopt/rm", &controller.YouVoteController{}, "POST:RmPollOpt")
	beego.Router("/vote", &controller.YouVoteController{}, "POST:Vote")
	beego.Router("/vote/list", &controller.YouVoteController{}, "GET:GetList")
}
