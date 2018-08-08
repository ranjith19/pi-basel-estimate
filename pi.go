package main

import "fmt"
import "math/rand"
import "math"


const N_TRIALS = 100000
const N_BOUND = 10000000

func gcd(x int, y int) int{
    a := x
    b := y
    if b > a {
        b, a = x, y
    }
    for {
        if a == b {
            return a
        } else {
            x = a - b
            y = b

            if x > y {
                a, b = x, y
            } else {
                b, a = x, y
            }
        }

    }
}


func get_1_if_coprime(x int, y int) int {
    if gcd(x, y) == 1 {
        return 1
    } else{
        return 0
    }
}

func get_int() int {
    return 1 + rand.Intn(N_BOUND)
}



func calculate() float64{
    var coprimes int = 0
    var i int = 0
    var x int
    var y int
    var fraction_co_prime float64
    var pi_estimate float64
    for i <= N_TRIALS {
        i = i + 1
        x, y = get_int(), get_int()
        coprimes += get_1_if_coprime(x, y)
        if coprimes > 0{
            fraction_co_prime = float64(coprimes) / float64(i)
            pi_estimate = math.Sqrt(6 / fraction_co_prime)
        }
    }
    return pi_estimate

}


func main() {
    i := 0
    for i < 10 {
        i = i + 1
        fmt.Println(calculate())
    }
}
