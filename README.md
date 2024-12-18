# Brainfuck

## What's this?

A program to execute [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) code.

Brainfuck is an [esoteric](https://en.wikipedia.org/wiki/Esoteric_programming_language) programming language: a humorously inconvenient language, consisting of just eight symbols, that is nevertheless real, i.e. [Turing complete](https://en.wikipedia.org/wiki/Turing_completeness), and theoretically capable of performing any computation that can be done by a more conventional programming language.

## Context

This is an exercise from my early days of learning Go, included here for fun, along with some extra tests and error handling.

Brainfuck captured by imagination when I first encountered it in the final exam of our 4-week intensive intro to Go. With 6 minutes to spare, I didn't manage to write an interpreter during the exam, but the concept made me laugh, so I couldn't resist figuring it out afterwards.

# Usage

Assuming you have [Go](https://go.dev/) installed, clone this repo and navigate into it, then enter `go run . "Insert Brainfuck instructions here!"` in your terminal to build and run in one step. Optionally follow this with a positive number, representing a timeout in seconds: `go run . "Insert Brainfuck instructions here!" 10`.

Alternatively, compile the source code once and for all with `go build`, then run the resulting executable file with `./brainfuck "Insert Brainfuck instructions here!"` on macOS or Linux, or `brainfuck.exe "Insert Brainfuck instructions here!"` on Windows.

See [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck#Language_design) for a guide to the eight commands in the Brainfuck instruction set.

## Example programs

```
Hello World!
"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

AAA
"+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.>-]"

ABC
"+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.+>-]"

A
"+>++++++++[<++++++++>-]<."
```

In the above example, the 2nd register acts as a counter, allowing us to run the loop 8 times. Inside the loop, we increment the 1st register by 8. When the counter reaches 0, the loop ends and we move the back to the print the value in the first register. ASCII 'A' is 1 + 8 \* 8 = 65. Notice that 8 + 8 = 16. If we want to express 64 as a different product, we'd have to double the number of pluses in one place for every halving in the other: 16 + 4 = 20, 32 + 2 = 34. In general, the closer we can get both factors to the square root of the target quantity the better. (For rectangles of a given area, the square is the rectangle with the smallest perimeter.)

```
AAA
"+>++++++++[<++++++++>-]+++[<.>-]"

ABC
">++++++++[<++++++++>-]+++[<+.>-]"

123456789
">+++++++[<+++++++>-]+++++++++[<.+>-]"

0123456789
">+++++++<->[<+++++++>-]++++++++++[<.+>-]"

I LOVE YOU
"++++++++[>+++++++++>++++++++>++++<<<-]>+.>>.<<+++.+++.+++++++.>+++++.>.<<+++.----------.++++++."
```
