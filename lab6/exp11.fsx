
open System
open System.Text.RegularExpressions

// 1. Reverse a string and check palindrome
let reverseString (s: string) =
    s |> Seq.rev |> Seq.toArray |> String

let checkPalindrome (s: string) =
    let reversed = reverseString s
    reversed = s

// 2. Transform string using pattern matching
let transformString (s: string) =
    match s with
    | x when Regex.IsMatch(x, @"^[a-zA-Z]+$") ->
        x.ToUpper()
    | x when Regex.IsMatch(x, @"^[0-9]+$") ->
        let number = int x
        string (number * 2)
    | _ -> s

// 3. Validate email using Regex
let isValidEmail (email: string) =
    Regex.IsMatch(email, @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$")

// Testing palindrome
printfn "----- PALINDROME CHECK -----"
["madam"; "hello"; "level"]
|> List.iter (fun s ->
    printfn "%s -> Reverse: %s, Palindrome: %b" s (reverseString s) (checkPalindrome s))

// Testing string transformation
printfn "\n----- STRING TRANSFORMATION -----"
["hello"; "123"; "Hello123"]
|> List.iter (fun s ->
    printfn "%s -> %s" s (transformString s))

// Testing email validation
printfn "\n----- EMAIL VALIDATION -----"
["abc@gmail.com"; "test123@yahoo.in"; "invalid-email"]
|> List.iter (fun email ->
    printfn "%s -> %b" email (isValidEmail email))
