det a b c = b^2 - 4*a*c
quadsol1 a b c = (-b - sqrt (det a b c))/2*a
quadsol2 a b c = (-b + sqrt (det a b c))/2*a

--exercise1--
--index operator
third_a (a) = a!!2
--pattern matching
third_b (a:b:c:d) = c
--hailstone
hailstone a
   | even a       = a `div` 2
   | otherwise    = (a*3)+1


