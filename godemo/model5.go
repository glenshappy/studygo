package main
import (
	"fmt"
	"sync"
)
var waitgroup sync.WaitGroup;
func Afunction(shownum int) {
	fmt.Println(shownum);
	waitgroup.Done();
}

func main() {
	for i:=0;i<10;i++ {
		waitgroup.Add(1);
		go Afunction(i);
	}
	waitgroup.Wait();
}
