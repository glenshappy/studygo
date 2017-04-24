package main

import (
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
)

func echoHandler(ws *websocket.Conn) {
    msg := make([]byte, 512)
    n, err := ws.Read(msg)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Receive: %s\n", msg[:n])

    send_msg := "[" + string(msg[:n]) + "]"
    m, err := ws.Write([]byte(send_msg))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Send: %s\n", msg[:m])
}

func main() {
    http.Handle("/echo", websocket.Handler(echoHandler))
    http.Handle("/", http.FileServer(http.Dir(".")))

    err := http.ListenAndServe(":8033", nil)

    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
