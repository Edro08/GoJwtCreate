package health

import (
	"github.com/dimiro1/health"
	"net/http"
	"reflect"
	"testing"
)

func TestChecker_CheckHandlerCustom(t *testing.T) {
	type fields struct {
		serverName string
	}
	handler := health.NewHandler()
	handler.AddInfo("service", "")
	handler.AddInfo("endpoint", "health")

	tests := []struct {
		name   string
		fields fields
		want   http.Handler
	}{
		{
			name: "Test Health Check handler Custom Successful",
			fields: fields{
				serverName: "",
			},
			want: handler,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &Checker{
				serverName: tt.fields.serverName,
			}
			if got := ch.CheckHandlerCustom(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckHandlerCustom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHealthChecker(t *testing.T) {
	type args struct {
		serverName string
	}

	tests := []struct {
		name string
		args args
		want *Checker
	}{
		{
			name: "Test New Health Check Successful",
			args: args{
				serverName: "",
			},
			want: NewHealthChecker(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealthChecker(tt.args.serverName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
