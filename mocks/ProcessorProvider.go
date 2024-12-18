// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	core "github.com/johnfercher/maroto/v2/pkg/core"
	mock "github.com/stretchr/testify/mock"

	node "github.com/johnfercher/go-tree/node"

	processorprovider "github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"

	propsmapper "github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

// ProcessorProvider is an autogenerated mock type for the ProcessorProvider type
type ProcessorProvider struct {
	mock.Mock
}

type ProcessorProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *ProcessorProvider) EXPECT() *ProcessorProvider_Expecter {
	return &ProcessorProvider_Expecter{mock: &_m.Mock}
}

// AddFooter provides a mock function with given fields: footer
func (_m *ProcessorProvider) AddFooter(footer ...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error) {
	_va := make([]interface{}, len(footer))
	for _i := range footer {
		_va[_i] = footer[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddFooter")
	}

	var r0 processorprovider.ProcessorProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error)); ok {
		return rf(footer...)
	}
	if rf, ok := ret.Get(0).(func(...processorprovider.ProviderComponent) processorprovider.ProcessorProvider); ok {
		r0 = rf(footer...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProcessorProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(...processorprovider.ProviderComponent) error); ok {
		r1 = rf(footer...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessorProvider_AddFooter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFooter'
type ProcessorProvider_AddFooter_Call struct {
	*mock.Call
}

// AddFooter is a helper method to define mock.On call
//   - footer ...processorprovider.ProviderComponent
func (_e *ProcessorProvider_Expecter) AddFooter(footer ...interface{}) *ProcessorProvider_AddFooter_Call {
	return &ProcessorProvider_AddFooter_Call{Call: _e.mock.On("AddFooter",
		append([]interface{}{}, footer...)...)}
}

func (_c *ProcessorProvider_AddFooter_Call) Run(run func(footer ...processorprovider.ProviderComponent)) *ProcessorProvider_AddFooter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]processorprovider.ProviderComponent, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(processorprovider.ProviderComponent)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_AddFooter_Call) Return(_a0 processorprovider.ProcessorProvider, _a1 error) *ProcessorProvider_AddFooter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessorProvider_AddFooter_Call) RunAndReturn(run func(...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error)) *ProcessorProvider_AddFooter_Call {
	_c.Call.Return(run)
	return _c
}

// AddHeader provides a mock function with given fields: header
func (_m *ProcessorProvider) AddHeader(header ...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error) {
	_va := make([]interface{}, len(header))
	for _i := range header {
		_va[_i] = header[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddHeader")
	}

	var r0 processorprovider.ProcessorProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error)); ok {
		return rf(header...)
	}
	if rf, ok := ret.Get(0).(func(...processorprovider.ProviderComponent) processorprovider.ProcessorProvider); ok {
		r0 = rf(header...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProcessorProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(...processorprovider.ProviderComponent) error); ok {
		r1 = rf(header...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessorProvider_AddHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddHeader'
type ProcessorProvider_AddHeader_Call struct {
	*mock.Call
}

// AddHeader is a helper method to define mock.On call
//   - header ...processorprovider.ProviderComponent
func (_e *ProcessorProvider_Expecter) AddHeader(header ...interface{}) *ProcessorProvider_AddHeader_Call {
	return &ProcessorProvider_AddHeader_Call{Call: _e.mock.On("AddHeader",
		append([]interface{}{}, header...)...)}
}

func (_c *ProcessorProvider_AddHeader_Call) Run(run func(header ...processorprovider.ProviderComponent)) *ProcessorProvider_AddHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]processorprovider.ProviderComponent, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(processorprovider.ProviderComponent)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_AddHeader_Call) Return(_a0 processorprovider.ProcessorProvider, _a1 error) *ProcessorProvider_AddHeader_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessorProvider_AddHeader_Call) RunAndReturn(run func(...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error)) *ProcessorProvider_AddHeader_Call {
	_c.Call.Return(run)
	return _c
}

