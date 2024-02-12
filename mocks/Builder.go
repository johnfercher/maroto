// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	config "github.com/johnfercher/maroto/v2/pkg/config"
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"

	extension "github.com/johnfercher/maroto/v2/pkg/consts/extension"

	mock "github.com/stretchr/testify/mock"

	orientation "github.com/johnfercher/maroto/v2/pkg/consts/orientation"

	pagesize "github.com/johnfercher/maroto/v2/pkg/consts/pagesize"

	props "github.com/johnfercher/maroto/v2/pkg/props"

	protection "github.com/johnfercher/maroto/v2/pkg/consts/protection"

	time "time"
)

// Builder is an autogenerated mock type for the Builder type
type Builder struct {
	mock.Mock
}

type Builder_Expecter struct {
	mock *mock.Mock
}

func (_m *Builder) EXPECT() *Builder_Expecter {
	return &Builder_Expecter{mock: &_m.Mock}
}

// Build provides a mock function with given fields:
func (_m *Builder) Build() *entity.Config {
	ret := _m.Called()

	var r0 *entity.Config
	if rf, ok := ret.Get(0).(func() *entity.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Config)
		}
	}

	return r0
}

// Builder_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type Builder_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
func (_e *Builder_Expecter) Build() *Builder_Build_Call {
	return &Builder_Build_Call{Call: _e.mock.On("Build")}
}

func (_c *Builder_Build_Call) Run(run func()) *Builder_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Builder_Build_Call) Return(_a0 *entity.Config) *Builder_Build_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_Build_Call) RunAndReturn(run func() *entity.Config) *Builder_Build_Call {
	_c.Call.Return(run)
	return _c
}

// WithAuthor provides a mock function with given fields: author, isUTF8
func (_m *Builder) WithAuthor(author string, isUTF8 bool) config.Builder {
	ret := _m.Called(author, isUTF8)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(string, bool) config.Builder); ok {
		r0 = rf(author, isUTF8)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithAuthor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithAuthor'
type Builder_WithAuthor_Call struct {
	*mock.Call
}

// WithAuthor is a helper method to define mock.On call
//   - author string
//   - isUTF8 bool
func (_e *Builder_Expecter) WithAuthor(author interface{}, isUTF8 interface{}) *Builder_WithAuthor_Call {
	return &Builder_WithAuthor_Call{Call: _e.mock.On("WithAuthor", author, isUTF8)}
}

func (_c *Builder_WithAuthor_Call) Run(run func(author string, isUTF8 bool)) *Builder_WithAuthor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *Builder_WithAuthor_Call) Return(_a0 config.Builder) *Builder_WithAuthor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithAuthor_Call) RunAndReturn(run func(string, bool) config.Builder) *Builder_WithAuthor_Call {
	_c.Call.Return(run)
	return _c
}

// WithBackgroundImage provides a mock function with given fields: _a0, _a1
func (_m *Builder) WithBackgroundImage(_a0 []byte, _a1 extension.Type) config.Builder {
	ret := _m.Called(_a0, _a1)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func([]byte, extension.Type) config.Builder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithBackgroundImage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithBackgroundImage'
type Builder_WithBackgroundImage_Call struct {
	*mock.Call
}

// WithBackgroundImage is a helper method to define mock.On call
//   - _a0 []byte
//   - _a1 extension.Type
func (_e *Builder_Expecter) WithBackgroundImage(_a0 interface{}, _a1 interface{}) *Builder_WithBackgroundImage_Call {
	return &Builder_WithBackgroundImage_Call{Call: _e.mock.On("WithBackgroundImage", _a0, _a1)}
}

func (_c *Builder_WithBackgroundImage_Call) Run(run func(_a0 []byte, _a1 extension.Type)) *Builder_WithBackgroundImage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(extension.Type))
	})
	return _c
}

