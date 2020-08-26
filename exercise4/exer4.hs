-- Hailstone, Again
--Support Functions
--hailstone
hailstone :: Int -> Int
hailstone a
   | a == 0            = 0
   | even a            = a `div` 2
   | otherwise         = (a*3)+1

--hailSeq
hailSeq :: Int -> [Int]
hailSeq n = hailList [] n
   where
      hailList seq 0 = []
      hailList seq 1 = seq ++ [1]
      hailList seq n = (hailList (seq ++ [n]) (hailstone n))

--hailSeq'
hailSeq' :: Int -> [Int]
hailSeq' 0 = []
hailSeq' n = (takeWhile (>1) (iterate hailstone n)) ++ [1]

-- Joining Strings, Again
--join
join :: String -> [String] -> String
join str [] = []
join str list = (foldl (\x y -> x ++ y ++ str) [] (init list)) ++ last list

-- Merge Sort
--Support Functions
--merge
merge :: Ord a => [a] -> [a] -> [a]
merge [] [] = []
merge (next1:tail1) [] = next1 : tail1
merge [] (next2:tail2) = next2 : tail2
merge (next1:tail1) (next2:tail2)
   | next1 < next2         = next1 : (merge tail1 (next2:tail2))
   | otherwise             = next2 : (merge (next1:tail1) tail2)

--mergeSort
mergeSort :: Ord a => [a] -> [a]
mergeSort [] = []
mergeSort [a] = [a]
mergeSort list = merge (mergeSort firstHalf) (mergeSort lastHalf)
   where 
      firstHalf = take ((length list) `div` 2) list
      lastHalf = drop ((length list) `div` 2) list

-- Searching? Maybe?
--Support Functions
getIndex index target list
   | index >= length list             = index
   | (list !! index) == target        = index
   | otherwise                        = getIndex (index+1) target list 

--findElt
findElt :: Eq a => a -> [a] -> Maybe Int
findElt target list
   | index >= length list         = Nothing
   | otherwise                    = Just index
   where
      index = getIndex 0 target list