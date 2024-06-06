// Copyright 2022 Northern.tech AS
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

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// DeleteDevice provides a mock function with given fields: ctx, tenantID, deviceID
func (_m *App) DeleteDevice(ctx context.Context, tenantID string, deviceID string) error {
	ret := _m.Called(ctx, tenantID, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantID, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DownloadFile provides a mock function with given fields: ctx, userID, deviceID, path
func (_m *App) DownloadFile(ctx context.Context, userID string, deviceID string, path string) error {
	ret := _m.Called(ctx, userID, deviceID, path)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, userID, deviceID, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FreeUserSession provides a mock function with given fields: ctx, sessionID, sessionTypes
func (_m *App) FreeUserSession(ctx context.Context, sessionID string, sessionTypes []string) error {
	ret := _m.Called(ctx, sessionID, sessionTypes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, sessionID, sessionTypes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetControlRecorder provides a mock function with given fields: ctx, sessionID
func (_m *App) GetControlRecorder(ctx context.Context, sessionID string) io.Writer {
	ret := _m.Called(ctx, sessionID)

	var r0 io.Writer
	if rf, ok := ret.Get(0).(func(context.Context, string) io.Writer); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Writer)
		}
	}

	return r0
}

// GetDevice provides a mock function with given fields: ctx, tenantID, deviceID
func (_m *App) GetDevice(ctx context.Context, tenantID string, deviceID string) (*model.Device, error) {
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

// GetRecorder provides a mock function with given fields: ctx, sessionID
func (_m *App) GetRecorder(ctx context.Context, sessionID string) io.Writer {
	ret := _m.Called(ctx, sessionID)

	var r0 io.Writer
	if rf, ok := ret.Get(0).(func(context.Context, string) io.Writer); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Writer)
		}
	}

	return r0
}

// GetSessionRecording provides a mock function with given fields: ctx, id, w
func (_m *App) GetSessionRecording(ctx context.Context, id string, w io.Writer) error {
	ret := _m.Called(ctx, id, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Writer) error); ok {
		r0 = rf(ctx, id, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *App) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogUserSession provides a mock function with given fields: ctx, sess, sessionType
func (_m *App) LogUserSession(ctx context.Context, sess *model.Session, sessionType string) error {
	ret := _m.Called(ctx, sess, sessionType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Session, string) error); ok {
		r0 = rf(ctx, sess, sessionType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrepareUserSession provides a mock function with given fields: ctx, sess
func (_m *App) PrepareUserSession(ctx context.Context, sess *model.Session) error {
	ret := _m.Called(ctx, sess)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Session) error); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionDevice provides a mock function with given fields: ctx, tenantID, device
func (_m *App) ProvisionDevice(ctx context.Context, tenantID string, device *model.Device) error {
	ret := _m.Called(ctx, tenantID, device)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Device) error); ok {
		r0 = rf(ctx, tenantID, device)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveSessionRecording provides a mock function with given fields: ctx, id, sessionBytes
func (_m *App) SaveSessionRecording(ctx context.Context, id string, sessionBytes []byte) error {
	ret := _m.Called(ctx, id, sessionBytes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, id, sessionBytes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDeviceConnected provides a mock function with given fields: ctx, tenantID, deviceID
func (_m *App) SetDeviceConnected(ctx context.Context, tenantID string, deviceID string) (int64, error) {
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
func (_m *App) SetDeviceDisconnected(ctx context.Context, tenantID string, deviceID string, version int64) error {
	ret := _m.Called(ctx, tenantID, deviceID, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) error); ok {
		r0 = rf(ctx, tenantID, deviceID, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceStatus provides a mock function with given fields: ctx, tenantID, deviceID, status
func (_m *App) UpdateDeviceStatus(ctx context.Context, tenantID string, deviceID string, status string) error {
	ret := _m.Called(ctx, tenantID, deviceID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, tenantID, deviceID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadFile provides a mock function with given fields: ctx, userID, deviceID, path
func (_m *App) UploadFile(ctx context.Context, userID string, deviceID string, path string) error {
	ret := _m.Called(ctx, userID, deviceID, path)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, userID, deviceID, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
