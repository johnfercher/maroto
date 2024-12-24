// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	core "github.com/johnfercher/maroto/v2/pkg/core"
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	node "github.com/johnfercher/go-tree/node"
)

// ProviderComponent is an autogenerated mock type for the ProviderComponent type
type ProviderComponent struct {
	mock.Mock
}

type ProviderComponent_Expecter struct {
	mock *mock.Mock
}

func (_m *ProviderComponent) EXPECT() *ProviderComponent_Expecter {
	return &ProviderComponent_Expecter{mock: &_m.Mock}
}

// GetStructure provides a mock function with given fields:
func (_m *ProviderComponent) GetStructure() *node.Node[core.Structure] {
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

// ProviderComponent_GetStructure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStructure'
type ProviderComponent_GetStructure_Call struct {
	*mock.Call
}

// GetStructure is a helper method to define mock.On call
func (_e *ProviderComponent_Expecter) GetStructure() *ProviderComponent_GetStructure_Call {
	return &ProviderComponent_GetStructure_Call{Call: _e.mock.On("GetStructure")}
}

func (_c *ProviderComponent_GetStructure_Call) Run(run func()) *ProviderComponent_GetStructure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ProviderComponent_GetStructure_Call) Return(_a0 *node.Node[core.Structure]) *ProviderComponent_GetStructure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProviderComponent_GetStructure_Call) RunAndReturn(run func() *node.Node[core.Structure]) *ProviderComponent_GetStructure_Call {
	_c.Call.Return(run)
	return _c
}

// SetConfig provides a mock function with given fields: config
func (_m *ProviderComponent) SetConfig(config *entity.Config) {
	_m.Called(config)
}

// ProviderComponent_SetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfig'
type ProviderComponent_SetConfig_Call struct {
	*mock.Call
}

// SetConfig is a helper method to define mock.On call
//   - config *entity.Config
func (_e *ProviderComponent_Expecter) SetConfig(config interface{}) *ProviderComponent_SetConfig_Call {
	return &ProviderComponent_SetConfig_Call{Call: _e.mock.On("SetConfig", config)}
}

func (_c *ProviderComponent_SetConfig_Call) Run(run func(config *entity.Config)) *ProviderComponent_SetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Config))
	})
	return _c
}

func (_c *ProviderComponent_SetConfig_Call) Return() *ProviderComponent_SetConfig_Call {
	_c.Call.Return()
	return _c
}

func (_c *ProviderComponent_SetConfig_Call) RunAndReturn(run func(*entity.Config)) *ProviderComponent_SetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewProviderComponent creates a new instance of ProviderComponent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProviderComponent(t interface {
	mock.TestingT
	Cleanup(func())
},
) *ProviderComponent {
	mock := &ProviderComponent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
