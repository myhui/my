package home

import (
	"fmt"
	"my/models"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index() {
	fmt.Println(" start .. ")

	//首页广告轮换
	var carousels []*models.Carousel
	query := new(models.Carousel).Query().Filter("Ishide", 0)
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("rank").Limit(3, 0).All(&carousels)
	}
	fmt.Println("count :%i", len(carousels))

	//--

	this.Data["carousels"] = carousels
	this.display("index")
}

func (this *IndexController) List() {
	fmt.Println(" have .. ")

	this.Data["Website"] = "list.."
	this.Data["Email"] = "astaxie@gmail.com"

	this.display("list")
}

func (this *IndexController) Deatil() {
	fmt.Println(" have .. ")

	this.Data["Website"] = "list.."
	this.Data["Email"] = "astaxie@gmail.com"

	this.display("detail")
}
