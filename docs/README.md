# Maroto V2

Maroto is about to gain a version`v2.0.0`where the public API was re-done from scratch thinking to allow
a bunch of new features.

The main objectives of`v2.0.0`are:
1. [Improve usability](README.md?id=improve-usability)
2. Refactor on grid system;
2. Allow unit testing;
3. Allow parallelism;
4. Allow recursive Row/Col
5. Allow generation based on HTML/JSON;

You can learn more about maroto v2 following these topics:
* The default branch is [this](https://github.com/johnfercher/maroto/tree/v2).
* Discussions are being addressed in [this issue](https://github.com/johnfercher/maroto/issues/257).

## Migration Phase

1. We will not maintain the current version`v0.43.0`of maroto.
   * The current project`v0.43.0`will be released as`v1.0.0`version through the main branch, and this will mark the end of the old version.
   * We will let clear this move on README.md
2. We will start a new branch for version`v2.0.0`and start the development.
   * [Beta versions](https://go.dev/doc/modules/version-numbers) will be released as we achieve small deliverables.
   * The project will focus first in defining a design which allows us to deliver all objectives. After that, we will start to aggregate all features from the old version and last but not least we will add the new features.

## Improve Usability
The recursive design of maroto v1 became a limiter to improve features on the project. This [issue][old_row_issue] was spotted years ago, 
and to improve code quality and to facilitate the adding of new features we need to refactor the [maroto interface][old_maroto_interface], 
extracting interfaces to each feature inside maroto as (Row, Col, Text, QRCode, Image and etc). With the new interface 
we will be able to improve usability a lot, with that maroto will reach a whole new level in this new fresh step.

### New Interfaces
[filename](https://raw.githubusercontent.com/johnfercher/maroto/v2/pkg/v2/domain/domain.go ':include :type=code')

## Structure
<iframe frameborder="0" style="width:100%;height:600px;" src="https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1&title=marotov2-structure.drawio#Uhttps%3A%2F%2Fdrive.google.com%2Fuc%3Fid%3D1H-xFq-6DNg-V6aUWsFxM0VthUvA5ptWZ%26export%3Ddownload"></iframe>



[old_maroto_interface]: https://github.com/johnfercher/maroto/blob/master/pkg/pdf/pdf.go
[old_row_issue]: https://github.com/johnfercher/maroto/issues/55