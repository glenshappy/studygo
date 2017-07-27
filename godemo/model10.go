package main
import(
	"fmt"
)
type Human struct {
	height float32
	weight float32
}
type Man struct {
	Human
	sex string
}
func (c *Human) BMindex() (index float32){
	index=c.height*10+c.weight;
	return;
}
/*
func (c *Man) BMindex()  (index float32){
	index=c.height+c.weight
	return
} 
*/
func main() {
	man1:=Man{Human{2.2,3.2},"nan"}
	rs1:=man1.BMindex()
	fmt.Println("man height:",man1.height);
	fmt.Println("man weight:",man1.weight);
	fmt.Println(rs1);
}
