import RainbowAssign
import Data.Maybe as Maybe
import qualified Data.Map as Map

--type Hash = Int32 (convert Int32 to Int using fromEnum)
--type Passwd = String

pwLength, nLetters, width, height :: Int
filename :: FilePath
pwLength = 8            -- length of each password
nLetters = 5            -- number of letters to use in passwords: 5 -> a-e
width = 40              -- length of each chain in the table
height = 1000           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table



-- Hashing & Reducing



--Support Functions
--convertDecimal : convert integer base 10 to base 
convertDecimal :: Int -> Int -> [Int]
convertDecimal val 10 = [val]
convertDecimal val 1 = [val]
convertDecimal val base = reverse (cdList [] val base)
   where
      --cdList : returns a list of integers representing val in the desired base
      cdList acc val base
         | val == 0        = acc
         | val == -1       = acc
         | otherwise       = cdList (acc ++ [val `mod` base]) (val `div` base) base
--signBit : identify if Int is negative (return 1) or positive (return 0)
signBit :: Int -> Int
signBit val
   | val > -1         = 0
   | otherwise        = 1
--pad : pads convertedList with 0 if signbit is 0, pads convertedList with last index for letter range
pad :: Int -> Int -> [Int] -> [Int]
pad signbit toFill convertedList
   | signbit > 0         = (take toFill (repeat (nLetters-1))) ++ convertedList
   | otherwise
              = (take toFill (repeat 0)) ++ convertedList

--pwReduce
pwReduce :: Hash -> String
pwReduce hashed = step3
   where
      step1 = convertDecimal (fromEnum hashed) nLetters
      step2
         | (length step1) <= pwLength         = pad (signBit (fromEnum hashed)) (pwLength-length step1) step1
         | otherwise                          = drop ((length step1) - pwLength) step1
      step3 = [toEnum (i+97)::Char | i <- step2]



-- Building the Table



--rainbowTable
rainbowTable :: Int -> [Passwd] -> Map.Map Hash Passwd
rainbowTable width initPasswds = Map.fromList (zip (rainbowWidth width (map pwHash initPasswds)) initPasswds)--[(hash,password)]
   where
      rainbowWidth width [] = []
      rainbowWidth 0 hashes = hashes
      rainbowWidth width hashes = rainbowWidth (width-1) (map pwHash (map pwReduce hashes))



-- Creating, Reading, and Writing Tables



--generateTable
generateTable :: IO ()
generateTable = do
  table <- buildTable rainbowTable nLetters pwLength width height
  writeTable table filename

--test1
test1 = do
  table <- readTable filename
  return (Map.lookup (-1482089486) table)
--Put the expression you want to test in the parens on the last line: calling the test function will give you the result of evaluating that expression.



-- Reversing Hashes



--Support Functions
--extract : return val of Monad. use only when val exists
extract :: Maybe a -> a
extract (Just val) = val
--getPassword : retrieve password from initial password in a row in Rainbow Table
getPassword :: Map.Map Hash Passwd -> Int -> Passwd -> Hash -> Maybe Passwd
getPassword rtable col passw hash
   | col == 0           = verifyPassword passw hash
   | otherwise          = getPassword rtable (col-1) (pwReduce (pwHash passw)) hash
      where
         verifyPassword vp_passw vp_hash
            | (pwHash vp_passw) == vp_hash         = Just vp_passw
            | otherwise                            = Nothing

--findPassword
findPassword :: Map.Map Hash Passwd -> Int -> Hash -> Maybe Passwd
findPassword rtable col hash = confirmPassword rtable col hash hash
   where
      confirmPassword cp_table cp_col cp_hash original
         | row /= Nothing              = getPassword cp_table (cp_col) (extract row) original
         | cp_col == 0                 = Nothing
         | otherwise                   = confirmPassword cp_table (cp_col-1) (pwHash (pwReduce cp_hash)) original
            where
               row = Map.lookup cp_hash cp_table



-- Experimenting



--test2
test2 :: Int -> IO ([Passwd], Int)
test2 n = do
  table <- readTable filename
  pws <- randomPasswords nLetters pwLength n
  let hs = map pwHash pws
  let result = Maybe.mapMaybe (findPassword table width) hs
  return (result, length result)

main :: IO ()
main = do
  generateTable
  res <- test2 10000
  print res