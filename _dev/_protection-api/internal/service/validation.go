package service

import (
	"github.com/rs/zerolog/log"
	"time"
)

const (
	ALLOW_ACCESS  = "ALLOW"
	VERIFY_ACCESS = "VERIFY"
	BLOCK_ACCESS  = "BLOCK"
	BOT           = "bot"
	HUMAN         = "human"
)

type Signals struct {
	UserAgent     string            `json:"userAgent" validate:""`
	RealAddress   string            `json:"rAddress" validate:""`
	ProxyAddress  []string          `json:"proxy" validate:""`
	Referer       string            `json:"referer"`
	CustomHeaders map[string]string `json:"headers"`
	CustomCookies map[string]string `json:"cookies"`
	Method        string            `json:"method" validate:"required"`
	Path          string            `json:"path"`
	ContentLength int64             `json:"contentLength" validate:""`
	Query         string            `json:"query"`
	Time          time.Time         `json:"time" validate:"required"`
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
	action, judgment := ALLOW_ACCESS, HUMAN

	log.Error().Interface("signals", signals).Msg("request query")

	if signals.ContentLength == 21 {
		return &ValidationResponse{
			Judgment: BOT,
			Action:   BLOCK_ACCESS,
		}, nil
	}

	if signals.UserAgent == "" {
		return &ValidationResponse{
			Judgment: BOT,
			Action:   BLOCK_ACCESS,
		}, nil
	}

	if signals.Referer == "" {
		score = score - 1
	}

	if signals.ProxyAddress != nil && len(signals.ProxyAddress) > 0 {
		score = score - 1
	}

	if signals.Query != "" {
		return &ValidationResponse{
			Judgment: BOT,
			Action:   BLOCK_ACCESS,
		}, nil
	}

	if score < -2 {
		action, judgment = BLOCK_ACCESS, BOT
	}

	if score < 0 {
		action, judgment = VERIFY_ACCESS, BOT
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
