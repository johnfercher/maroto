// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	extension "github.com/johnfercher/maroto/v2/pkg/consts/extension"
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	props "github.com/johnfercher/maroto/v2/pkg/props"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

type Provider_Expecter struct {
	mock *mock.Mock
}

func (_m *Provider) EXPECT() *Provider_Expecter {
	return &Provider_Expecter{mock: &_m.Mock}
}

// AddBackgroundImageFromBytes provides a mock function with given fields: bytes, cell, prop, _a3
func (_m *Provider) AddBackgroundImageFromBytes(bytes []byte, cell *entity.Cell, prop *props.Rect, _a3 extension.Type) {
	_m.Called(bytes, cell, prop, _a3)
}

// Provider_AddBackgroundImageFromBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBackgroundImageFromBytes'
type Provider_AddBackgroundImageFromBytes_Call struct {
	*mock.Call
}

// AddBackgroundImageFromBytes is a helper method to define mock.On call
//   - bytes []byte
//   - cell *entity.Cell
//   - prop *props.Rect
//   - _a3 extension.Type
func (_e *Provider_Expecter) AddBackgroundImageFromBytes(bytes interface{}, cell interface{}, prop interface{}, _a3 interface{}) *Provider_AddBackgroundImageFromBytes_Call {
	return &Provider_AddBackgroundImageFromBytes_Call{Call: _e.mock.On("AddBackgroundImageFromBytes", bytes, cell, prop, _a3)}
}

func (_c *Provider_AddBackgroundImageFromBytes_Call) Run(run func(bytes []byte, cell *entity.Cell, prop *props.Rect, _a3 extension.Type)) *Provider_AddBackgroundImageFromBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(*entity.Cell), args[2].(*props.Rect), args[3].(extension.Type))
	})
	return _c
}

func (_c *Provider_AddBackgroundImageFromBytes_Call) Return() *Provider_AddBackgroundImageFromBytes_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddBackgroundImageFromBytes_Call) RunAndReturn(run func([]byte, *entity.Cell, *props.Rect, extension.Type)) *Provider_AddBackgroundImageFromBytes_Call {
	_c.Call.Return(run)
	return _c
}

// AddBarCode provides a mock function with given fields: code, cell, prop
func (_m *Provider) AddBarCode(code string, cell *entity.Cell, prop *props.Barcode) {
	_m.Called(code, cell, prop)
}

// Provider_AddBarCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBarCode'
type Provider_AddBarCode_Call struct {
	*mock.Call
}

// AddBarCode is a helper method to define mock.On call
//   - code string
//   - cell *entity.Cell
//   - prop *props.Barcode
func (_e *Provider_Expecter) AddBarCode(code interface{}, cell interface{}, prop interface{}) *Provider_AddBarCode_Call {
	return &Provider_AddBarCode_Call{Call: _e.mock.On("AddBarCode", code, cell, prop)}
}

func (_c *Provider_AddBarCode_Call) Run(run func(code string, cell *entity.Cell, prop *props.Barcode)) *Provider_AddBarCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entity.Cell), args[2].(*props.Barcode))
	})
	return _c
}

func (_c *Provider_AddBarCode_Call) Return() *Provider_AddBarCode_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddBarCode_Call) RunAndReturn(run func(string, *entity.Cell, *props.Barcode)) *Provider_AddBarCode_Call {
	_c.Call.Return(run)
	return _c
}

// AddImageFromBytes provides a mock function with given fields: bytes, cell, prop, _a3
func (_m *Provider) AddImageFromBytes(bytes []byte, cell *entity.Cell, prop *props.Rect, _a3 extension.Type) {
	_m.Called(bytes, cell, prop, _a3)
}

// Provider_AddImageFromBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddImageFromBytes'
type Provider_AddImageFromBytes_Call struct {
	*mock.Call
}

// AddImageFromBytes is a helper method to define mock.On call
//   - bytes []byte
//   - cell *entity.Cell
//   - prop *props.Rect
//   - _a3 extension.Type
func (_e *Provider_Expecter) AddImageFromBytes(bytes interface{}, cell interface{}, prop interface{}, _a3 interface{}) *Provider_AddImageFromBytes_Call {
	return &Provider_AddImageFromBytes_Call{Call: _e.mock.On("AddImageFromBytes", bytes, cell, prop, _a3)}
}

