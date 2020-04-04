package main

import(
	"fmt"
	"time"
)

func main(){
	i:=2
	fmt.Print("write ",i,"as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}
	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it not weekend")
	}
	fmt.
}