fun greet(name: String = "Student", message: String = "Welcome") {
    println("$message, $name!")
}

fun displayDetails(name: String, age: Int, city: String) {
    println("Name: $name")
    println("Age: $age")
    println("City: $city")
}

fun printNumbers(vararg numbers: Int) {
    println("Numbers:")
    for (num in numbers) {
        println(num)
    }
}

fun calculate(numbers: List<Int>, operation: (Int) -> Int) {
    for (num in numbers) {
        println(operation(num))
    }
}

fun main() {
    println("Default Arguments:")
    greet()
    greet("Tripti", "Hello")

    println("\nNamed Arguments:")
    displayDetails(city = "Faridabad", age = 20, name = "Tripti")

    println("\nVararg Function:")
    printNumbers(10, 20, 30, 40, 50)

    println("\nHigher Order Function:")
    val numbers = listOf(1, 2, 3, 4, 5)
    calculate(numbers) { it * it }
}