package main
//export PATH=$PATH:/usr/local/go/bin
//copy and paste exer10 directory into home/go/src
//go test exer10 -v

import (
	"fmt"
	//"math"
)

type Rational struct {
	n int
	d int
}
/*
func Abs(num int) int {
	if(num < 0){
		num = num*(-1)
	}
	return num
}

func (rat Rational) String() string {
	str1 := ""
	if(rat.n < 0 || rat.d < 0){
		if(!(rat.n < 0 && rat.d < 0)){
			str1 = "-"
		}
	}
	str2 := fmt.Sprintf("%v/%v", Abs(rat.n), Abs(rat.d))
	ret := str1 + str2
	return ret
}

func (rat *Rational) ToLowestTerms() {
	gcd := Gcd(rat.n, rat.d)
	rat.n = rat.n/gcd
	rat.d = rat.d/gcd
}
*/

func allRationals(results chan Rational) {
	rat := Rational{1,1}
	i := 1
	for {
		for rat.d < i {
			if((rat.n + rat.d) == i){
				results <- rat
			}
			rat.d += 1
		}
		rat.d = 1
		i += 1
	}
}

func main() {
	r := Rational{6, -9}
	//r.ToLowestTerms()
	fmt.Println(r)
}