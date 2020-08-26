////////////
//Examples//
////////////

//dynamically bound
//class galaxy extends phones {...}

//type inference
/*
var inferType = "my type is inferred"
- type does not need to be declared, the type "String" is inferred
*/

//immutability
/*
input: var mut = 5
input: mut = mut + 1
output: 6
input: val immut = 5
input: immut = immut + 1
output: error: reassignment to val
*/

//lazy evaluation
def ifStrict() : Boolean = {
	println("I'm strict")
	return false
}
def testLazy(){
	if(false && ifStrict()){
		println("this branch never used")
	}
	else {
		println("If not strict, I'm lazy")
	}
}

/*
input (using && in testLazy, the lazy AND): testLazy()
output: If not strict, I'm lazy
input (using & in testLazy, the strict AND): testLazy()
output: I'm strict\n If not strict, I'm lazy
*/

//pattern matching
def patternMatch(num: Int) : String = {
  num match {
    case 1 =>
      "case 1"
    case 2 =>
      "case 2"
    case _ =>
      "other cases"
  }
}
/*
input: patternMatch(1)
output: "case 1"
input: patternMatch(15)
output: "other cases"
*/

//currying
def curryAdd(a: Int)(b: Int)(c: Int) : Int = {
	return (a+b+c)
}
/*
input: var curryAdder = curryAdd(1)(1)(_)
output: Int => Int = <function1>
input: curryAdder(1)
output: 3
*/

//First-Class and Higher-Order Functions
/*
reference currying
*/

//callback
def cbPrint(callback: => Unit) {
	callback
}
/*
input: cbPrint(println("I'm a callback"))
output: I'm a callback
*/

//concurrency
//runnable
/*
class aThread extends Runnable {
	def threadCode {
		//define thread behavior here
	}
}
new Thread(new aThread).start //start thread
//futures
afuture = Future {
	val aThread = function()
	aThread
}
*/
/*
execute aThread in a separate thread, working to resolve future immediately, but not blocking the completion of the Future itself
*/