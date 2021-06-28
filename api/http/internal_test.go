// Copyright 2021 Northern.tech AS
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

package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mendersoftware/deviceconnect/app"
	app_mocks "github.com/mendersoftware/deviceconnect/app/mocks"
	nats_mocks "github.com/mendersoftware/deviceconnect/client/nats/mocks"
	"github.com/mendersoftware/deviceconnect/model"
)

var contextMatcher = mock.MatchedBy(func(_ context.Context) bool {
	return true
})

func TestProvision(t *testing.T) {
	testCases := []struct {
		Name               string
		TenantID           string
		Tenant             string
		ProvisionTenantErr error
		HTTPStatus         int
	}{
		{
			Name:       "ok",
			TenantID:   "1234",
			Tenant:     `{"tenant_id": "1234"}`,
			HTTPStatus: http.StatusCreated,
		},
		{
			Name:       "ko, empty payload",
			Tenant:     ``,
			HTTPStatus: http.StatusBadRequest,
		},
		{
			Name:       "ko, bad payload",
			Tenant:     `...`,
			HTTPStatus: http.StatusBadRequest,
		},
		{
			Name:       "ko, empty tenant ID",
			Tenant:     `{"tenant_id": ""}`,
			HTTPStatus: http.StatusBadRequest,
		},
		{
			Name:               "ko, error",
			TenantID:           "1234",
			Tenant:             `{"tenant_id": "1234"}`,
			ProvisionTenantErr: errors.New("error"),
			HTTPStatus:         http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			deviceConnectApp := &app_mocks.App{}
			if tc.TenantID != "" {
				deviceConnectApp.On("ProvisionTenant",
					mock.MatchedBy(func(_ context.Context) bool {
						return true
					}),
					&model.Tenant{TenantID: tc.TenantID},
				).Return(tc.ProvisionTenantErr)
			}

			router, _ := NewRouter(deviceConnectApp, nil)

			req, err := http.NewRequest("POST", APIURLInternalTenants, strings.NewReader(tc.Tenant))
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.HTTPStatus, w.Code)
			if tc.HTTPStatus == http.StatusNoContent {
				assert.Nil(t, w.Body.Bytes())
			}

			deviceConnectApp.AssertExpectations(t)
		})
	}
}

func TestInternalCheckUpdate(t *testing.T) {
	testCases := []struct {
		Name     string
		TenantID string
		DeviceID string

		GetDevice      *model.Device
		GetDeviceError error

		PublishErr error

		HTTPStatus int
	}{
		{
			Name:     "ok",
			DeviceID: "1234567890",

			GetDevice: &model.Device{
				ID:     "1234567890",
				Status: model.DeviceStatusConnected,
			},

			HTTPStatus: http.StatusAccepted,
		},
		{
			Name:     "ok, with tenantID",
			DeviceID: "1234567890",
			TenantID: "tenant_id",

			GetDevice: &model.Device{
				ID:     "1234567890",
				Status: model.DeviceStatusConnected,
			},

			HTTPStatus: http.StatusAccepted,
		},
		{
			Name:     "ko, not found",
			DeviceID: "1234567890",

			GetDeviceError: app.ErrDeviceNotFound,

			HTTPStatus: 404,
		},
		{
			Name:     "ko, other error",
			DeviceID: "1234567890",

			GetDeviceError: errors.New("error"),

			HTTPStatus: 400,
		},
		{
			Name:     "ko, device not connected",
			DeviceID: "1234567890",

			GetDevice: &model.Device{
				ID:     "1234567890",
				Status: model.DeviceStatusDisconnected,
			},

			HTTPStatus: http.StatusConflict,
		},
		{
			Name:     "ko, publish error",
			DeviceID: "1234567890",

			GetDevice: &model.Device{
				ID:     "1234567890",
				Status: model.DeviceStatusConnected,
			},

			PublishErr: errors.New("error"),

			HTTPStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			app := &app_mocks.App{}
			defer app.AssertExpectations(t)

			natsClient := &nats_mocks.Client{}
			defer natsClient.AssertExpectations(t)

			router, _ := NewRouter(app, natsClient)
			s := httptest.NewServer(router)
			defer s.Close()

			url := strings.Replace(APIURLInternalDevicesIDCheckUpdate, ":tenantId", tc.TenantID, 1)
			url = strings.Replace(url, ":deviceId", tc.DeviceID, 1)
			req, err := http.NewRequest("POST", "http://localhost"+url, nil)

			app.On("GetDevice",
				contextMatcher,
				tc.TenantID,
				tc.DeviceID,
			).Return(tc.GetDevice, tc.GetDeviceError)

			if tc.GetDeviceError == nil && tc.GetDevice != nil &&
				tc.GetDevice.Status == model.DeviceStatusConnected {
				natsClient.On("Publish",
					contextMatcher,
					mock.AnythingOfType("string"),
					mock.AnythingOfType("[]uint8"),
				).Return(tc.PublishErr)
			}
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.HTTPStatus, w.Code)
		})
	}
}

