package main
import "fmt"
import "runtime"
func main(){
	fmt.Println("numcpus:",runtime.NumCPU());
}