func (_c *Provider_AddImageFromBytes_Call) Run(run func(bytes []byte, cell *entity.Cell, prop *props.Rect, _a3 extension.Type)) *Provider_AddImageFromBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(*entity.Cell), args[2].(*props.Rect), args[3].(extension.Type))
	})
	return _c
}

func (_c *Provider_AddImageFromBytes_Call) Return() *Provider_AddImageFromBytes_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddImageFromBytes_Call) RunAndReturn(run func([]byte, *entity.Cell, *props.Rect, extension.Type)) *Provider_AddImageFromBytes_Call {
	_c.Call.Return(run)
	return _c
}

// AddImageFromFile provides a mock function with given fields: value, cell, prop
func (_m *Provider) AddImageFromFile(value string, cell *entity.Cell, prop *props.Rect) {
	_m.Called(value, cell, prop)
}

// Provider_AddImageFromFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddImageFromFile'
type Provider_AddImageFromFile_Call struct {
	*mock.Call
}

// AddImageFromFile is a helper method to define mock.On call
//   - value string
//   - cell *entity.Cell
//   - prop *props.Rect
func (_e *Provider_Expecter) AddImageFromFile(value interface{}, cell interface{}, prop interface{}) *Provider_AddImageFromFile_Call {
	return &Provider_AddImageFromFile_Call{Call: _e.mock.On("AddImageFromFile", value, cell, prop)}
}

func (_c *Provider_AddImageFromFile_Call) Run(run func(value string, cell *entity.Cell, prop *props.Rect)) *Provider_AddImageFromFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entity.Cell), args[2].(*props.Rect))
	})
	return _c
}

func (_c *Provider_AddImageFromFile_Call) Return() *Provider_AddImageFromFile_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddImageFromFile_Call) RunAndReturn(run func(string, *entity.Cell, *props.Rect)) *Provider_AddImageFromFile_Call {
	_c.Call.Return(run)
	return _c
}

// AddLine provides a mock function with given fields: cell, prop
func (_m *Provider) AddLine(cell *entity.Cell, prop *props.Line) {
	_m.Called(cell, prop)
}

// Provider_AddLine_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddLine'
type Provider_AddLine_Call struct {
	*mock.Call
}

// AddLine is a helper method to define mock.On call
//   - cell *entity.Cell
//   - prop *props.Line
func (_e *Provider_Expecter) AddLine(cell interface{}, prop interface{}) *Provider_AddLine_Call {
	return &Provider_AddLine_Call{Call: _e.mock.On("AddLine", cell, prop)}
}

func (_c *Provider_AddLine_Call) Run(run func(cell *entity.Cell, prop *props.Line)) *Provider_AddLine_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Cell), args[1].(*props.Line))
	})
	return _c
}

func (_c *Provider_AddLine_Call) Return() *Provider_AddLine_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddLine_Call) RunAndReturn(run func(*entity.Cell, *props.Line)) *Provider_AddLine_Call {
	_c.Call.Return(run)
	return _c
}

// AddMatrixCode provides a mock function with given fields: code, cell, prop
func (_m *Provider) AddMatrixCode(code string, cell *entity.Cell, prop *props.Rect) {
	_m.Called(code, cell, prop)
}

// Provider_AddMatrixCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddMatrixCode'
type Provider_AddMatrixCode_Call struct {
	*mock.Call
}

// AddMatrixCode is a helper method to define mock.On call
//   - code string
//   - cell *entity.Cell
//   - prop *props.Rect
func (_e *Provider_Expecter) AddMatrixCode(code interface{}, cell interface{}, prop interface{}) *Provider_AddMatrixCode_Call {
	return &Provider_AddMatrixCode_Call{Call: _e.mock.On("AddMatrixCode", code, cell, prop)}
}

func (_c *Provider_AddMatrixCode_Call) Run(run func(code string, cell *entity.Cell, prop *props.Rect)) *Provider_AddMatrixCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entity.Cell), args[2].(*props.Rect))
	})
	return _c
}

func (_c *Provider_AddMatrixCode_Call) Return() *Provider_AddMatrixCode_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddMatrixCode_Call) RunAndReturn(run func(string, *entity.Cell, *props.Rect)) *Provider_AddMatrixCode_Call {
	_c.Call.Return(run)
	return _c
}

// AddPageNumber provides a mock function with given fields: current, total, pg, cell
func (_m *Provider) AddPageNumber(current int, total int, pg *props.PageNumber, cell *entity.Cell) {
	_m.Called(current, total, pg, cell)
}

