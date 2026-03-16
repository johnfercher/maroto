# Unit Testing

maroto provides a dedicated `pkg/test` package that lets you write deterministic unit tests for your PDF-generating code. Instead of comparing binary PDF output, the test package serialises the document's **component tree** to JSON and compares it against a saved fixture. This makes tests fast, readable, and diff-friendly.

## How it works

1. Build your maroto document normally.
2. In your test, call `test.New(m)` to wrap the document.
3. Call `.Assert(t, "fixture-name")` to compare the component tree against the JSON file at the path defined in `.maroto.yml`.
4. On the first run (or when you want to update the fixture), call `.Save("fixture-name")` to write the JSON file.

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/test#New)
* [method : Assert](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/test#MarotoTest.Assert)
* [method : Equals](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/test#MarotoTest.Equals)
* [method : Save](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/test#MarotoTest.Save)

## Configuration
To allow unit testing of PDFs, you must create the`.maroto.yml`file on the root project folder. The field`test_path`define which [folder](https://github.com/johnfercher/maroto/tree/v2/test/maroto)
will keep the json files to unit test the components tree.

[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/.maroto.yml ':include :type=code')

## Code
[filename](../../assets/examples/unittests/v2/main_test.go ':include :type=code')

## Test file
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/example_unit_test.json ':include :type=code')