func (_c *Builder_WithBackgroundImage_Call) Return(_a0 config.Builder) *Builder_WithBackgroundImage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithBackgroundImage_Call) RunAndReturn(run func([]byte, extension.Type) config.Builder) *Builder_WithBackgroundImage_Call {
	_c.Call.Return(run)
	return _c
}

// WithCompression provides a mock function with given fields: compression
func (_m *Builder) WithCompression(compression bool) config.Builder {
	ret := _m.Called(compression)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(bool) config.Builder); ok {
		r0 = rf(compression)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithCompression_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithCompression'
type Builder_WithCompression_Call struct {
	*mock.Call
}

// WithCompression is a helper method to define mock.On call
//   - compression bool
func (_e *Builder_Expecter) WithCompression(compression interface{}) *Builder_WithCompression_Call {
	return &Builder_WithCompression_Call{Call: _e.mock.On("WithCompression", compression)}
}

func (_c *Builder_WithCompression_Call) Run(run func(compression bool)) *Builder_WithCompression_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *Builder_WithCompression_Call) Return(_a0 config.Builder) *Builder_WithCompression_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithCompression_Call) RunAndReturn(run func(bool) config.Builder) *Builder_WithCompression_Call {
	_c.Call.Return(run)
	return _c
}

// WithCreationDate provides a mock function with given fields: _a0
func (_m *Builder) WithCreationDate(_a0 time.Time) config.Builder {
	ret := _m.Called(_a0)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(time.Time) config.Builder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithCreationDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithCreationDate'
type Builder_WithCreationDate_Call struct {
	*mock.Call
}

// WithCreationDate is a helper method to define mock.On call
//   - _a0 time.Time
func (_e *Builder_Expecter) WithCreationDate(_a0 interface{}) *Builder_WithCreationDate_Call {
	return &Builder_WithCreationDate_Call{Call: _e.mock.On("WithCreationDate", _a0)}
}

func (_c *Builder_WithCreationDate_Call) Run(run func(_a0 time.Time)) *Builder_WithCreationDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Time))
	})
	return _c
}

func (_c *Builder_WithCreationDate_Call) Return(_a0 config.Builder) *Builder_WithCreationDate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithCreationDate_Call) RunAndReturn(run func(time.Time) config.Builder) *Builder_WithCreationDate_Call {
	_c.Call.Return(run)
	return _c
}

// WithCreator provides a mock function with given fields: creator, isUTF8
func (_m *Builder) WithCreator(creator string, isUTF8 bool) config.Builder {
	ret := _m.Called(creator, isUTF8)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(string, bool) config.Builder); ok {
		r0 = rf(creator, isUTF8)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithCreator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithCreator'
type Builder_WithCreator_Call struct {
	*mock.Call
}

// WithCreator is a helper method to define mock.On call
//   - creator string
//   - isUTF8 bool
func (_e *Builder_Expecter) WithCreator(creator interface{}, isUTF8 interface{}) *Builder_WithCreator_Call {
	return &Builder_WithCreator_Call{Call: _e.mock.On("WithCreator", creator, isUTF8)}
}

func (_c *Builder_WithCreator_Call) Run(run func(creator string, isUTF8 bool)) *Builder_WithCreator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *Builder_WithCreator_Call) Return(_a0 config.Builder) *Builder_WithCreator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithCreator_Call) RunAndReturn(run func(string, bool) config.Builder) *Builder_WithCreator_Call {
	_c.Call.Return(run)
	return _c
}

// WithCustomFonts provides a mock function with given fields: _a0
func (_m *Builder) WithCustomFonts(_a0 []*entity.CustomFont) config.Builder {
	ret := _m.Called(_a0)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func([]*entity.CustomFont) config.Builder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithCustomFonts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithCustomFonts'
type Builder_WithCustomFonts_Call struct {
	*mock.Call
}

// WithCustomFonts is a helper method to define mock.On call
//   - _a0 []*entity.CustomFont
func (_e *Builder_Expecter) WithCustomFonts(_a0 interface{}) *Builder_WithCustomFonts_Call {
	return &Builder_WithCustomFonts_Call{Call: _e.mock.On("WithCustomFonts", _a0)}
}

func (_c *Builder_WithCustomFonts_Call) Run(run func(_a0 []*entity.CustomFont)) *Builder_WithCustomFonts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*entity.CustomFont))
	})
	return _c
}

