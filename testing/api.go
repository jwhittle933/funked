package testing

import (
	"errors"
	"fmt"
)

var (
	GenericError = errors.New("generic error")
)

type Expected struct {
	ExpectedCalls int
	Calls         int
}

func (e Expected) String() string {
	return fmt.Sprintf("Expected calls: %d. Recieved: %d", e.ExpectedCalls, e.Calls)
}

func (e *Expected) Incr() {
	e.Calls++
}

type Tester interface {
	TestName() string
	ShouldError() bool
}

type TestCase struct {
	Name           string
	ExpectingError bool
}

func (t TestCase) TestName() string {
	return t.Name
}

func (t TestCase) ShouldError() bool {
	return t.ExpectingError
}
