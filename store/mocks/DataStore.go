// Copyright 2023 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.2.2. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	model "github.com/mendersoftware/deviceconnect/model"
)

// DataStore is an autogenerated mock type for the DataStore type
type DataStore struct {
	mock.Mock
}

// AllocateSession provides a mock function with given fields: ctx, sess
func (_m *DataStore) AllocateSession(ctx context.Context, sess *model.Session) error {
	ret := _m.Called(ctx, sess)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Session) error); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *DataStore) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDevice provides a mock function with given fields: ctx, tenantID, deviceID
func (_m *DataStore) DeleteDevice(ctx context.Context, tenantID string, deviceID string) error {
	ret := _m.Called(ctx, tenantID, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantID, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSession provides a mock function with given fields: ctx, sessionID
func (_m *DataStore) DeleteSession(ctx context.Context, sessionID string) (*model.Session, error) {
	ret := _m.Called(ctx, sessionID)

	var r0 *model.Session
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Session); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevice provides a mock function with given fields: ctx, tenantID, deviceID
func (_m *DataStore) GetDevice(ctx context.Context, tenantID string, deviceID string) (*model.Device, error) {
	ret := _m.Called(ctx, tenantID, deviceID)

	var r0 *model.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Device); ok {
		r0 = rf(ctx, tenantID, deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenantID, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSession provides a mock function with given fields: ctx, sessionID
func (_m *DataStore) GetSession(ctx context.Context, sessionID string) (*model.Session, error) {
	ret := _m.Called(ctx, sessionID)

	var r0 *model.Session
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Session); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertControlRecording provides a mock function with given fields: ctx, sessionID, sessionBytes
func (_m *DataStore) InsertControlRecording(ctx context.Context, sessionID string, sessionBytes []byte) error {
	ret := _m.Called(ctx, sessionID, sessionBytes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, sessionID, sessionBytes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertSessionRecording provides a mock function with given fields: ctx, sessionID, sessionBytes
func (_m *DataStore) InsertSessionRecording(ctx context.Context, sessionID string, sessionBytes []byte) error {
	ret := _m.Called(ctx, sessionID, sessionBytes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, sessionID, sessionBytes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields: ctx
func (_m *DataStore) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionDevice provides a mock function with given fields: ctx, tenantID, deviceID
func (_m *DataStore) ProvisionDevice(ctx context.Context, tenantID string, deviceID string) error {
	ret := _m.Called(ctx, tenantID, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantID, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDeviceConnected provides a mock function with given fields: ctx, tenantID, deviceID
func (_m *DataStore) SetDeviceConnected(ctx context.Context, tenantID string, deviceID string) (int64, error) {
	ret := _m.Called(ctx, tenantID, deviceID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, tenantID, deviceID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenantID, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDeviceDisconnected provides a mock function with given fields: ctx, tenantID, deviceID, version
func (_m *DataStore) SetDeviceDisconnected(ctx context.Context, tenantID string, deviceID string, version int64) error {
	ret := _m.Called(ctx, tenantID, deviceID, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) error); ok {
		r0 = rf(ctx, tenantID, deviceID, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertDeviceStatus provides a mock function with given fields: ctx, tenantID, deviceID, status
func (_m *DataStore) UpsertDeviceStatus(ctx context.Context, tenantID string, deviceID string, status string) error {
	ret := _m.Called(ctx, tenantID, deviceID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, tenantID, deviceID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteSessionRecords provides a mock function with given fields: ctx, sessionID, w
func (_m *DataStore) WriteSessionRecords(ctx context.Context, sessionID string, w io.Writer) error {
	ret := _m.Called(ctx, sessionID, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Writer) error); ok {
		r0 = rf(ctx, sessionID, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
