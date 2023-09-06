package main
import "fmt"
func main() {
	var i int32 = 12
	var j float64 = float64(i)
	var k int64 = int64(j)
	fmt.Printf("i = %v,j = %v,k = %v",i,j,k)
}