func TestInternalSendInventory(t *testing.T) {
	testCases := []struct {
		Name     string
		TenantID string
		DeviceID string

		GetDevice      *model.Device
		GetDeviceError error

		PublishErr error

		HTTPStatus int
	}{
		{
			Name:     "ok",
			DeviceID: "1234567890",

			GetDevice: &model.Device{
				ID:     "1234567890",
				Status: model.DeviceStatusConnected,
			},

			HTTPStatus: http.StatusAccepted,
		},
		{
			Name:     "ok, with tenantID",
			DeviceID: "1234567890",
			TenantID: "tenant_id",

			GetDevice: &model.Device{
				ID:     "1234567890",
				Status: model.DeviceStatusConnected,
			},

			HTTPStatus: http.StatusAccepted,
		},
		{
			Name:     "ko, not found",
			DeviceID: "1234567890",

			GetDeviceError: app.ErrDeviceNotFound,

			HTTPStatus: 404,
		},
		{
			Name:     "ko, other error",
			DeviceID: "1234567890",

			GetDeviceError: errors.New("error"),

			HTTPStatus: 400,
		},
		{
			Name:     "ko, device not connected",
			DeviceID: "1234567890",

			GetDevice: &model.Device{
				ID:     "1234567890",
				Status: model.DeviceStatusDisconnected,
			},

			HTTPStatus: http.StatusConflict,
		},
		{
			Name:     "ko, publish error",
			DeviceID: "1234567890",

			GetDevice: &model.Device{
				ID:     "1234567890",
				Status: model.DeviceStatusConnected,
			},

			PublishErr: errors.New("error"),

			HTTPStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			app := &app_mocks.App{}
			defer app.AssertExpectations(t)

			natsClient := &nats_mocks.Client{}
			defer natsClient.AssertExpectations(t)

			router, _ := NewRouter(app, natsClient)
			s := httptest.NewServer(router)
			defer s.Close()

			url := strings.Replace(APIURLInternalDevicesIDSendInventory, ":tenantId", tc.TenantID, 1)
			url = strings.Replace(url, ":deviceId", tc.DeviceID, 1)
			req, err := http.NewRequest("POST", "http://localhost"+url, nil)

			app.On("GetDevice",
				mock.MatchedBy(func(_ context.Context) bool {
					return true
				}),
				tc.TenantID,
				tc.DeviceID,
			).Return(tc.GetDevice, tc.GetDeviceError)

			if tc.GetDeviceError == nil && tc.GetDevice != nil &&
				tc.GetDevice.Status == model.DeviceStatusConnected {
				natsClient.On("Publish",
					contextMatcher,
					mock.AnythingOfType("string"),
					mock.AnythingOfType("[]uint8"),
				).Return(tc.PublishErr)
			}
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.HTTPStatus, w.Code)
		})
	}
}
