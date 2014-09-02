package main
import(
	"fmt"
	"time"
)
func ready(w string,sec int64){
	secs := time.Duration(sec)*time.Second
	time.Sleep(secs)
	fmt.Println(w,"is ready")
}
func num()int{
	a := 10
	return a
}
func main(){
	go ready("roin",5)
	go ready("lgl",5)
	fmt.Println("i am waiting")
	secs := time.Duration(2)*time.Second 
	time.Sleep(5*secs)
	b := num()
	fmt.Println(b)
}