// Provider_AddPageNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPageNumber'
type Provider_AddPageNumber_Call struct {
	*mock.Call
}

// AddPageNumber is a helper method to define mock.On call
//   - current int
//   - total int
//   - pg *props.PageNumber
//   - cell *entity.Cell
func (_e *Provider_Expecter) AddPageNumber(current interface{}, total interface{}, pg interface{}, cell interface{}) *Provider_AddPageNumber_Call {
	return &Provider_AddPageNumber_Call{Call: _e.mock.On("AddPageNumber", current, total, pg, cell)}
}

func (_c *Provider_AddPageNumber_Call) Run(run func(current int, total int, pg *props.PageNumber, cell *entity.Cell)) *Provider_AddPageNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].(*props.PageNumber), args[3].(*entity.Cell))
	})
	return _c
}

func (_c *Provider_AddPageNumber_Call) Return() *Provider_AddPageNumber_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddPageNumber_Call) RunAndReturn(run func(int, int, *props.PageNumber, *entity.Cell)) *Provider_AddPageNumber_Call {
	_c.Call.Return(run)
	return _c
}

// AddQrCode provides a mock function with given fields: code, cell, rect
func (_m *Provider) AddQrCode(code string, cell *entity.Cell, rect *props.Rect) {
	_m.Called(code, cell, rect)
}

// Provider_AddQrCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddQrCode'
type Provider_AddQrCode_Call struct {
	*mock.Call
}

// AddQrCode is a helper method to define mock.On call
//   - code string
//   - cell *entity.Cell
//   - rect *props.Rect
func (_e *Provider_Expecter) AddQrCode(code interface{}, cell interface{}, rect interface{}) *Provider_AddQrCode_Call {
	return &Provider_AddQrCode_Call{Call: _e.mock.On("AddQrCode", code, cell, rect)}
}

func (_c *Provider_AddQrCode_Call) Run(run func(code string, cell *entity.Cell, rect *props.Rect)) *Provider_AddQrCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entity.Cell), args[2].(*props.Rect))
	})
	return _c
}

func (_c *Provider_AddQrCode_Call) Return() *Provider_AddQrCode_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddQrCode_Call) RunAndReturn(run func(string, *entity.Cell, *props.Rect)) *Provider_AddQrCode_Call {
	_c.Call.Return(run)
	return _c
}

// AddText provides a mock function with given fields: text, cell, prop
func (_m *Provider) AddText(text string, cell *entity.Cell, prop *props.Text) {
	_m.Called(text, cell, prop)
}

// Provider_AddText_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddText'
type Provider_AddText_Call struct {
	*mock.Call
}

// AddText is a helper method to define mock.On call
//   - text string
//   - cell *entity.Cell
//   - prop *props.Text
func (_e *Provider_Expecter) AddText(text interface{}, cell interface{}, prop interface{}) *Provider_AddText_Call {
	return &Provider_AddText_Call{Call: _e.mock.On("AddText", text, cell, prop)}
}

func (_c *Provider_AddText_Call) Run(run func(text string, cell *entity.Cell, prop *props.Text)) *Provider_AddText_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entity.Cell), args[2].(*props.Text))
	})
	return _c
}

func (_c *Provider_AddText_Call) Return() *Provider_AddText_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_AddText_Call) RunAndReturn(run func(string, *entity.Cell, *props.Text)) *Provider_AddText_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCol provides a mock function with given fields: width, height, config, prop
func (_m *Provider) CreateCol(width float64, height float64, config *entity.Config, prop *props.Cell) {
	_m.Called(width, height, config, prop)
}

// Provider_CreateCol_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCol'
type Provider_CreateCol_Call struct {
	*mock.Call
}

// CreateCol is a helper method to define mock.On call
//   - width float64
//   - height float64
//   - config *entity.Config
//   - prop *props.Cell
func (_e *Provider_Expecter) CreateCol(width interface{}, height interface{}, config interface{}, prop interface{}) *Provider_CreateCol_Call {
	return &Provider_CreateCol_Call{Call: _e.mock.On("CreateCol", width, height, config, prop)}
}

func (_c *Provider_CreateCol_Call) Run(run func(width float64, height float64, config *entity.Config, prop *props.Cell)) *Provider_CreateCol_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64), args[1].(float64), args[2].(*entity.Config), args[3].(*props.Cell))
	})
	return _c
}

