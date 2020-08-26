package exer9

import (
	"math"
	"math/rand"
	"time"
)

type sums struct {
	mean float64
	stddev float64
}

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.
	randArr := make([]int, length)
	randGen := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	for i := 0; i < length; i++ {
		randArr[i] = (randGen.Int()) % maxInt
	}
	return randArr
}

func MeanStddevCalc(arr []int, result chan sums) {
	arrSums := sums{mean: 0, stddev: 0}
	for _, elem := range arr {
		arrSums.mean = arrSums.mean + float64(elem)
		arrSums.stddev = arrSums.stddev + math.Pow(float64(elem),2)
	}
	result <- arrSums
}
func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size!")
	}
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.
	arrChan := make(chan sums)
	arrSums := sums{mean: 0, stddev: 0}
	arrLen := len(arr)
	chunkLen := arrLen/chunks
	for i := 0; i < chunks; i++ {
		go MeanStddevCalc(arr[(chunkLen*i):(chunkLen*(i+1))], arrChan)
	}
	var tempSums sums
	for j := 0; j < chunks; j++ {
		tempSums = <-arrChan
		arrSums.mean = arrSums.mean + float64(tempSums.mean)
		arrSums.stddev = arrSums.stddev + float64(tempSums.stddev)
	}
	arrSums.mean = arrSums.mean/float64(arrLen)
	arrSums.stddev = math.Sqrt( (arrSums.stddev/float64(arrLen)) - math.Pow((arrSums.mean),2) )
	return arrSums.mean, arrSums.stddev
}