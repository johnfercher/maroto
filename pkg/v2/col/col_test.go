package col

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/internal/mocks"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should success create a new col", func(t *testing.T) {
		// arrange
		size := 12
		// act
		col := New(size)

		// assert
		assert.NotNil(t, col)
		assert.Equal(t, size, col.size)
		assert.Equal(t, "col", col._type.String())
	})
}

func TestCol_GetType(t *testing.T) {
	t.Run("should success get type", func(t *testing.T) {
		// arrange
		size := 12
		col := New(size)

		// act
		_type := col.GetType()

		// assert
		assert.Equal(t, "col", _type)
	})
}

func TestCol_Add(t *testing.T) {
	t.Run("should add component when is a valid child", func(t *testing.T) {
		// arrange
		size := 12
		col := New(size)
		txt := text.New("test")

		// act
		col.Add(txt)

		// assert
		assert.Equal(t, 1, len(col.components))
		assert.Equal(t, txt, col.components[0])
	})
	t.Run("should not add component when is not a valid child", func(t *testing.T) {
		// arrange
		size := 12
		col := New(size)
		invalid := &invalidComponent{_type: v2.Page}

		// act
		col.Add(invalid)

		// assert
		assert.Equal(t, 0, len(col.components))
	})
}

func TestCol_Render(t *testing.T) {
	t.Run("should render col with 100% parent width when col size is 12", func(t *testing.T) {
		// arrange
		col := New(12)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 80.0
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert

	})
	t.Run("should render col with 91% parent width when col size is 11", func(t *testing.T) {
		// arrange
		col := New(11)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 72.8
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert
	})
	t.Run("should render col with 83% parent width when col size is 10", func(t *testing.T) {
		// arrange
		col := New(10)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 66.4
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert
	})
	t.Run("should render col with 75% parent width when col size is 9", func(t *testing.T) {
		// arrange
		col := New(9)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 60.0
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert
	})
	t.Run("should render col with 66% parent width when col size is 8", func(t *testing.T) {
		// arrange
		col := New(8)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 53.333333333333336
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert
	})
	t.Run("should render col with 60% parent width when col size is 7", func(t *testing.T) {
		// arrange
		col := New(7)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 48.0
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert
	})
	t.Run("should render col with 50% parent width when col size is 6", func(t *testing.T) {
		// arrange
		col := New(6)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 40.0
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert

	})
	t.Run("should render col with 40% parent width when col size is 5", func(t *testing.T) {
		// arrange
		col := New(5)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 33.333333333333336
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert
	})
	t.Run("should render col with 30% parent width when col size is 4", func(t *testing.T) {
		// arrange
		col := New(4)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 26.666666666666664
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert

	})
	t.Run("should render col with 25% parent width when col size is 3", func(t *testing.T) {
		// arrange
		col := New(3)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 20.0
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert

	})
	t.Run("should render col with 20% parent width when col size is 2", func(t *testing.T) {
		// arrange
		col := New(2)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 13.333333333333332
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert

	})
	t.Run("should render col with 10% parent width when col size is 1", func(t *testing.T) {
		// arrange
		col := New(1)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)

		expectedSizeX := 6.666666666666666
		pdf.EXPECT().CellFormat(expectedSizeX, ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")

		// act
		col.Render(pdf, ctx)

		// assert

	})
	t.Run("should render child components with width matching col width", func(t *testing.T) {
		// arrange
		col := New(12)
		child := text.New("test")
		col.Add(child)
		pdf := mocks.NewFpdf(t)

		ctx := context.NewRootContext(100, 100, &context.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 10,
		}).WithDimension(80, 80)
		pdf.EXPECT().SetDrawColor(255, 0, 0)
		pdf.EXPECT().CellFormat(ctx.GetXOffset(), ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")
		pdf.EXPECT().Text(ctx.GetXOffset(), ctx.GetYOffset(), "test")
		// act
		col.Render(pdf, ctx)

		// assert

	})
}

// region Test Support
type invalidComponent struct {
	_type v2.DocumentType
}

func (i *invalidComponent) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	return
}

func (i *invalidComponent) GetType() string {
	return i._type.String()
}

func (i *invalidComponent) Add(_ ...v2.Component) {
	return
}

// endregion
