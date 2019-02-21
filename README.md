# Go Playground

A lot of the more basic experiments are me following along with [Go by Example](https://gobyexample.com).

More practical examples:

- `append.go`: testing the `flags` package and file IO with `io/ioutil`, a CLI that appends a new line to a given file
- `env.go`: prints out environments variables with some nice formatting, including expanding out each path in the `$PATH` environment variable in its own line, so it's easier to read than `env`'s output'
- `picalc.go`: Monte-Carlo approximation of Pi demonstrating Go's PRNG and parallelism through goroutines
- `seer.go`: A simple CLI to search through large CSV files for substring matches and display the matching rows readably
