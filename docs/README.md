# Maroto V2

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#template-engines) 
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfercher/maroto)](https://goreportcard.com/report/github.com/johnfercher/maroto)
[![CI](https://github.com/johnfercher/maroto/actions/workflows/goci.yml/badge.svg)](https://github.com/johnfercher/maroto/actions/workflows/goci.yml)
[![Lint](https://github.com/johnfercher/maroto/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/johnfercher/maroto/actions/workflows/golangci-lint.yml)
[![Codecov](https://img.shields.io/codecov/c/github/johnfercher/maroto)](https://codecov.io/gh/johnfercher/maroto)
[![Visits Badge](https://badges.pufler.dev/visits/johnfercher/maroto)](https://badges.pufler.dev)
[![Stars Badge](https://img.shields.io/github/stars/johnfercher/maroto.svg?style=social&label=Stars)](https://github.com/johnfercher/maroto/stargazers)

#### Maroto`v2.0.0-beta.1`is here! Try out:

* Installation with`go get`:

```bash
go get github.com/johnfercher/maroto/v2@v2.0.0-beta.1
```

The public API was completely redesigned with the aim of enhancing the 
library in various aspects. The main objectives of`v2.0.0`are:

1. [Improve usability](README.md?id=improve-usability);
2. [Allow unit testing](README.md?id=unit-testing);
3. [Add built-in metrics](README.md?id=built-in-metrics);
4. [Improve execution time](README.md?id=execution-time-improvement);
5. Allow recursive Row/Col; **(on roadmap)**
6. Allow generation based on HTML/JSON; **(on roadmap)**

## Migration

1. We will no longer maintain the current version`v1.0.0`of maroto.
   - The last version`v0.43.0`was release as`v1.0.0`through the main branch, marking the end of the old version.
   - The old`v1`code was move to a`v1`branch.
2. The master branch now keeps the`v2`code, being the default implementation now.
   - [Beta versions](https://go.dev/doc/modules/version-numbers) will be released as we achieve small deliverables.
   - There still some issues not solved from`v1`, but`v2`already solved more than 20 issues from`v1`.

## Code Example
This is part of the [simplest example](v2/examples/simplest?id=simplest).

[filename](assets/examples/simplest/v2/main.go ':include :type=code')

## PDF Example
This is part of the [billing example](v2/examples/billing?id=billing).

```pdf
	assets/pdf/billingv2.pdf
```

## Conventions

The way that maroto`v2`works is similar to the old version. The concept of components still exists, but now they are more 
rigorously controlled by maroto, meaning that`v2`will be less tightly integrated with [gofpdf][gofpdf]. This constraint will 
enable maroto`v2`to utilize other PDF providers, such as [pdfcpu][pdfcpu], in the future and even to generate other types 
of documents.

!> **gofpdf** has been archived by the owner on Nov 13, 2021.

### Structure
In maroto`v2`, everything is a **component**. When you add a **row** to the document, you are essentially adding a
**component**. Similarly, when you add an **image** to a **col**, you are also introducing a **component**. The 
functioning of maroto`v2`is straightforward: it involves constructing a **components tree** and processing the 
tree by incorporating the **components** into the document.

<iframe frameborder="0" style="width:100%;height:600px;" src="https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1&title=marotov2-structure.drawio#Uhttps%3A%2F%2Fdrive.google.com%2Fuc%3Fid%3D1H-xFq-6DNg-V6aUWsFxM0VthUvA5ptWZ%26export%3Ddownload"></iframe>
<div style="text-align: center;">Components tree</div>

### Runtime

Maroto`v2`now features a distinct separation into two runtime phases: **Declaration Phase** and **Generation Phase**.

1. **Declaration Phase:** During this phase, you will create the maroto instance, add various elements such as pages, rows, cols, images, and so on.
   - When calling the add methods, maroto will not make any changes to the document. Instead, it will solely construct the **components tree**.
2. **Generation Phase:** This phase is triggered when the `Generate() (Document, error)` method is called.
   - In this phase, maroto will traverse the **components tree** structure, compute the grid dimensions, and add components to the document.

## Improve Usability
The recursive design of maroto`v1`has proven to be a limiting factor in enhancing project features. This 
[issue][old_row_issue] was identified years ago. To enhance code quality and streamline the addition of new features, 
a refactor of the [maroto interface][old_maroto_interface] is necessary. This involves extracting interfaces for each 
feature within maroto, such as (Row, Col, Text, QRCode, Image, etc). With the new interfaces, usability will be greatly
improved, enabling maroto to reach a whole new level in this fresh new step.

### New Interfaces
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/pkg/core/core.go ':include :type=code')

## Unit Testing
In maroto`v2`, it will be possible to write unit tests by analyzing the **components tree**. To facilitate the 
writing of unit tests, we are in the process of creating a dedicated test package.

For an example, refer to [this link](v2/features/unittests?id=unit-testing).

## Built-in Metrics
This new version of maroto introduces an **optional decorator** that provides metrics for nearly all operations 
performed by the library. When the decorator is enabled, maroto will populate the **report** struct within 
the **document** response.

The **report** struct contains the following information. For a complete example, refer 
to [this link](v2/basics?id=using-metrics-decorator).

[filename](../assets/text/report.txt ':include :type=code')

## Execution Time Improvement
In Maroto`v2`, numerous performance enhancements have been implemented. The core algorithm is now **at least 
twice as fast as V1**. This disparity becomes even more remarkable when parallel generation is enabled. The 
subsequent results were achieved by generating a PDF with **100 pages** encompassing **all components supported** 
by Maroto`v2`.

[filename](../assets/text/parallel.txt ':include :type=code')

The PDF generated was a custom version of ([billing example](v2/examples/billing?id=billing)), with **100 pages**.
The pages are merged using [pdfcpu][pdfcpu]. For a complete example, please refer to
[this link](v2/features/parallelism?id=parallelism).


[gofpdf]: https://github.com/jung-kurt/gofpdf
[pdfcpu]: https://github.com/pdfcpu/pdfcpu
[old_maroto_interface]: https://github.com/johnfercher/maroto/blob/master/pkg/pdf/pdf.go
[old_row_issue]: https://github.com/johnfercher/maroto/issues/55
