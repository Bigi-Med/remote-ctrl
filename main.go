package main

import (
    "fmt"
    "net"
    "os"
    "strings"
    "remote/httpd"
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
func router(corps httpd.Httpp){
    switch corps.Path{
    case "/shutdown":
        shutdown()
    case "/lock":
        lock()
    case "/url":
        url(corps)
    case "/health":
        health()
    default:
        fmt.Println("How did u get here")
    }
}

func shutdown(){
    fmt.Println("Shutdown")  
}

func lock(){
    fmt.Println("Lock")  
}

func url(corps httpd.Httpp){
    fmt.Println("Url")  
    fmt.Println(corps.Body)
}

func health(){
    fmt.Println("Health")  
}

func parser(conn net.Conn){
    data := make([]byte,1024)
    var body string ;
    conn.Read(data)
    dataStr := string(data)
    slotedData := strings.Split(dataStr,"\r\n")
    for i := 0 ; i<len(slotedData); i++{
        if strings.TrimSpace(slotedData[i]) == ""{
            body = slotedData[i+1]
        }
    }
    requestLine := strings.Split(slotedData[0]," ")
    path := requestLine[1]
    switch path{
        case "/shutdown":   
        corps := httpd.Httpp{
            Path: "/shutdown",
            Body: "",
        }
        router(corps)
    case "/lock":
        corps := httpd.Httpp{
            Path: "/lock",
            Body: body,
        }
        router(corps)
    case "/url":
        corps := httpd.Httpp{
            Path: "/url",
            Body: body,
        }
        router(corps)
    case "/health":
        corps := httpd.Httpp{
            Path: "/health",
            Body: "",
        }
        router(corps)
    default:
        fmt.Println("tf u going?")

    }

 }
