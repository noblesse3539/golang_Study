package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	q2 := months[4:7]
	summer := months[6:9]
	t := summer[0:]
	fmt.Println(q2)
	fmt.Println(summer)
	fmt.Printf("len : %d \ncap : %d\n", len(summer), cap(summer))
	fmt.Println(t)

}