// AddPages provides a mock function with given fields: pages
func (_m *ProcessorProvider) AddPages(pages ...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error) {
	_va := make([]interface{}, len(pages))
	for _i := range pages {
		_va[_i] = pages[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddPages")
	}

	var r0 processorprovider.ProcessorProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error)); ok {
		return rf(pages...)
	}
	if rf, ok := ret.Get(0).(func(...processorprovider.ProviderComponent) processorprovider.ProcessorProvider); ok {
		r0 = rf(pages...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProcessorProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(...processorprovider.ProviderComponent) error); ok {
		r1 = rf(pages...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessorProvider_AddPages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPages'
type ProcessorProvider_AddPages_Call struct {
	*mock.Call
}

// AddPages is a helper method to define mock.On call
//   - pages ...processorprovider.ProviderComponent
func (_e *ProcessorProvider_Expecter) AddPages(pages ...interface{}) *ProcessorProvider_AddPages_Call {
	return &ProcessorProvider_AddPages_Call{Call: _e.mock.On("AddPages",
		append([]interface{}{}, pages...)...)}
}

func (_c *ProcessorProvider_AddPages_Call) Run(run func(pages ...processorprovider.ProviderComponent)) *ProcessorProvider_AddPages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]processorprovider.ProviderComponent, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(processorprovider.ProviderComponent)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_AddPages_Call) Return(_a0 processorprovider.ProcessorProvider, _a1 error) *ProcessorProvider_AddPages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessorProvider_AddPages_Call) RunAndReturn(run func(...processorprovider.ProviderComponent) (processorprovider.ProcessorProvider, error)) *ProcessorProvider_AddPages_Call {
	_c.Call.Return(run)
	return _c
}

