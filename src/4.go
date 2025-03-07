 
 // Function to compute the factorial of a number
 func Factorial(n int) int {
    if n == 0 {
        return 1
    } else {
        return n * Factorial(n-1)
    }
}

fmt.Println("The factorial of 5 is:", Factorial(5)) // prints "The factorial of 5 is: 120"
