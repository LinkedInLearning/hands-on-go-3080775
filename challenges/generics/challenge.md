# Challenge: Refactoring with Generics

This challenge will help cement what we've learned about generics in this course by having us refactor some non-generic functions into their generic counterparts.

Open up `challenges/generics/begin/main.go` and take a look at the pre-defined print functions `printString`, `printInt`, and `printBool`.

Your first task is to write a generic `printAny` function that can work with any of the types the non-generic versions support.

Next, note the `sum` function and its use of `interface{}` and necessary type switch to sum the numbers it's invoked with. 

Your second task is to create a generic `sumAny` function that uses a custom type constraint you'll define called `numeric`. The `numeric` type constraint need to have a type set of `~int` and `~float64`. The `sumAny` function should be able to sum any number of arguments passed in (i.e. it's variadic) and return a sum of those numeric values.

## Acceptance Criteria
1. You have a `printAny` that you can invoke with a `string`, an `int`, and a `bool`.
2. You have a `sumAny` generic function parameterized with a custom `numeric` type constraint.
3. You can invoke each of the above generic function in your `main` and print their output to console.

## Sample output

Calling your functions like the following

```
printString("Hello")
printInt(42)
printBool(true)
fmt.Println(sumAny(1, 2, 3))
```

...should yield

```
Hello
42
true
6
```

Pause the video here to work on the challenge before we review it. Good luck!