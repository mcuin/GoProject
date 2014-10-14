package main

import "fmt"

func loop() {
	var largestPrime int
	var number = 600851475143
	for i := 3; i < number; i += 2 {
		for j := 3; j < i; j += 2 {
			if i % j == 0 {
				break
			} else {
				largestPrime = i
			}
		}	
	}
	fmt.Println(largestPrime)
}

func main() {
	loop()
}
