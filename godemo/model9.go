package main
import(
	"fmt"
	"time"
)
func main() {
c:=make(chan int64,5)
//defer close(c)
timeout:=make(chan bool)
//defer close(timeout)
go func(){
	time.Sleep(time.Second*10)
//	timeout <- true
}()
select {
	case <-timeout:
		fmt.Println("timeout....");
	case <-c:
		fmt.Println("Read a date");
}
}