func (_c *Provider_CreateCol_Call) Return() *Provider_CreateCol_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_CreateCol_Call) RunAndReturn(run func(float64, float64, *entity.Config, *props.Cell)) *Provider_CreateCol_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRow provides a mock function with given fields: height
func (_m *Provider) CreateRow(height float64) {
	_m.Called(height)
}

// Provider_CreateRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRow'
type Provider_CreateRow_Call struct {
	*mock.Call
}

// CreateRow is a helper method to define mock.On call
//   - height float64
func (_e *Provider_Expecter) CreateRow(height interface{}) *Provider_CreateRow_Call {
	return &Provider_CreateRow_Call{Call: _e.mock.On("CreateRow", height)}
}

func (_c *Provider_CreateRow_Call) Run(run func(height float64)) *Provider_CreateRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *Provider_CreateRow_Call) Return() *Provider_CreateRow_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_CreateRow_Call) RunAndReturn(run func(float64)) *Provider_CreateRow_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateBytes provides a mock function with given fields:
func (_m *Provider) GenerateBytes() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenerateBytes")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Provider_GenerateBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateBytes'
type Provider_GenerateBytes_Call struct {
	*mock.Call
}

// GenerateBytes is a helper method to define mock.On call
func (_e *Provider_Expecter) GenerateBytes() *Provider_GenerateBytes_Call {
	return &Provider_GenerateBytes_Call{Call: _e.mock.On("GenerateBytes")}
}

func (_c *Provider_GenerateBytes_Call) Run(run func()) *Provider_GenerateBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Provider_GenerateBytes_Call) Return(_a0 []byte, _a1 error) *Provider_GenerateBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Provider_GenerateBytes_Call) RunAndReturn(run func() ([]byte, error)) *Provider_GenerateBytes_Call {
	_c.Call.Return(run)
	return _c
}

// GetDimensionsByImage provides a mock function with given fields: file
func (_m *Provider) GetDimensionsByImage(file string) (*entity.Dimensions, error) {
	ret := _m.Called(file)

	if len(ret) == 0 {
		panic("no return value specified for GetDimensionsByImage")
	}

	var r0 *entity.Dimensions
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Dimensions, error)); ok {
		return rf(file)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Dimensions); ok {
		r0 = rf(file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Dimensions)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Provider_GetDimensionsByImage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDimensionsByImage'
type Provider_GetDimensionsByImage_Call struct {
	*mock.Call
}

// GetDimensionsByImage is a helper method to define mock.On call
//   - file string
func (_e *Provider_Expecter) GetDimensionsByImage(file interface{}) *Provider_GetDimensionsByImage_Call {
	return &Provider_GetDimensionsByImage_Call{Call: _e.mock.On("GetDimensionsByImage", file)}
}

func (_c *Provider_GetDimensionsByImage_Call) Run(run func(file string)) *Provider_GetDimensionsByImage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Provider_GetDimensionsByImage_Call) Return(_a0 *entity.Dimensions, _a1 error) *Provider_GetDimensionsByImage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Provider_GetDimensionsByImage_Call) RunAndReturn(run func(string) (*entity.Dimensions, error)) *Provider_GetDimensionsByImage_Call {
	_c.Call.Return(run)
	return _c
}

// GetDimensionsByImageByte provides a mock function with given fields: bytes, _a1
func (_m *Provider) GetDimensionsByImageByte(bytes []byte, _a1 extension.Type) (*entity.Dimensions, error) {
	ret := _m.Called(bytes, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetDimensionsByImageByte")
	}

	var r0 *entity.Dimensions
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, extension.Type) (*entity.Dimensions, error)); ok {
		return rf(bytes, _a1)
	}
	if rf, ok := ret.Get(0).(func([]byte, extension.Type) *entity.Dimensions); ok {
		r0 = rf(bytes, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Dimensions)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, extension.Type) error); ok {
		r1 = rf(bytes, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Provider_GetDimensionsByImageByte_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDimensionsByImageByte'
type Provider_GetDimensionsByImageByte_Call struct {
	*mock.Call
}

// GetDimensionsByImageByte is a helper method to define mock.On call
//   - bytes []byte
//   - _a1 extension.Type
func (_e *Provider_Expecter) GetDimensionsByImageByte(bytes interface{}, _a1 interface{}) *Provider_GetDimensionsByImageByte_Call {
	return &Provider_GetDimensionsByImageByte_Call{Call: _e.mock.On("GetDimensionsByImageByte", bytes, _a1)}
}

func (_c *Provider_GetDimensionsByImageByte_Call) Run(run func(bytes []byte, _a1 extension.Type)) *Provider_GetDimensionsByImageByte_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(extension.Type))
	})
	return _c
}

