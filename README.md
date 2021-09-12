<h1 style="text-align: center">Funked: A Functional Library for Go</h1>

![Build](https://github.com/jwhittle933/funked/actions/workflows/ci.yml/badge.svg)
<a href="https://pkg.go.dev/github.com/jwhittle933/funked"><img src="https://pkg.go.dev/badge/github.com/jwhittle933/funked.svg" alt="Go Reference"></a>

Funked, as the name implies, is a functional library for Go. But `functional` has multiple meanings in this context. On the one hand, "functional" refers to functional programming techniques, principles such as immutability, referential transparency (pure functions), and function composition. On the other hand, "functional" refers to that more abstract quality of code that indicates that it works, is useful, and make certain aspects of writing it more enjoyable or less mentally taxing. This library aims at both, and has taken inspiration from a number of other languages, most notably Rust and Elixir.

## Libraries
Funked is really just a wrapper for smaller libraries. You can get the whole thing with `go get github.com/jwhittle933/funked/...` or just one, with `go get github.com/jwhittle933/funked/async/process`, or even a collection with `go get github.com/jwhittle933/funked/slices/...`.

### Slices
The Go has been alive with talks of generics for a number of years, and course seems to have been plotted for inclusion. Until that time, writing `Filter`, `Map`, and other functional collection helpers is a very manual process. The `slices` package helps to alleviate that for some primitive types, namely `int`, `string`, and `byte` (`uint8`). See the [README](./slices/README.md) for `slices` for an overview of using the API.

## Async
The `async` package is designed for starting "managed" go routines, async processes that provide a way to monitor activity and communicate. 