package fpdf

import (
	"io"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// Fpdf is an extension of gofpdf interface which can expose more useful methods.
type Fpdf interface {
	AddFont(familyStr, styleStr, fileStr string)
	AddFontFromBytes(familyStr, styleStr string, jsonFileBytes, zFileBytes []byte)
	AddFontFromReader(familyStr, styleStr string, r io.Reader)
	AddLayer(name string, visible bool) (layerID int)
	AddLink() int
	AddPage()
	AddPageFormat(orientationStr string, size gofpdf.SizeType)
	AddSpotColor(nameStr string, c, m, y, k byte)
	AliasNbPages(aliasStr string)
	ArcTo(x, y, rx, ry, degRotate, degStart, degEnd float64)
	Arc(x, y, rx, ry, degRotate, degStart, degEnd float64, styleStr string)
	BeginLayer(id int)
	Beziergon(points []gofpdf.PointType, styleStr string)
	Bookmark(txtStr string, level int, y float64)
	CellFormat(w, h float64, txtStr, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string)
	Cellf(w, h float64, fmtStr string, args ...interface{})
	Cell(w, h float64, txtStr string)
	Circle(x, y, r float64, styleStr string)
	ClearError()
	ClipCircle(x, y, r float64, outline bool)
	ClipEllipse(x, y, rx, ry float64, outline bool)
	ClipEnd()
	ClipPolygon(points []gofpdf.PointType, outline bool)
	ClipRect(x, y, w, h float64, outline bool)
	ClipRoundedRect(x, y, w, h, r float64, outline bool)
	ClipText(x, y float64, txtStr string, outline bool)
	Close()
	ClosePath()
	CreateTemplateCustom(corner gofpdf.PointType, size gofpdf.SizeType, fn func(*gofpdf.Tpl)) gofpdf.Template
	CreateTemplate(fn func(*gofpdf.Tpl)) gofpdf.Template
	CurveBezierCubicTo(cx0, cy0, cx1, cy1, x, y float64)
	CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1 float64, styleStr string)
	CurveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1 float64, styleStr string)
	CurveTo(cx, cy, x, y float64)
	Curve(x0, y0, cx, cy, x1, y1 float64, styleStr string)
	DrawPath(styleStr string)
	Ellipse(x, y, rx, ry, degRotate float64, styleStr string)
	EndLayer()
	Err() bool
	Error() error
	GetAlpha() (alpha float64, blendModeStr string)
	GetAutoPageBreak() (auto bool, margin float64)
	GetCellMargin() float64
	GetConversionRatio() float64
	GetDrawColor() (int, int, int)
	GetDrawSpotColor() (name string, c, m, y, k byte)
	GetFillColor() (int, int, int)
	GetFillSpotColor() (name string, c, m, y, k byte)
	GetFontDesc(familyStr, styleStr string) gofpdf.FontDescType
	GetFontSize() (ptSize, unitSize float64)
	GetImageInfo(imageStr string) (info *gofpdf.ImageInfoType)
	GetLineWidth() float64
	GetMargins() (left, top, right, bottom float64)
	GetPageSizeStr(sizeStr string) (size gofpdf.SizeType)
	GetPageSize() (width, height float64)
	GetStringWidth(s string) float64
	GetTextColor() (int, int, int)
	GetTextSpotColor() (name string, c, m, y, k byte)
	GetX() float64
	GetXY() (float64, float64)
	GetY() float64
	HTMLBasicNew() (html gofpdf.HTMLBasicType)
	Image(imageNameStr string, x, y, w, h float64, flow bool, tp string, link int, linkStr string)
	ImageOptions(imageNameStr string, x, y, w, h float64, flow bool, options gofpdf.ImageOptions, link int, linkStr string)
	ImageTypeFromMime(mimeStr string) (tp string)
	LinearGradient(x, y, w, h float64, r1, g1, b1, r2, g2, b2 int, x1, y1, x2, y2 float64)
	LineTo(x, y float64)
	Line(x1, y1, x2, y2 float64)
	LinkString(x, y, w, h float64, linkStr string)
	Link(x, y, w, h float64, link int)
	Ln(h float64)
	MoveTo(x, y float64)
	MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool)
	Ok() bool
	OpenLayerPane()
	OutputAndClose(w io.WriteCloser) error
	OutputFileAndClose(fileStr string) error
	Output(w io.Writer) error
	PageCount() int
	PageNo() int
	PageSize(pageNum int) (wd, ht float64, unitStr string)
	PointConvert(pt float64) (u float64)
	PointToUnitConvert(pt float64) (u float64)
	Polygon(points []gofpdf.PointType, styleStr string)
	RadialGradient(x, y, w, h float64, r1, g1, b1, r2, g2, b2 int, x1, y1, x2, y2, r float64)
	RawWriteBuf(r io.Reader)
	RawWriteStr(str string)
	Rect(x, y, w, h float64, styleStr string)
	RegisterAlias(alias, replacement string)
	RegisterImage(fileStr, tp string) (info *gofpdf.ImageInfoType)
	RegisterImageOptions(fileStr string, options gofpdf.ImageOptions) (info *gofpdf.ImageInfoType)
	RegisterImageOptionsReader(imgName string, options gofpdf.ImageOptions, r io.Reader) (info *gofpdf.ImageInfoType)
	RegisterImageReader(imgName, tp string, r io.Reader) (info *gofpdf.ImageInfoType)
	SetAcceptPageBreakFunc(fnc func() bool)
	SetAlpha(alpha float64, blendModeStr string)
	SetAuthor(authorStr string, isUTF8 bool)
	SetAutoPageBreak(auto bool, margin float64)
	SetCatalogSort(flag bool)
	SetCellMargin(margin float64)
	SetCompression(compress bool)
	SetCreationDate(tm time.Time)
	SetCreator(creatorStr string, isUTF8 bool)
	SetDashPattern(dashArray []float64, dashPhase float64)
	SetDisplayMode(zoomStr, layoutStr string)
	SetDrawColor(r, g, b int)
	SetDrawSpotColor(nameStr string, tint byte)
	SetError(err error)
	SetErrorf(fmtStr string, args ...interface{})
	SetFillColor(r, g, b int)
	SetFillSpotColor(nameStr string, tint byte)
	SetFont(familyStr, styleStr string, size float64)
	SetFontLoader(loader gofpdf.FontLoader)
	SetFontLocation(fontDirStr string)
	SetFontSize(size float64)
	SetFontStyle(styleStr string)
	SetFontUnitSize(size float64)
	SetFooterFunc(fnc func())
	SetFooterFuncLpi(fnc func(lastPage bool))
	SetHeaderFunc(fnc func())
	SetHeaderFuncMode(fnc func(), homeMode bool)
	SetHomeXY()
	SetJavascript(script string)
	SetKeywords(keywordsStr string, isUTF8 bool)
	SetLeftMargin(margin float64)
	SetLineCapStyle(styleStr string)
	SetLineJoinStyle(styleStr string)
	SetLineWidth(width float64)
	SetLink(link int, y float64, page int)
	SetMargins(left, top, right float64)
	SetPageBoxRec(t string, pb gofpdf.PageBox)
	SetPageBox(t string, x, y, wd, ht float64)
	SetPage(pageNum int)
	SetProtection(actionFlag byte, userPassStr, ownerPassStr string)
	SetRightMargin(margin float64)
	SetSubject(subjectStr string, isUTF8 bool)
	SetTextColor(r, g, b int)
	SetTextSpotColor(nameStr string, tint byte)
	SetTitle(titleStr string, isUTF8 bool)
	SetTopMargin(margin float64)
	SetXmpMetadata(xmpStream []byte)
	SetX(x float64)
	SetXY(x, y float64)
	SetY(y float64)
	SplitLines(txt []byte, w float64) [][]byte
	String() string
	SVGBasicWrite(sb *gofpdf.SVGBasicType, scale float64)
	Text(x, y float64, txtStr string)
	TransformBegin()
	TransformEnd()
	TransformMirrorHorizontal(x float64)
	TransformMirrorLine(angle, x, y float64)
	TransformMirrorPoint(x, y float64)
	TransformMirrorVertical(y float64)
	TransformRotate(angle, x, y float64)
	TransformScale(scaleWd, scaleHt, x, y float64)
	TransformScaleX(scaleWd, x, y float64)
	TransformScaleXY(s, x, y float64)
	TransformScaleY(scaleHt, x, y float64)
	TransformSkew(angleX, angleY, x, y float64)
	TransformSkewX(angleX, x, y float64)
	TransformSkewY(angleY, x, y float64)
	Transform(tm gofpdf.TransformMatrix)
	TransformTranslate(tx, ty float64)
	TransformTranslateX(tx float64)
	TransformTranslateY(ty float64)
	UnicodeTranslatorFromDescriptor(cpStr string) (rep func(string) string)
	UnitToPointConvert(u float64) (pt float64)
	UseTemplateScaled(t gofpdf.Template, corner gofpdf.PointType, size gofpdf.SizeType)
	UseTemplate(t gofpdf.Template)
	WriteAligned(width, lineHeight float64, textStr, alignStr string)
	Writef(h float64, fmtStr string, args ...interface{})
	Write(h float64, txtStr string)
	WriteLinkID(h float64, displayStr string, linkID int)
	WriteLinkString(h float64, displayStr, targetStr string)
	AddUTF8Font(familyStr, styleStr, fileStr string)
}