func (_c *Provider_GetDimensionsByImageByte_Call) Return(_a0 *entity.Dimensions, _a1 error) *Provider_GetDimensionsByImageByte_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Provider_GetDimensionsByImageByte_Call) RunAndReturn(run func([]byte, extension.Type) (*entity.Dimensions, error)) *Provider_GetDimensionsByImageByte_Call {
	_c.Call.Return(run)
	return _c
}

// GetDimensionsByMatrixCode provides a mock function with given fields: code
func (_m *Provider) GetDimensionsByMatrixCode(code string) (*entity.Dimensions, error) {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for GetDimensionsByMatrixCode")
	}

	var r0 *entity.Dimensions
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Dimensions, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Dimensions); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Dimensions)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Provider_GetDimensionsByMatrixCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDimensionsByMatrixCode'
type Provider_GetDimensionsByMatrixCode_Call struct {
	*mock.Call
}

// GetDimensionsByMatrixCode is a helper method to define mock.On call
//   - code string
func (_e *Provider_Expecter) GetDimensionsByMatrixCode(code interface{}) *Provider_GetDimensionsByMatrixCode_Call {
	return &Provider_GetDimensionsByMatrixCode_Call{Call: _e.mock.On("GetDimensionsByMatrixCode", code)}
}

func (_c *Provider_GetDimensionsByMatrixCode_Call) Run(run func(code string)) *Provider_GetDimensionsByMatrixCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Provider_GetDimensionsByMatrixCode_Call) Return(_a0 *entity.Dimensions, _a1 error) *Provider_GetDimensionsByMatrixCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Provider_GetDimensionsByMatrixCode_Call) RunAndReturn(run func(string) (*entity.Dimensions, error)) *Provider_GetDimensionsByMatrixCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetDimensionsByQrCode provides a mock function with given fields: code
func (_m *Provider) GetDimensionsByQrCode(code string) (*entity.Dimensions, error) {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for GetDimensionsByQrCode")
	}

	var r0 *entity.Dimensions
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Dimensions, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Dimensions); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Dimensions)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Provider_GetDimensionsByQrCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDimensionsByQrCode'
type Provider_GetDimensionsByQrCode_Call struct {
	*mock.Call
}

// GetDimensionsByQrCode is a helper method to define mock.On call
//   - code string
func (_e *Provider_Expecter) GetDimensionsByQrCode(code interface{}) *Provider_GetDimensionsByQrCode_Call {
	return &Provider_GetDimensionsByQrCode_Call{Call: _e.mock.On("GetDimensionsByQrCode", code)}
}

func (_c *Provider_GetDimensionsByQrCode_Call) Run(run func(code string)) *Provider_GetDimensionsByQrCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Provider_GetDimensionsByQrCode_Call) Return(_a0 *entity.Dimensions, _a1 error) *Provider_GetDimensionsByQrCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Provider_GetDimensionsByQrCode_Call) RunAndReturn(run func(string) (*entity.Dimensions, error)) *Provider_GetDimensionsByQrCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetFontHeight provides a mock function with given fields: prop
func (_m *Provider) GetFontHeight(prop *props.Font) float64 {
	ret := _m.Called(prop)

	if len(ret) == 0 {
		panic("no return value specified for GetFontHeight")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func(*props.Font) float64); ok {
		r0 = rf(prop)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Provider_GetFontHeight_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFontHeight'
type Provider_GetFontHeight_Call struct {
	*mock.Call
}

// GetFontHeight is a helper method to define mock.On call
//   - prop *props.Font
func (_e *Provider_Expecter) GetFontHeight(prop interface{}) *Provider_GetFontHeight_Call {
	return &Provider_GetFontHeight_Call{Call: _e.mock.On("GetFontHeight", prop)}
}

func (_c *Provider_GetFontHeight_Call) Run(run func(prop *props.Font)) *Provider_GetFontHeight_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*props.Font))
	})
	return _c
}

func (_c *Provider_GetFontHeight_Call) Return(_a0 float64) *Provider_GetFontHeight_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Provider_GetFontHeight_Call) RunAndReturn(run func(*props.Font) float64) *Provider_GetFontHeight_Call {
	_c.Call.Return(run)
	return _c
}

