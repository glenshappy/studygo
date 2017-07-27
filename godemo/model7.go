package main;
import (
	"fmt"
	"strconv"
)
func main(){
	urls:=[12]string{};
	for i:=0;i<12;i++ {
		urls[i] = "url:"+strconv.Itoa(i);
	}
	fmt.Printf("urls:%v----%T",urls,urls);
}
