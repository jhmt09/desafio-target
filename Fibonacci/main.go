package main

import "fmt"

func fibonacci(n int) []int {
    fib := []int{0, 1}

    for i := 2; i < n; i++ {
        fib = append(fib, fib[i-1]+fib[i-2])
    }

    return fib
}

func contains(slice []int, num int) bool {
    for _, n := range slice {
        if n == num {
            return true
        }
    }
    return false
}

func main() {
    var n int
    fmt.Print("Digite um número: ")
    fmt.Scan(&n)

    fib := fibonacci(n)

    if contains(fib, n) {
        fmt.Printf("%d pertence à sequência de Fibonacci %v\n", n, fib)
    } else {
        fmt.Printf("%d não pertence à sequência de Fibonacci %v\n", n, fib)
    }
}
