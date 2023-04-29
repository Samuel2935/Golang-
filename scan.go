package main
import "fmt"
func main(){
	fmt.Printf("num\n")
	num, err := fmt.Scan(&num)
	fmt.Printf(&num)
}