# Brainfuck

- [What's this](#whats-this)
- [Context](#context)
- [Usage](#usage)
- [Structure](#structure)
- [Sample programs](#sample-programs)

## What's this?

A program to execute [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) code.

Invented by Urban Müller in 1993, Brainfuck is an [esoteric](https://en.wikipedia.org/wiki/Esoteric_programming_language) programming language: a humorously inconvenient language, consisting of just eight single-character commands, that is nevertheless [Turing complete](https://en.wikipedia.org/wiki/Turing_completeness), and hence theoretically capable of performing any computation that can be done by a more conventional programming language.

## Context

The core of this program is an exercise from my early days at 01Founders, included here for fun, along with some extra error handling, tests, and a timeout option.

## Usage

Assuming you have [Go](https://go.dev/) installed, open a terminal and enter the following three commands. (The `$` here represents the terminal prompt rather than part of what you type. Your terminal may vary.)

```bash
$ git clone https://github.com/pjtunstall/brainfuck
$ cd brainfuck
$ go run . ">++++++++[<++++++++>-]+++[<+.>-]"
Output: ABC
```

... or Brainfuck instructions of your choice enclosed in double quotes. The third command above will build and run the program in one step. Alternatively, you can compile the source code just once in advance with `go build`, then run the resulting executable file any number of times with `./brainfuck` on macOS or Linux, or `brainfuck.exe` on Windows, followed by your Brainfuck instructions.

Optionally you can add a positive number after the instructions, representing a timeout in seconds:

```bash
$ ./brainfuck "+[]" 1
Error: timed out after 1 seconds
```

You can also specify a timeout after the `go run .` syntax, but, in that case, there will be a short delay while the program compiles. The seconds are only counted from when the program starts running, not from when it starts compiling.

To run all tests, `go test ./...`. In case you want to run them again without changing the code, e.g. to repeat the fuzz test, enter `go test -count=1 ./...` to make sure they really do run again rather than relying on cached results from earlier trials.

See [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck#Language_design) for a guide to the eight commands of the Brainfuck instruction set.

## Structure

```
brainfuck/
├── brainfuck.go (main)
├── parse_args.go
└── bf/
    ├── wrapped_interpret.go
    ├── interpret.go
    ├── open_bracket.go
    ├── close_bracket.go
    └── out_of_range_error.go
```

The `main` function in `brainfuck.go` calls `parge_args`, then `bf.WrappedInterpret`.

`bf.WrappedInterpret` handles the timeout logic, in case the timeout logic is selected. It passes a `stopInterpretingChan` channel to `bf.Interpret.go`. When a timeout signal is sent, it's received in the interpreting loop and `bf.Interpret.go` returns an empty result and a timeout error. `bf.WrappedInterpret` uses a second channel, `stopWaitingChan`, to make sure that these values are received before returning them to `main`.

Two helper functions for `bf.Interpret.go` (`open_bracket` and `close_bracket`) take care of brackets, which, in Brainfuck, cause the intruction pointer to jump according to certain conditions. These helpers return the new position of the instruction pointer or, if necessary, a `bf.OutOfRangeError`, letting the user know if the instruction pointer went beyond the last instruction or before the first, its last position, and the likely reason: a missing open or close bracket.

## Sample programs

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

An infinite loop:

```
"+[]"
```
