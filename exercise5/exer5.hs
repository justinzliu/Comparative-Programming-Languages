-- Built-In Functions
--myIterate
myIterate :: (a -> a) -> a -> [a]
myIterate f a = a : myIterate f (f a)

--myTakeWhile
myTakeWhile :: (a -> Bool) -> [a] -> [a]
myTakeWhile f [] = []
myTakeWhile f (next:tail) = myTake [] f (next:tail)
   where
      myTake acc f [] = acc
      myTake acc f (next:tail)
         | (f next) == True         = myTake (acc ++ [next]) f tail
         | otherwise                = acc

-- Pascal's Triangle
--pascal
pascal :: Int -> [Int]
pascal n
   | n == 0         = [1]
   | otherwise      = [1] ++ prev_list ++ [1]
      where
         prev = pascal (n - 1)
         listTup = zip prev (tail prev)
         prev_list = map (\(x,y) -> x + y) listTup

-- Pointfree Addition
--addPair
addPair :: Num a => (a, a) -> a
addPair = uncurry (+)

-- Pointfree Filtering
--withoutZeros
withoutZeros :: (Eq a, Num a) => [a] -> [a]
withoutZeros = filter (/=0)


-- Exploring Fibonacci
--fib
fib :: Int -> Int
fib n
   | n == 0         = 0
   | n == 1         = 1
   | otherwise      = (fib (n - 1)) + (fib (n - 2))

--fibs
fibs = map fib [0..]

-- Something Else
--things
things :: [Integer]
things = 0 : 1 : zipWith (+) things (tail things)