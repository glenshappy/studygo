package main
import (
	"fmt"
	"time"
)
func main() {
	t1:=time.Now();
	var maxnum int64=1000000000000;
	var sum,i int64;
	for i=1;i<=maxnum;i++ {
		sum+=i
	}
	duration:=time.Since(t1);
	fmt.Println("duration:",duration);
	fmt.Println("sum:",sum);
}
