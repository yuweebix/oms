// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// MockCache is an autogenerated mock type for the cache type
type MockCache struct {
	mock.Mock
}

type MockCache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCache) EXPECT() *MockCache_Expecter {
	return &MockCache_Expecter{mock: &_m.Mock}
}

// CreateOrder provides a mock function with given fields: ctx, o
func (_m *MockCache) CreateOrder(ctx context.Context, o *models.Order) {
	_m.Called(ctx, o)
}

// MockCache_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type MockCache_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - o *models.Order
func (_e *MockCache_Expecter) CreateOrder(ctx interface{}, o interface{}) *MockCache_CreateOrder_Call {
	return &MockCache_CreateOrder_Call{Call: _e.mock.On("CreateOrder", ctx, o)}
}

func (_c *MockCache_CreateOrder_Call) Run(run func(ctx context.Context, o *models.Order)) *MockCache_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Order))
	})
	return _c
}

func (_c *MockCache_CreateOrder_Call) Return() *MockCache_CreateOrder_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_CreateOrder_Call) RunAndReturn(run func(context.Context, *models.Order)) *MockCache_CreateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteOrder provides a mock function with given fields: ctx, o
func (_m *MockCache) DeleteOrder(ctx context.Context, o *models.Order) {
	_m.Called(ctx, o)
}

// MockCache_DeleteOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOrder'
type MockCache_DeleteOrder_Call struct {
	*mock.Call
}

// DeleteOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - o *models.Order
func (_e *MockCache_Expecter) DeleteOrder(ctx interface{}, o interface{}) *MockCache_DeleteOrder_Call {
	return &MockCache_DeleteOrder_Call{Call: _e.mock.On("DeleteOrder", ctx, o)}
}

func (_c *MockCache_DeleteOrder_Call) Run(run func(ctx context.Context, o *models.Order)) *MockCache_DeleteOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Order))
	})
	return _c
}

func (_c *MockCache_DeleteOrder_Call) Return() *MockCache_DeleteOrder_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_DeleteOrder_Call) RunAndReturn(run func(context.Context, *models.Order)) *MockCache_DeleteOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrder provides a mock function with given fields: ctx, o
func (_m *MockCache) GetOrder(ctx context.Context, o *models.Order) *models.Order {
	ret := _m.Called(ctx, o)

	if len(ret) == 0 {
		panic("no return value specified for GetOrder")
	}

	var r0 *models.Order
	if rf, ok := ret.Get(0).(func(context.Context, *models.Order) *models.Order); ok {
		r0 = rf(ctx, o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	return r0
}

// MockCache_GetOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrder'
type MockCache_GetOrder_Call struct {
	*mock.Call
}

// GetOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - o *models.Order
func (_e *MockCache_Expecter) GetOrder(ctx interface{}, o interface{}) *MockCache_GetOrder_Call {
	return &MockCache_GetOrder_Call{Call: _e.mock.On("GetOrder", ctx, o)}
}

func (_c *MockCache_GetOrder_Call) Run(run func(ctx context.Context, o *models.Order)) *MockCache_GetOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Order))
	})
	return _c
}

func (_c *MockCache_GetOrder_Call) Return(_a0 *models.Order) *MockCache_GetOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_GetOrder_Call) RunAndReturn(run func(context.Context, *models.Order) *models.Order) *MockCache_GetOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrders provides a mock function with given fields: ctx, userID, limit, offset, isStored
func (_m *MockCache) GetOrders(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool) []*models.Order {
	ret := _m.Called(ctx, userID, limit, offset, isStored)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, uint64, bool) []*models.Order); ok {
		r0 = rf(ctx, userID, limit, offset, isStored)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	return r0
}

// MockCache_GetOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrders'
type MockCache_GetOrders_Call struct {
	*mock.Call
}

// GetOrders is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint64
//   - limit uint64
//   - offset uint64
//   - isStored bool
func (_e *MockCache_Expecter) GetOrders(ctx interface{}, userID interface{}, limit interface{}, offset interface{}, isStored interface{}) *MockCache_GetOrders_Call {
	return &MockCache_GetOrders_Call{Call: _e.mock.On("GetOrders", ctx, userID, limit, offset, isStored)}
}

