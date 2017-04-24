package main;
import(
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
)
func echoHandler(ws *websocket.Conn){
	readData:=make([]byte,512)
	n,err:=ws.Read(readData)
	if err!=nil {
		fmt.Println("read_error:"+err.Error());
	}
	fmt.Println("from client:",string(readData));
	writeData:="["+string(readData[:n])+"]";
	_,err1:=ws.Write([]byte(writeData))
	if err1!=nil {
		fmt.Println("write_error:",err1.Error());
	}
	fmt.Println("send to client:",string(writeData));
}
func main(){
	http.Handle("/echo",websocket.Handler(echoHandler))
	http.Handle("/",http.FileServer(http.Dir(".")))
	err:=http.ListenAndServe(":8033",nil)
	if err!=nil {
		fmt.Println("ListenAndServe:"+err.Error())
	}
}
