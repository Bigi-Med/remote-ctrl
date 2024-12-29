package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	//"remote"
)

func main(){
    /*comd := "loginctl"   
    arg1 := "lock-session"

    cmd := exec.Command(comd,arg1)

    stdout, stderr := cmd.Output()

    if stderr != nil {
     fmt.Println(stderr)   
     return
    }

    fmt.Println(string(stdout))*/

    sock, err := net.Listen("tcp","0.0.0.0:8080")

    if err!=nil {
        fmt.Println(err)    
        os.Exit(1)
    }

    fmt.Println("Listening in PORT: 8080")

    for {
        conn, err := sock.Accept()

        if err != nil{
            fmt.Println(err)
            os.Exit(1)
        }

        parser(conn)
    }

}
 func router(conn net.Conn){

 }

 func parser(conn net.Conn){
    data := make([]byte,1024)
    conn.Read(data)
    dataStr := string(data)
    requestLine := strings.Split(dataStr," ")
    path := requestLine[1]
    switch path{
        case "/shutdown":   
            fmt.Println("shutdown")
        case "/lock":
            fmt.Println("lock")
        case "/url":
            fmt.Println("url")
        case "/health":
            fmt.Println("health")
        default:
            fmt.Println("tf u going?")

    }

 }
