// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"
	mock "github.com/stretchr/testify/mock"

	props "github.com/johnfercher/maroto/v2/pkg/props"
)

// Math is an autogenerated mock type for the Math type
type Math struct {
	mock.Mock
}

type Math_Expecter struct {
	mock *mock.Mock
}

func (_m *Math) EXPECT() *Math_Expecter {
	return &Math_Expecter{mock: &_m.Mock}
}

// GetInnerCenterCell provides a mock function with given fields: inner, outer, percent
func (_m *Math) GetInnerCenterCell(inner *entity.Dimensions, outer *entity.Dimensions, percent float64) *entity.Cell {
	ret := _m.Called(inner, outer, percent)

	if len(ret) == 0 {
		panic("no return value specified for GetInnerCenterCell")
	}

	var r0 *entity.Cell
	if rf, ok := ret.Get(0).(func(*entity.Dimensions, *entity.Dimensions, float64) *entity.Cell); ok {
		r0 = rf(inner, outer, percent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Cell)
		}
	}

	return r0
}

// Math_GetInnerCenterCell_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInnerCenterCell'
type Math_GetInnerCenterCell_Call struct {
	*mock.Call
}

// GetInnerCenterCell is a helper method to define mock.On call
//   - inner *entity.Dimensions
//   - outer *entity.Dimensions
//   - percent float64
func (_e *Math_Expecter) GetInnerCenterCell(inner interface{}, outer interface{}, percent interface{}) *Math_GetInnerCenterCell_Call {
	return &Math_GetInnerCenterCell_Call{Call: _e.mock.On("GetInnerCenterCell", inner, outer, percent)}
}

func (_c *Math_GetInnerCenterCell_Call) Run(run func(inner *entity.Dimensions, outer *entity.Dimensions, percent float64)) *Math_GetInnerCenterCell_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Dimensions), args[1].(*entity.Dimensions), args[2].(float64))
	})
	return _c
}

func (_c *Math_GetInnerCenterCell_Call) Return(_a0 *entity.Cell) *Math_GetInnerCenterCell_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Math_GetInnerCenterCell_Call) RunAndReturn(run func(*entity.Dimensions, *entity.Dimensions, float64) *entity.Cell) *Math_GetInnerCenterCell_Call {
	_c.Call.Return(run)
	return _c
}

// GetInnerNonCenterCell provides a mock function with given fields: inner, outer, prop
func (_m *Math) GetInnerNonCenterCell(inner *entity.Dimensions, outer *entity.Dimensions, prop *props.Rect) *entity.Cell {
	ret := _m.Called(inner, outer, prop)

	if len(ret) == 0 {
		panic("no return value specified for GetInnerNonCenterCell")
	}

	var r0 *entity.Cell
	if rf, ok := ret.Get(0).(func(*entity.Dimensions, *entity.Dimensions, *props.Rect) *entity.Cell); ok {
		r0 = rf(inner, outer, prop)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Cell)
		}
	}

	return r0
}

// Math_GetInnerNonCenterCell_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInnerNonCenterCell'
type Math_GetInnerNonCenterCell_Call struct {
	*mock.Call
}

// GetInnerNonCenterCell is a helper method to define mock.On call
//   - inner *entity.Dimensions
//   - outer *entity.Dimensions
//   - prop *props.Rect
func (_e *Math_Expecter) GetInnerNonCenterCell(inner interface{}, outer interface{}, prop interface{}) *Math_GetInnerNonCenterCell_Call {
	return &Math_GetInnerNonCenterCell_Call{Call: _e.mock.On("GetInnerNonCenterCell", inner, outer, prop)}
}

func (_c *Math_GetInnerNonCenterCell_Call) Run(run func(inner *entity.Dimensions, outer *entity.Dimensions, prop *props.Rect)) *Math_GetInnerNonCenterCell_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Dimensions), args[1].(*entity.Dimensions), args[2].(*props.Rect))
	})
	return _c
}

func (_c *Math_GetInnerNonCenterCell_Call) Return(_a0 *entity.Cell) *Math_GetInnerNonCenterCell_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Math_GetInnerNonCenterCell_Call) RunAndReturn(run func(*entity.Dimensions, *entity.Dimensions, *props.Rect) *entity.Cell) *Math_GetInnerNonCenterCell_Call {
	_c.Call.Return(run)
	return _c
}

// NewMath creates a new instance of Math. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMath(t interface {
	mock.TestingT
	Cleanup(func())
}) *Math {
	mock := &Math{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
