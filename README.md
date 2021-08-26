<h1 style="text-align: center">Funked: A Functional Library for Go</h1>

![Build](https://github.com/jwhittle933/funked/actions/workflows/ci.yml/badge.svg)
<a href="https://pkg.go.dev/github.com/jwhittle933/funked"><img src="https://pkg.go.dev/badge/github.com/jwhittle933/funked.svg" alt="Go Reference"></a>

Each package in `Funked` observes the same or very similar API. The packages here are designed to create "handles" on
common slice types for easy use and composition. Nothing here is revolutionary: its goal is ease of use and convenience. 

Why create a package for different slice types? Others have tried to create functional libraries using runtime reflection, but the extreme performance hit and use of `interface{}` in every function signature is simply not ideal for real applications. Take, for example, a large CLI application. Runtime reflection would slow down argument and flag parsing by up to a second (depending on the number of args and flags) and coercing `interface{}` to concrete types fails the most basic test of clear and idiomatic Go (and anyone who's written Go for an extensive amount of time knows that this is only to be done sparingly). Go 2 has a spec for generics. Until then, we have to implement this type of functionality by hand for every type that needs it.

## APIs
For examples of the APIs presented below, see the [examples](./examples) directory.

### Bare API
The libraries here share a common API, but there are several ways to use it. The first is the "bare" functions:

```go
intslice.Map([]int{0, 1, 2, 3}, func (i int, _ int, _ []int) []int {
    return i+1
})
```
Similarly, there are `Filter`, `Some`, `Every`, `Find`, `FindIndex`, `IndexOf`, and others. Check the packages for the most up-to-date api.

## Slice API
The Slice API starts with an `<intslice | stringslice | byteslice>.Slice`, and existing slices can be wrapped in this type by using the `From` function. The Bare API is then exposed via methods on the slice:
```go
stringslice.From([]int{0, 1, 2, 3, 4}).Filter(func(s string, _ int, _ []string) bool {
	return len(s) > 5
})
```
This api is useful if you need to persist slices on your application structures, i.e., database models, JSON schemas, etc. This API is much larger than the other APIs, and has a companion interface in every package: `<Type>Slicer`.

## Fn API
The Fn (function) API begins with functions and applies it to a slice"
```go
intslice.BoolFn(func(i int, _ int, _ []int) bool {
    return i > 10	
}).Some([]int{5, 44, 33, 1005, 34})
```
This api is useful if you want to persist common functions used in mapping/filtering/etc. and reuse them. For example:
```go
type MyStructure struct {
    // BoolFunked interface is satisfied by the Slice API and Fn API
    filter intslice.BoolFunked
    // IntFunked interface is satisfied by the Slice API and Fn API
    mapper intslice.IntFunked
}

func New(bfn intslice.BoolFn, ifn intslice.IntFn) *MyStructure {
    return &MyStructure{filter: bfn, mapper: ifn}
}

func main() {
    myStructure := New(intslice.BoolFn(gt(50)), intslice.IntFn(starBy(100)))
    myStructure.mapper.Map([]int{2, 3, 4})
    myStructure.filter.Filter([]int{200, 30, 4003})
}

func gt(delta int) intslice.BoolFn {
    return func(i int, _ int, _ []int) bool {
        return i > delta
    }   
}

func starBy(delta int) intslice.IntFn {
    return func(i int, _ int, _ []int) int {
        return i * delta
    }   
}
```
Using the funked interfaces, the Slice API and Fn API can be dependency injected (or just instantiated) as part of your application structures.

## Composition API
The Composition api allows you to compose the Fn API into reusable pipelines for mapping, filtering, etc. Each package has a `BoolComped` and `<Type>Comped` interface, e.g., `intslice.BoolComped` and `intslice.IntComped`.

## [`stringslice`](./slices/stringslice)
This package adds functional helpers to `[]string`.

## [`intslice`](./slices/intslice)
This package adds functional helpers to `[]int`.
