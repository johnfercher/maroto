# Text

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#New)
* [constructor : NewCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#NewCol)
* [constructor : NewRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#NewRow)
* [props : Text](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Text)
* [component : Text](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#Text)

## Overview

Using the Text component, you can add personalized texts to your PDF. Possible customizations include: font, color, style and margin. The component also allows you to add links and use different styles in the same text.
Different builders are offered to facilitate use, they are: 

* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#New)
* [constructor : NewCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#NewCol)
* [constructor : NewRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#NewRow)
* [constructor : NewCustomText](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#NewCustomText)

In naughty texts can contain one or more styles, if your text does not need to mix several styles, you can use the constructors **New**, **NewCol** or **NewRow**; However, if you need to merge different styles, this can be done with the **NewCustomText** constructor, unlike the previous constructors, you will need to pass a list of [Sub texts](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/core/entity#SubText), each subtext will have its own [style](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#SubText)


## Code Example
[filename](../../assets/examples/textgrid/v2/main.go ':include :type=code')

## PDF Generated
```pdf
    assets/pdf/textgridv2.pdf
```

## Time Execution
[filename](../../assets/text/textgridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/textgrid.json  ':include :type=code')