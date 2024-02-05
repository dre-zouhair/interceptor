package builder

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/dre-zouhair/interceptor/internal/protectioncli"
)

func Test_concreteBuilder_BuildCookiesSignals(t *testing.T) {
	type args struct {
		cookies []*http.Cookie
	}
	tests := []struct {
		name string
		args args
		want protectioncli.Signals
	}{
		{
			name: "case 1",
			args: args{
				cookies: []*http.Cookie{
					{
						Name:  "cookie",
						Value: "cookie",
					},
					nil,
				},
			},
			want: protectioncli.Signals{
				CookiesNamesLent:  []int{6},
				CookiesValuesLent: []int{6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSignalsBuilder()
			if got := c.BuildCookiesSignals(tt.args.cookies).GetSignals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildCookiesSignals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concreteBuilder_BuildCustomHeadersSignals(t *testing.T) {

	type args struct {
		headers http.Header
		keys    []string
	}

	headers := http.Header{}
	headers.Add("header-key", "value")

	tests := []struct {
		name string
		args args
		want protectioncli.Signals
	}{
		{
			"case 1",
			args{
				headers: headers,
				keys:    []string{"header-key"},
			},
			protectioncli.Signals{
				CustomHeaders: map[string]string{"header-key": "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSignalsBuilder()
			if got := c.BuildCustomHeadersSignals(tt.args.headers, tt.args.keys).GetSignals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildCustomHeadersSignals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concreteBuilder_BuildHeadersSignals(t *testing.T) {

	type args struct {
		headers http.Header
	}

	headers := http.Header{}
	headers.Add("Accept-Language", "Language")
	headers.Add("User-Agent", "Agent")
	headers.Add("Referer", "Referer")

	tests := []struct {
		name string
		args args
		want protectioncli.Signals
	}{
		{
			name: "case 1",
			args: args{
				headers: headers,
			},
			want: protectioncli.Signals{
				AcceptLanguage: "Language",
				UserAgent:      "Agent",
				Referer:        "Referer",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSignalsBuilder()
			if got := c.BuildHeadersSignals(tt.args.headers).GetSignals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildHeadersSignals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concreteBuilder_BuildRealRemoteAddr(t *testing.T) {

	type args struct {
		headers    http.Header
		remoteAddr string
	}

	headers := http.Header{}
	headers.Add("X-Forwarded-For", "123.4.4.4")
	headers.Add("X-Forwarded-For", "123.4.4.5")

	tests := []struct {
		name string
		args args
		want protectioncli.Signals
	}{
		{
			"case 1",
			args{
				headers:    headers,
				remoteAddr: "",
			},
			protectioncli.Signals{
				RealAddress:  "123.4.4.4",
				ProxyAddress: []string{"123.4.4.5"},
			},
		},
		{
			"case 2",
			args{
				headers:    http.Header{},
				remoteAddr: "123.1.1.4",
			},
			protectioncli.Signals{
				RealAddress: "123.1.1.4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSignalsBuilder()
			if got := c.BuildRealRemoteAddr(tt.args.headers, tt.args.remoteAddr).GetSignals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRealRemoteAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
