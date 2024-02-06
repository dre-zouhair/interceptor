package configuration

type ProtectionAPIConfig struct {
	ProtectionEndpoint string
	ProtectionToken    string
}

type ProcessorConfig struct {
	CustomHeaderSignals []string
	CustomHeaderCookies []string
}

type ProtectionMiddlewareConfig struct {
	ServerPort          string
	ProtectionFailMode  string
	ProtectionAPIConfig ProtectionAPIConfig
	ProcessorConfig     ProcessorConfig
	ForwardEndPoint     string
}
