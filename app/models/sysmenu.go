// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

package models

//系统菜单
import "fmt"
//import "time"
import "admin/utils"
import "html/template"
import "github.com/revel/revel"

type Sysmenu struct {
	Id          int64       `xorm:"pk"`
	En_name     string      `xorm:"char(100)"`
	Cn_name     string      `xorm:"char(100)"`
	Remark      string      `xorm:"varchar(200)"`
	Opusername  string      `xorm:"char(100)"`
	Optime      int64       `xorm:"int(11)"`
	Pid         int64       `xorm:"int(11)"`
	Status      int8        `xorm:"tinyint(1)"`
}

//根据Id获取信息
func (u *Sysmenu) GetById(Id int64) *Sysmenu {

	user := new(Sysmenu)
	has, err := DB_Read.Table("sysmenu").Id(Id).Get(user)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
//		user_group := new(User_Group)
//		user.UserGroup = user_group.GetById(user.Groupid)
	}

	return user
}

//获取菜单列表
func (u *Sysmenu) GetSysmenuList(search string, Page int64, Perpage int64) (user_arr []*Sysmenu, html template.HTML, where map[string]interface{}) {
	//初始化菜单
	user_list := []*Sysmenu{}

	//查询条件
	var WhereStr string = " 1 AND "

	if len(search) > 0 {
		//解码
		where = utils.DecodeSegment(search)

		revel.WARN.Println(where)

		if where["start_time"] != "" {
			WhereStr += " `optime` >='" + fmt.Sprintf("%s", where["start_time"]) + " 00:00:00' AND "
		}

		if where["end_time"] != "" {
			WhereStr += " `optime` <='" + fmt.Sprintf("%s", where["end_time"]) + " 23:59:59' AND "
		}


		if where["type"] != "" && where["keyword"] != "" {

			if where["type"] == "1" {
				//菜单名
				WhereStr += " `cn_name` ='" + fmt.Sprintf("%s", where["keyword"]) + "' AND "
			} else if where["type"] == "2" {
				//操作人
				WhereStr += " `opusername` =" + fmt.Sprintf("%s", where["keyword"]) + " AND "
			} 
		}
	}

	WhereStr += " 1 "

	//查询总数
	user := new(Sysmenu)
	Total, err := DB_Read.Table("sysmenu").Where(WhereStr).Count(user)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	if len(search) > 0 {
		Pager.SubPage_link = "/Authorize/" + search + "/"
	} else {
		Pager.SubPage_link = "/Authorize/"
	}

	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	DB_Read.Table("sysmenu").Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&user_list)

	if len(user_list) > 0 {
//		user_group := new(User_Group)
//		for i, v := range user_list {
		//	tm:=time.Unix(v.Optime,0);
		//	user_list[i].Optimeformat = tm.Format("2006-01-02 03:04:05 PM");
//		}
	}

	return user_list, pages, where
}
