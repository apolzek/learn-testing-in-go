# learn-testing-in-go

A hands on collection of small examples that show how to test Go code in
different ways. Each folder is a standalone Go module that focuses on one
testing technique. You can read the theory here and then open the matching
folder to run the code.

## Why testing matters

Tests are code that checks your code. They give you confidence that a function
behaves the way you expect, and they warn you when a future change breaks
something that used to work. In Go testing is part of the standard toolchain,
so you do not need any extra framework to get started. You write test files,
you run `go test`, and the tool reports what passed and what failed.

A few ideas that show up across every example in this repo:

* A test file ends with `_test.go` and lives in the same folder as the code it
  tests.
* A test function starts with `Test`, takes one argument of type `*testing.T`,
  and reports failures by calling methods like `t.Errorf`.
* A benchmark function starts with `Benchmark` and takes `*testing.B`.
* You run everything with the `go test` command.

## How to run the examples

Every folder is its own module. Move into the folder you want and run the
test command shown in that section.

```bash
cd testing-with-standard-library
go test ./...
```

## Folder guide

| Folder | What it teaches |
| --- | --- |
| `testing-with-standard-library` | Plain tests with the built in `testing` package |
| `testing-with-testify` | Cleaner assertions with the testify library |
| `testing-with-goconvey` | Readable, nested specs with GoConvey |
| `testing-with-mock` | Replacing real dependencies with mocks |
| `testing-benchmark` | Measuring how fast code runs |
| `integration-test` | Testing two services talking to each other |
| `load-testing` | Sending heavy traffic with Vegeta |
| `performance-profiling` | Finding slow code with pprof and trace |

## Standard library testing

Folder: `testing-with-standard-library`

### Theory

The `testing` package ships with Go. To test a function you write a function
named `TestSomething` that receives `*testing.T`. Inside it you call the code,
compare the real result with the value you expect, and call `t.Errorf` when
they do not match. A failed assertion marks the test as failed but lets the
rest of the function keep running.

A common and recommended pattern is the table driven test. You build a slice of
cases, each one holding inputs and the expected output, and then you loop over
the slice. This keeps the logic in one place and makes it easy to add new cases.

### Practice

A simple test compares one result against one expected value:

```go
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
```

A table driven test runs many cases through the same logic:

```go
func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{10, -5, 5},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("For %d+%d, expected %d, but got %d", test.a, test.b, test.expected, result)
		}
	}
}
```

Run it:

```bash
go test ./...
```

Add `-v` to see the name and result of every test:

```bash
go test -v ./...
```

## Testify

Folder: `testing-with-testify`

### Theory

The standard library works well, but writing `if` checks for every assertion
gets repetitive. Testify is a popular library that gives you ready made
assertions such as `assert.Equal`, `assert.True`, and `assert.Nil`. The result
is shorter test code and clearer failure messages. Testify does not replace the
`testing` package. It works on top of it, so you still write `Test` functions
and still run `go test`.

### Practice

The same addition test with testify:

```go
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	assert.Equal(t, 5, result, "Expected 5, but got %d", result)
}
```

Run it:

```bash
go test ./...
```

## GoConvey

Folder: `testing-with-goconvey`

### Theory

GoConvey lets you write tests as readable specifications. You describe behavior
with `Convey` blocks and check values with `So`. The blocks can be nested, which
helps you express context in plain language, for example "given a sum function,
when I add two numbers, the result should be five". GoConvey also offers a web UI
that watches your files and reruns tests automatically.

### Practice

```go
import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSum(t *testing.T) {
	Convey("Sum function should add two numbers", t, func() {
		result := Sum(2, 3)
		Convey("The result should be 5", func() {
			So(result, ShouldEqual, 5)
		})
	})
}
```

Run it:

```bash
go test ./...
```

## Mocking

Folder: `testing-with-mock`

### Theory

Sometimes the code you want to test depends on something slow or external, like
a database or a remote API. You do not want a unit test to hit a real database.
A mock is a fake version of that dependency that you control during the test. In
Go the key to mocking is the interface. If your function accepts an interface
instead of a concrete type, you can pass the real implementation in production
and a mock in your tests.

In this example the calculator depends on a `DBService` interface. The real
program passes a struct that returns a value, and the test passes a mock built
with testify that you program to return whatever you want.

### Practice

The code depends on an interface, not a concrete type:

```go
type DBService interface {
	GetSum() int
}

func GetSumFromDB(db DBService) int {
	return db.GetSum()
}
```

