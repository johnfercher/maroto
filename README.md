Base on [johnfercher/maroto](https://github.com/johnfercher/maroto)

Add  function Base on [gofpdf](https://github.com/jung-kurt/gofpdf) 
1. AddUTF8Font to support chinese 
1. Add  SetProtection to support protection pdf 

使用更加优雅的方式去书写PDF。不再为每个点定位而苦恼

站在巨人的肩膀

AddUTF8Font 方法，支持自己设定字体:比如使用 [NotoSansSC-Regular.ttf](https://github.com/jsntn/webfonts/blob/master/NotoSansSC-Regular.ttf) 就可以支持中文的展示。

SetProtection 方法，支持PDF设置密码。

maroto的布局是一行12个格子，然后向每个格子每个格子的添加数据。


Example
```go
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.AddUTF8Font("NotoSansSC", "", "./NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "I", "./NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "B", "./NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "BI", "./NotoSansSC-Regular.ttf")
	m.SetPageMargins(10, 35, 10)
	textProps := props.Text{
		Top:    0.5,
		Family: "NotoSansSC",
		Style:  consts.Bold,
		Align:  consts.Left,
	}
	m.Row(5, func() {
		m.ColSpace(5)
		m.Col(2, func() {
			m.Text("世界你好",textProps)
		})
		m.ColSpace(5)
	})
	m.SetProtection(gofpdf.CnProtectPrint, "123", "abc")
	err := m.OutputFileAndClose("世界你好.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
```