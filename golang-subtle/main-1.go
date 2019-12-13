package main

import "crypto/subtle"
import "fmt"
import "time"

func main() {
	var x int64
	var y int64
	x = 99999999999999999
	y = 99999999999999999
	start := time.Now()
	result := subtle.ConstantTimeEq(int32(x), int32(y)) & subtle.ConstantTimeEq(int32(x>>32), int32(y>>32))
	fmt.Println(time.Since(start))
	fmt.Println(result)
}