// CreateBarCode provides a mock function with given fields: value, props
func (_m *ProcessorProvider) CreateBarCode(value string, props ...*propsmapper.Barcode) processorprovider.ProviderComponent {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateBarCode")
	}

	var r0 processorprovider.ProviderComponent
	if rf, ok := ret.Get(0).(func(string, ...*propsmapper.Barcode) processorprovider.ProviderComponent); ok {
		r0 = rf(value, props...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	return r0
}

// ProcessorProvider_CreateBarCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBarCode'
type ProcessorProvider_CreateBarCode_Call struct {
	*mock.Call
}

// CreateBarCode is a helper method to define mock.On call
//   - value string
//   - props ...*propsmapper.Barcode
func (_e *ProcessorProvider_Expecter) CreateBarCode(value interface{}, props ...interface{}) *ProcessorProvider_CreateBarCode_Call {
	return &ProcessorProvider_CreateBarCode_Call{Call: _e.mock.On("CreateBarCode",
		append([]interface{}{value}, props...)...)}
}

func (_c *ProcessorProvider_CreateBarCode_Call) Run(run func(value string, props ...*propsmapper.Barcode)) *ProcessorProvider_CreateBarCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*propsmapper.Barcode, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*propsmapper.Barcode)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateBarCode_Call) Return(_a0 processorprovider.ProviderComponent) *ProcessorProvider_CreateBarCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProcessorProvider_CreateBarCode_Call) RunAndReturn(run func(string, ...*propsmapper.Barcode) processorprovider.ProviderComponent) *ProcessorProvider_CreateBarCode_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCol provides a mock function with given fields: size, components
func (_m *ProcessorProvider) CreateCol(size int, components ...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error) {
	_va := make([]interface{}, len(components))
	for _i := range components {
		_va[_i] = components[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, size)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCol")
	}

	var r0 processorprovider.ProviderComponent
	var r1 error
	if rf, ok := ret.Get(0).(func(int, ...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error)); ok {
		return rf(size, components...)
	}
	if rf, ok := ret.Get(0).(func(int, ...processorprovider.ProviderComponent) processorprovider.ProviderComponent); ok {
		r0 = rf(size, components...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	if rf, ok := ret.Get(1).(func(int, ...processorprovider.ProviderComponent) error); ok {
		r1 = rf(size, components...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessorProvider_CreateCol_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCol'
type ProcessorProvider_CreateCol_Call struct {
	*mock.Call
}

// CreateCol is a helper method to define mock.On call
//   - size int
//   - components ...processorprovider.ProviderComponent
func (_e *ProcessorProvider_Expecter) CreateCol(size interface{}, components ...interface{}) *ProcessorProvider_CreateCol_Call {
	return &ProcessorProvider_CreateCol_Call{Call: _e.mock.On("CreateCol",
		append([]interface{}{size}, components...)...)}
}

func (_c *ProcessorProvider_CreateCol_Call) Run(run func(size int, components ...processorprovider.ProviderComponent)) *ProcessorProvider_CreateCol_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]processorprovider.ProviderComponent, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(processorprovider.ProviderComponent)
			}
		}
		run(args[0].(int), variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateCol_Call) Return(_a0 processorprovider.ProviderComponent, _a1 error) *ProcessorProvider_CreateCol_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessorProvider_CreateCol_Call) RunAndReturn(run func(int, ...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error)) *ProcessorProvider_CreateCol_Call {
	_c.Call.Return(run)
	return _c
}

// CreateImage provides a mock function with given fields: path, props
func (_m *ProcessorProvider) CreateImage(path string, props ...*propsmapper.Rect) (processorprovider.ProviderComponent, error) {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateImage")
	}

	var r0 processorprovider.ProviderComponent
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...*propsmapper.Rect) (processorprovider.ProviderComponent, error)); ok {
		return rf(path, props...)
	}
	if rf, ok := ret.Get(0).(func(string, ...*propsmapper.Rect) processorprovider.ProviderComponent); ok {
		r0 = rf(path, props...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...*propsmapper.Rect) error); ok {
		r1 = rf(path, props...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessorProvider_CreateImage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateImage'
type ProcessorProvider_CreateImage_Call struct {
	*mock.Call
}

// CreateImage is a helper method to define mock.On call
//   - path string
//   - props ...*propsmapper.Rect
func (_e *ProcessorProvider_Expecter) CreateImage(path interface{}, props ...interface{}) *ProcessorProvider_CreateImage_Call {
	return &ProcessorProvider_CreateImage_Call{Call: _e.mock.On("CreateImage",
		append([]interface{}{path}, props...)...)}
}

func (_c *ProcessorProvider_CreateImage_Call) Run(run func(path string, props ...*propsmapper.Rect)) *ProcessorProvider_CreateImage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*propsmapper.Rect, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*propsmapper.Rect)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateImage_Call) Return(_a0 processorprovider.ProviderComponent, _a1 error) *ProcessorProvider_CreateImage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessorProvider_CreateImage_Call) RunAndReturn(run func(string, ...*propsmapper.Rect) (processorprovider.ProviderComponent, error)) *ProcessorProvider_CreateImage_Call {
	_c.Call.Return(run)
	return _c
}

// CreateLine provides a mock function with given fields: props
func (_m *ProcessorProvider) CreateLine(props ...*propsmapper.Line) processorprovider.ProviderComponent {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateLine")
	}

	var r0 processorprovider.ProviderComponent
	if rf, ok := ret.Get(0).(func(...*propsmapper.Line) processorprovider.ProviderComponent); ok {
		r0 = rf(props...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	return r0
}

// ProcessorProvider_CreateLine_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLine'
type ProcessorProvider_CreateLine_Call struct {
	*mock.Call
}

// CreateLine is a helper method to define mock.On call
//   - props ...*propsmapper.Line
func (_e *ProcessorProvider_Expecter) CreateLine(props ...interface{}) *ProcessorProvider_CreateLine_Call {
	return &ProcessorProvider_CreateLine_Call{Call: _e.mock.On("CreateLine",
		append([]interface{}{}, props...)...)}
}

func (_c *ProcessorProvider_CreateLine_Call) Run(run func(props ...*propsmapper.Line)) *ProcessorProvider_CreateLine_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*propsmapper.Line, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(*propsmapper.Line)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateLine_Call) Return(_a0 processorprovider.ProviderComponent) *ProcessorProvider_CreateLine_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProcessorProvider_CreateLine_Call) RunAndReturn(run func(...*propsmapper.Line) processorprovider.ProviderComponent) *ProcessorProvider_CreateLine_Call {
	_c.Call.Return(run)
	return _c
}

