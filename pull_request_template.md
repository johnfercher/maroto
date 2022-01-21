<!-- Please follow the PR naming pattern. -->
<!-- For features: feature/name -->
<!-- For fixes: fix/name -->

**Description**
<!-- Please, describe how this PR will be useful. If it has any tricky technical detail, please explain too. -->

**Related Issue**
<!-- If it has any issue related to this PR, please add a reference here. -->

**Checklist**
> check with "x", if applied to your change

- [ ] All methods associated with structs has ```func (s *struct) method() {}``` name style. <!-- If applied -->
- [ ] Wrote unit tests for new/changed features. <!-- If applied -->
- [ ] Updated docs/doc.go <!-- If applied -->
- [ ] Updated pkg/pdf/example_test.go <!-- If applied -->
- [ ] Updated README.md <!-- If applied -->
- [ ] Updated all examples inside internal/examples <!-- If applied -->
- [ ] New public methods/structs/interfaces has comments upside them explaining they responsibilities <!-- If applied -->
- [ ] Executed `make dod` with none issues pointed out by `golangci-lint` and `goreportcard-cli`