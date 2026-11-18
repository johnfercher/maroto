# Protection

`WithProtection` applies PDF password protection and permission restrictions. You can require a password to open the document, restrict printing, copying, or modifying, or combine multiple restrictions.

## Permission flags (`protection.Type`)

| Constant | Restricted action |
|----------|-------------------|
| `protection.None` | No restrictions |
| `protection.Print` | Printing |
| `protection.Modify` | Document modification |
| `protection.Copy` | Copying text and graphics |
| `protection.AnnotForms` | Annotating and filling forms |

Flags can be combined with `|`: `protection.Print | protection.Copy` restricts both printing and copying.

## Usage notes

- An empty user password disables the open password: a document with only an owner password can be opened without a password, but restrictions apply to regular users.
- An empty owner password is replaced by a random one that is not returned, so nobody gets owner access. Set an owner password if you need to lift the restrictions later.
- PDF encryption strength depends on the underlying library. maroto encrypts with 40-bit RC4 (standard security handler revision 2) via gopdf.
- PDF protection is not a strong security guarantee — determined users with the right tools can bypass it. Use it as a deterrent, not a security boundary.

## GoDoc
* [builder : WithProtection](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithProtection)
* [protection : Type](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/consts/protection)

## Code Example
[filename](../../assets/examples/protection/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/protectionv2.pdf
```
## Time Execution
[filename](../../assets/text/protectionv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/protection.json  ':include :type=code')