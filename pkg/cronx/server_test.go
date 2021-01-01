package cronx

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServerController_APIJobs(t *testing.T) {
	type fields struct {
		CommandController *CommandController
	}
	tests := []struct {
		name    string
		target  string
		fields  fields
		expect  int
		wantErr bool
	}{
		{
			name:   "Success",
			target: "/api/jobs",
			fields: fields{
				CommandController: &CommandController{
					CreatedTime: time.Now(),
					Location:    time.Local,
				},
			},
			expect:  http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			ctrl := &ServerController{
				CommandController: tt.fields.CommandController,
			}
			if assert.NoError(t, ctrl.APIJobs(c)) {
				assert.Equal(t, tt.expect, rec.Code)
			}
		})
	}
}

func TestServerController_HealthCheck(t *testing.T) {
	type fields struct {
		CommandController *CommandController
	}
	tests := []struct {
		name    string
		target  string
		fields  fields
		expect  int
		wantErr bool
	}{
		{
			name:   "Success",
			target: "/",
			fields: fields{
				CommandController: &CommandController{
					CreatedTime: time.Now(),
					Location:    time.Local,
				},
			},
			expect:  http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			ctrl := &ServerController{
				CommandController: tt.fields.CommandController,
			}
			if assert.NoError(t, ctrl.HealthCheck(c)) {
				assert.Equal(t, tt.expect, rec.Code)
			}
		})
	}
}

func TestServerController_Jobs(t *testing.T) {
	type fields struct {
		CommandController *CommandController
	}
	tests := []struct {
		name    string
		target  string
		fields  fields
		expect  int
		wantErr bool
	}{
		{
			name:   "Success",
			target: "/jobs",
			fields: fields{
				CommandController: &CommandController{
					CreatedTime: time.Now(),
					Location:    time.Local,
				},
			},
			expect:  http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			ctrl := &ServerController{
				CommandController: tt.fields.CommandController,
			}
			if assert.NoError(t, ctrl.Jobs(c)) {
				assert.Equal(t, tt.expect, rec.Code)
			}
		})
	}
}