// CreateMatrixCode provides a mock function with given fields: value, props
func (_m *ProcessorProvider) CreateMatrixCode(value string, props ...*propsmapper.Rect) processorprovider.ProviderComponent {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMatrixCode")
	}

	var r0 processorprovider.ProviderComponent
	if rf, ok := ret.Get(0).(func(string, ...*propsmapper.Rect) processorprovider.ProviderComponent); ok {
		r0 = rf(value, props...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	return r0
}

// ProcessorProvider_CreateMatrixCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMatrixCode'
type ProcessorProvider_CreateMatrixCode_Call struct {
	*mock.Call
}

// CreateMatrixCode is a helper method to define mock.On call
//   - value string
//   - props ...*propsmapper.Rect
func (_e *ProcessorProvider_Expecter) CreateMatrixCode(value interface{}, props ...interface{}) *ProcessorProvider_CreateMatrixCode_Call {
	return &ProcessorProvider_CreateMatrixCode_Call{Call: _e.mock.On("CreateMatrixCode",
		append([]interface{}{value}, props...)...)}
}

func (_c *ProcessorProvider_CreateMatrixCode_Call) Run(run func(value string, props ...*propsmapper.Rect)) *ProcessorProvider_CreateMatrixCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*propsmapper.Rect, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*propsmapper.Rect)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateMatrixCode_Call) Return(_a0 processorprovider.ProviderComponent) *ProcessorProvider_CreateMatrixCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProcessorProvider_CreateMatrixCode_Call) RunAndReturn(run func(string, ...*propsmapper.Rect) processorprovider.ProviderComponent) *ProcessorProvider_CreateMatrixCode_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePage provides a mock function with given fields: components
func (_m *ProcessorProvider) CreatePage(components ...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error) {
	_va := make([]interface{}, len(components))
	for _i := range components {
		_va[_i] = components[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePage")
	}

	var r0 processorprovider.ProviderComponent
	var r1 error
	if rf, ok := ret.Get(0).(func(...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error)); ok {
		return rf(components...)
	}
	if rf, ok := ret.Get(0).(func(...processorprovider.ProviderComponent) processorprovider.ProviderComponent); ok {
		r0 = rf(components...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	if rf, ok := ret.Get(1).(func(...processorprovider.ProviderComponent) error); ok {
		r1 = rf(components...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessorProvider_CreatePage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePage'
type ProcessorProvider_CreatePage_Call struct {
	*mock.Call
}

// CreatePage is a helper method to define mock.On call
//   - components ...processorprovider.ProviderComponent
func (_e *ProcessorProvider_Expecter) CreatePage(components ...interface{}) *ProcessorProvider_CreatePage_Call {
	return &ProcessorProvider_CreatePage_Call{Call: _e.mock.On("CreatePage",
		append([]interface{}{}, components...)...)}
}

func (_c *ProcessorProvider_CreatePage_Call) Run(run func(components ...processorprovider.ProviderComponent)) *ProcessorProvider_CreatePage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]processorprovider.ProviderComponent, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(processorprovider.ProviderComponent)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreatePage_Call) Return(_a0 processorprovider.ProviderComponent, _a1 error) *ProcessorProvider_CreatePage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessorProvider_CreatePage_Call) RunAndReturn(run func(...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error)) *ProcessorProvider_CreatePage_Call {
	_c.Call.Return(run)
	return _c
}

// CreateQrCode provides a mock function with given fields: value, props
func (_m *ProcessorProvider) CreateQrCode(value string, props ...*propsmapper.Rect) processorprovider.ProviderComponent {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateQrCode")
	}

	var r0 processorprovider.ProviderComponent
	if rf, ok := ret.Get(0).(func(string, ...*propsmapper.Rect) processorprovider.ProviderComponent); ok {
		r0 = rf(value, props...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	return r0
}

// ProcessorProvider_CreateQrCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateQrCode'
type ProcessorProvider_CreateQrCode_Call struct {
	*mock.Call
}

// CreateQrCode is a helper method to define mock.On call
//   - value string
//   - props ...*propsmapper.Rect
func (_e *ProcessorProvider_Expecter) CreateQrCode(value interface{}, props ...interface{}) *ProcessorProvider_CreateQrCode_Call {
	return &ProcessorProvider_CreateQrCode_Call{Call: _e.mock.On("CreateQrCode",
		append([]interface{}{value}, props...)...)}
}

func (_c *ProcessorProvider_CreateQrCode_Call) Run(run func(value string, props ...*propsmapper.Rect)) *ProcessorProvider_CreateQrCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*propsmapper.Rect, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*propsmapper.Rect)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateQrCode_Call) Return(_a0 processorprovider.ProviderComponent) *ProcessorProvider_CreateQrCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProcessorProvider_CreateQrCode_Call) RunAndReturn(run func(string, ...*propsmapper.Rect) processorprovider.ProviderComponent) *ProcessorProvider_CreateQrCode_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRow provides a mock function with given fields: height, components
func (_m *ProcessorProvider) CreateRow(height float64, components ...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error) {
	_va := make([]interface{}, len(components))
	for _i := range components {
		_va[_i] = components[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, height)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRow")
	}

	var r0 processorprovider.ProviderComponent
	var r1 error
	if rf, ok := ret.Get(0).(func(float64, ...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error)); ok {
		return rf(height, components...)
	}
	if rf, ok := ret.Get(0).(func(float64, ...processorprovider.ProviderComponent) processorprovider.ProviderComponent); ok {
		r0 = rf(height, components...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	if rf, ok := ret.Get(1).(func(float64, ...processorprovider.ProviderComponent) error); ok {
		r1 = rf(height, components...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessorProvider_CreateRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRow'
type ProcessorProvider_CreateRow_Call struct {
	*mock.Call
}

// CreateRow is a helper method to define mock.On call
//   - height float64
//   - components ...processorprovider.ProviderComponent
func (_e *ProcessorProvider_Expecter) CreateRow(height interface{}, components ...interface{}) *ProcessorProvider_CreateRow_Call {
	return &ProcessorProvider_CreateRow_Call{Call: _e.mock.On("CreateRow",
		append([]interface{}{height}, components...)...)}
}

func (_c *ProcessorProvider_CreateRow_Call) Run(run func(height float64, components ...processorprovider.ProviderComponent)) *ProcessorProvider_CreateRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]processorprovider.ProviderComponent, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(processorprovider.ProviderComponent)
			}
		}
		run(args[0].(float64), variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateRow_Call) Return(_a0 processorprovider.ProviderComponent, _a1 error) *ProcessorProvider_CreateRow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessorProvider_CreateRow_Call) RunAndReturn(run func(float64, ...processorprovider.ProviderComponent) (processorprovider.ProviderComponent, error)) *ProcessorProvider_CreateRow_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSignature provides a mock function with given fields: value, props
func (_m *ProcessorProvider) CreateSignature(value string, props ...*propsmapper.Signature) processorprovider.ProviderComponent {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSignature")
	}

	var r0 processorprovider.ProviderComponent
	if rf, ok := ret.Get(0).(func(string, ...*propsmapper.Signature) processorprovider.ProviderComponent); ok {
		r0 = rf(value, props...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	return r0
}

// ProcessorProvider_CreateSignature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSignature'
type ProcessorProvider_CreateSignature_Call struct {
	*mock.Call
}

// CreateSignature is a helper method to define mock.On call
//   - value string
//   - props ...*propsmapper.Signature
func (_e *ProcessorProvider_Expecter) CreateSignature(value interface{}, props ...interface{}) *ProcessorProvider_CreateSignature_Call {
	return &ProcessorProvider_CreateSignature_Call{Call: _e.mock.On("CreateSignature",
		append([]interface{}{value}, props...)...)}
}

func (_c *ProcessorProvider_CreateSignature_Call) Run(run func(value string, props ...*propsmapper.Signature)) *ProcessorProvider_CreateSignature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*propsmapper.Signature, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*propsmapper.Signature)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateSignature_Call) Return(_a0 processorprovider.ProviderComponent) *ProcessorProvider_CreateSignature_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProcessorProvider_CreateSignature_Call) RunAndReturn(run func(string, ...*propsmapper.Signature) processorprovider.ProviderComponent) *ProcessorProvider_CreateSignature_Call {
	_c.Call.Return(run)
	return _c
}

