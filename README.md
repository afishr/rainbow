# 🌈 Rainbow 

Rainbow Programming Language

## WIP

Currently only lexer and parser are implemented. In the near future here can appear the intepreter 👀

## Syntax example

```
var fib = fn(x) {
  if (!x)
    return 0;

  if (x < 3)
    return 1;

  return fib(x-1) + fib(x-2);
}

var result = fib(10);
print("The 10th number in Fibonacci sequence is " + result);
```

## REPL

You can play with rainbow by runing `$ go run main.go` and see how your entered statements are parsed in real time

## Credits

[Thorsten Ball - Writing An Interpreter In Go](https://interpreterbook.com/)

[Pratt Parsers: Expression Parsing Made Easy](http://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)

```
  _ __ __ _(_)_ __ | |__   _____      __
 | '__/ _` | | '_ \| '_ \ / _ \ \ /\ / /
 | | | (_| | | | | | |_) | (_) \ V  V / 
 |_|  \__,_|_|_| |_|_.__/ \___/ \_/\_/ 
```
