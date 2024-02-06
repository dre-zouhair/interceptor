package builder

import (
	"net/http"
	"slices"
	"strings"

	"github.com/dre-zouhair/interceptor/internal/analyser"
)

type ISignalsBuilder interface {
	BuildHeadersSignals(headers http.Header) ISignalsBuilder
	BuildCustomHeadersSignals(headers http.Header, keys []string) ISignalsBuilder
	BuildCustomCookiesSignals(cookies []*http.Cookie, keys []string) ISignalsBuilder
	BuildRealRemoteAddr(headers http.Header, remoteAddr string) ISignalsBuilder
	GetSignals() analyser.Signals
}

type concreteBuilder struct {
	signals analyser.Signals
}

func NewSignalsBuilder() ISignalsBuilder {
	return &concreteBuilder{}
}

func (c concreteBuilder) BuildHeadersSignals(headers http.Header) ISignalsBuilder {
	c.signals.UserAgent = headers.Get("User-Agent")
	c.signals.Referer = headers.Get("Referer")

	return c
}

func (c concreteBuilder) BuildCustomHeadersSignals(headers http.Header, keys []string) ISignalsBuilder {
	c.signals.CustomHeaders = make(map[string]string)
	for _, name := range keys {
		c.signals.CustomHeaders[name] = headers.Get(name)
	}
	return c
}

func (c concreteBuilder) BuildRealRemoteAddr(headers http.Header, remoteAddr string) ISignalsBuilder {
	xForwardedFor := headers.Values("X-Forwarded-For")
	if xForwardedFor != nil && len(xForwardedFor) != 0 {
		c.signals.ProxyAddress = xForwardedFor[1:]
		c.signals.RealAddress = strings.TrimSpace(xForwardedFor[0])
	} else {
		c.signals.RealAddress = remoteAddr
	}
	return c
}

func (c concreteBuilder) BuildCustomCookiesSignals(cookies []*http.Cookie, keys []string) ISignalsBuilder {
	c.signals.CustomCookies = make(map[string]string)
	for _, cookie := range cookies {
		if cookie == nil {
			continue
		}
		if slices.Contains(keys, cookie.Name) {
			c.signals.CustomCookies[cookie.Name] = cookie.Value
		}
	}
	return c
}

func (c concreteBuilder) GetSignals() analyser.Signals {
	return c.signals
}