func (_c *Builder_WithCustomFonts_Call) Return(_a0 config.Builder) *Builder_WithCustomFonts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithCustomFonts_Call) RunAndReturn(run func([]*entity.CustomFont) config.Builder) *Builder_WithCustomFonts_Call {
	_c.Call.Return(run)
	return _c
}

// WithDebug provides a mock function with given fields: on
func (_m *Builder) WithDebug(on bool) config.Builder {
	ret := _m.Called(on)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(bool) config.Builder); ok {
		r0 = rf(on)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithDebug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithDebug'
type Builder_WithDebug_Call struct {
	*mock.Call
}

// WithDebug is a helper method to define mock.On call
//   - on bool
func (_e *Builder_Expecter) WithDebug(on interface{}) *Builder_WithDebug_Call {
	return &Builder_WithDebug_Call{Call: _e.mock.On("WithDebug", on)}
}

func (_c *Builder_WithDebug_Call) Run(run func(on bool)) *Builder_WithDebug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *Builder_WithDebug_Call) Return(_a0 config.Builder) *Builder_WithDebug_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithDebug_Call) RunAndReturn(run func(bool) config.Builder) *Builder_WithDebug_Call {
	_c.Call.Return(run)
	return _c
}

// WithDefaultFont provides a mock function with given fields: font
func (_m *Builder) WithDefaultFont(font *props.Font) config.Builder {
	ret := _m.Called(font)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(*props.Font) config.Builder); ok {
		r0 = rf(font)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithDefaultFont_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithDefaultFont'
type Builder_WithDefaultFont_Call struct {
	*mock.Call
}

// WithDefaultFont is a helper method to define mock.On call
//   - font *props.Font
func (_e *Builder_Expecter) WithDefaultFont(font interface{}) *Builder_WithDefaultFont_Call {
	return &Builder_WithDefaultFont_Call{Call: _e.mock.On("WithDefaultFont", font)}
}

func (_c *Builder_WithDefaultFont_Call) Run(run func(font *props.Font)) *Builder_WithDefaultFont_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*props.Font))
	})
	return _c
}

func (_c *Builder_WithDefaultFont_Call) Return(_a0 config.Builder) *Builder_WithDefaultFont_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithDefaultFont_Call) RunAndReturn(run func(*props.Font) config.Builder) *Builder_WithDefaultFont_Call {
	_c.Call.Return(run)
	return _c
}

// WithDimensions provides a mock function with given fields: width, height
func (_m *Builder) WithDimensions(width float64, height float64) config.Builder {
	ret := _m.Called(width, height)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(float64, float64) config.Builder); ok {
		r0 = rf(width, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithDimensions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithDimensions'
type Builder_WithDimensions_Call struct {
	*mock.Call
}

// WithDimensions is a helper method to define mock.On call
//   - width float64
//   - height float64
func (_e *Builder_Expecter) WithDimensions(width interface{}, height interface{}) *Builder_WithDimensions_Call {
	return &Builder_WithDimensions_Call{Call: _e.mock.On("WithDimensions", width, height)}
}

func (_c *Builder_WithDimensions_Call) Run(run func(width float64, height float64)) *Builder_WithDimensions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64), args[1].(float64))
	})
	return _c
}

func (_c *Builder_WithDimensions_Call) Return(_a0 config.Builder) *Builder_WithDimensions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithDimensions_Call) RunAndReturn(run func(float64, float64) config.Builder) *Builder_WithDimensions_Call {
	_c.Call.Return(run)
	return _c
}

