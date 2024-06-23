// Program to print numbers for natural numbers 1 + 2 + 3 + ... +n

package main
import "fmt"

func main() {
  var n, sum = 10, 0
  
  for i := 1 ; i <= n; i++ {
    sum += i    // sum = sum + i  
  }

  fmt.Println("sum =", sum)
}

// for i := 1; i <= n; i++: This loop will iterate from i = 1 to i = n (which is 10).
// Iteration 1: i = 1, sum = 0 + 1 = 1
// Iteration 2: i = 2, sum = 1 + 2 = 3
// Iteration 3: i = 3, sum = 3 + 3 = 6
// Iteration 4: i = 4, sum = 6 + 4 = 10
// Iteration 5: i = 5, sum = 10 + 5 = 15
// Iteration 6: i = 6, sum = 15 + 6 = 21
// Iteration 7: i = 7, sum = 21 + 7 = 28
// Iteration 8: i = 8, sum = 28 + 8 = 36
// Iteration 9: i = 9, sum = 36 + 9 = 45
// Iteration 10: i = 10, sum = 45 + 10 = 55