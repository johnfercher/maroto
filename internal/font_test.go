package internal_test

/*func TestNewFont(t *testing.T) {
	// Arrange
	size := 10.0
	family := consts.Arial
	style := consts.Bold

	fpdf := &mocks.Fpdf{}
	fpdf.On("SetFont", family, string(style), size)
	fontstyle := internal.NewFont(fpdf, size, family, style)

	assert.NotNil(t, fontstyle)
	assert.Equal(t, fmt.Sprintf("%T", fontstyle), "*internal.fontstyle")
	assert.Equal(t, family, fontstyle.GetFamily())
	assert.Equal(t, style, fontstyle.GetStyle())
	assert.Equal(t, size, fontstyle.GetSize())
	assert.Equal(t, color.Color{Red: 0, Green: 0, Blue: 0}, fontstyle.GetColor())
}

func TestFont_GetSetFamily(t *testing.T) {
	cases := []struct {
		name        string
		family      string
		fpdf        func() *mocks.Fpdf
		assertCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertFont  func(t *testing.T, family string)
	}{
		{
			"PdfMaroto.Arial",
			consts.Arial,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "arial", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.Arial)
			},
		},
		{
			"PdfMaroto.Helvetica",
			consts.Helvetica,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "helvetica", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.Helvetica)
			},
		},
		{
			"PdfMaroto.Symbol",
			consts.Symbol,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "symbol", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.Symbol)
			},
		},
		{
			"PdfMaroto.ZapBats",
			consts.ZapBats,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "zapfdingbats", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.ZapBats)
			},
		},
		{
			"PdfMaroto.Courier",
			consts.Courier,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "courier", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.Courier)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		fontstyle := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

		// Act
		fontstyle.SetFamily(c.family)

		// Assert
		c.assertCalls(t, fpdf)
		c.assertFont(t, fontstyle.GetFamily())
	}
}

func TestFont_GetSetStyle(t *testing.T) {
	cases := []struct {
		name        string
		style       fontstyle.Type
		fpdf        func() *mocks.Fpdf
		assertCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertStyle func(t *testing.T, style fontstyle.Type)
	}{
		{
			"PdfMaroto.Normal",
			consts.Normal,
			func() *mocks.Fpdf {
				size := 10.0
				family := consts.Arial
				style := consts.Bold

				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", family, string(style), size)
				fpdf.On("SetFontStyle", mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				fpdf.AssertCalled(t, "SetFontStyle", "")
			},
			func(t *testing.T, style fontstyle.Type) {
				assert.Equal(t, style, consts.Normal)
			},
		},
		{
			"PdfMaroto.Bold",
			consts.Bold,
			func() *mocks.Fpdf {
				size := 10.0
				family := consts.Arial
				style := consts.Bold

				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", family, string(style), size)
				fpdf.On("SetFontStyle", mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				fpdf.AssertCalled(t, "SetFontStyle", "B")
			},
			func(t *testing.T, style fontstyle.Type) {
				assert.Equal(t, style, consts.Bold)
			},
		},
		{
			"PdfMaroto.Italic",
			consts.Italic,
			func() *mocks.Fpdf {
				size := 10.0
				family := consts.Arial
				style := consts.Bold

				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", family, string(style), size)
				fpdf.On("SetFontStyle", mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				fpdf.AssertCalled(t, "SetFontStyle", "I")
			},
			func(t *testing.T, style fontstyle.Type) {
				assert.Equal(t, style, consts.Italic)
			},
		},
		{
			"PdfMaroto.BoldItalic",
			consts.BoldItalic,
			func() *mocks.Fpdf {
				size := 10.0
				family := consts.Arial
				style := consts.Bold

				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", family, string(style), size)
				fpdf.On("SetFontStyle", mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				fpdf.AssertCalled(t, "SetFontStyle", "BI")
			},
			func(t *testing.T, style fontstyle.Type) {
				assert.Equal(t, style, consts.BoldItalic)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		fontstyle := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

		// Act
		fontstyle.SetStyle(c.style)

		// Assert
		c.assertCalls(t, fpdf)
		c.assertStyle(t, fontstyle.GetStyle())
	}
}

func TestFont_GetSetSize(t *testing.T) {
	// Arrange
	size := 10.0
	family := consts.Arial
	style := consts.Bold

	fpdf := &mocks.Fpdf{}
	fpdf.On("SetFont", family, string(style), size)
	fpdf.On("SetFontSize", mock.Anything)
	fontstyle := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

	// Act
	fontstyle.SetSize(16)

	// Assert
	fpdf.AssertNumberOfCalls(t, "SetFontSize", 1)
	fpdf.MethodCalled("SetFontSize", 16)
	assert.Equal(t, fontstyle.GetSize(), 16.0)
}

func TestFont_GetSetFont(t *testing.T) {
	cases := []struct {
		name        string
		family      string
		style       fontstyle.Type
		size        float64
		fpdf        func() *mocks.Fpdf
		assertCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertFont  func(t *testing.T, family string, style fontstyle.Type, size float64)
	}{
		{
			"PdfMaroto.Arial, PdfMaroto.Normal, 16",
			consts.Arial,
			consts.Normal,
			16.0,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "arial", "", 16.0)
			},
			func(t *testing.T, family string, style fontstyle.Type, size float64) {
				assert.Equal(t, family, consts.Arial)
				assert.Equal(t, style, consts.Normal)
				assert.Equal(t, 16, int(size))
			},
		},
		{
			"PdfMaroto.Helvetica, PdfMaroto.Bold, 13",
			consts.Helvetica,
			consts.Bold,
			13,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "helvetica", "B", 13.0)
			},
			func(t *testing.T, family string, style fontstyle.Type, size float64) {
				assert.Equal(t, family, consts.Helvetica)
				assert.Equal(t, style, consts.Bold)
				assert.Equal(t, 13, int(size))
			},
		},
		{
			"PdfMaroto.Symbol, PdfMaroto.Italic, 10",
			consts.Symbol,
			consts.Italic,
			10,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "symbol", "I", 10.0)
			},
			func(t *testing.T, family string, style fontstyle.Type, size float64) {
				assert.Equal(t, family, consts.Symbol)
				assert.Equal(t, style, consts.Italic)
				assert.Equal(t, 10, int(size))
			},
		},
		{
			"PdfMaroto.ZapBats, PdfMaroto.BoldItalic, 5",
			consts.ZapBats,
			consts.BoldItalic,
			5,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "zapfdingbats", "BI", 5.0)
			},
			func(t *testing.T, family string, style fontstyle.Type, size float64) {
				assert.Equal(t, family, consts.ZapBats)
				assert.Equal(t, style, consts.BoldItalic)
				assert.Equal(t, 5, int(size))
			},
		},
		{
			"PdfMaroto.Courier, PdfMaroto.Normal, 12",
			consts.Courier,
			consts.Normal,
			12,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 2)
				fpdf.AssertCalled(t, "SetFont", "courier", "", 12.0)
			},
			func(t *testing.T, family string, style fontstyle.Type, size float64) {
				assert.Equal(t, family, consts.Courier)
				assert.Equal(t, style, consts.Normal)
				assert.Equal(t, 12, int(size))
			},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		fontstyle := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

		// Act
		fontstyle.SetFont(c.family, c.style, c.size)
		family, style, size := fontstyle.GetFont()

		// Assert
		c.assertCalls(t, fpdf)
		c.assertFont(t, family, style, size)
	}
}

func TestFont_GetScaleFactor(t *testing.T) {
	// Arrange
	size := 10.0
	family := consts.Arial
	style := consts.Bold

	fpdf := &mocks.Fpdf{}
	fpdf.On("SetFont", family, string(style), size)
	fpdf.On("GetFontSize").Return(1.0, 1.0)
	sut := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

	// Act
	scalarFactor := sut.GetScaleFactor()

	// Assert
	assert.InDelta(t, scalarFactor, 2.83, 0.1)
}

func TestFont_GetSetColor(t *testing.T) {
	cases := []struct {
		name        string
		fontColor   color.Color
		Fpdf        func() *mocks.Fpdf
		assertCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertFont  func(t *testing.T, fontColor color.Color)
	}{
		{
			"Without custom color",
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				size := 10.0
				family := consts.Arial
				style := consts.Bold

				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", family, string(style), size)
				fpdf.On("SetTextColor", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "SetTextColor", 1)
				Fpdf.AssertCalled(t, "SetTextColor", 0, 0, 0)
			},
			func(t *testing.T, fontColor color.Color) {
				assert.Equal(t, fontColor.Red, 0)
				assert.Equal(t, fontColor.Green, 0)
				assert.Equal(t, fontColor.Blue, 0)
			},
		},
		{
			"With custom color",
			color.Color{Red: 20, Green: 20, Blue: 20},
			func() *mocks.Fpdf {
				size := 10.0
				family := consts.Arial
				style := consts.Bold

				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", family, string(style), size)
				fpdf.On("SetTextColor", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "SetTextColor", 1)
				Fpdf.AssertCalled(t, "SetTextColor", 20, 20, 20)
			},
			func(t *testing.T, fontColor color.Color) {
				assert.Equal(t, fontColor.Red, 20)
				assert.Equal(t, fontColor.Green, 20)
				assert.Equal(t, fontColor.Blue, 20)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := c.Fpdf()
		fontstyle := internal.NewFont(Fpdf, 10, consts.Arial, consts.Bold)

		// Act
		fontstyle.SetColor(c.fontColor)
		fontColor := fontstyle.GetColor()

		// Assert
		c.assertCalls(t, Fpdf)
		c.assertFont(t, fontColor)
	}
}*/
