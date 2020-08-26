///////////////////
//Sample Programs//
///////////////////

//Instructions//
/*
Download scala
navigate to local folder
scalac sample.scala
scala samples
*/


//Samples Test Suite
object samples {
	def main(args: Array[String]) {
		println("map test")
		println("expected: 11,30,13,14,65")
		println("actual: " + map(add,10,Array(1,20,3,4,55)).mkString(",") + "\n")

		println("censor test")
		println("expected: This is a testi CENSORED testi. I really CENSORED this CENSORED works!")
		println("actual: " + censor("This is a testi test testi. I really hope this test works!", Array("test","hope")) + "\n")

		println("strict test:")
		testLazy(1)
		println("lazy test:")
		testLazy(5)
	}

	//Map Implementation in Scala
	/*
	callbacks are powerful functions that allow for functions to be passed into other functions as variables.
	When combined with currying and partial function characteristics, we can easily write the foldl equivalent from Haskell
	*/

	def add(num1: Int)(num2: Int) : Int = {
		num1 + num2
	}

	def map(callback: Int => (Int => Int), num: Int, list: Array[Int]) : Array[Int] = {
		val partialFunc = callback(num) //extra line to illustrate the ability to partially apply a function and store it as a variable
		for(it <- 0 to list.length-1){
			list(it) = partialFunc(list(it))
		}
		return list
	}

	//Censor - a program that takes a string (str) and a list of words to censor out of str
	/*
	Since Scala lacks any notion of pointers and referencing, we should expect programs heaviliy involved with string manipulation to be more costly in memory relative to C and C++.
	Censor relies heavily on support functions that manipulate strings, as such, strings are passed in and out of functions often.
	Since referencing and pointer manipulation is not available in Scala, strings are instead copied over and over again, resulting in a high memory cost
	*/

	def censor(str: String, cenList: Array[String]) : String = {
		var strSplit = str.split(" ")
		var it = 0
		for(i <- strSplit){
			for(j <- cenList){
				if(i == j){
					strSplit(it) = "CENSORED"
				}
			}
			it = it + 1
		}
		var newStr = strSplit.head
		for(i <- strSplit.drop(1)){
			newStr = newStr + " " + i
		}
		return newStr
	}

	//Lazy test - program illustrates the use of the strick & operator and the && lazy operator
	/*
	program illustrates the use of the strick & operator and the && lazy operator.
	The efficiency gained from avoiding possibly costly evaluations may be reaped through Scala's incorporation of lazy operators and lazy / val keyword pairings
	*/
	def ifStrict() : Boolean = {
		println("I'm strict")
		return false
	}
	def testLazy(caseNum: Int) : Unit = {
  		caseNum match {
  			//strict case
  			case 1 => 
  				if(false & ifStrict()){
					println("this branch never used")
				}
				else {
					println("I'm definitely strict")
				}
			//lazy case
  			case _ =>
				if(false && ifStrict()){
					println("this branch never used")
				}
				else {
					println("If not strict, I'm lazy")
				}
		}

	}
}
