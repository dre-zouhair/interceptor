package service

import "time"

const (
	ALLOW_ACCESS  = "ALLOW"
	VERIFY_ACCESS = "VERIFY"
	BLOCK_ACCESS  = "BLOCK"
)

type Signals struct {
	ContentLength     int64             `json:"contentLength" validate:""`
	AcceptLanguage    string            `json:"acceptLanguage" validate:""`
	UserAgent         string            `json:"userAgent" validate:"required"`
	RealAddress       string            `json:"rAddress" validate:"required"`
	ProxyAddress      []string          `json:"proxy" validate:""`
	Referer           string            `json:"referer"`
	CustomHeaders     map[string]string `json:"headers"`
	CustomCookies     map[string]string `json:"cookies"`
	CookiesNamesLent  []int             `json:"cookiesNamesLen"`
	CookiesValuesLent []int             `json:"cookiesValuesLent"`
	Method            string            `json:"method" validate:"required"`
	Path              string            `json:"path"`
	Query             string            `json:"query"`
	Time              time.Time         `json:"time" validate:"required"`
}

type validationService struct {
}

func NewValidationService() IValidationService {
	return &validationService{}
}

type IValidationService interface {
	Validate(signals Signals) (*ValidationResponse, error)
}

func (s validationService) Validate(signals Signals) (*ValidationResponse, error) {
	score := 0
	action, judgment := ALLOW_ACCESS, "human"

	if signals.UserAgent == "" {
		score = score - 1
	}

	if signals.AcceptLanguage == "" {
		score = score - 1
	}

	if signals.ProxyAddress != nil && len(signals.ProxyAddress) > 0 {
		score = score - 1
	}

	if score < -1 {
		action, judgment = BLOCK_ACCESS, "bot"
	}

	if score < 0 {
		action, judgment = VERIFY_ACCESS, "bot"
	}

	return &ValidationResponse{
		Judgment: judgment,
		Action:   action,
	}, nil
}

type ValidationResponse struct {
	Judgment string `json:"judgment"`
	Action   string `json:"action"`
}
