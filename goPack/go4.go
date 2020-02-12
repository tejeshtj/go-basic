package main
import "fmt"
	func main(){
		var name string="tejesh tj"
		fmt.Println(len(name))
		fmt.Println(name," is a king ")

		k:=1
		age:=17
		for j:=1;j<=10;j++{ // this is how we use normal for loop
			fmt.Println(j)
		}

		for k<=10{ // this if for while loop
			fmt.Println(k)
			k++
		}

		for i := 0; i <5; i++ {

			for j:=1;j<=i;j++{
			fmt.Printf("*")
		}
			fmt.Println()
		}

		if age>18{
			fmt.Println("go  and vote idiot")		
		}else{
			fmt.Println(" shut up go home")
		}

		switch age{
		case 16:fmt.Println("prepare for college")
	    case 17:fmt.Println("prepare for exams")
	default:fmt.Println("you are jus living")
		}
	}