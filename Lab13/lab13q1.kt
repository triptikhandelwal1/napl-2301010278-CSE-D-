fun main() {

    println("Hello Kotlin!!!")

    print("Enter department code: ")
    val deptCode = readLine()!!.uppercase()

    print("Enter years of experience: ")
    val experience = readLine()!!.toInt()

    print("Enter employee salary: ")
    val salary = readLine()!!.toDouble()

    val bonusPercentage = when {
        deptCode == "IT" && experience > 5 -> 20
        deptCode == "HR" && experience > 3 -> 15
        deptCode == "FINANCE" && experience > 4 -> 18
        else -> 10
    }

    val bonus = salary * bonusPercentage / 100

    println("Bonus Percentage: $bonusPercentage%")
    println("Employee Bonus: $bonus")

        print("Enter an integer: ")
    val number = readLine()!!.toInt()

    var n = kotlin.math.abs(number)
    var sum = 0
    var count = 0

    if (n == 0) {
        count = 1
    }

    while (n > 0) {
        val digit = n % 10
        sum += digit
        count++
        n /= 10
    }

    val average = sum.toDouble() / count

    println("Average of digits: $average")

}