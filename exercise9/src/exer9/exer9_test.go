package exer9

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
	"math"
)

func TestRandomArrays(t *testing.T) {
	length := 1000
	maxint := 100
	arr1 := RandomArray(length, maxint)
	assert.Equal(t, length, len(arr1))
	for _, v := range arr1 {
		assert.True(t, 0 <= v, "contains a negative integer")
		assert.True(t, v < maxint, "contains an integer >=maxint")
	}

	// check that different calls return different results
	arr2 := RandomArray(length, maxint)
	assert.False(t, reflect.DeepEqual(arr1, arr2))
}

func TestArrayStatParallel(t *testing.T) {
	length := 30000000
	maxint := 10000
	arr2 := RandomArray(length, maxint)

	// call MeanStddev single-threaded
	start := time.Now()
	mean2, stddev2 := MeanStddev(arr2, 1)
	end := time.Now()
	dur1 := end.Sub(start)

	// now turn on cuncurrency and make sure we get the same results, but faster
	start = time.Now()
	mean3, stddev3 := MeanStddev(arr2, 3)
	end = time.Now()
	dur2 := end.Sub(start)

	speedup := float64(dur1) / float64(dur2)
	assert.True(t, speedup > 1.25, "Running MeanStddev with concurrency didn't speed up as expected. Sped up by %g.", speedup)
	assert.Equal(t, float32(mean2), float32(mean3)) // compare as float32 to avoid rounding differences
	assert.Equal(t, float32(stddev2), float32(stddev3))
}

// copied from https://golang.org/src/math/rand/rand.go?s=7456:7506#L225 for Go <1.10 compatibility
func shuffle(n int, swap func(i, j int)) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	if n < 0 {
		panic("invalid argument to shuffle")
	}

	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

func benchmarkSorting(b *testing.B, sort func(arr []float64)) {
	const length = 1000
	arr := make([]float64, length)
	for i := 0; i < length; i++ {
		arr[i] = float64(i)
	}

	// run the benchmark
	for n := 0; n < b.N; n++ {
		shuffle(length, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
		sort(arr)
	}
}

func BenchmarkInsertionSort(b *testing.B) { benchmarkSorting(b, InsertionSort) }
func BenchmarkQuickSort(b *testing.B)     { benchmarkSorting(b, QuickSort) }
func BenchmarkFloat64s(b *testing.B)      { benchmarkSorting(b, sort.Float64s) }

//Writing Tests
func TestPointScale(t *testing.T) {
	test1 := NewPoint(1,0)
	test1.Scale(5)

	test2 := NewPoint(5,10)
	test2.Scale(0)

	assert.Equal(t, float32(test1.x), float32(5))
	assert.Equal(t, float32(test1.y), float32(0))
	assert.Equal(t, float32(test2.x), float32(0))
	assert.Equal(t, float32(test2.y), float32(0))
}
func TestPointRotate(t *testing.T) {
	test1 := NewPoint(1,0)
	test1.Rotate(math.Pi / 2)

	test2 := NewPoint(0,1)
	test2.Rotate(math.Pi / 2)

	//truncate point value to int
	assert.Equal(t, float32(int(test1.x)), float32(0))
	assert.Equal(t, float32(int(test1.y)), float32(1))
	assert.Equal(t, float32(int(test2.x)), float32(-1))
	assert.Equal(t, float32(int(test2.y)), float32(0))
}
//Simple, nonconcurrent MeanStddev
func SimpleMeanStddev(arr []int) (mean, stddev float64) {
	arrSums := sums{mean: 0, stddev: 0}
	arrLen := len(arr)
	for _, elem := range arr {
		arrSums.mean = arrSums.mean + float64(elem)
		arrSums.stddev = arrSums.stddev + math.Pow(float64(elem),2)
	}
	arrSums.mean = arrSums.mean/float64(arrLen)
	arrSums.stddev = math.Sqrt( (arrSums.stddev/float64(arrLen)) - math.Pow((arrSums.mean),2) )
	return arrSums.mean, arrSums.stddev
}
func TestMeanStddev(t *testing.T) {
	arrTest := RandomArray(24,5)
	testSums := sums{mean: 0, stddev: 0}
	testSums.mean, testSums.stddev = MeanStddev(arrTest, 6)
	corrSums := sums{mean: 0, stddev: 0}
	corrSums.mean, corrSums.stddev = SimpleMeanStddev(arrTest)

	assert.Equal(t, float32(testSums.mean), float32(corrSums.mean))
	assert.Equal(t, float32(testSums.stddev), float32(corrSums.stddev))
}