func (_c *MockCache_GetOrders_Call) Run(run func(ctx context.Context, userID uint64, limit uint64, offset uint64, isStored bool)) *MockCache_GetOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64), args[3].(uint64), args[4].(bool))
	})
	return _c
}

func (_c *MockCache_GetOrders_Call) Return(_a0 []*models.Order) *MockCache_GetOrders_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_GetOrders_Call) RunAndReturn(run func(context.Context, uint64, uint64, uint64, bool) []*models.Order) *MockCache_GetOrders_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrdersForDelivery provides a mock function with given fields: ctx, orderIDs
func (_m *MockCache) GetOrdersForDelivery(ctx context.Context, orderIDs []uint64) []*models.Order {
	ret := _m.Called(ctx, orderIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetOrdersForDelivery")
	}

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(context.Context, []uint64) []*models.Order); ok {
		r0 = rf(ctx, orderIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	return r0
}

// MockCache_GetOrdersForDelivery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrdersForDelivery'
type MockCache_GetOrdersForDelivery_Call struct {
	*mock.Call
}

// GetOrdersForDelivery is a helper method to define mock.On call
//   - ctx context.Context
//   - orderIDs []uint64
func (_e *MockCache_Expecter) GetOrdersForDelivery(ctx interface{}, orderIDs interface{}) *MockCache_GetOrdersForDelivery_Call {
	return &MockCache_GetOrdersForDelivery_Call{Call: _e.mock.On("GetOrdersForDelivery", ctx, orderIDs)}
}

func (_c *MockCache_GetOrdersForDelivery_Call) Run(run func(ctx context.Context, orderIDs []uint64)) *MockCache_GetOrdersForDelivery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint64))
	})
	return _c
}

func (_c *MockCache_GetOrdersForDelivery_Call) Return(_a0 []*models.Order) *MockCache_GetOrdersForDelivery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_GetOrdersForDelivery_Call) RunAndReturn(run func(context.Context, []uint64) []*models.Order) *MockCache_GetOrdersForDelivery_Call {
	_c.Call.Return(run)
	return _c
}

// GetReturns provides a mock function with given fields: ctx, limit, offset
func (_m *MockCache) GetReturns(ctx context.Context, limit uint64, offset uint64) []*models.Order {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetReturns")
	}

	var r0 []*models.Order
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) []*models.Order); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Order)
		}
	}

	return r0
}

// MockCache_GetReturns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReturns'
type MockCache_GetReturns_Call struct {
	*mock.Call
}

// GetReturns is a helper method to define mock.On call
//   - ctx context.Context
//   - limit uint64
//   - offset uint64
func (_e *MockCache_Expecter) GetReturns(ctx interface{}, limit interface{}, offset interface{}) *MockCache_GetReturns_Call {
	return &MockCache_GetReturns_Call{Call: _e.mock.On("GetReturns", ctx, limit, offset)}
}

func (_c *MockCache_GetReturns_Call) Run(run func(ctx context.Context, limit uint64, offset uint64)) *MockCache_GetReturns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64))
	})
	return _c
}

func (_c *MockCache_GetReturns_Call) Return(_a0 []*models.Order) *MockCache_GetReturns_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCache_GetReturns_Call) RunAndReturn(run func(context.Context, uint64, uint64) []*models.Order) *MockCache_GetReturns_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOrder provides a mock function with given fields: ctx, o
func (_m *MockCache) UpdateOrder(ctx context.Context, o *models.Order) {
	_m.Called(ctx, o)
}

// MockCache_UpdateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOrder'
type MockCache_UpdateOrder_Call struct {
	*mock.Call
}

// UpdateOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - o *models.Order
func (_e *MockCache_Expecter) UpdateOrder(ctx interface{}, o interface{}) *MockCache_UpdateOrder_Call {
	return &MockCache_UpdateOrder_Call{Call: _e.mock.On("UpdateOrder", ctx, o)}
}

func (_c *MockCache_UpdateOrder_Call) Run(run func(ctx context.Context, o *models.Order)) *MockCache_UpdateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Order))
	})
	return _c
}

func (_c *MockCache_UpdateOrder_Call) Return() *MockCache_UpdateOrder_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCache_UpdateOrder_Call) RunAndReturn(run func(context.Context, *models.Order)) *MockCache_UpdateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCache creates a new instance of MockCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCache {
	mock := &MockCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
