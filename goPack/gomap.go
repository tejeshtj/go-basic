package main
import "fmt"
	func main(){

		peopleDetail:=make(map[string] int)//map with string key ND INT VALUE

		peopleDetail["teju"]=22
		peopleDetail['vaishu']=22

		fmt.Println(peopleDetail)
		fmt.Println(len(peopleDetail))

		// map inside a map

		studentDetail :map[string] map[string] string{
			"tejesh":map[string]string{
				"nickname":"tj",
				"age":"twenty three",
			},

			"vaishnavi":map[string]string{
				"nickname":"dora",
				"age":"fifteen",
			},
		}

		if temp,hero:=studentDetail["tejesh"];hero{
			fmt.Println(temp["nickname"],temp['age'])
		}
	}
