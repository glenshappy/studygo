package controllers;
type authorize struct {
	*revel.Controller
}
func (c *authorize) Index(){
	title:="这是一个新版权限管理系统"
	c.Render(title);
	c.RenderTemplate("index");
}
func (c *authorize) EditAuthorize(){
	fmt.Println(3444);
}
