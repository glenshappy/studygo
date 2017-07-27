package main;
import(
	"fmt"
)
func producer(c chan int64, max int) string {
	var str string;
	defer 	close(c)
	for i:=0;i<max;i++ {
		select {
			case c <- 0:
				str+="0"
			case c <- 1:
				str+="1"
		}
	}
	return str;
}
func main() {
	c:=make(chan int64,8)
	rs:=producer(c,8);
	fmt.Println(rs);
}
