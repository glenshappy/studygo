package main
import "fmt"
func main() {
	x:="NN33中文abcdefg"
	xbytes:=[]byte(x)
	xbytes[0] = 'T'
	fmt.Println(string(xbytes));
}