// Builder_WithDisableAutoPageBreak_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithDisableAutoPageBreak'
type Builder_WithDisableAutoPageBreak_Call struct {
	*mock.Call
}

// WithDisableAutoPageBreak is a helper method to define mock.On call
//   - disabled bool
func (_e *Builder_Expecter) WithDisableAutoPageBreak(disabled interface{}) *Builder_WithDisableAutoPageBreak_Call {
	return &Builder_WithDisableAutoPageBreak_Call{Call: _e.mock.On("WithDisableAutoPageBreak", disabled)}
}

func (_c *Builder_WithDisableAutoPageBreak_Call) Run(run func(disabled bool)) *Builder_WithDisableAutoPageBreak_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *Builder_WithDisableAutoPageBreak_Call) Return(_a0 config.Builder) *Builder_WithDisableAutoPageBreak_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithDisableAutoPageBreak_Call) RunAndReturn(run func(bool) config.Builder) *Builder_WithDisableAutoPageBreak_Call {
	_c.Call.Return(run)
	return _c
}

// WithMargins provides a mock function with given fields: left, top, right
func (_m *Builder) WithMargins(left float64, top float64, right float64) config.Builder {
	ret := _m.Called(left, top, right)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(float64, float64, float64) config.Builder); ok {
		r0 = rf(left, top, right)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithMargins_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithMargins'
type Builder_WithMargins_Call struct {
	*mock.Call
}

// WithMargins is a helper method to define mock.On call
//   - left float64
//   - top float64
//   - right float64
func (_e *Builder_Expecter) WithMargins(left interface{}, top interface{}, right interface{}) *Builder_WithMargins_Call {
	return &Builder_WithMargins_Call{Call: _e.mock.On("WithMargins", left, top, right)}
}

func (_c *Builder_WithMargins_Call) Run(run func(left float64, top float64, right float64)) *Builder_WithMargins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64), args[1].(float64), args[2].(float64))
	})
	return _c
}

func (_c *Builder_WithMargins_Call) Return(_a0 config.Builder) *Builder_WithMargins_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithMargins_Call) RunAndReturn(run func(float64, float64, float64) config.Builder) *Builder_WithMargins_Call {
	_c.Call.Return(run)
	return _c
}

// WithMaxGridSize provides a mock function with given fields: maxGridSize
func (_m *Builder) WithMaxGridSize(maxGridSize int) config.Builder {
	ret := _m.Called(maxGridSize)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(int) config.Builder); ok {
		r0 = rf(maxGridSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithMaxGridSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithMaxGridSize'
type Builder_WithMaxGridSize_Call struct {
	*mock.Call
}

// WithMaxGridSize is a helper method to define mock.On call
//   - maxGridSize int
func (_e *Builder_Expecter) WithMaxGridSize(maxGridSize interface{}) *Builder_WithMaxGridSize_Call {
	return &Builder_WithMaxGridSize_Call{Call: _e.mock.On("WithMaxGridSize", maxGridSize)}
}

func (_c *Builder_WithMaxGridSize_Call) Run(run func(maxGridSize int)) *Builder_WithMaxGridSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Builder_WithMaxGridSize_Call) Return(_a0 config.Builder) *Builder_WithMaxGridSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithMaxGridSize_Call) RunAndReturn(run func(int) config.Builder) *Builder_WithMaxGridSize_Call {
	_c.Call.Return(run)
	return _c
}

// WithOrientation provides a mock function with given fields: _a0
func (_m *Builder) WithOrientation(_a0 orientation.Type) config.Builder {
	ret := _m.Called(_a0)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(orientation.Type) config.Builder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithOrientation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithOrientation'
type Builder_WithOrientation_Call struct {
	*mock.Call
}

// WithOrientation is a helper method to define mock.On call
//   - _a0 orientation.Type
func (_e *Builder_Expecter) WithOrientation(_a0 interface{}) *Builder_WithOrientation_Call {
	return &Builder_WithOrientation_Call{Call: _e.mock.On("WithOrientation", _a0)}
}

func (_c *Builder_WithOrientation_Call) Run(run func(_a0 orientation.Type)) *Builder_WithOrientation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(orientation.Type))
	})
	return _c
}

