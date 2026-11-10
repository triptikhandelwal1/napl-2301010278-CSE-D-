
#r "nuget: MongoDB.Driver"

open MongoDB.Bson
open MongoDB.Driver

let client = MongoClient("mongodb://localhost:27017")
let database = client.GetDatabase("Lab8DB")
let collection = database.GetCollection<BsonDocument>("Students")

let filter =
    Builders<BsonDocument>.Filter.Eq<string>("_id", "lab8_student_1")

let student = BsonDocument()
student.Add("_id", BsonString("lab8_student_1"))
student.Add("name", BsonString("Rahul"))
student.Add("age", BsonInt32(20))

collection.DeleteOne(filter) |> ignore

collection.InsertOne(student)

printfn "After CREATE:"
let created = collection.Find(filter).FirstOrDefault()
printfn "%O" created

printfn "\nAfter READ:"
let read = collection.Find(filter).FirstOrDefault()
printfn "%O" read

let update =
    Builders<BsonDocument>.Update.Set("age", BsonInt32(21))

collection.UpdateOne(filter, update) |> ignore

printfn "\nAfter UPDATE:"
let updated = collection.Find(filter).FirstOrDefault()
printfn "%O" updated

collection.DeleteOne(filter) |> ignore

printfn "\nAfter DELETE:"
let deleted = collection.Find(filter).FirstOrDefault()

if isNull deleted then
    printfn "No student found."
else
    printfn "%O" deleted
