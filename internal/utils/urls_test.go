package utils

import (
	"reflect"
	"testing"
)

func TestBuildForwardURL(t *testing.T) {
	type args struct {
		baseURL string
		path    string
		query   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"case 1",
			args{
				"",
				"",
				"",
			},
			"",
			true,
		},
		{
			"case 2",
			args{
				"www.test.com",
				"",
				"",
			},
			"http://www.test.com",
			false,
		},
		{
			"case 3",
			args{
				"www.test.com",
				"/api/v1/ping",
				"data=6753478",
			},
			"http://www.test.com/api/v1/ping?data=6753478",
			false,
		},
		{
			"case 4",
			args{
				"http://127.0.0.1:178",
				"",
				"",
			},
			"http://127.0.0.1:178",
			false,
		},
		{
			"case 5",
			args{
				"http://123.23.2.2",
				"",
				"",
			},
			"http://123.23.2.2",
			false,
		},
		{
			"case 5",
			args{
				"123.23.2.2",
				"",
				"",
			},
			"http://123.23.2.2",
			false,
		},
		{
			"case 5",
			args{
				"http://123.23.2.2:8989",
				"",
				"",
			},
			"http://123.23.2.2:8989",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildURL(tt.args.baseURL, tt.args.path, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("BuildURL() got = %v, want %v", got.String(), tt.want)
			}
			if got == nil && !reflect.DeepEqual("", tt.want) {
				t.Errorf("BuildURL() got = %v, want %v", got.String(), tt.want)
			}
		})
	}
}
