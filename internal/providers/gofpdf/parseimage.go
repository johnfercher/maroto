package gofpdf

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

var (
	ErrInvalidImageFormat     = errors.New("invalid image format")
	ErrCannotDecode16BitImage = errors.New("cannot decode 16 bit image")
	ErrCannotEncodeImage      = errors.New("cannot encode image")
)

func FromBytes(bytes []byte, ext extension.Type) (*entity.Image, error) {
	if !ext.IsValid() {
		return nil, ErrInvalidImageFormat
	}

	data := bytes
	if ext == extension.Png {
		if depth, ok := pngBitDepth(bytes); ok && depth == 16 {
			converted, err := normalize16BitPNG(bytes)
			if err != nil {
				return nil, err
			}
			data = converted
		}
	}

	return &entity.Image{
		Bytes:     data,
		Extension: ext,
	}, nil
}

func pngBitDepth(data []byte) (int, bool) {
	const pngSignature = "\x89PNG\r\n\x1a\n"
	if len(data) < 26 || string(data[:8]) != pngSignature {
		return 0, false
	}
	if string(data[12:16]) != "IHDR" {
		return 0, false
	}
	return int(data[24]), true
}

// normalize16BitPNG re-encodes a 16-bit PNG as 8-bit NRGBA. gofpdf rejects 16-bit samples.
func normalize16BitPNG(data []byte) ([]byte, error) {
	src, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return data, fmt.Errorf("%w: %w", ErrCannotDecode16BitImage, err)
	}

	b := src.Bounds()
	dst := image.NewNRGBA(b)
	draw.Draw(dst, b, src, b.Min, draw.Src)

	var buf bytes.Buffer
	err = png.Encode(&buf, dst)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotEncodeImage, err)
	}
	return buf.Bytes(), nil
}