type fpdf struct {
	Gofpdf *gofpdf.Fpdf
}

// NewWrapper is the constructor of the wrapper based on gofpdf interface.
func NewWrapper(pdf *gofpdf.Fpdf) *fpdf {
	return &fpdf{
		Gofpdf: pdf,
	}
}

func (s fpdf) AddFont(familyStr, styleStr, fileStr string) {
	s.Gofpdf.AddFont(familyStr, styleStr, fileStr)
}

func (s fpdf) AddUTF8Font(familyStr, styleStr, fileStr string) {
	s.Gofpdf.AddUTF8Font(familyStr, styleStr, fileStr)
}

func (s fpdf) AddFontFromBytes(familyStr, styleStr string, jsonFileBytes, zFileBytes []byte) {
	s.Gofpdf.AddFontFromBytes(familyStr, styleStr, jsonFileBytes, zFileBytes)
}

func (s fpdf) AddFontFromReader(familyStr, styleStr string, r io.Reader) {
	s.Gofpdf.AddFontFromReader(familyStr, styleStr, r)
}

func (s fpdf) AddLayer(name string, visible bool) (layerID int) {
	return s.Gofpdf.AddLayer(name, visible)
}

func (s fpdf) AddLink() int {
	return s.Gofpdf.AddLink()
}

