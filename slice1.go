package main
import(
	"fmt"
)
func printSlice(s string,x []int){
	fmt.Printf("%s len=%d cap=%d %v \n",
		s,len(x),cap(x),x)
}
func main(){
	a := make([]int,5)
	printSlice("a",a)
	fmt.Println(a,len(a),cap(a))

}
