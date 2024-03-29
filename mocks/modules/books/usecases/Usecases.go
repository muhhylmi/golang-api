// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	web "golang-api/modules/books/models/web"

	wrapper "golang-api/utils/wrapper"
)

// Usecases is an autogenerated mock type for the Usecases type
type Usecases struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: ctx, payload
func (_m *Usecases) CreateBook(ctx context.Context, payload *web.RequestCreateBook) wrapper.Result {
	ret := _m.Called(ctx, payload)

	var r0 wrapper.Result
	if rf, ok := ret.Get(0).(func(context.Context, *web.RequestCreateBook) wrapper.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(wrapper.Result)
	}

	return r0
}

// CreateBookByGrpc provides a mock function with given fields: ctx, payload
func (_m *Usecases) CreateBookByGrpc(ctx context.Context, payload *web.RequestCreateBook) wrapper.Result {
	ret := _m.Called(ctx, payload)

	var r0 wrapper.Result
	if rf, ok := ret.Get(0).(func(context.Context, *web.RequestCreateBook) wrapper.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(wrapper.Result)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: ctx, payload
func (_m *Usecases) DeleteBook(ctx context.Context, payload *web.RequestDeleteBook) wrapper.Result {
	ret := _m.Called(ctx, payload)

	var r0 wrapper.Result
	if rf, ok := ret.Get(0).(func(context.Context, *web.RequestDeleteBook) wrapper.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(wrapper.Result)
	}

	return r0
}

// GetBook provides a mock function with given fields: ctx
func (_m *Usecases) GetBook(ctx context.Context) wrapper.Result {
	ret := _m.Called(ctx)

	var r0 wrapper.Result
	if rf, ok := ret.Get(0).(func(context.Context) wrapper.Result); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(wrapper.Result)
	}

	return r0
}

// GetBookSheetData provides a mock function with given fields: ctx, payload
func (_m *Usecases) GetBookSheetData(ctx context.Context, payload *web.GetBookSheetRequest) wrapper.Result {
	ret := _m.Called(ctx, payload)

	var r0 wrapper.Result
	if rf, ok := ret.Get(0).(func(context.Context, *web.GetBookSheetRequest) wrapper.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(wrapper.Result)
	}

	return r0
}

// GetDetailBook provides a mock function with given fields: ctx, payload
func (_m *Usecases) GetDetailBook(ctx context.Context, payload *web.RequestDetailBook) wrapper.Result {
	ret := _m.Called(ctx, payload)

	var r0 wrapper.Result
	if rf, ok := ret.Get(0).(func(context.Context, *web.RequestDetailBook) wrapper.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(wrapper.Result)
	}

	return r0
}

// UpdateBook provides a mock function with given fields: ctx, payload
func (_m *Usecases) UpdateBook(ctx context.Context, payload *web.RequestUpdateBook) wrapper.Result {
	ret := _m.Called(ctx, payload)

	var r0 wrapper.Result
	if rf, ok := ret.Get(0).(func(context.Context, *web.RequestUpdateBook) wrapper.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(wrapper.Result)
	}

	return r0
}

// NewUsecases creates a new instance of Usecases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsecases(t interface {
	mock.TestingT
	Cleanup(func())
}) *Usecases {
	mock := &Usecases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
