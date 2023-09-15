## Fibonacci numbers

A 'hello world' for Golang which introduces various features:

- the standard library _fmt_ and getting data from stdin
- make a _slice_ (that is, a variable length array, with size set at run time)
- function calls (with functions organised in their own file for ease of readability)
- _for_ loops
- string formatting 
- recursion

It also demonstrates pitfalls in algorithm design. The example compares two implementations of computing Fibonacci: the first (efficient) implementation simply stores the successive values in an array for use with the next two values. (This is a very simple example of a dynamic program.) The second (slow!) implementation is recursive, which looks 'cool' in code but is hopelessly naive, requiring 2^n function calls for the n-th value.

(Including both implementations was motivated by a recent Medium post about Rust - extolling the speed of Rust by demonstrating it on the naive Fibonacci algorithm!)

Run with:

 go mod init example.com/fibonacci
 go build .
 ./fibonacci

Try values in the range 40-50.