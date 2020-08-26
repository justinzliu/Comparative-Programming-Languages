import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate

-- Merging --

--merge
merge :: Ord a => [a] -> [a] -> [a]
merge [] [] = []
merge (next1:tail1) [] = next1 : tail1
merge [] (next2:tail2) = next2 : tail2
merge (next1:tail1) (next2:tail2)
   | next1 < next2         = next1 : (merge tail1 (next2:tail2))
   | otherwise             = next2 : (merge (next1:tail1) tail2)

-- Tail Recursive Hailstone --

--hailstone
hailstone :: Int -> Int
hailstone a
   | even a       = a `div` 2
   | otherwise    = (a*3)+1

--hailLen
hailLen :: Int -> Int
hailLen n = hailTail 0 n
   where
      hailTail a 1 = a
      hailTail a n = hailTail (a+1) (hailstone n)

-- Factorial --

--fact
fact :: Int -> Int
fact n
   | n == 0        = 1
   | n == 1        = 1
   | otherwise     = n * (fact (n-1))

--fact'
fact' :: Int -> Int
fact' n = foldl (*) 1 [1..n]

-- Haskell Library and Dates --

--daysInYear
daysInYear :: Integer -> [Day]
daysInYear y = [jan1..dec31]
   where 
      jan1 = fromGregorian y 1 1
      dec31 = fromGregorian y 12 31

--isFriday
isFriday :: Day -> Bool
isFriday day
   | snd (mondayStartWeek day) == 5    = True
   | otherwise                         = False

--isPrimeDay
--divsors
divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]
--getDay
getDay :: (Integer, Int, Int) -> Int
getDay (y,m,d) = d

isPrimeDay :: Day -> Bool
isPrimeDay day
   | divisors (getDay (toGregorian day)) == []   = True
   | otherwise                                   = False

--primeFridays
primeFridays :: Integer -> [Day]
primeFridays year = [x | x <- (daysInYear year), (isPrimeDay x), (isFriday x)]
