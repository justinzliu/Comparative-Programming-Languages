--hailstone
hailstone :: Int -> Int
hailstone a
   | even a       = a `div` 2
   | otherwise    = (a*3)+1

--haillen
hailLen :: Int -> Int
hailLen a
   | a == 1       = 0
   | otherwise    = (hailLen (hailstone a)) + 1

-- divsors and primes
divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]
primes :: Int -> [Int]
primes n = [i | i <- [2..n], divisors i == []]

-- Joining Strings
join :: [Char] -> [[Char]] -> [Char]
join str [] = ""
join str [last] = last
join str (next:tail) = next ++ str ++ (join str tail)

-- Pythagorean Triples
-- received guidance from https://stackoverflow.com/questions/48354060/pythagorean-triple-in-haskell
pythagorean :: Int -> [(Int, Int, Int)]
pythagorean cmax = [(a,b,c) | c <- [1..cmax], a <- [1..cmax-1], b <- [1..a], (a^2) + (b^2) == (c^2)]
