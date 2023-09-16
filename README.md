# goformath

Tutorial code examples to teach myself Go.

IN PROGRESS!!!

| Go concept | Example code |
|------------|--------------|
| _Files_ ||
| code organisation | fibonacci |
| get data from stdin | fibonacci |
| open file or stdio | wordcount |
| bufio scanner | wordcount |
| read data from file using _fmt.Fscanln_ | gaussian |
| read data from file using _bufio_ scanner | tsp |
| io.Writer | trees |
| _Arrays/slices_ ||
| make a slice  with length read from stdin | fibonacci |
| initialising a constant-length array | sieve |
| static arrays (can't be declared const; use of [...]) | pfactor |
| slices and append() | pfactor |
| append slices with ... | gaussian |
| _Loops_ ||
| for loop | fibonacci |
| multi-condition for-loop "for a,b := X,Y; ... ; a,b = P,Q { }" | pfactor |
| infinite loop for {} | pfactor |
| recursion | gcd |
| switch | trees |
| range | trees |
| _Variables_ ||
| constants (like #define in C) | sieve |
| cast int as float64 | pfactor |
| runes vs strings | trees |
| passing function type as argument to another function | gaussian |
| structs and methods | trees |
| _Command line_ ||
| command line options and flag package | gcd |
| command line args | sieve |
| cmd line arguments with flag package | coinflip |
| --help | gcd |
| usage check and os.Exit() | gcd |
| _Formatting_ ||
| string formatting a la C printf() | fibonacci |
| strings package | wordcount |
| strconv package | gcd |
| convert string to int, float64 etc | sieve |
| %v string format for general 'value' | bingcd |
| _General_ ||
| math package | pfactor |
| use of math/rand package | coinflip |
| bit ops | bingcd |
| timing | bingcd |
| parallelisation with goroutines | gaussian |
| use of a buffered channel | gaussian |
| profiling using pprof | gaussian |