func (s fpdf) AddPage() {
	s.Gofpdf.AddPage()
}

func (s fpdf) AddPageFormat(orientationStr string, size gofpdf.SizeType) {
	s.Gofpdf.AddPageFormat(orientationStr, size)
}

func (s fpdf) AddSpotColor(nameStr string, c, m, y, k byte) {
	s.Gofpdf.AddSpotColor(nameStr, c, m, y, k)
}

func (s fpdf) AliasNbPages(aliasStr string) {
	s.Gofpdf.AliasNbPages(aliasStr)
}

func (s fpdf) ArcTo(x, y, rx, ry, degRotate, degStart, degEnd float64) {
	s.Gofpdf.ArcTo(x, y, rx, ry, degRotate, degStart, degEnd)
}

func (s fpdf) Arc(x, y, rx, ry, degRotate, degStart, degEnd float64, styleStr string) {
	s.Gofpdf.Arc(x, y, rx, ry, degRotate, degStart, degEnd, styleStr)
}

func (s fpdf) BeginLayer(id int) {
	s.Gofpdf.BeginLayer(id)
}

func (s fpdf) Beziergon(points []gofpdf.PointType, styleStr string) {
	s.Gofpdf.Beziergon(points, styleStr)
}

func (s fpdf) Bookmark(txtStr string, level int, y float64) {
	s.Gofpdf.Bookmark(txtStr, level, y)
}

func (s fpdf) CellFormat(w, h float64, txtStr, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string) {
	s.Gofpdf.CellFormat(w, h, txtStr, borderStr, ln, alignStr, fill, link, linkStr)
}

func (s fpdf) Cellf(w, h float64, fmtStr string, args ...interface{}) {
	s.Gofpdf.Cellf(w, h, fmtStr, args...)
}

func (s fpdf) Cell(w, h float64, txtStr string) {
	s.Gofpdf.Cell(w, h, txtStr)
}

func (s fpdf) Circle(x, y, r float64, styleStr string) {
	s.Gofpdf.Circle(x, y, r, styleStr)
}

func (s fpdf) ClearError() {
	s.Gofpdf.ClearError()
}

func (s fpdf) ClipCircle(x, y, r float64, outline bool) {
	s.Gofpdf.ClipCircle(x, y, r, outline)
}

