package main
import(
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //import your used driver
)

//Model Struct
type User struct{
	ID int
	Name string `orm:"size(100)"`
}

func init(){
	// set default database
	orm.RegisterDataBase("default","mysql","root:root@/go_content?charset=utf8",30);
	//regiser model
	orm.RegisterModel(new(User));
	//create table
	orm.RunSyncdb("default",false,true);
}

func main(){
    o:=orm.NewOrm();
    user:=User{Name:"wanjn"}
    //insert
    id,err:=o.Insert(&user);
    fmt.Printf("ID:%d,ERR:%v\n",id,err);
}
