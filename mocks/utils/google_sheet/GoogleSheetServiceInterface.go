// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	sheets "google.golang.org/api/sheets/v4"
)

// GoogleSheetServiceInterface is an autogenerated mock type for the GoogleSheetServiceInterface type
type GoogleSheetServiceInterface struct {
	mock.Mock
}

// GetSheetData provides a mock function with given fields: sheetId, sheetRange
func (_m *GoogleSheetServiceInterface) GetSheetData(sheetId string, sheetRange string) (*sheets.ValueRange, error) {
	ret := _m.Called(sheetId, sheetRange)

	var r0 *sheets.ValueRange
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*sheets.ValueRange, error)); ok {
		return rf(sheetId, sheetRange)
	}
	if rf, ok := ret.Get(0).(func(string, string) *sheets.ValueRange); ok {
		r0 = rf(sheetId, sheetRange)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sheets.ValueRange)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(sheetId, sheetRange)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSheetDataWithFilter provides a mock function with given fields: sheetId, sheetRange
func (_m *GoogleSheetServiceInterface) GetSheetDataWithFilter(sheetId string, sheetRange string) (*sheets.ValueRange, []int, error) {
	ret := _m.Called(sheetId, sheetRange)

	var r0 *sheets.ValueRange
	var r1 []int
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (*sheets.ValueRange, []int, error)); ok {
		return rf(sheetId, sheetRange)
	}
	if rf, ok := ret.Get(0).(func(string, string) *sheets.ValueRange); ok {
		r0 = rf(sheetId, sheetRange)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sheets.ValueRange)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) []int); ok {
		r1 = rf(sheetId, sheetRange)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]int)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(sheetId, sheetRange)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewGoogleSheetServiceInterface creates a new instance of GoogleSheetServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGoogleSheetServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GoogleSheetServiceInterface {
	mock := &GoogleSheetServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}