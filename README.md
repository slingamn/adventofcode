# Advent of Code Solutions in Golang

First of all, don't do [Advent of Code](https://adventofcode.com). It's a waste of time. Definitely don't do it in real time.

That said, if you do Advent of Code for whatever reason (peer pressure in my case), don't do it in Go. Take [Kevin Yap's advice](https://kevinyap.ca/2019/12/going-fast-in-advent-of-code/) and do it in Python.

That said, if you love Go, Go is not the worst possible choice for puzzling; the language's limitations will rarely impose a decisive barrier. These are my solutions in Go. In 2020, I imperfectly followed the following conventions:

1. `2020/22/a.go` and `2020/22/b.go` (for example) are solutions for 2020 day 22, parts 1 and 2 respectively. They are only minimally cleaned up from the time of the solution.
2. Sometimes there's a more elaborately refactored or optimized solution as `b_final.go` (even more rarely, as `a_final.go` if parts 1 and 2 were sufficiently different).

All solutions are stdlib-only, but require a 64-bit architecture (more on this below).

## Some notes on puzzling in Go

Here are some observations on doing puzzles in Go that may or may not be helpful.

### Why use Go?

As mentioned previously, don't use Go. But if you are using Go, here are some perks:

* The compiler will significantly protect you from yourself
* Your unoptimized solutions will typically be an order of magnitude faster than unoptimized dynamic language solutions
* You can do often do low-level performance hacks to squeeze out an additional order of magnitude
* Very occasionally, this may give you the elbow room to get away with shenanigans
* Somewhat less occasionally, the problems will give you the chance to micro-benchmark different Go implementation techniques

I am, however, skeptical that puzzling in Go provides meaningful practice at rapidly writing production-grade Go. Puzzling (with or without the competitive aspect) fosters a variety of bad habits.

### Stubs

`stub.go` is my problem stub. It contains accumulated library code, together with a basic problem harness (it reads newline-delimited input from stdin into a `[]string`, then calls a stub `func solve(input []string) (result int)`). I started each day in 2020 from a copy of the stub, adding to it as I went. Then after solving the problem, I would delete unused stub code from my solutions.

I find it appealing to produce stdlib-only solutions, and the Go compiler is fast enough that this won't slow you down significantly (on my system, the stub can currently be compiled in about 200 milliseconds).

### Error handling

Production Go will not typically use `panic()` for error handling. However, I strongly recommend it for puzzles; it speeds up debugging considerably. You can also use `panic()` where you would `assert` in other languages, as a way of checking your work.

If so desired, these panics can easily be cleaned up (converted to error values, or eliminated) in a subsequent refactoring step.

### Portability

Many AoC problems use integers larger than 32 bits (one infamously induces overflows on 64-bit integers as well). Go's [int type](https://golang.org/ref/spec#Numeric_types) is word-sized, i.e., 32 bits on a 32-bit architecture and 64 bits on a 64-bit architecture. For full portability, you'd have to explicitly use `int64` and `uint64` types in your code, but this will slow you down --- particularly because the `int` type occurs "naturally" in Go as the type of `len()` values, or as the index type when enumerating a slice with `for index, value := range mySlice`, and in a puzzling context it's natural to do arithmetic operations on these values together with your data.

I recommend using `int` everywhere and assuming a 64-bit architecture. If you want to go back and convert to portable code, it should be fairly straightforward: start by parsing the input into `int64` instead of `int`, then fix the resulting compiler errors.

### Data structures

* Slices make very efficient queues and/or deques.  Watch out for the `append` gotcha: `append` may reallocate the array, so you can't use slices as drop-in replacements for Python lists. If you pass a slice to a function you should return the slice by value to ensure that your changes correctly propagate back into the caller.
* Maps are useful despite their disappointing performance. Unfortunately, the only variable-length data type that can be a map key in Go is `string`; there's no equivalent to Python's variable-length tuples. You can use deliberately oversized fixed-size arrays as an alternative (I once used `[2][50]int`), or build a string.
* [container/heap](https://golang.org/pkg/container/heap/) should make an effective priority queue, although I've never seen a problem that required it
* Anything too fancy is probably a wrong turn (the problems are designed not to require obscure algorithms or data structures)

## Interesting solutions

Most of my solutions are pretty dull, but there are a few that may be of interest:

<details>
  <summary>here be spoilers!</summary>

  * `2019/18/a.go`: 2019 day 18 ("Many-Worlds Interpretation") part 1 using some tricks with embedded structs
  * `2019/22/b.go`: 2019 day 22 ("Slam Shuffle") part 2, in portable code that avoids integer overflow
  * `2020/17/b_final.go`: 2020 day 17 ("Conway Cubes") parts 1 and 2, fully parameterized by dimension (no nested for loops), and with minimal dynamic allocations (using a stack-allocated iterator type)
  * `2020/18/a_final.go` and `b_final.go`: 2020 day 18 ("Operation Order"), O(n) with no dynamic allocations
  * `2020/22/b_final.go`: 2020 day 22 ("Crab Combat") part 2, using a trick with `strings.Builder` to efficiently hash slices
  * `2020/23/b.go`: 2020 day 23 ("Crab Cups") part 2 without pointers
</details>
