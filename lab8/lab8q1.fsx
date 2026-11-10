type Shape() =
    abstract member Area : unit -> float
    default _.Area() = 0.0

type Circle(radius: float) =
    inherit Shape()
    override _.Area() = System.Math.PI * radius * radius

type Rectangle(length: float, width: float) =
    inherit Shape()
    override _.Area() = length * width

let circle = Circle(5.0)
let rectangle = Rectangle(4.0, 6.0)

printfn "Circle Area: %.2f" (circle.Area())
printfn "Rectangle Area: %.2f" (rectangle.Area())

let shapes: Shape list = [
    circle :> Shape
    rectangle :> Shape
]

printfn "\nPolymorphism:"

for shape in shapes do
    printfn "Area: %.2f" (shape.Area())