func (_c *Builder_WithOrientation_Call) Return(_a0 config.Builder) *Builder_WithOrientation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithOrientation_Call) RunAndReturn(run func(orientation.Type) config.Builder) *Builder_WithOrientation_Call {
	_c.Call.Return(run)
	return _c
}

// WithPageNumber provides a mock function with given fields: pattern, place
func (_m *Builder) WithPageNumber(pattern string, place props.Place) config.Builder {
	ret := _m.Called(pattern, place)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(string, props.Place) config.Builder); ok {
		r0 = rf(pattern, place)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithPageNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithPageNumber'
type Builder_WithPageNumber_Call struct {
	*mock.Call
}

// WithPageNumber is a helper method to define mock.On call
//   - pattern string
//   - place props.Place
func (_e *Builder_Expecter) WithPageNumber(pattern interface{}, place interface{}) *Builder_WithPageNumber_Call {
	return &Builder_WithPageNumber_Call{Call: _e.mock.On("WithPageNumber", pattern, place)}
}

func (_c *Builder_WithPageNumber_Call) Run(run func(pattern string, place props.Place)) *Builder_WithPageNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(props.Place))
	})
	return _c
}

func (_c *Builder_WithPageNumber_Call) Return(_a0 config.Builder) *Builder_WithPageNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithPageNumber_Call) RunAndReturn(run func(string, props.Place) config.Builder) *Builder_WithPageNumber_Call {
	_c.Call.Return(run)
	return _c
}

// WithPageSize provides a mock function with given fields: size
func (_m *Builder) WithPageSize(size pagesize.Type) config.Builder {
	ret := _m.Called(size)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(pagesize.Type) config.Builder); ok {
		r0 = rf(size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithPageSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithPageSize'
type Builder_WithPageSize_Call struct {
	*mock.Call
}

// WithPageSize is a helper method to define mock.On call
//   - size pagesize.Type
func (_e *Builder_Expecter) WithPageSize(size interface{}) *Builder_WithPageSize_Call {
	return &Builder_WithPageSize_Call{Call: _e.mock.On("WithPageSize", size)}
}

func (_c *Builder_WithPageSize_Call) Run(run func(size pagesize.Type)) *Builder_WithPageSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(pagesize.Type))
	})
	return _c
}

func (_c *Builder_WithPageSize_Call) Return(_a0 config.Builder) *Builder_WithPageSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithPageSize_Call) RunAndReturn(run func(pagesize.Type) config.Builder) *Builder_WithPageSize_Call {
	_c.Call.Return(run)
	return _c
}

// WithProtection provides a mock function with given fields: protectionType, userPassword, ownerPassword
func (_m *Builder) WithProtection(protectionType protection.Type, userPassword string, ownerPassword string) config.Builder {
	ret := _m.Called(protectionType, userPassword, ownerPassword)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(protection.Type, string, string) config.Builder); ok {
		r0 = rf(protectionType, userPassword, ownerPassword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithProtection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithProtection'
type Builder_WithProtection_Call struct {
	*mock.Call
}

// WithProtection is a helper method to define mock.On call
//   - protectionType protection.Type
//   - userPassword string
//   - ownerPassword string
func (_e *Builder_Expecter) WithProtection(protectionType interface{}, userPassword interface{}, ownerPassword interface{}) *Builder_WithProtection_Call {
	return &Builder_WithProtection_Call{Call: _e.mock.On("WithProtection", protectionType, userPassword, ownerPassword)}
}

func (_c *Builder_WithProtection_Call) Run(run func(protectionType protection.Type, userPassword string, ownerPassword string)) *Builder_WithProtection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(protection.Type), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Builder_WithProtection_Call) Return(_a0 config.Builder) *Builder_WithProtection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithProtection_Call) RunAndReturn(run func(protection.Type, string, string) config.Builder) *Builder_WithProtection_Call {
	_c.Call.Return(run)
	return _c
}