func (s fpdf) ClipEllipse(x, y, rx, ry float64, outline bool) {
	s.Gofpdf.ClipEllipse(x, y, rx, ry, outline)
}

func (s fpdf) ClipEnd() {
	s.Gofpdf.ClipEnd()
}

func (s fpdf) ClipPolygon(points []gofpdf.PointType, outline bool) {
	s.Gofpdf.ClipPolygon(points, outline)
}

func (s fpdf) ClipRect(x, y, w, h float64, outline bool) {
	s.Gofpdf.ClipRect(x, y, w, h, outline)
}

func (s fpdf) ClipRoundedRect(x, y, w, h, r float64, outline bool) {
	s.Gofpdf.ClipRoundedRect(x, y, w, h, r, outline)
}

func (s fpdf) ClipText(x, y float64, txtStr string, outline bool) {
	s.Gofpdf.ClipText(x, y, txtStr, outline)
}

func (s fpdf) Close() {
	s.Gofpdf.Close()
}

func (s fpdf) ClosePath() {
	s.Gofpdf.ClosePath()
}

func (s fpdf) CreateTemplateCustom(corner gofpdf.PointType, size gofpdf.SizeType, fn func(*gofpdf.Tpl)) gofpdf.Template {
	return s.Gofpdf.CreateTemplateCustom(corner, size, fn)
}

func (s fpdf) CreateTemplate(fn func(*gofpdf.Tpl)) gofpdf.Template {
	return s.Gofpdf.CreateTemplate(fn)
}

func (s fpdf) CurveBezierCubicTo(cx0, cy0, cx1, cy1, x, y float64) {
	s.Gofpdf.CurveBezierCubicTo(cx0, cy0, cx1, cy1, x, y)
}

func (s fpdf) CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1 float64, styleStr string) {
	s.Gofpdf.CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1, styleStr)
}

func (s fpdf) CurveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1 float64, styleStr string) {
	s.Gofpdf.CurveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1, styleStr)
}

func (s fpdf) CurveTo(cx, cy, x, y float64) {
	s.Gofpdf.CurveTo(cx, cy, x, y)
}

func (s fpdf) Curve(x0, y0, cx, cy, x1, y1 float64, styleStr string) {
	s.Gofpdf.Curve(x0, y0, cx, cy, x1, y1, styleStr)
}

func (s fpdf) DrawPath(styleStr string) {
	s.Gofpdf.DrawPath(styleStr)
}

func (s fpdf) Ellipse(x, y, rx, ry, degRotate float64, styleStr string) {
	s.Gofpdf.Ellipse(x, y, rx, ry, degRotate, styleStr)
}

func (s fpdf) EndLayer() {
	s.Gofpdf.EndLayer()
}

func (s fpdf) Err() bool {
	return s.Gofpdf.Err()
}

func (s fpdf) Error() error {
	return s.Gofpdf.Error()
}

func (s fpdf) GetAlpha() (alpha float64, blendModeStr string) {
	return s.Gofpdf.GetAlpha()
}

func (s fpdf) GetAutoPageBreak() (auto bool, margin float64) {
	return s.Gofpdf.GetAutoPageBreak()
}

func (s fpdf) GetCellMargin() float64 {
	return s.Gofpdf.GetCellMargin()
}

func (s fpdf) GetConversionRatio() float64 {
	return s.Gofpdf.GetConversionRatio()
}

func (s fpdf) GetDrawColor() (int, int, int) {
	return s.Gofpdf.GetDrawColor()
}

func (s fpdf) GetDrawSpotColor() (name string, c, m, y, k byte) {
	return s.Gofpdf.GetDrawSpotColor()
}

func (s fpdf) GetFillColor() (int, int, int) {
	return s.Gofpdf.GetFillColor()
}

func (s fpdf) GetFillSpotColor() (name string, c, m, y, k byte) {
	return s.Gofpdf.GetFillSpotColor()
}

func (s fpdf) GetFontDesc(familyStr, styleStr string) gofpdf.FontDescType {
	return s.Gofpdf.GetFontDesc(familyStr, styleStr)
}

func (s fpdf) GetFontSize() (ptSize, unitSize float64) {
	return s.Gofpdf.GetFontSize()
}

func (s fpdf) GetImageInfo(imageStr string) (info *gofpdf.ImageInfoType) {
	return s.Gofpdf.GetImageInfo(imageStr)
}

