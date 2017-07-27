package main
import "fmt"
type Cee float64
type Student struct {
	name string
	Cee
	age int
        score float64
}
func main() {
	o:=Student{"wjn",2.33333,33,2.2}
	fmt.Println(o);
	fmt.Println(o.Cee);
}
