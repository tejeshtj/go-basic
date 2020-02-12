package main
import "fmt"
	func main(){

		var a int=5
		var b float64=3.141414
		const pi=25.23524

		x,y :=14,15

		fmt.Println(a," ",b," ",pi," ",x," ",y)
		fmt.Println("sum of a+b = ",a+b);
		fmt.Println("sum of a*b = ",a*b);
		fmt.Println("sum of a/b = ",a/b);
		fmt.Println("sum of a-b = ",a-b);

		val1:=true
		val2:=false
		fmt.Println(val1&&val2);
		fmt.Println(val1||val2);
		fmt.Println(!val2);

		fmt.Println("address of x : ",&x)
	}