package tempconv0

import (
	"fmt"
)

func main() {
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; String을 명시적으로 호출할 필요 없음
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; String을 호출하지 않음
	fmt.Println(float64(c)) // "100"; String을 호출하지 않
}