func (s fpdf) GetLineWidth() float64 {
	return s.Gofpdf.GetLineWidth()
}

func (s fpdf) GetMargins() (left, top, right, bottom float64) {
	return s.Gofpdf.GetMargins()
}

func (s fpdf) GetPageSizeStr(sizeStr string) (size gofpdf.SizeType) {
	return s.Gofpdf.GetPageSizeStr(sizeStr)
}

func (s fpdf) GetPageSize() (width, height float64) {
	return s.Gofpdf.GetPageSize()
}

func (s fpdf) GetStringWidth(stringWidth string) float64 {
	return s.Gofpdf.GetStringWidth(stringWidth)
}

func (s fpdf) GetTextColor() (int, int, int) {
	return s.Gofpdf.GetTextColor()
}

func (s fpdf) GetTextSpotColor() (name string, c, m, y, k byte) {
	return s.Gofpdf.GetTextSpotColor()
}

func (s fpdf) GetX() float64 {
	return s.Gofpdf.GetX()
}

func (s fpdf) GetXY() (float64, float64) {
	return s.Gofpdf.GetXY()
}

func (s fpdf) GetY() float64 {
	return s.Gofpdf.GetY()
}

func (s fpdf) HTMLBasicNew() (html gofpdf.HTMLBasicType) {
	return s.Gofpdf.HTMLBasicNew()
}

func (s fpdf) Image(imageNameStr string, x, y, w, h float64, flow bool, tp string, link int, linkStr string) {
	s.Gofpdf.Image(imageNameStr, x, y, w, h, flow, tp, link, linkStr)
}

func (s fpdf) ImageOptions(imageNameStr string, x, y, w, h float64, flow bool, options gofpdf.ImageOptions, link int, linkStr string) {
	s.Gofpdf.ImageOptions(imageNameStr, x, y, w, h, flow, options, link, linkStr)
}

func (s fpdf) ImageTypeFromMime(mimeStr string) (tp string) {
	return s.Gofpdf.ImageTypeFromMime(mimeStr)
}

func (s fpdf) LinearGradient(x, y, w, h float64, r1, g1, b1, r2, g2, b2 int, x1, y1, x2, y2 float64) {
	s.Gofpdf.LinearGradient(x, y, w, h, r1, g1, b1, r2, g2, b2, x1, y1, x2, y2)
}

func (s fpdf) LineTo(x, y float64) {
	s.Gofpdf.LineTo(x, y)
}

func (s fpdf) Line(x1, y1, x2, y2 float64) {
	s.Gofpdf.Line(x1, y1, x2, y2)
}

func (s fpdf) LinkString(x, y, w, h float64, linkStr string) {
	s.Gofpdf.LinkString(x, y, w, h, linkStr)
}

func (s fpdf) Link(x, y, w, h float64, link int) {
	s.Gofpdf.Link(x, y, w, h, link)
}

func (s fpdf) Ln(h float64) {
	s.Gofpdf.Ln(h)
}

func (s fpdf) MoveTo(x, y float64) {
	s.Gofpdf.MoveTo(x, y)
}

func (s fpdf) MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool) {
	s.Gofpdf.MultiCell(w, h, txtStr, borderStr, alignStr, fill)
}

func (s fpdf) Ok() bool {
	return s.Gofpdf.Ok()
}

func (s fpdf) OpenLayerPane() {
	s.Gofpdf.OpenLayerPane()
}

func (s fpdf) OutputAndClose(w io.WriteCloser) error {
	return s.Gofpdf.OutputAndClose(w)
}

func (s fpdf) OutputFileAndClose(fileStr string) error {
	return s.Gofpdf.OutputFileAndClose(fileStr)
}

func (s fpdf) Output(w io.Writer) error {
	return s.Gofpdf.Output(w)
}

func (s fpdf) PageCount() int {
	return s.Gofpdf.PageCount()
}

func (s fpdf) PageNo() int {
	return s.Gofpdf.PageNo()
}

func (s fpdf) PageSize(pageNum int) (wd, ht float64, unitStr string) {
	return s.Gofpdf.PageSize(pageNum)
}

func (s fpdf) PointConvert(pt float64) (u float64) {
	return s.Gofpdf.PointConvert(pt)
}

func (s fpdf) PointToUnitConvert(pt float64) (u float64) {
	return s.Gofpdf.PointToUnitConvert(pt)
}

