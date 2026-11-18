
type Student = {
    id: int
    name: string
    marks: int
}

let students = [
    { id = 1; name = "Rahul"; marks = 85 }
    { id = 2; name = "Priya"; marks = 72 }
    { id = 3; name = "Aman"; marks = 32 }
]

let displayStudents students =
    students |> List.iter (fun s ->
        printfn "ID: %d Name: %s Marks: %d" s.id s.name s.marks
    )

let getNames students =
    students |> List.map (fun s -> s.name)

let passedStudents students =
    students |> List.filter (fun s -> s.marks >= 40)

let getGrade marks =
    match marks with
    | m when m >= 80 -> "Distinction"
    | m when m >= 60 -> "First Division"
    | m when m >= 40 -> "Pass"
    | _ -> "Fail"

let addStudent student students =
    students @ [student]

let updateStudent id newMarks students =
    students |> List.map (fun s ->
        if s.id = id then
            { s with marks = newMarks }
        else
            s
    )

let deleteStudent id students =
    students |> List.filter (fun s -> s.id <> id)

[<EntryPoint>]
let main argv =
    printfn "STUDENT MANAGEMENT SYSTEM"

    printfn "\nAll Students:"
    displayStudents students

    printfn "\nStudent Names:"
    getNames students |> List.iter (printfn "%s")

    printfn "\nPassed Students:"
    displayStudents (passedStudents students)

    printfn "\nStudent Grades:"
    students |> List.iter (fun s ->
        printfn "%s: %s" s.name (getGrade s.marks)
    )

    let newStudent = { id = 4; name = "Neha"; marks = 90 }

    let list1 = addStudent newStudent students
    printfn "\nAfter Adding:"
    displayStudents list1

    let list2 = updateStudent 2 88 list1
    printfn "\nAfter Updating:"
    displayStudents list2

    let list3 = deleteStudent 3 list2
    printfn "\nAfter Deleting:"
    displayStudents list3

    0
