
#r "nuget: Microsoft.Data.Sqlite, 10.0.14"

open Microsoft.Data.Sqlite

let connection = new SqliteConnection("Data Source=students.db")
connection.Open()

let createTable = connection.CreateCommand()
createTable.CommandText <- "DROP TABLE IF EXISTS Students; CREATE TABLE Students (Id INTEGER PRIMARY KEY, Name TEXT, Age INTEGER)"
createTable.ExecuteNonQuery() |> ignore

let insert = connection.CreateCommand()
insert.CommandText <- "INSERT INTO Students (Id, Name, Age) VALUES (1, 'Rahul', 20)"
insert.ExecuteNonQuery() |> ignore

let showStudents title =
    printfn "\n%s" title
    use command = connection.CreateCommand()
    command.CommandText <- "SELECT * FROM Students"
    use reader = command.ExecuteReader()

    while reader.Read() do
        printfn "ID: %d, Name: %s, Age: %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

showStudents "After CREATE:"

showStudents "After READ:"

let update = connection.CreateCommand()
update.CommandText <- "UPDATE Students SET Age = 21 WHERE Id = 1"
update.ExecuteNonQuery() |> ignore

showStudents "After UPDATE:"

let delete = connection.CreateCommand()
delete.CommandText <- "DELETE FROM Students WHERE Id = 1"
delete.ExecuteNonQuery() |> ignore

showStudents "After DELETE:"

connection.Close()
