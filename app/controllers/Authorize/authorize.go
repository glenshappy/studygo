package controllers;
import (
	"github.com/revel/revel"
	"fmt"
	"admin/app/models"
	"strconv"
	"reflect"
)
type Authorize struct {
	*revel.Controller
}
func (c *Authorize) Index(sysmenumodel *models.Sysmenu) revel.Result {
	var search string = c.Params.Query.Get("search");
	var pageParam string = c.Params.Get("page");
	var page int64=1;
	if len(pageParam)>0 {
		page,_ = strconv.ParseInt(pageParam,10,64)
	}
	fmt.Println("pageparam:",pageParam);
	var perpage int64 = 10;
	title:="这是一个新版权限管理系统"
	c.Render(title);
	lists,pages,where:=sysmenumodel.GetSysmenuList(search,page,perpage);
	fmt.Println("ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss");
	fmt.Println("lists:",lists,"pages:",pages,"where:",where);
	fmt.Println("liststype:",reflect.TypeOf(lists));
	menu_list:=lists
	c.Render(menu_list);
	c.Render(pages);
	return c.RenderTemplate("Authorize/index.html");
}
func (c *Authorize) EditAuthorize() revel.Result{
	var menu_idstr string = c.Params.Get("id");
	var menu_id int64=0;
	var ac = c.Params.Get("ac");
	if len(menu_idstr)>0 {
		menu_id,_ = strconv.ParseInt(menu_idstr,10,64);
	}
	if c.Request.Method=="GET" && menu_id!=0 {
		ac = "edit"
		c.Render(ac);
		return c.RenderTemplate("Authorize/manage/edit.html");
	}
	ac="add";
	return c.RenderTemplate("Authorize/manage/edit.html");
}
