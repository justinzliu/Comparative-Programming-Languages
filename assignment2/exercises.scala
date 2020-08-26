//Instructions//
/*
Download scala
navigate to local folder
scalac exercises.scala
scala exercises
*/

/////////////////
//Main Function//
/////////////////

//Exercises Test Suite
object exercises {
   def main(args: Array[String]) {
   	  println("divisors test")
      println("expected: 2,3,5,6,10,15")
      println("actual: " + divisors(30).mkString(",") + "\n")
      println("expected: 2,4,8,16,32")
      println("actual: " + divisors(64).mkString(",") + "\n")
      println("expected: ")
      println("actual: " + divisors(127).mkString(",") + "\n")

      println("primes test")
      println("expected: 2,3,5,7")
      println("actual: " + primes(7).mkString(",") + "\n")
      println("expected: 2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97")
      println("actual: " + primes(100).mkString(",") + "\n")

      println("join test")
      println("expected: one, two, three")
      println("actual: " + join(" ,", Array[String]("one","two","three")) + "\n")
      println("expected: 1+2+3")
      println("actual: " + join("+", Array[String]("1","2","3")) + "\n")
      println("expected: abc")
      println("actual: " + join("X", Array[String]("abc")) + "\n")
      println("expected: ")
      println("actual: " + join("X", Array[String]()) + "\n")
      
      println("pythagorean triples test")
      println("expected: (3,4,5),(6,8,10)")
      println("actual: " + pythagorean(10).mkString(",") + "\n")
      println("expected: (3,4,5),(6,8,10),(5,12,13),(9,12,15),(8,15,17),(12,16,20),(15,20,25),(7,24,25),(10,24,26),(20,21,29),(18,24,30)")
      println("actual: " + pythagorean(30).mkString(",") + "\n")
      
      println("hailSeq test")
      println("expected: 1")
      println("actual: " + hailSeq(1).mkString(",") + "\n")
      println("expected: 11,34,17,52,26,13,40,20,10,5,16,8,4,2,1")
      println("actual: " + hailSeq(11).mkString(",") + "\n")
      println("expected: 6,3,10,5,16,8,4,2,1")
      println("actual: " + hailSeq(6).mkString(",") + "\n")

      println("mergesort test")
      println("expected: 1,2,3,4,5,6,7,8,9")
      println("actual: " + mergesort(Array[Int](1,9,3,2,7,6,4,8,5)).mkString(",") + "\n")
      println("expected: 1,2,3,4,5,6,7,8,9,10")
      println("actual: " + mergesort(Array[Int](6,2,4,8,9,5,3,1,7,10)).mkString(",") + "\n")
      println("expected: ")
      println("actual: " + mergesort(Array[Int]()).mkString(",") + "\n")
      println("expected: 4")
      println("actual: " + mergesort(Array[Int](4)).mkString(",") + "\n")

      println("fib test")
      println("expected: 0")
      println("actual: " + fib(0) + "\n")
      println("expected: 1")
      println("actual: " + fib(1) + "\n")
      println("expected: 1")
      println("actual: " + fib(2) + "\n")
      println("expected: 2")
      println("actual: " + fib(3) + "\n")
      println("expected: 55")
      println("actual: " + fib(10) + "\n")
      println("expected: 6765")
      println("actual: " + fib(20) + "\n")
   }


	/////////////////////
	//Support Functions//
	/////////////////////

	def isEven(num: Int) : Boolean = {
		var ret = false
		if(num%2 == 0){
			ret = true
		}
		return ret
	}

	def getSmaller(num1: Int, num2: Int) : (Int,Int) = {
		var ret = (2,num2)
		if(num1 < num2){ret = (1,num1)}
		return ret
	}

	def hailstone(num: Int) : Int = {
		var ret : Int = 1
		if(isEven(num)){
			ret = num/2
		}
		else {
			ret = (num*3)+1
		}
		return ret
	}

	def hailLen(num: Int) : Int = {
		var len = 0
		if(num == 1){len = 1}
		var currHail = num
		while(currHail > 1){
			len += 1
			currHail = hailstone(currHail)
		}
		return len
	}


	/////////////
	//Exercises//
	/////////////

	//divisors
	def divisors(num: Int) : Array[Int] = {
		var valList = (2 to num/2).toArray
		var divideList = for(divides <- valList if (num % divides) == 0) yield divides
		return divideList
	}

	//primes
	def primes(num: Int) : Array[Int] = {
		var valList = (2 to num).toArray
		var primesList = for(primes <- valList if divisors(primes).isEmpty) yield primes
		return primesList
	}

	//join
	def join(str: String, strArr: Array[String]) : String = {
		var joined = ""
		if (!strArr.isEmpty) {
			for(i <- 0 to (strArr.length-2)){
				joined += strArr(i) + str
			}
			joined += strArr(strArr.length-1)
		}
		return joined
	}

	//pythagorean
	def pythagorean(num: Int) : Array[(Int, Int, Int)] = {
		var pyList = Array[(Int, Int, Int)]()
		for(i <- 2 to num) {
			var c = i*i
			for(j <- 1 to num){
				var a = j*j
				for(k <- j to num){
					var b = k*k
					if((a+b) == c){
						pyList = pyList :+ (j,k,i)
					}
				}
			}
		}
		return pyList
	}

	//hailSeq
	def hailSeq(num: Int) : Array[Int] = {
		var len = hailLen(num)
		var currHail = num
		var hailList = Array[Int](num)
		if(num == 1){len=0}
		while(len > 0){
			currHail = hailstone(currHail)
			hailList = hailList :+ currHail
			len = len-1
		}
		return hailList
	}

	//mergesort
	def merge(firsthalf: Array[Int], lasthalf: Array[Int]) : Array[Int] = {
		var fh_it, lh_it = 0
		var fh_max = firsthalf.length
		var lh_max = lasthalf.length
		var merged = Array[Int]();
		while(fh_it < fh_max && lh_it < lh_max){
			if(firsthalf(fh_it) < lasthalf(lh_it)){
				merged = merged :+ firsthalf(fh_it)
				fh_it = fh_it + 1
			}
			else {
				merged = merged :+ lasthalf(lh_it)
				lh_it = lh_it + 1
			}
		}
		if(fh_it == fh_max){
			while(lh_it < lh_max){
				merged = merged :+ lasthalf(lh_it)
				lh_it = lh_it + 1
			}
		}
		if(lh_it == lh_max){
			while(fh_it < fh_max){
				merged = merged :+ firsthalf(fh_it)
				fh_it = fh_it + 1
			}
		}
		return merged
	}
	def mergesort(list: Array[Int]) : Array[Int] = {
		if(list.isEmpty) {return Array[Int]()}
		if(list.length == 1) {return list}
		var halfpoint = (list.length/2)
		var merged = merge(mergesort(list.take(halfpoint)), mergesort(list.drop(halfpoint)))
		return merged
	}

	//fib
	def fib(num: Int) : Int = {
		var dynFib = Array[Int](0,1)
		var currNum = 2
		while(currNum <= num){
			dynFib = dynFib :+ (dynFib(currNum-1) + dynFib(currNum-2))
			currNum = currNum + 1
		}
		return dynFib(num)
	}

}