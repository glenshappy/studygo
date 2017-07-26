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
func (c *Authorize) EditAuthorize(sysmenumodel *models.Sysmenu) revel.Result{
	if c.Request.Method=="GET" {
		var menu_idstr string = c.Params.Get("id");
	        var menu_id int64=0;
        	var ac = c.Params.Get("ac");
	        if len(menu_idstr)>0 {
                	menu_id,_ = strconv.ParseInt(menu_idstr,10,64);
        	}
	        if c.Request.Method=="GET" && menu_id!=0 {
        	        ac = "edit"
	                c.Render(ac);
        	}
	        ac="add";
                return c.RenderTemplate("Authorize/manage/edit.html");
	}else{
		var cn_name string= c.Params.Form.Get("cn_name");
	        var en_name string= c.Params.Form.Get("en_name");
        	var remark string = c.Params.Form.Get("remark");
	        sysmenumodel.Cn_name = cn_name;
        	sysmenumodel.En_name = en_name;
	        sysmenumodel.Remark = remark;
        	var ret map[string]string;
	        rs:=sysmenumodel.Save();
        	if rs==true {
                	ret["status"] = "200";
	                ret["msg"] = "成功";
        	        return c.RenderJSON(ret);
	        }else{
        	        ret["status"] = "201";
                	ret["msg"] = "失败"
        	        return  c.RenderJSON(ret);
        	}
	}
}
func (c *Authorize) Add(sysmenumodel *models.Sysmenu) revel.Result {
	   var cn_name string= c.Params.Get("cn_name");
           var en_name string= c.Params.Get("en_name");
           var remark string = c.Params.Get("remark");
           sysmenumodel.Cn_name = cn_name;
           sysmenumodel.En_name = en_name;
           sysmenumodel.Remark = remark;
           var ret map[string]string=make(map[string]string);
           rs:=sysmenumodel.Save();
           if rs==true {
                   ret["status"] = "200";
                   ret["msg"] = "成功";
                   return c.RenderJSON(ret);
           }else{
                   ret["status"] = "201";
                   ret["msg"] = "失败"
                   return  c.RenderJSON(ret);
           }
}
func (c *Authorize) Edit(sysmenumodel *models.Sysmenu) revel.Result {
           var menu_id_str string = c.Params.Form.Get("menu_id");
	   var menu_id int64;
           var cn_name string= c.Params.Form.Get("cn_name");
           var en_name string= c.Params.Form.Get("en_name");
           var remark string = c.Params.Form.Get("remark");
           sysmenumodel.Cn_name = cn_name;
           sysmenumodel.En_name = en_name;
           sysmenumodel.Remark = remark;
		title:="菜单编辑";
	   if len(menu_id_str)>0 {
		menu_id1,err:=strconv.ParseInt(menu_id_str,10,64);
		if err!=nil {
			revel.WARN.Println(err);
		}
		menu_id=menu_id1;
	   }else{
		c.Render(title,menu_id);
	   }
           var ret map[string]string;
           rs:=sysmenumodel.Edit(menu_id);
           if rs==true {
                   ret["status"] = "200";
                   ret["msg"] = "成功";
                   return c.RenderJSON(ret);
           }else{
                   ret["status"] = "201";
                   ret["msg"] = "失败"
                   return  c.RenderJSON(ret);
           }
}