// WithSubject provides a mock function with given fields: subject, isUTF8
func (_m *Builder) WithSubject(subject string, isUTF8 bool) config.Builder {
	ret := _m.Called(subject, isUTF8)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(string, bool) config.Builder); ok {
		r0 = rf(subject, isUTF8)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithSubject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithSubject'
type Builder_WithSubject_Call struct {
	*mock.Call
}

// WithSubject is a helper method to define mock.On call
//   - subject string
//   - isUTF8 bool
func (_e *Builder_Expecter) WithSubject(subject interface{}, isUTF8 interface{}) *Builder_WithSubject_Call {
	return &Builder_WithSubject_Call{Call: _e.mock.On("WithSubject", subject, isUTF8)}
}

func (_c *Builder_WithSubject_Call) Run(run func(subject string, isUTF8 bool)) *Builder_WithSubject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *Builder_WithSubject_Call) Return(_a0 config.Builder) *Builder_WithSubject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithSubject_Call) RunAndReturn(run func(string, bool) config.Builder) *Builder_WithSubject_Call {
	_c.Call.Return(run)
	return _c
}

// WithTitle provides a mock function with given fields: title, isUTF8
func (_m *Builder) WithTitle(title string, isUTF8 bool) config.Builder {
	ret := _m.Called(title, isUTF8)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(string, bool) config.Builder); ok {
		r0 = rf(title, isUTF8)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithTitle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTitle'
type Builder_WithTitle_Call struct {
	*mock.Call
}

// WithTitle is a helper method to define mock.On call
//   - title string
//   - isUTF8 bool
func (_e *Builder_Expecter) WithTitle(title interface{}, isUTF8 interface{}) *Builder_WithTitle_Call {
	return &Builder_WithTitle_Call{Call: _e.mock.On("WithTitle", title, isUTF8)}
}

func (_c *Builder_WithTitle_Call) Run(run func(title string, isUTF8 bool)) *Builder_WithTitle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *Builder_WithTitle_Call) Return(_a0 config.Builder) *Builder_WithTitle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithTitle_Call) RunAndReturn(run func(string, bool) config.Builder) *Builder_WithTitle_Call {
	_c.Call.Return(run)
	return _c
}

// WithWorkerPoolSize provides a mock function with given fields: poolSize
func (_m *Builder) WithWorkerPoolSize(poolSize int) config.Builder {
	ret := _m.Called(poolSize)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(int) config.Builder); ok {
		r0 = rf(poolSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithWorkerPoolSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithWorkerPoolSize'
type Builder_WithWorkerPoolSize_Call struct {
	*mock.Call
}

// WithWorkerPoolSize is a helper method to define mock.On call
//   - poolSize int
func (_e *Builder_Expecter) WithWorkerPoolSize(poolSize interface{}) *Builder_WithWorkerPoolSize_Call {
	return &Builder_WithWorkerPoolSize_Call{Call: _e.mock.On("WithWorkerPoolSize", poolSize)}
}

func (_c *Builder_WithWorkerPoolSize_Call) Run(run func(poolSize int)) *Builder_WithWorkerPoolSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Builder_WithWorkerPoolSize_Call) Return(_a0 config.Builder) *Builder_WithWorkerPoolSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithWorkerPoolSize_Call) RunAndReturn(run func(int) config.Builder) *Builder_WithWorkerPoolSize_Call {
	_c.Call.Return(run)
	return _c
}

// NewBuilder creates a new instance of Builder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBuilder(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Builder {
	mock := &Builder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
