package main
import "fmt"
	func main(){
		x:=10
		y:=5
		changeValuePassingValue(x)
		fmt.Println(x)

		changeValuePassingMemoryLocation(&y)
		fmt.Println(y)


	}
	func changeValuePassingValue(x int){

		x=83;
		fmt.Println("x value in pass value func = ",x)

	}

	func changeValuePassingMemoryLocation(y *int){

		*y=15;
		fmt.Println("y value in pass address func = ",y)
	}