// GetLinesQuantity provides a mock function with given fields: text, textProp, colWidth
func (_m *Provider) GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int {
	ret := _m.Called(text, textProp, colWidth)

	if len(ret) == 0 {
		panic("no return value specified for GetLinesQuantity")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(string, *props.Text, float64) int); ok {
		r0 = rf(text, textProp, colWidth)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Provider_GetLinesQuantity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLinesQuantity'
type Provider_GetLinesQuantity_Call struct {
	*mock.Call
}

// GetLinesQuantity is a helper method to define mock.On call
//   - text string
//   - textProp *props.Text
//   - colWidth float64
func (_e *Provider_Expecter) GetLinesQuantity(text interface{}, textProp interface{}, colWidth interface{}) *Provider_GetLinesQuantity_Call {
	return &Provider_GetLinesQuantity_Call{Call: _e.mock.On("GetLinesQuantity", text, textProp, colWidth)}
}

func (_c *Provider_GetLinesQuantity_Call) Run(run func(text string, textProp *props.Text, colWidth float64)) *Provider_GetLinesQuantity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*props.Text), args[2].(float64))
	})
	return _c
}

func (_c *Provider_GetLinesQuantity_Call) Return(_a0 int) *Provider_GetLinesQuantity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Provider_GetLinesQuantity_Call) RunAndReturn(run func(string, *props.Text, float64) int) *Provider_GetLinesQuantity_Call {
	_c.Call.Return(run)
	return _c
}

// SetCompression provides a mock function with given fields: compression
func (_m *Provider) SetCompression(compression bool) {
	_m.Called(compression)
}

// Provider_SetCompression_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetCompression'
type Provider_SetCompression_Call struct {
	*mock.Call
}

// SetCompression is a helper method to define mock.On call
//   - compression bool
func (_e *Provider_Expecter) SetCompression(compression interface{}) *Provider_SetCompression_Call {
	return &Provider_SetCompression_Call{Call: _e.mock.On("SetCompression", compression)}
}

func (_c *Provider_SetCompression_Call) Run(run func(compression bool)) *Provider_SetCompression_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *Provider_SetCompression_Call) Return() *Provider_SetCompression_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_SetCompression_Call) RunAndReturn(run func(bool)) *Provider_SetCompression_Call {
	_c.Call.Return(run)
	return _c
}

// SetMetadata provides a mock function with given fields: metadata
func (_m *Provider) SetMetadata(metadata *entity.Metadata) {
	_m.Called(metadata)
}

// Provider_SetMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetMetadata'
type Provider_SetMetadata_Call struct {
	*mock.Call
}

// SetMetadata is a helper method to define mock.On call
//   - metadata *entity.Metadata
func (_e *Provider_Expecter) SetMetadata(metadata interface{}) *Provider_SetMetadata_Call {
	return &Provider_SetMetadata_Call{Call: _e.mock.On("SetMetadata", metadata)}
}

func (_c *Provider_SetMetadata_Call) Run(run func(metadata *entity.Metadata)) *Provider_SetMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Metadata))
	})
	return _c
}

func (_c *Provider_SetMetadata_Call) Return() *Provider_SetMetadata_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_SetMetadata_Call) RunAndReturn(run func(*entity.Metadata)) *Provider_SetMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// SetProtection provides a mock function with given fields: protection
func (_m *Provider) SetProtection(protection *entity.Protection) {
	_m.Called(protection)
}

// Provider_SetProtection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetProtection'
type Provider_SetProtection_Call struct {
	*mock.Call
}

// SetProtection is a helper method to define mock.On call
//   - protection *entity.Protection
func (_e *Provider_Expecter) SetProtection(protection interface{}) *Provider_SetProtection_Call {
	return &Provider_SetProtection_Call{Call: _e.mock.On("SetProtection", protection)}
}

func (_c *Provider_SetProtection_Call) Run(run func(protection *entity.Protection)) *Provider_SetProtection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Protection))
	})
	return _c
}

func (_c *Provider_SetProtection_Call) Return() *Provider_SetProtection_Call {
	_c.Call.Return()
	return _c
}

func (_c *Provider_SetProtection_Call) RunAndReturn(run func(*entity.Protection)) *Provider_SetProtection_Call {
	_c.Call.Return(run)
	return _c
}

// NewProvider creates a new instance of Provider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProvider(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Provider {
	mock := &Provider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
