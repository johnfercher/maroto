// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	core "github.com/johnfercher/maroto/v2/pkg/core"
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	node "github.com/johnfercher/go-tree/node"
)

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

type Node_Expecter struct {
	mock *mock.Mock
}

func (_m *Node) EXPECT() *Node_Expecter {
	return &Node_Expecter{mock: &_m.Mock}
}

// GetStructure provides a mock function with given fields:
func (_m *Node) GetStructure() *node.Node[core.Structure] {
	ret := _m.Called()

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

// Node_GetStructure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStructure'
type Node_GetStructure_Call struct {
	*mock.Call
}

// GetStructure is a helper method to define mock.On call
func (_e *Node_Expecter) GetStructure() *Node_GetStructure_Call {
	return &Node_GetStructure_Call{Call: _e.mock.On("GetStructure")}
}

func (_c *Node_GetStructure_Call) Run(run func()) *Node_GetStructure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Node_GetStructure_Call) Return(_a0 *node.Node[core.Structure]) *Node_GetStructure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Node_GetStructure_Call) RunAndReturn(run func() *node.Node[core.Structure]) *Node_GetStructure_Call {
	_c.Call.Return(run)
	return _c
}

// SetConfig provides a mock function with given fields: config
func (_m *Node) SetConfig(config *entity.Config) {
	_m.Called(config)
}

// Node_SetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfig'
type Node_SetConfig_Call struct {
	*mock.Call
}

// SetConfig is a helper method to define mock.On call
//   - config *entity.Config
func (_e *Node_Expecter) SetConfig(config interface{}) *Node_SetConfig_Call {
	return &Node_SetConfig_Call{Call: _e.mock.On("SetConfig", config)}
}

func (_c *Node_SetConfig_Call) Run(run func(config *entity.Config)) *Node_SetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Config))
	})
	return _c
}

func (_c *Node_SetConfig_Call) Return() *Node_SetConfig_Call {
	_c.Call.Return()
	return _c
}

func (_c *Node_SetConfig_Call) RunAndReturn(run func(*entity.Config)) *Node_SetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewNode creates a new instance of Node. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNode(t interface {
	mock.TestingT
	Cleanup(func())
}) *Node {
	mock := &Node{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
