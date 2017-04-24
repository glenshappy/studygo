package controllers;
import (
	"github.com/revel/revel"
	"fmt"
)
type Authorize struct {
	*revel.Controller
}
func (c *Authorize) Index() revel.Result {
	title:="这是一个新版权限管理系统"
	c.Render(title);
	return c.RenderTemplate("Authorize/index.html");
}
func (c *Authorize) EditAuthorize() revel.Result{
	fmt.Println(3444);
	data:=make(map[string]string)
	data["name"]="wjn"
	data["age"] = "30";
	return c.RenderJSON(data);
}
