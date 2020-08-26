package exer10

func Fibonacci(n uint) uint {
	return Fib(n, 0)
}

func Fibby(n uint) uint {
	if(n == 0){return 0}
	if(n == 1){return 1}
	return (Fibby(n-1) + Fibby(n-2))
}

func FibbyPar(n uint, cutoff uint, channel chan uint) {
	channel <- Fib(n-1,cutoff)
	channel <- Fib(n-2,cutoff)
}

func Fib(n uint, cutoff uint) uint {
	var fibNum uint = 0
	//cutoff = 10 is a reasonable default
	if(n < cutoff || n < 2){
		fibNum = Fibby(n)
	} else{
		fibChan := make(chan uint)
		go FibbyPar(n, cutoff, fibChan)
		fibNum = <-fibChan + <-fibChan
	}
	return fibNum
}