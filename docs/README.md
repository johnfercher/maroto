# Maroto V2

!> **Working in progress**

Maroto is about to gain a version`v2.0.0`where the public API was re-done from scratch thinking to improve the lib
in many aspects. The main objectives of`v2.0.0`are:

1. [Improve usability](README.md?id=improve-usability);
2. [Allow unit testing](README.md?id=unit-testing);
3. [Add built-in metrics](README.md?id=built-in-metrics);
4. [Allow parallelism](README.md??id=parallelism);
5. Allow recursive Row/Col; **(on roadmap)**
6. Allow generation based on HTML/JSON; **(on roadmap)**

You can learn more about maroto`v2`following these topics:
* The default branch is [this](https://github.com/johnfercher/maroto/tree/v2).
* Discussions are being addressed in [this issue](https://github.com/johnfercher/maroto/issues/257).

## Conventions & Algorithm

The way that maroto`v2`works is similar to the old version. The concept of components still exists, but they are more
strictly maintained on the code to make the new version less dependent on [gofpdf][gofpdf], this will allow maroto`v2` to use
other pdf providers as [pdfcpu][pdfcpu] in the future and even to generate other kind of documents.

!> **gofpdf** has been archived by the owner on Nov 13, 2021.

### Structure
In maroto`v2`everything is a **component**, when you add a **row** to the document, you are adding a **component**, when you add an **image** to a **col** you
are also adding a **component**. The way that maroto`v2`works is simply by creating a **component tree** and computing the tree by writing
the **components** inside the document.

<iframe frameborder="0" style="width:100%;height:600px;" src="https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1&title=marotov2-structure.drawio#Uhttps%3A%2F%2Fdrive.google.com%2Fuc%3Fid%3D1H-xFq-6DNg-V6aUWsFxM0VthUvA5ptWZ%26export%3Ddownload"></iframe>
<div style="text-align: center;">Components tree</div>

### Runtime
Maroto`v2` now has a clear separation in two runtime phases: **Declaration Phase** and **Generation Phase**.
1. **Declaration phase:** during this phase you will create the maroto instance, add a bunch of pages, rows, cols, images and etc.
   * When calling the add methods, maroto will do nothing in a document. It will only create the components tree.
2. **Generation phase:** this occurs when is called the `Generate() (Document, error)` method.
   * During this phase maroto will walk through the components tree structure computing the grid dimensions and add components to the document.

## Migration

1. We will not maintain the current version`v0.43.0`of maroto.
   * The current project`v0.43.0`will be released as`v1.0.0`version through the main branch, and this will mark the end of the old version.
   * We will let clear this move on README.md
2. We will start a new branch for version`v2.0.0`and start the development.
   * [Alpha/Beta versions](https://go.dev/doc/modules/version-numbers) will be released as we achieve small deliverables.
   * The project will focus first in defining a design which allows us to deliver all objectives. After that, we will start to aggregate all features from the old version and last but not least we will add the new features.

## Improve Usability
The recursive design of maroto`v1`became a limiter to improve features on the project. This [issue][old_row_issue] was spotted years ago, 
and to improve code quality and to facilitate the adding of new features we need to refactor the [maroto interface][old_maroto_interface], 
extracting interfaces to each feature inside maroto as (Row, Col, Text, QRCode, Image and etc). With the new interface 
we will be able to improve usability a lot, with that maroto will reach a whole new level in this new fresh step.

### New Interfaces
[filename](https://raw.githubusercontent.com/johnfercher/maroto/v2/pkg/v2/domain/domain.go ':include :type=code')

## Unit Testing
In maroto`v2`will be possible to write unit tests by analyzing the **component tree**. For this purpose we are creating a
test package to facilitate the writing of unit tests. 

You can see an example [here](v2/tests.md).

## Built-in Metrics
This new version of maroto provides an **optional decorator** which metrifies almost all operations realized by the lib.
When enabling the decorator, maroto will fill the **report** struct on the **document** response.

The **report** struct has the follow informations. You can see the full example [here](v2/basics?id=using-metrics-decorator).

[filename](../assets/text/report.txt ':include :type=code')

## Parallelism
There are several performance improvements in maroto`v2`, but more expressive is the possibility to generate chunks
of pages in a parallel way. These are the results by generating a PDF with **100 pages** with **all components supported** by maroto`v2`. 

[filename](../assets/text/parallel.txt ':include :type=code')

The pages are merged using [pdfcpu][pdfcpu] and you can see full the example [here](v2/configbuilder.md?id=parallelism).

[gofpdf]: https://github.com/jung-kurt/gofpdf
[pdfcpu]: https://github.com/pdfcpu/pdfcpu
[old_maroto_interface]: https://github.com/johnfercher/maroto/blob/master/pkg/pdf/pdf.go
[old_row_issue]: https://github.com/johnfercher/maroto/issues/55