// CreateText provides a mock function with given fields: value, props
func (_m *ProcessorProvider) CreateText(value string, props ...*propsmapper.Text) processorprovider.ProviderComponent {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateText")
	}

	var r0 processorprovider.ProviderComponent
	if rf, ok := ret.Get(0).(func(string, ...*propsmapper.Text) processorprovider.ProviderComponent); ok {
		r0 = rf(value, props...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(processorprovider.ProviderComponent)
		}
	}

	return r0
}

// ProcessorProvider_CreateText_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateText'
type ProcessorProvider_CreateText_Call struct {
	*mock.Call
}

// CreateText is a helper method to define mock.On call
//   - value string
//   - props ...*propsmapper.Text
func (_e *ProcessorProvider_Expecter) CreateText(value interface{}, props ...interface{}) *ProcessorProvider_CreateText_Call {
	return &ProcessorProvider_CreateText_Call{Call: _e.mock.On("CreateText",
		append([]interface{}{value}, props...)...)}
}

func (_c *ProcessorProvider_CreateText_Call) Run(run func(value string, props ...*propsmapper.Text)) *ProcessorProvider_CreateText_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*propsmapper.Text, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*propsmapper.Text)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *ProcessorProvider_CreateText_Call) Return(_a0 processorprovider.ProviderComponent) *ProcessorProvider_CreateText_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProcessorProvider_CreateText_Call) RunAndReturn(run func(string, ...*propsmapper.Text) processorprovider.ProviderComponent) *ProcessorProvider_CreateText_Call {
	_c.Call.Return(run)
	return _c
}

