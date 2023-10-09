// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	fontstyle "github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	repository "github.com/johnfercher/maroto/v2/pkg/repository"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// AddUTF8Font provides a mock function with given fields: family, style, file
func (_m *Repository) AddUTF8Font(family string, style fontstyle.Type, file string) repository.Repository {
	ret := _m.Called(family, style, file)

	var r0 repository.Repository
	if rf, ok := ret.Get(0).(func(string, fontstyle.Type, string) repository.Repository); ok {
		r0 = rf(family, style, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Repository)
		}
	}

	return r0
}

// Repository_AddUTF8Font_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUTF8Font'
type Repository_AddUTF8Font_Call struct {
	*mock.Call
}

// AddUTF8Font is a helper method to define mock.On call
//   - family string
//   - style fontstyle.Type
//   - file string
func (_e *Repository_Expecter) AddUTF8Font(family interface{}, style interface{}, file interface{}) *Repository_AddUTF8Font_Call {
	return &Repository_AddUTF8Font_Call{Call: _e.mock.On("AddUTF8Font", family, style, file)}
}

func (_c *Repository_AddUTF8Font_Call) Run(run func(family string, style fontstyle.Type, file string)) *Repository_AddUTF8Font_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(fontstyle.Type), args[2].(string))
	})
	return _c
}

func (_c *Repository_AddUTF8Font_Call) Return(_a0 repository.Repository) *Repository_AddUTF8Font_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_AddUTF8Font_Call) RunAndReturn(run func(string, fontstyle.Type, string) repository.Repository) *Repository_AddUTF8Font_Call {
	_c.Call.Return(run)
	return _c
}

// Load provides a mock function with given fields:
func (_m *Repository) Load() ([]*entity.CustomFont, error) {
	ret := _m.Called()

	var r0 []*entity.CustomFont
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entity.CustomFont, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entity.CustomFont); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.CustomFont)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type Repository_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
func (_e *Repository_Expecter) Load() *Repository_Load_Call {
	return &Repository_Load_Call{Call: _e.mock.On("Load")}
}

func (_c *Repository_Load_Call) Run(run func()) *Repository_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_Load_Call) Return(_a0 []*entity.CustomFont, _a1 error) *Repository_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Load_Call) RunAndReturn(run func() ([]*entity.CustomFont, error)) *Repository_Load_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
