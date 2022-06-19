# GoExercises
Go exercises for different levels

Any help is welcome.

## How it works?

Clone the repository,

Go to any exercise, 

Implement functions (do not change test file), 

Run any test you want, by the help of this example commands:

`go test go_exercises/fizzbuzz`

`go test go_exercises/even_odd`

`go test go_exercises/sumdiv`

`go test go_exercises/primesq`

`go test go_exercises/passgen`

In case of Fibonacci and Factorial, etc, you may not write all the functions, it's ok, as long as you run these commands instead of commands above.

`go test -timeout 30s -run ^TestFunctionName$ go_exercises/fibonacci -v`

for example:

`go test -timeout 30s -run ^TestFibonacciLoop$ go_exercises/fibonacci -v`
