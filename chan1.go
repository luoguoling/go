package main
import(
    "fmt"
)
var (
    c chan int
)
func write(){
    data := 0
    fmt.Println("等待写入数据",data)
    c <- data
    fmt.Println("写入数据成功")
}
func read(){
    fmt.Println("等待取出数据")
    data := <- c
    fmt.Println("取出数据成功",data)
}
func main(){
    c = make(chan int,8)
    go write()
    read()
    fmt.Println("---------")
    go read()
    write()

}
