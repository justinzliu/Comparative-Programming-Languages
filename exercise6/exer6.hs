import Data.Ratio

-- Rational Numbers
--rationalSum
rationalSum :: Int -> [(Ratio Int)]
rationalSum a = rationalList [] [1..(a-1)]
   where
      rationalList accum [] = accum
      rationalList accum (head:tail) = accum ++ [(%) head (a - head)] ++ (rationalList accum tail)

-- Lowest Terms Only
--rationalSumLowest
rationalSumLowest :: Int -> [(Ratio Int)]
rationalSumLowest a = rationalList [] [1..(a-1)]
   where
      rationalList accum [] = accum
      rationalList accum (head:tail)
         | (gcd head a) == 1      = accum ++ [(%) head (a - head)] ++ (rationalList accum tail)
         | otherwise              = accum ++ (rationalList accum tail)

-- All Rational Numbers
--rationals
rationals :: [(Ratio Int)]
rationals = rationalsList [] [1..]
   where
      rationalsList accum [] = accum
      rationalsList accum (head:tail) = rationalSumLowest head ++ (rationalsList accum tail)

-- Input/Output
--sumFile
sumFile :: IO ()
sumFile = do
   content <- readFile "input.txt"
   let contentList = (lines :: String -> [String]) content
   let intList = map (read :: String -> Int) contentList
   let sumF = sum intList
   putStrLn (show sumF)

-- Quiz 1 Testing

anyNegative [] = False
anyNegative (x:xs) = (x<0) || anyNegative xs

myTake :: Int -> [a] -> [a]
myTake 0 list = []
myTake _ [] = []
myTake num (head:tail)
   | num == 0         = [head]
   | otherwise        = head : myTake (num-1) (tail)

positionSum :: [Int] -> Int
positionSum list = positionSummer 0 1 list
   where
      positionSummer sum index [] = sum
      positionSummer sum index (head:tail) = positionSummer (sum + (index * head)) (index+1) tail