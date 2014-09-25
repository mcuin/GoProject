//Sum of multiples of 3 & 5 from Project Euler

package main

import "fmt"

func loop() {
	var j int
	for i := 3; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			j += i
		}
	}
	fmt.Println(j)
}

func main() {
	loop()
}
