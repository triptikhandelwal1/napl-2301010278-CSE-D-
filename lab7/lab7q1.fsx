
open System

let task1 = async {
    printfn "Task 1 started"
    do! Async.Sleep(3000)
    printfn "Task 1 completed"
    return "Result from Task 1"
}

let task2 = async {
    printfn "Task 2 started"
    do! Async.Sleep(1000)
    printfn "Task 2 completed"
    return "Result from Task 2"
}

let task3 = async {
    printfn "Task 3 started"
    do! Async.Sleep(2000)
    printfn "Task 3 completed"
    return "Result from Task 3"
}

printfn "Starting parallel execution...\n"

let results =
    [| task1; task2; task3 |]
    |> Async.Parallel
    |> Async.RunSynchronously

printfn "\nAll tasks completed!"
printfn "Collecting results:\n"

results
|> Array.iteri (fun i result ->
    printfn "Task %d: %s" (i + 1) result)

printfn "\nProgram finished successfully."
