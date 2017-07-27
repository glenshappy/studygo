package main
import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //import your used driver
       "reflect"
	"time"
)

//Model Struct
type User struct{
	ID int
	Name string `orm:"size(100)"`
}

func init(){
	//set db deriver
        orm.RegisterDriver("mysql",orm.DRSqlite);
	fmt.Printf("DRSqlite:%v\n",orm.DRSqlite);
	// set default database
	orm.RegisterDataBase("default","mysql","root:root@/go_content?charset=utf8",30);
	//regiser model
	orm.RegisterModel(new(User));
	//create table
	orm.RunSyncdb("default",false,true);
}

func main(){
    var qs orm.QuerySeter
    orm.Debug=true
    o:=orm.NewOrm();
    qs=o.QueryTable("user");
    fmt.Printf("type:%v;------qs:%v\n",reflect.TypeOf(o),qs);
    user:=User{Name:"wanjn"}
    u:=User{ID:user.ID}
    err:=o.Read(&u)
    fmt.Printf("ERR:%v\n",err);
    var maps []orm.Params
    num,err:=o.Raw("select * from user where name='wanjn'").Values(&maps);
    for _,term:=range maps{
       fmt.Println(term["i_d"],":",term["name"],":",num);
    }
    maps=nil
    time.Sleep(time.Second*13);
    fmt.Println(11);
    user=User{Name:"wanxh"}
    u=User{ID:user.ID}
    err=o.Read(&u);
    num,err=o.Raw("select * from user where name='wanxh'").Values(&maps)
    for _,term:=range maps {
	fmt.Println(term["i_d"],":",term["name"],":",num);
    }
    time.Sleep(time.Second*10);
}
