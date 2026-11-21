class Student(
    private var rollNo: Int,
    private var name: String,
    private var course: String,
    private var percentage: Double
) {

    public fun acceptData() {
        print("Enter Roll Number: ")
        rollNo = readLine()!!.toInt()

        print("Enter Name: ")
        name = readLine()!!

        print("Enter Course: ")
        course = readLine()!!

        print("Enter Percentage: ")
        percentage = readLine()!!.toDouble()
    }

    public fun displayDetails() {
        println("\nStudent Details")
        println("----------------------")
        println("Roll Number: $rollNo")
        println("Name: $name")
        println("Course: $course")
        println("Percentage: $percentage%")
    }
}

fun main() {
    val student1 = Student(0, "", "", 0.0)
    val student2 = Student(0, "", "", 0.0)

    println("Enter details for Student 1")
    student1.acceptData()
    student1.displayDetails()

    println("\nEnter details for Student 2")
    student2.acceptData()
    student2.displayDetails()
}