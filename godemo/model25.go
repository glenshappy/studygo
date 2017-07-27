package main
import "runtime"
import "fmt"
func main() {
	done:=false
	go func(){
		done=true
		fmt.Println(9999999);
	}()
	for !done {
		fmt.Println("not done");
		runtime.Gosched();
	}
	fmt.Println("done!");
}
