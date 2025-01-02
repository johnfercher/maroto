// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mappers "github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	mock "github.com/stretchr/testify/mock"
)

// GenerateComponent is an autogenerated mock type for the GenerateComponent type
type GenerateComponent struct {
	mock.Mock
}

type GenerateComponent_Expecter struct {
	mock *mock.Mock
}

func (_m *GenerateComponent) EXPECT() *GenerateComponent_Expecter {
	return &GenerateComponent_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: document, sourceKey
func (_m *GenerateComponent) Execute(document interface{}, sourceKey string) (mappers.OrderedComponents, error) {
	ret := _m.Called(document, sourceKey)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 mappers.OrderedComponents
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, string) (mappers.OrderedComponents, error)); ok {
		return rf(document, sourceKey)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string) mappers.OrderedComponents); ok {
		r0 = rf(document, sourceKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mappers.OrderedComponents)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string) error); ok {
		r1 = rf(document, sourceKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateComponent_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type GenerateComponent_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - document interface{}
//   - sourceKey string
func (_e *GenerateComponent_Expecter) Execute(document interface{}, sourceKey interface{}) *GenerateComponent_Execute_Call {
	return &GenerateComponent_Execute_Call{Call: _e.mock.On("Execute", document, sourceKey)}
}

func (_c *GenerateComponent_Execute_Call) Run(run func(document interface{}, sourceKey string)) *GenerateComponent_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(string))
	})
	return _c
}

func (_c *GenerateComponent_Execute_Call) Return(_a0 mappers.OrderedComponents, _a1 error) *GenerateComponent_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GenerateComponent_Execute_Call) RunAndReturn(run func(interface{}, string) (mappers.OrderedComponents, error)) *GenerateComponent_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewGenerateComponent creates a new instance of GenerateComponent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGenerateComponent(t interface {
	mock.TestingT
	Cleanup(func())
},
) *GenerateComponent {
	mock := &GenerateComponent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
