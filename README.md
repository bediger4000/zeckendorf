# Zeckendorf Representation


I read a blog post titled
[The golden ratio as a number base](https://apieceofthepi.substack.com/p/the-golden-ratio-as-a-number-base).
It had an interesting statement:

> According to Zeckendorf's Theorem, every positive integer can be
> represented in a unique way as a sum of distinct, non-consecutive Fibonacci
> numbers.

The Wikipedia page for [Zeckendorf's Theorem](https://en.wikipedia.org/wiki/Zeckendorf's_theorem)
has a slightly tighter wording for that statement.

## Build and run

```
$ go build z1.go
$ ./z1 7734
6765, 610, 233, 89, 34, 3
sum 7734
```

## Fibonacci Encoding

The Wikipedia page for Zeckendorf's Theorem mentions that
Zeckendorf Representation is closely related to
[Fibonacci Encoding](https://en.wikipedia.org/wiki/Fibonacci_coding)

> The Fibonacci code word for a particular integer is exactly the integer's
> Zeckendorf representation with the order of its digits reversed and an
> additional "1" appended to the end. 

## Fibonacci Numbers

To do the Fibonacci Coding, I used a slice literal.
To find values for the literal,
I used this iterative method of calculating Fibonacci numbers.

```
 1	func fibonacciNumbers() {
 2	    var x uint64
 3	    var y uint64
 4	    x, y = 1, 1
 5	    fmt.Printf("1,\n1,\n")
 6	
 7	    for {
 8	        nxt := x + y
 9	        if nxt < y {
10	            // overflow
11	            break
12	        }
13	        fmt.Printf("%d,\n", nxt)
14	        x, y = y, nxt
15	    }
16	}
```

For my first attempt, variables `x` and `y` had type `int`.
The function printed 92 numbers, beginning with 1, 1, ...
and ending with 4660046610375530309, 7540113804746346429

I took advantage of my computer's (x86_64) integer arithmetic
to end the loop of lines 7-15.
Line 8 has an integer overflow. Variable `nxt` can and does end up negative.

When I changed the type of `x` and `y` on lines 2 and 3 from `int` to `int64`,
I got one more Fibonacci number out of the loop, 
12200160415121876738, the sum of 4660046610375530309 and 7540113804746346429
I did not change the loop end condition on line 19.
For `uint64`, the overflow is to a smaller _positive_ value.
I thought this was interesting.

The next Fibonacci Number after 12200160415121876738 is 19740274219868223167,
which users of Linux' `bc` arbitrary precision calculator can easily find.
That Fibonacci Number is greater than 2<sup>64</sup> or 18446744073709551615
