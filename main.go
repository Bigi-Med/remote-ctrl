package main

import(
 "fmt"   
 "os/exec"
)

func main(){
    comd := "echo"   
    arg1 := "Hello World"

    cmd := exec.Command(comd,arg1)

    stdout, stderr := cmd.Output()

    if stderr != nil {
     fmt.Println(stderr)   
     return
    }

    fmt.Println(string(stdout))
}
