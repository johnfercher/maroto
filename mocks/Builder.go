// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	config "github.com/johnfercher/maroto/pkg/v2/config"
	mock "github.com/stretchr/testify/mock"

	provider "github.com/johnfercher/maroto/pkg/v2/provider"
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
func (_m *Builder) Build() *config.Maroto {
	ret := _m.Called()

	var r0 *config.Maroto
	if rf, ok := ret.Get(0).(func() *config.Maroto); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.Maroto)
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

func (_c *Builder_Build_Call) Return(_a0 *config.Maroto) *Builder_Build_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_Build_Call) RunAndReturn(run func() *config.Maroto) *Builder_Build_Call {
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

// WithDimensions provides a mock function with given fields: dimensions
func (_m *Builder) WithDimensions(dimensions *config.Dimensions) config.Builder {
	ret := _m.Called(dimensions)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(*config.Dimensions) config.Builder); ok {
		r0 = rf(dimensions)
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
//   - dimensions *config.Dimensions
func (_e *Builder_Expecter) WithDimensions(dimensions interface{}) *Builder_WithDimensions_Call {
	return &Builder_WithDimensions_Call{Call: _e.mock.On("WithDimensions", dimensions)}
}

func (_c *Builder_WithDimensions_Call) Run(run func(dimensions *config.Dimensions)) *Builder_WithDimensions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*config.Dimensions))
	})
	return _c
}

func (_c *Builder_WithDimensions_Call) Return(_a0 config.Builder) *Builder_WithDimensions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithDimensions_Call) RunAndReturn(run func(*config.Dimensions) config.Builder) *Builder_WithDimensions_Call {
	_c.Call.Return(run)
	return _c
}

// WithFont provides a mock function with given fields: font
func (_m *Builder) WithFont(font *config.Font) config.Builder {
	ret := _m.Called(font)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(*config.Font) config.Builder); ok {
		r0 = rf(font)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithFont_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithFont'
type Builder_WithFont_Call struct {
	*mock.Call
}

// WithFont is a helper method to define mock.On call
//   - font *config.Font
func (_e *Builder_Expecter) WithFont(font interface{}) *Builder_WithFont_Call {
	return &Builder_WithFont_Call{Call: _e.mock.On("WithFont", font)}
}

func (_c *Builder_WithFont_Call) Run(run func(font *config.Font)) *Builder_WithFont_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*config.Font))
	})
	return _c
}

func (_c *Builder_WithFont_Call) Return(_a0 config.Builder) *Builder_WithFont_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithFont_Call) RunAndReturn(run func(*config.Font) config.Builder) *Builder_WithFont_Call {
	_c.Call.Return(run)
	return _c
}

// WithMargins provides a mock function with given fields: margins
func (_m *Builder) WithMargins(margins *config.Margins) config.Builder {
	ret := _m.Called(margins)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(*config.Margins) config.Builder); ok {
		r0 = rf(margins)
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
//   - margins *config.Margins
func (_e *Builder_Expecter) WithMargins(margins interface{}) *Builder_WithMargins_Call {
	return &Builder_WithMargins_Call{Call: _e.mock.On("WithMargins", margins)}
}

func (_c *Builder_WithMargins_Call) Run(run func(margins *config.Margins)) *Builder_WithMargins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*config.Margins))
	})
	return _c
}

func (_c *Builder_WithMargins_Call) Return(_a0 config.Builder) *Builder_WithMargins_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithMargins_Call) RunAndReturn(run func(*config.Margins) config.Builder) *Builder_WithMargins_Call {
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

// WithPageSize provides a mock function with given fields: size
func (_m *Builder) WithPageSize(size config.PageSize) config.Builder {
	ret := _m.Called(size)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(config.PageSize) config.Builder); ok {
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
//   - size config.PageSize
func (_e *Builder_Expecter) WithPageSize(size interface{}) *Builder_WithPageSize_Call {
	return &Builder_WithPageSize_Call{Call: _e.mock.On("WithPageSize", size)}
}

func (_c *Builder_WithPageSize_Call) Run(run func(size config.PageSize)) *Builder_WithPageSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(config.PageSize))
	})
	return _c
}

func (_c *Builder_WithPageSize_Call) Return(_a0 config.Builder) *Builder_WithPageSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithPageSize_Call) RunAndReturn(run func(config.PageSize) config.Builder) *Builder_WithPageSize_Call {
	_c.Call.Return(run)
	return _c
}

// WithProvider provides a mock function with given fields: providerType
func (_m *Builder) WithProvider(providerType provider.Type) config.Builder {
	ret := _m.Called(providerType)

	var r0 config.Builder
	if rf, ok := ret.Get(0).(func(provider.Type) config.Builder); ok {
		r0 = rf(providerType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Builder)
		}
	}

	return r0
}

// Builder_WithProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithProvider'
type Builder_WithProvider_Call struct {
	*mock.Call
}

// WithProvider is a helper method to define mock.On call
//   - providerType provider.Type
func (_e *Builder_Expecter) WithProvider(providerType interface{}) *Builder_WithProvider_Call {
	return &Builder_WithProvider_Call{Call: _e.mock.On("WithProvider", providerType)}
}

func (_c *Builder_WithProvider_Call) Run(run func(providerType provider.Type)) *Builder_WithProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(provider.Type))
	})
	return _c
}

func (_c *Builder_WithProvider_Call) Return(_a0 config.Builder) *Builder_WithProvider_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Builder_WithProvider_Call) RunAndReturn(run func(provider.Type) config.Builder) *Builder_WithProvider_Call {
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
}) *Builder {
	mock := &Builder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
