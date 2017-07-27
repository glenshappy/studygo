package main
import(
	"fmt"
)
func main() {
	one:=0;
	one,two:=1,2
	one,two=two,one
	fmt.Println(one,two);
}