func (s fpdf) Polygon(points []gofpdf.PointType, styleStr string) {
	s.Gofpdf.Polygon(points, styleStr)
}

func (s fpdf) RadialGradient(x, y, w, h float64, r1, g1, b1, r2, g2, b2 int, x1, y1, x2, y2, r float64) {
	s.Gofpdf.RadialGradient(x, y, w, h, r1, g1, b1, r2, g2, b2, x1, y1, x2, y2, r)
}

func (s fpdf) RawWriteBuf(r io.Reader) {
	s.Gofpdf.RawWriteBuf(r)
}

func (s fpdf) RawWriteStr(str string) {
	s.Gofpdf.RawWriteStr(str)
}

func (s fpdf) Rect(x, y, w, h float64, styleStr string) {
	s.Gofpdf.Rect(x, y, w, h, styleStr)
}

func (s fpdf) RegisterAlias(alias, replacement string) {
	s.Gofpdf.RegisterAlias(alias, replacement)
}

func (s fpdf) RegisterImage(fileStr, tp string) (info *gofpdf.ImageInfoType) {
	return s.Gofpdf.RegisterImage(fileStr, tp)
}

func (s fpdf) RegisterImageOptions(fileStr string, options gofpdf.ImageOptions) (info *gofpdf.ImageInfoType) {
	return s.Gofpdf.RegisterImageOptions(fileStr, options)
}

func (s fpdf) RegisterImageOptionsReader(imgName string, options gofpdf.ImageOptions, r io.Reader) (info *gofpdf.ImageInfoType) {
	return s.Gofpdf.RegisterImageOptionsReader(imgName, options, r)
}

func (s fpdf) RegisterImageReader(imgName, tp string, r io.Reader) (info *gofpdf.ImageInfoType) {
	return s.Gofpdf.RegisterImageReader(imgName, tp, r)
}

func (s fpdf) SetAcceptPageBreakFunc(fnc func() bool) {
	s.Gofpdf.SetAcceptPageBreakFunc(fnc)
}

func (s fpdf) SetAlpha(alpha float64, blendModeStr string) {
	s.Gofpdf.SetAlpha(alpha, blendModeStr)
}

func (s fpdf) SetAuthor(authorStr string, isUTF8 bool) {
	s.Gofpdf.SetAuthor(authorStr, isUTF8)
}

func (s fpdf) SetAutoPageBreak(auto bool, margin float64) {
	s.Gofpdf.SetAutoPageBreak(auto, margin)
}

func (s fpdf) SetCatalogSort(flag bool) {
	s.Gofpdf.SetCatalogSort(flag)
}

func (s fpdf) SetCellMargin(margin float64) {
	s.Gofpdf.SetCellMargin(margin)
}

func (s fpdf) SetCompression(compress bool) {
	s.Gofpdf.SetCompression(compress)
}

func (s fpdf) SetCreationDate(tm time.Time) {
	s.Gofpdf.SetCreationDate(tm)
}

func (s fpdf) SetCreator(creatorStr string, isUTF8 bool) {
	s.Gofpdf.SetCreator(creatorStr, isUTF8)
}

func (s fpdf) SetDashPattern(dashArray []float64, dashPhase float64) {
	s.Gofpdf.SetDashPattern(dashArray, dashPhase)
}

func (s fpdf) SetDisplayMode(zoomStr, layoutStr string) {
	s.Gofpdf.SetDisplayMode(zoomStr, layoutStr)
}

func (s fpdf) SetDrawColor(r, g, b int) {
	s.Gofpdf.SetDrawColor(r, g, b)
}

func (s fpdf) SetDrawSpotColor(nameStr string, tint byte) {
	s.Gofpdf.SetDrawSpotColor(nameStr, tint)
}

func (s fpdf) SetError(err error) {
	s.Gofpdf.SetError(err)
}

func (s fpdf) SetErrorf(fmtStr string, args ...interface{}) {
	s.Gofpdf.SetErrorf(fmtStr, args...)
}

func (s fpdf) SetFillColor(r, g, b int) {
	s.Gofpdf.SetFillColor(r, g, b)
}

func (s fpdf) SetFillSpotColor(nameStr string, tint byte) {
	s.Gofpdf.SetFillSpotColor(nameStr, tint)
}

func (s fpdf) SetFont(familyStr, styleStr string, size float64) {
	s.Gofpdf.SetFont(familyStr, styleStr, size)
}