The test passes a mock that you program in advance:

```go
type MockDBService struct {
	mock.Mock
}

func (m *MockDBService) GetSum() int {
	args := m.Called()
	return args.Int(0)
}

func TestGetSumFromDB(t *testing.T) {
	mockDB := new(MockDBService)
	mockDB.On("GetSum").Return(10)

	result := GetSumFromDB(mockDB)

	mockDB.AssertExpectations(t)
	if result != 10 {
		t.Errorf("Expected result to be 10, but got %d", result)
	}
}
```

Run it:

```bash
go test ./calculator/ -v
```

## Benchmarks

Folder: `testing-benchmark`

### Theory

A benchmark measures how fast a piece of code runs. A benchmark function starts
with `Benchmark` and takes `*testing.B`. Inside you run the code in a loop that
goes from zero to `b.N`. Go decides the value of `b.N` for you. It keeps raising
the count until the timing is stable, then it reports the average time per
operation in nanoseconds.

### Practice

```go
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
```

Run it:

```bash
go test -bench=.
```

A sample line of output looks like this. It says the loop ran many times and
each call took a fraction of a nanosecond:

```text
BenchmarkAdd-4   1000000000   0.6201 ns/op
```

## Integration testing

Folder: `integration-test`

### Theory

A unit test checks one function in isolation. An integration test checks that
several parts work together. In this example two HTTP endpoints, one for adding
and one for multiplying, are exposed by the same service. The test starts by
sending real HTTP requests, reads the JSON responses, and confirms the numbers
are correct. Because the test talks to a running server, you start the service
first and then run the test against it.

### Practice

First start the service:

```bash
go run add_service.go
```

You can poke it by hand with curl:

```bash
curl -vv -X POST http://localhost:8080/multiply -H "Content-Type: application/json" -d '{"a": 2, "b": 3}'
```

The test sends requests and checks each response:

```go
data := map[string]int{"a": 2, "b": 3}
jsonData, _ := json.Marshal(data)

addResp, _ := http.Post(addURL, "application/json", bytes.NewBuffer(jsonData))
addBody, _ := io.ReadAll(addResp.Body)

var addResult map[string]int
json.Unmarshal(addBody, &addResult)
So(addResult["result"], ShouldEqual, 5)
```

With the service running, run the test in another terminal:

```bash
go test ./...
```

## Load testing

Folder: `load-testing`

### Theory

Load testing answers a different question than unit tests. Instead of asking "is
the answer correct", it asks "how does the service behave under heavy traffic".
You send many requests per second and watch latency, throughput, and error rate.
This example uses Vegeta, a command line tool that fires requests at a target and
prints a report. You need Vegeta installed and the service from the integration
example running.

### Practice

Start the service, then attack it for ten seconds and read the report:

```bash
echo "POST http://localhost:8080/multiply
@data.json" | vegeta attack -duration=10s | tee results.bin | vegeta report
```

Run at a fixed rate of one hundred requests per second for thirty seconds and
draw a chart of the results:

```bash
echo "POST http://localhost:8080/multiply
@data.json" | vegeta attack -rate=100 -duration=30s > attack-results.bin

cat attack-results.bin | vegeta plot > plot.html
```

Open `plot.html` in a browser to see latency over time.

## Performance profiling

Folder: `performance-profiling`

### Theory

When code is slower than you want, profiling tells you where the time goes. The
Go runtime can record a CPU profile while your program runs. It samples the call
stack many times per second and shows which functions used the most time. A
trace captures an even more detailed timeline of what the program did. You then
open these files with `go tool pprof` and `go tool trace` to explore them.

### Practice

The program starts a CPU profile and a trace, does some work, then stops them:

```go
f, _ := os.Create("profile.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

Run the program to produce the profile files:

```bash
go run main.go
```

Inspect the CPU profile:

```bash
go tool pprof profile.prof
```

Inspect the execution trace in your browser:

```bash
go tool trace trace.out
```

## Summary

| Technique | Question it answers |
| --- | --- |
| Unit test | Does this function return the right value |
| Table driven test | Does it work for many inputs |
| Testify and GoConvey | Can I write the same checks more clearly |
| Mock | Can I test code that depends on a database or API |
| Benchmark | How fast is this code |
| Integration test | Do the parts work together |
| Load test | How does it behave under heavy traffic |
| Profiling | Where is the slow code |

Pick a folder, read the section above, and run the commands. Each example is
small on purpose so you can focus on one idea at a time.
