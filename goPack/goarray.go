package main
import "fmt"
	func main(){
		var evenNum[5] int

		evenNum[0]=0
		evenNum[1]=2
		evenNum[2]=4
		evenNum[3]=6
		evenNum[4]=8
		fmt.Println(evenNum[2])


		 oddNum :=[]int{1,3,5,7,9} // this can be also defined by mentioning array size in[] 
		 fmt.Println(oddNum[2])

		 	for _,value :=range evenNum{// this does not require any index it uses range of array
		 		fmt.Println(value)
		 	}
		 	for i,value :=range evenNum{ // this prints along with index
		 		fmt.Println(value,i)
		 	}
		 	slicedAray:=evenNum[0:2] // slice and store in array from index 0 till 2
		 	fmt.Println(slicedAray)
		 	/*slicedAray:=evenNum[0:] this will slice from index 0 to end
			slicedAray:=evenNum[:3] this will lice from start till 3 but not 3*/

			array2:=make([]int,1,2) // make an int array with 1,2
			copy(array2,slicedAray) // copy one array into another
			fmt.Println(slicedAray)

			array3:=append(slicedAray,0,2,7,7,7,7)
			fmt.Println(array3)	
	}