func (s fpdf) SetFontLoader(loader gofpdf.FontLoader) {
	s.Gofpdf.SetFontLoader(loader)
}

func (s fpdf) SetFontLocation(fontDirStr string) {
	s.Gofpdf.SetFontLocation(fontDirStr)
}

func (s fpdf) SetFontSize(size float64) {
	s.Gofpdf.SetFontSize(size)
}

func (s fpdf) SetFontStyle(styleStr string) {
	s.Gofpdf.SetFontStyle(styleStr)
}

func (s fpdf) SetFontUnitSize(size float64) {
	s.Gofpdf.SetFontUnitSize(size)
}

func (s fpdf) SetFooterFunc(fnc func()) {
	s.Gofpdf.SetFooterFunc(fnc)
}

func (s fpdf) SetFooterFuncLpi(fnc func(lastPage bool)) {
	s.Gofpdf.SetFooterFuncLpi(fnc)
}

func (s fpdf) SetHeaderFunc(fnc func()) {
	s.Gofpdf.SetHeaderFunc(fnc)
}

func (s fpdf) SetHeaderFuncMode(fnc func(), homeMode bool) {
	s.Gofpdf.SetHeaderFuncMode(fnc, homeMode)
}

func (s fpdf) SetHomeXY() {
	s.Gofpdf.SetHomeXY()
}

func (s fpdf) SetJavascript(script string) {
	s.Gofpdf.SetJavascript(script)
}

func (s fpdf) SetKeywords(keywordsStr string, isUTF8 bool) {
	s.Gofpdf.SetKeywords(keywordsStr, isUTF8)
}

func (s fpdf) SetLeftMargin(margin float64) {
	s.Gofpdf.SetLeftMargin(margin)
}

func (s fpdf) SetLineCapStyle(styleStr string) {
	s.Gofpdf.SetLineCapStyle(styleStr)
}

func (s fpdf) SetLineJoinStyle(styleStr string) {
	s.Gofpdf.SetLineJoinStyle(styleStr)
}

func (s fpdf) SetLineWidth(width float64) {
	s.Gofpdf.SetLineWidth(width)
}

func (s fpdf) SetLink(link int, y float64, page int) {
	s.Gofpdf.SetLink(link, y, page)
}

func (s fpdf) SetMargins(left, top, right float64) {
	s.Gofpdf.SetMargins(left, top, right)
}

func (s fpdf) SetPageBoxRec(t string, pb gofpdf.PageBox) {
	s.Gofpdf.SetPageBoxRec(t, pb)
}

func (s fpdf) SetPageBox(t string, x, y, wd, ht float64) {
	s.Gofpdf.SetPageBox(t, x, y, wd, ht)
}

func (s fpdf) SetPage(pageNum int) {
	s.Gofpdf.SetPage(pageNum)
}

func (s fpdf) SetProtection(actionFlag byte, userPassStr, ownerPassStr string) {
	s.Gofpdf.SetProtection(actionFlag, userPassStr, ownerPassStr)
}

func (s fpdf) SetRightMargin(margin float64) {
	s.Gofpdf.SetRightMargin(margin)
}

func (s fpdf) SetSubject(subjectStr string, isUTF8 bool) {
	s.Gofpdf.SetSubject(subjectStr, isUTF8)
}

func (s fpdf) SetTextColor(r, g, b int) {
	s.Gofpdf.SetTextColor(r, g, b)
}

func (s fpdf) SetTextSpotColor(nameStr string, tint byte) {
	s.Gofpdf.SetTextSpotColor(nameStr, tint)
}

func (s fpdf) SetTitle(titleStr string, isUTF8 bool) {
	s.Gofpdf.SetTitle(titleStr, isUTF8)
}

func (s fpdf) SetTopMargin(margin float64) {
	s.Gofpdf.SetTopMargin(margin)
}

func (s fpdf) SetXmpMetadata(xmpStream []byte) {
	s.Gofpdf.SetXmpMetadata(xmpStream)
}

func (s fpdf) SetX(x float64) {
	s.Gofpdf.SetX(x)
}

func (s fpdf) SetXY(x, y float64) {
	s.Gofpdf.SetXY(x, y)
}

func (s fpdf) SetY(y float64) {
	s.Gofpdf.SetY(y)
}

