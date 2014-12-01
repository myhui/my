package admin

type MainController struct {
	baseController
}

func (this *MainController) Index() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.display("mian")
}
