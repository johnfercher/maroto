# Generate Unit Tests

Generate unit tests for a Go file in this project following the project's established conventions.

## Instructions

Read the target file(s) provided by the user. Then read 1-2 existing test files in the same package for context. Generate a complete `_test.go` file following ALL rules below.

---

## Mandatory conventions (from PR checklist)

### 1. Package and imports
- Use the **external test package**: `package <pkg>_test` (not `package <pkg>`).
- Import `"testing"`, `"fmt"`, `"github.com/stretchr/testify/assert"`.
- Import mocks from `github.com/johnfercher/maroto/v2/mocks` when dependencies need to be mocked.
- Import `github.com/johnfercher/maroto/v2/internal/fixture` for pre-built prop values when available.

### 2. Test function structure
Every exported function or method needs a test function:
- **Constructors** → `TestNew<Type>(t *testing.T)` — verify the returned value is not nil and has the correct concrete type using `fmt.Sprintf("%T", sut)`.
- **Methods** → `Test<Type>_<MethodName>(t *testing.T)` — one parent function, many subtests.

### 3. Parallelism (mandatory everywhere)
```go
func TestFoo_Bar(t *testing.T) {
    t.Parallel()                          // ← always first line of parent
    t.Run("when X, should Y", func(t *testing.T) {
        t.Parallel()                      // ← always first line of all subtests
        // Arrange
        ...
        // Act
        ...
        // Assert
        ...
    })
}
```

### 4. Subtest naming pattern
Use `"when <condition>, should <expected outcome>"` — always lowercase, always present tense:
```go
t.Run("when value is zero, should use default", ...)
t.Run("when color is nil, should not call SetDrawColor", ...)
t.Run("when checked is true, should draw X mark lines", ...)
```

### 5. AAA comments inside every subtest
```go
// Arrange
...
// Act
result := sut.Method(...)
// Assert
assert.Equal(t, expected, result)
```
- Omit `// Assert` only when the test relies entirely on mock expectations.
- Omit `// Arrange` when there is no steps done between `// Arrange` and `// Act`
- Use `// Act & Assert` instead of separated `// Act` and `// Assert` when `// Act` and `// Assert` can
be placed in one line.

### 6. System under test variable name
Always name the instance being tested `sut`:
```go
sut := gofpdf.NewCheckbox(fpdf, font)
```

### 7. Mocks
- Create with: `dep := mocks.NewDependency(t)` (never `new(...)` or `&mocks.Dependency{}`).
- Set expectations with the type-safe EXPECT API:
  ```go
  dep.EXPECT().MethodName(arg1, arg2).Return(val1, val2)
  ```
- Try avoid using `mock.Anything`, it's only allowed when is not possible to build the arguments.
  When you have to use a `mock.Anything` warn the user that you have to. And make explicit on commits
  and pull requests.
  ```go
  dep.EXPECT().MethodName(mock.Anythin, arg2).Return(val1, val2) // mock.Anything have to be avoided 
  ```
- Assert call count only when it's important, example: when dependency is called inside a for loop.
  ```go
  dep.AssertNumberOfCalls(t, "MethodName", 1)
  ```
- For methods expected **not** to be called, do not set an EXPECT — testify/mock will fail 
  automatically if an unexpected call occurs. If a dependency have zero mock definitions called
  remove this dependency and pass it as `nil` on `sut` constructor.

### 8. Fixture package
- Prefer `fixture.CheckboxProp()`, `fixture.TextProp()`, etc. over building props inline whenever a fixture exists for that type. Use inline props only when the test needs specific values that differ from the fixture.
- If one inline fixture is used in more than one file, move it to `package fixture` and reuse it from there.

### 9. Float64 precision
When computing expected float64 values that involve division (e.g., `80.0/100.0`), use values that are exactly representable in IEEE 754 (powers of 2: 0.5, 0.25, etc.) or write out the exact decimal to avoid `9.999999999999998` mismatches.

