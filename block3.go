package main
import "fmt"


func main(){

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {continue}; {
	// 		fmt.Printf("hi\n")
	// 	}
	// }

	// i := 0
	// for i < 10 {
	// i++
	// 	if i == 5 {continue}; 
	// 	{
			
	// 		fmt.Printf("Hakimi\n")
	// 	}
	// }

	var name string
	var unit string
	var amount int
	var temp string
// pointer is variable that stores the address of an object in the memory of the variable
	fmt.Scanf("%s %s %d %s", &name, &unit, &amount, &temp)
	fmt.Printf("%d %s of land is owned by %s\n", amount, unit, name)
}
