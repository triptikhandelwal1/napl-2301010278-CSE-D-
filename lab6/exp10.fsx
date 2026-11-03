
open System

let isPrime n =
    let rec check divisor =
        if divisor * divisor > n then
            true
        elif n % divisor = 0 then
            false
        else
            check (divisor + 1)

    if n < 2 then false
    else check 2

printf "Enter an integer: "
let number = int (Console.ReadLine())

if isPrime number then
    printfn "%d is a Prime Number" number
else
    printfn "%d is not a Prime Number" number
