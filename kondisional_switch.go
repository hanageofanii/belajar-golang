package main
import "fmt"

func main(){
	var point = 3
	switch point{
	case 8:
		fmt.Println("perfect")
	case 5,6,7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}