// Generate provides a mock function with given fields:
func (_m *ProcessorProvider) Generate() (core.Document, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Generate")
	}

	var r0 core.Document
	var r1 error
	if rf, ok := ret.Get(0).(func() (core.Document, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() core.Document); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Document)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessorProvider_Generate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Generate'
type ProcessorProvider_Generate_Call struct {
	*mock.Call
}

// Generate is a helper method to define mock.On call
func (_e *ProcessorProvider_Expecter) Generate() *ProcessorProvider_Generate_Call {
	return &ProcessorProvider_Generate_Call{Call: _e.mock.On("Generate")}
}

func (_c *ProcessorProvider_Generate_Call) Run(run func()) *ProcessorProvider_Generate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ProcessorProvider_Generate_Call) Return(_a0 core.Document, _a1 error) *ProcessorProvider_Generate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProcessorProvider_Generate_Call) RunAndReturn(run func() (core.Document, error)) *ProcessorProvider_Generate_Call {
	_c.Call.Return(run)
	return _c
}

// GetStructure provides a mock function with given fields:
func (_m *ProcessorProvider) GetStructure() *node.Node[core.Structure] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStructure")
	}

	var r0 *node.Node[core.Structure]
	if rf, ok := ret.Get(0).(func() *node.Node[core.Structure]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*node.Node[core.Structure])
		}
	}

	return r0
}

// ProcessorProvider_GetStructure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStructure'
type ProcessorProvider_GetStructure_Call struct {
	*mock.Call
}

// GetStructure is a helper method to define mock.On call
func (_e *ProcessorProvider_Expecter) GetStructure() *ProcessorProvider_GetStructure_Call {
	return &ProcessorProvider_GetStructure_Call{Call: _e.mock.On("GetStructure")}
}

func (_c *ProcessorProvider_GetStructure_Call) Run(run func()) *ProcessorProvider_GetStructure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ProcessorProvider_GetStructure_Call) Return(_a0 *node.Node[core.Structure]) *ProcessorProvider_GetStructure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProcessorProvider_GetStructure_Call) RunAndReturn(run func() *node.Node[core.Structure]) *ProcessorProvider_GetStructure_Call {
	_c.Call.Return(run)
	return _c
}

// NewProcessorProvider creates a new instance of ProcessorProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProcessorProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProcessorProvider {
	mock := &ProcessorProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
