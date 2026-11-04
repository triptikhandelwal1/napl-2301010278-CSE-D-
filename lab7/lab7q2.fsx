
let list1 = [1; 2; 3; 4; 5]
let list2 = [2; 4; 6; 8]

printfn "----- LIST.MAP (SQUARE) -----"

let square1 = List.map (fun x -> x * x) list1
let square2 = List.map (fun x -> x * x) list2

printfn "List 1: %A -> %A" list1 square1
printfn "List 2: %A -> %A" list2 square2

printfn "\n----- LIST.FILTER (EVEN NUMBERS) -----"

let even1 = List.filter (fun x -> x % 2 = 0) list1
let even2 = List.filter (fun x -> x % 2 = 0) list2

printfn "List 1: %A -> %A" list1 even1
printfn "List 2: %A -> %A" list2 even2

printfn "\n----- LIST.FOLD (SUM) -----"

let sum1 = List.fold (fun acc x -> acc + x) 0 list1
let sum2 = List.fold (fun acc x -> acc + x) 0 list2

printfn "Sum of List 1: %d" sum1
printfn "Sum of List 2: %d" sum2

printfn "\nProgram completed successfully."
