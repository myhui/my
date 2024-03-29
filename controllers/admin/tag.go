package admin

import (
	"my/models"
	"strconv"
	"strings"
)

type TagController struct {
	baseController
}

func (this *TagController) Index() {
	act := this.GetString("act")
	switch act {
	case "batch":
		this.batch()
	default:
		this.list()
	}
}

//标签列表
func (this *TagController) list() {
	var page int64
	var pagesize int64 = 10
	var list []*models.Tag
	var tag models.Tag

	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize

	count, _ := tag.Query().Count()
	if count > 0 {
		tag.Query().OrderBy("-count").Limit(pagesize, offset).All(&list)
	}

	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(page, count, pagesize, "/admin/tag?page=%d").ToString()
	this.display("tag_list")
}

//批量操作
func (this *TagController) batch() {
	ids := this.GetStrings("ids[]")
	op := this.GetString("op")

	idarr := make([]int64, 0)
	for _, v := range ids {
		if id, _ := strconv.Atoi(v); id > 0 {
			idarr = append(idarr, int64(id))
		}
	}

	switch op {
	case "upcount": //更新统计
		for _, id := range idarr {
			obj := models.Tag{Id: id}
			if obj.Read() == nil {
				obj.UpCount()
			}
		}
	case "merge": //合并到
		toname := strings.TrimSpace(this.GetString("toname"))
		if toname != "" && len(idarr) > 0 {
			tag := new(models.Tag)
			tag.Name = toname
			if tag.Read("name") != nil {
				tag.Count = 0
				tag.Insert()
			}
			for _, id := range idarr {
				obj := models.Tag{Id: id}
				if obj.Read() == nil {
					obj.MergeTo(tag)
					obj.Delete()
				}
			}
			tag.UpCount()
		}
	case "delete": //批量删除
		for _, id := range idarr {
			obj := models.Tag{Id: id}
			if obj.Read() == nil {
				obj.Delete()
			}
		}
	}

	this.Redirect("/admin/tag", 302)
}
