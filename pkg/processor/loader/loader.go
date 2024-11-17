package loader

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type loader struct{}

func NewLoader() *loader {
	return &loader{}
}

// Load takes the path/url/uri of an asset (image, font)
// and returns its contents.
func (l *loader) Load(path string) ([]byte, error) {
	ext := l.GetExt(path)
	if _, ok := validExts[ext]; !ok {
		return nil, errors.Wrap(ErrUnsupportedExtension, ext)
	}

	uri, err := url.Parse(path)
	if err != nil {
		return nil, errors.Wrap(ErrInvalidPath, path)
	}

	loadFn, ok := loadFuncs[uri.Scheme]
	if !ok {
		return nil, errors.Wrap(ErrUnsupportedProtocol, uri.Scheme)
	}

	r, err := loadFn(uri.String())
	if err != nil {
		return nil, errors.Wrap(ErrAccessFail, err.Error())
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(ErrReadFail, err.Error())
	}

	return data, nil
}

func (l *loader) GetExt(path string) string {
	toks := strings.Split(path, ".")
	if len(toks) < 2 {
		return ""
	}
	return toks[len(toks)-1]
}

var validExts = map[string]struct{}{
	"png":  {},
	"jpg":  {},
	"svg":  {},
	"jpeg": {},
	"ttf":  {},
}

var loadFuncs = map[string]func(string) (io.ReadCloser, error){
	"http":  loadHTTP,
	"https": loadHTTP,
	"file":  loadLocal,
	"":      loadLocal,
}

func loadLocal(path string) (io.ReadCloser, error) {
	path = strings.TrimPrefix(path, "file://")
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(absolutePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func loadHTTP(path string) (io.ReadCloser, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

var (
	ErrUnsupportedProtocol  = errors.New("unsupported protocol")
	ErrUnsupportedExtension = errors.New("unsupported extension")
	ErrInvalidPath          = errors.New("invalid path")
	ErrAccessFail           = errors.New("failed to access asset")
	ErrReadFail             = errors.New("failed to read asset")
)
