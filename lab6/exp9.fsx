
open System

printfn "Enter marks for four subjects:"

let marks =
    [ for i in 1..4 do
        printf "Enter marks for Subject %d: " i
        let m = float (Console.ReadLine())
        yield m ]

let total = List.sum marks
let percentage = total / 4.0

printfn "\nTotal Marks = %.2f" total
printfn "Percentage = %.2f%%" percentage

let division =
    match percentage with
    | p when p >= 60.0 -> "First Division"
    | p when p >= 50.0 -> "Second Division"
    | p when p >= 40.0 -> "Third Division"
    | _ -> "Fail"

printfn "Division = %s" division