func (s fpdf) SplitLines(txt []byte, w float64) [][]byte {
	return s.Gofpdf.SplitLines(txt, w)
}

func (s fpdf) String() string {
	return s.Gofpdf.String()
}

func (s fpdf) SVGBasicWrite(sb *gofpdf.SVGBasicType, scale float64) {
	s.Gofpdf.SVGBasicWrite(sb, scale)
}

func (s fpdf) Text(x, y float64, txtStr string) {
	s.Gofpdf.Text(x, y, txtStr)
}

func (s fpdf) TransformBegin() {
	s.Gofpdf.TransformBegin()
}

func (s fpdf) TransformEnd() {
	s.Gofpdf.TransformEnd()
}

func (s fpdf) TransformMirrorHorizontal(x float64) {
	s.Gofpdf.TransformMirrorHorizontal(x)
}

func (s fpdf) TransformMirrorLine(angle, x, y float64) {
	s.Gofpdf.TransformMirrorLine(angle, x, y)
}

func (s fpdf) TransformMirrorPoint(x, y float64) {
	s.Gofpdf.TransformMirrorPoint(x, y)
}

func (s fpdf) TransformMirrorVertical(y float64) {
	s.Gofpdf.TransformMirrorVertical(y)
}

func (s fpdf) TransformRotate(angle, x, y float64) {
	s.Gofpdf.TransformRotate(angle, x, y)
}

func (s fpdf) TransformScale(scaleWd, scaleHt, x, y float64) {
	s.Gofpdf.TransformScale(scaleWd, scaleHt, x, y)
}

func (s fpdf) TransformScaleX(scaleWd, x, y float64) {
	s.Gofpdf.TransformScaleX(scaleWd, x, y)
}

func (s fpdf) TransformScaleXY(scale, x, y float64) {
	s.Gofpdf.TransformScaleXY(scale, x, y)
}

func (s fpdf) TransformScaleY(scaleHt, x, y float64) {
	s.Gofpdf.TransformScaleY(scaleHt, x, y)
}

func (s fpdf) TransformSkew(angleX, angleY, x, y float64) {
	s.Gofpdf.TransformSkew(angleX, angleY, x, y)
}

func (s fpdf) TransformSkewX(angleX, x, y float64) {
	s.Gofpdf.TransformSkewY(angleX, x, y)
}

func (s fpdf) TransformSkewY(angleY, x, y float64) {
	s.Gofpdf.TransformSkewY(angleY, x, y)
}

func (s fpdf) Transform(tm gofpdf.TransformMatrix) {
	s.Gofpdf.Transform(tm)
}

func (s fpdf) TransformTranslate(tx, ty float64) {
	s.Gofpdf.TransformTranslate(tx, ty)
}

func (s fpdf) TransformTranslateX(tx float64) {
	s.Gofpdf.TransformTranslateX(tx)
}

func (s fpdf) TransformTranslateY(ty float64) {
	s.Gofpdf.TransformTranslateY(ty)
}

func (s fpdf) UnicodeTranslatorFromDescriptor(cpStr string) (rep func(string) string) {
	return s.Gofpdf.UnicodeTranslatorFromDescriptor(cpStr)
}

func (s fpdf) UnitToPointConvert(u float64) (pt float64) {
	return s.Gofpdf.UnitToPointConvert(u)
}

func (s fpdf) UseTemplateScaled(t gofpdf.Template, corner gofpdf.PointType, size gofpdf.SizeType) {
	s.Gofpdf.UseTemplateScaled(t, corner, size)
}

func (s fpdf) UseTemplate(t gofpdf.Template) {
	s.Gofpdf.UseTemplate(t)
}

func (s fpdf) WriteAligned(width, lineHeight float64, textStr, alignStr string) {
	s.Gofpdf.WriteAligned(width, lineHeight, textStr, alignStr)
}

func (s fpdf) Writef(h float64, fmtStr string, args ...interface{}) {
	s.Gofpdf.Writef(h, fmtStr, args...)
}

func (s fpdf) Write(h float64, txtStr string) {
	s.Gofpdf.Write(h, txtStr)
}

func (s fpdf) WriteLinkID(h float64, displayStr string, linkID int) {
	s.Gofpdf.WriteLinkID(h, displayStr, linkID)
}

func (s fpdf) WriteLinkString(h float64, displayStr, targetStr string) {
	s.Gofpdf.WriteLinkString(h, displayStr, targetStr)
}
