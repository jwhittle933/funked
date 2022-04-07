<h1 style="text-align: center">Slices</h1>

This package is designed to create "handles" on slices for easy use and composition. Nothing here is revolutionary: its goal is ease of use and convenience.

## APIs
For examples of the APIs presented below, see the [examples](./examples) directory.

### Naked API
The libraries here share a common API, but there are several ways to use it. The first are the "naked" functions. These are transparent functions that operate on a slice.

```go
slices.Map([]int{0, 1, 2, 3}, func (i int, _ int, _ []int) []int {
    return i+1
})
```
Similarly, there are `Filter`, `Some`, `Every`, `Find`, `FindIndex`, `IndexOf`, and others. Check the package for the most up-to-date api.

## Slice API
The Slice API starts with an `slice.Slice[T]`, and existing slices can be wrapped in this type by using the `From` function. The Naked API is then exposed via methods on the slice:
```go
slice.From([]int{0, 1, 2, 3, 4}).Filter(func(s string, _ int, _ []string) bool {
	return len(s) > 5
})
```
This api is useful if you need to persist slices on your application structures, i.e., database models, JSON schemas, etc. This API is much larger than the other APIs, and has a companion interface: `Slicer[T]`.

## Fn API
The Fn (function) API begins with functions and applies it to a slice.
```go
slice.BoolFn(func(i int, _ int, _ []int) bool {
    return i > 10	
}).Some([]int{5, 44, 33, 1005, 34})
```
This api is useful if you want to persist common functions used in mapping/filtering/etc. and reuse them. For example:
```go
type MyStructure struct {
    // BoolFunked interface is satisfied by the Slice API and Fn API
    filter slice.BoolFunked[int]
    // IntFunked interface is satisfied by the Slice API and Fn API
    mapper slice.TFunked[int]
}

func New(bfn slice.BoolFn[int], ifn slice.TFn[int]) *MyStructure {
    return &MyStructure{filter: bfn, mapper: ifn}
}

func main() {
    myStructure := New(gt(50), starBy(100))
    myStructure.mapper.Map([]int{2, 3, 4})
    myStructure.filter.Filter([]int{200, 30, 4003})
}

func gt(delta int) slice.BoolFn[int] {
    return func(i int, _ int, _ []int) bool {
        return i > delta
    }   
}

func times(delta int) slice.TFn[int] {
    return func(i int, _ int, _ []int) int {
        return i * delta
    }   
}
```
Using the funked interfaces, the Slice API and Fn API can be dependency injected (or just instantiated) as part of your application structures.

## Composition API
The Composition api allows you to compose the Fn API into reusable pipelines for mapping, filtering, etc.
```go
slice.
    TFn(prependDashes).
    And(doubleWithDashes).
    And(appendDashes).
    Map(strs)
```
