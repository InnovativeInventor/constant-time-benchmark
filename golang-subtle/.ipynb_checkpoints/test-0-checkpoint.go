package main

import "fmt"
import "time"

func main() {
	var x int64
	var y int64
	x = 11111111111111111
	y = 99999999999999999
	start := time.Now()
	result := 0
	if x == y {
		result = 1
	}
	fmt.Println(time.Since(start))
	fmt.Println(result)
}
