main :: IO ()
main = do
    inputs <- lines <$> readFile "input.txt"
    let nums = map read inputs :: [Int]
    let result = [a*b*c | a <- nums, b <- nums, c <- nums, a+b+c == 2020]
    print result


-- main :: IO ()
-- main = do
--     handle <- openFile "input.txt" ReadMode
--     contents <- hGetContents handle
--     putStr contents
--     ans <- findit contents 0 1
--     putStr ans
--     hClose handle

-- findit contents j k = do
--     res <- contents[j] + contents[k]
--     if res == 2020 then
--         res
--     else if k < length contents then
--         findit contents j k+1
--     else if j + 2 < length contents then
--         findit contents (j+1) (j+2)
--     else
--         0