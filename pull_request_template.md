<!-- Please follow the PR naming pattern. -->
<!-- For features: feature/name -->
<!-- For fixes: fix/name -->

**Description**
<!-- Please, describe how this PR will be useful. If it has any tricky technical detail, please explain too. -->

**Related Issue**
<!-- If it has any issue related to this PR, please add a reference here. -->

**Checklist**
> check with "x", **ONLY IF APPLIED** to your change

- [ ] All methods associated with structs has ```func (<first letter of struct> *struct) method() {}``` name style. <!-- If applied -->
- [ ] Wrote unit tests for new/changed features. <!-- If applied -->
- [ ] Followed the unit test ```when,should``` naming pattern. <!-- If applied -->
- [ ] All mocks created with ```m := mocks.NewConstructor(t)```. <!-- If applied -->
- [ ] All mocks using ```m.EXPECT()``` method to mock methods. <!-- If applied -->
- [ ] Updated docs/doc.go and docs/* <!-- If applied -->
- [ ] Updated example_test.go <!-- If applied -->
- [ ] Updated README.md <!-- If applied -->
- [ ] Executed `make examples` to update all examples inside docs/examples. <!-- If applied -->
- [ ] New public methods/structs/interfaces has comments upside them explaining they responsibilities <!-- If applied -->
- [ ] Executed `make dod` with none issues pointed out by `golangci-lint`