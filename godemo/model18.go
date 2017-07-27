package main
import(
	"fmt"
)
func main() {
	m:=make(map[string]int,2)
	m["ss"] = 1;
	m["cc"]=2;
	m["dd"]=22;
	fmt.Println(m);
	fmt.Println(len(m));
}