### 10. Quantity of testing cases on subsets
- When generating tests, make sure to count the **Cyclomatic Complexity** and generate the correct amount of subset case tests
(see the reference for more info about cyclomatic complexity).
- When working with mathematical operations like: `+`, `-`, `*`, `/`, `math.Pow()` or any usage of the lib `math` you can create
more testing cases/subsets to make sure that the math operation is correctly tested, vary values between negative and positive values.

### 11. Concept of multiple
- When testing a method that receives a slice/array of any type `[]int`, `[]string`, `[]any`, or a variadic argument 
(see the reference for more info about variadic arguments) you have to create at least `3` tests:
    1. One subset passing `nil`.
    2. Other subset passing empty, like: `[]string{}`.
    3. Last subset passing at least `3` values, like: `[]int{1, 2, 5}`

---

## Template for a constructor test

```go
func TestNew<Type>(t *testing.T) {
    t.Parallel()
    // Act
    sut := pkg.New<Type>(nil, nil)

    // Assert
    assert.NotNil(t, sut)
    assert.Equal(t, "*pkg.<Type>", fmt.Sprintf("%T", sut))
}
```

## Template for a method test

- Should not place empty new lines `\n` between `t.Run` on a function.

```go
func Test<Type>_<Method>(t *testing.T) {
    t.Parallel()
    t.Run("when <condition1>, should <outcome1>", func(t *testing.T) {
        t.Parallel()
        // Arrange
        dep := mocks.New<Dep>(t)
        dep.EXPECT().<MockedMethod>(<args>).Return(<vals>)

        sut := pkg.New<Type>(dep)

        // Act
        result, err := sut.<Method>(<args>)

        // Assert
        assert.Nil(t, err)
        assert.Equal(t, expected, result)
    })
    t.Run("when <condition2>, should <outcome2>", func(t *testing.T) {
        t.Parallel()
        // Arrange
        dep := mocks.New<Dep>(t)
        dep.EXPECT().<MockedMethod>(<args>).Return(<vals>)
        
        sut := pkg.New<Type>(dep)
        
        // Act
        result, err := sut.<Method>(<args>)
        
        // Assert
        assert.Nil(t, err)
        assert.Equal(t, expected, result)
    })
}
```

---

## What to generate

Given the file `$ARGUMENTS`:

1. Read the target source file.
2. Read 1-2 existing `_test.go` files in the same package as reference.
3. For every **exported** function, method, or constructor, generate subtests that cover:
   - The **happy path** (valid inputs, expected output).
   - **Edge cases and branches**: nil inputs, zero values, negative values, error returns, each `if` branch.
   - Any **interaction with mocked dependencies** (each mock method that should or should not be called).
4. Write the complete test file at `<source_file_without_.go>_test.go` in the same directory.
5. Verify the test file compiles by checking imports match what is used.

## References

### Cyclomatic Complexity
Cyclomatic complexity is a software metric used to indicate the complexity of a program. It is a quantitative 
measure of the number of linearly independent paths through a program's source code. It was developed by 
Thomas J. McCabe, Sr. in 1976.

Cyclomatic complexity is computed using the control-flow graph of the program. The nodes of the graph 
correspond to indivisible groups of commands of a program, and a directed edge connects two nodes if 
the second command might be executed immediately after the first command. Cyclomatic complexity may 
also be applied to individual functions, modules, methods, or classes within a program.

One testing strategy, called basis path testing by McCabe who first proposed it, is to test each 
linearly independent path through the program. In this case, the number of test cases will equal 
the cyclomatic complexity of the program.

If you need more information, this is the original source: https://en.wikipedia.org/wiki/Cyclomatic_complexity

### Variadic Arguments (vararg)
Variadic functions in Go allow you to pass a variable number of arguments to a function. This feature is useful
when you don’t know beforehand how many arguments you will pass. A variadic function accepts multiple arguments 
of the same type and can be called with any number of arguments, including none.

#### Example
```golang
package main
import "fmt"

// Variadic function to calculate sum
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
    fmt.Println("Sum of 4, 5:", sum(4, 5))
    fmt.Println("Sum of no numbers:", sum())
}
```