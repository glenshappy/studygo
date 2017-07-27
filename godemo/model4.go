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
    var r orm.RawSeter
    r = o.Raw("select * from user where i_d = ? and name = ?",1,"wanjn");
    o.Using("default1");
    dr := o.Driver();
    fmt.Println(dr.Name()=="default"); //true
    fmt.Println(dr.Type() == orm.DRMySQL);
    fmt.Printf("query seter:%v\n",r);
}
