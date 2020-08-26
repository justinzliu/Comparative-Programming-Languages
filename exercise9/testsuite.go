package main
//export PATH=$PATH:/usr/local/go/bin
//copy and paste exer9 directory into home/go/src
//go test exer9 -v

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//Point Structure
type Point struct {
	x  float64
	y  float64
}
func NewPoint(newx float64, newy float64) Point {
	var point Point
	point.x = newx
	point.y = newy
	return point
}
func (p Point) String() string {
	stringed := fmt.Sprintf("(%v, %v)", p.x, p.y)
	return stringed
}
func (p Point) Norm() float64 {
	return (math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2)))
}

//Scale
func (p *Point) Scale(factor int) {
	p.x = p.x * float64(factor)
	p.y = p.y * float64(factor)
}

//Rotate
func (p *Point) Rotate(angle float64) {
	var x = p.x
	var y = p.y
	p.x = (x * math.Cos(angle)) - (y * math.Sin(angle)) 
	p.y = (x * math.Sin(angle)) + (y * math.Cos(angle))
}

//Random Arrays
func RandomArray(length int, maxInt int) []int {
	randArr := make([]int, length)
	randGen := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	for i := 0; i < length; i++ {
		randArr[i] = (randGen.Int()) % maxInt
	}
	return randArr
}

//Array Summary Stats
type sums struct {
	mean float64
	stddev float64
}
func MeanStddevCalc(arr []int, result chan sums) {
	arrSums := sums{mean: 0, stddev: 0}
	for _, elem := range arr {
		arrSums.mean = arrSums.mean + float64(elem)
		arrSums.stddev = arrSums.stddev + math.Pow(float64(elem),2)
	}
	result <- arrSums
	//close(result)
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
	//testing
	/*
	arrSums_test := sums{mean: 0, stddev: 0}
	for _, elem := range arr {
		arrSums_test.mean = arrSums_test.mean + float64(elem)
		arrSums_test.stddev = arrSums_test.stddev + math.Pow(float64(elem),2)
	}
	arrSums_test.mean = arrSums_test.mean/float64(arrLen)
	arrSums_test.stddev = math.Sqrt( (arrSums_test.stddev/float64(arrLen)) - math.Pow((arrSums_test.mean),2) )
	fmt.Println("arrTest correct is: ", arrSums_test)
	*/

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
//Insertion Sort
func InsertionSort(arr []float64) {
	for i, elem := range arr {
		sortedIndex := i-1
		target := elem
		targetIndex := i
		for j := sortedIndex; j >= 0; j-- {
			if target < arr[j] {
				arr[targetIndex], arr[j] = arr[j], arr[targetIndex]
				targetIndex = j
			}
		}
	}
}
//Quick Sort
func QuickSort(arr []float64) {
	arrLen := len(arr)
	if arrLen > 1 {
		pivotIndex := arrLen/2
		leftIndex, rightIndex := 0, arrLen-1
		arr[pivotIndex], arr[rightIndex] = arr[rightIndex], arr[pivotIndex]
		for i := range arr {
			if arr[i] < arr[rightIndex] {
				arr[leftIndex], arr[i] = arr[i], arr[leftIndex]
				leftIndex = leftIndex + 1
			}
		}
		arr[leftIndex], arr[rightIndex] = arr[rightIndex], arr[leftIndex]
		QuickSort(arr[:leftIndex])
		QuickSort(arr[leftIndex+1:])
	}
}


func main() {
	//Scale test
	p1 := NewPoint(1,2)
	p1.Scale(5)
	fmt.Println(p1.String())
	//Rotate test
	p2 := NewPoint(1,0)
	p2.Rotate(math.Pi / 2)
	fmt.Println(p2.String())
	p2.Rotate(math.Pi / 2)
	fmt.Println(p2.String())
	//Random Array test
	arrTest := RandomArray(24,5)
	fmt.Println("arrTest is: ", arrTest)
	//MeanStddev test
	mean, stddev := MeanStddev(arrTest,6)
	fmt.Println("mean is:", mean, "\nstddev is:", stddev)
	//Insertion Sort test
	arrTest2 := []float64{55.0,22.15,5.0,88.1,1.1}
	InsertionSort(arrTest2)
	fmt.Println(arrTest2)
	//Quicksort test
	arrTest3 := []float64{55.0,22.15,5.0,88.1,1.1}
	QuickSort(arrTest3)
	fmt.Println(arrTest3)
}	