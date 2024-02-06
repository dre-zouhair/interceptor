package analyser

import (
	"bytes"
	"encoding/json"
	configuration "github.com/dre-zouhair/interceptor/config"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type protectionCli struct {
	conf   configuration.ProtectionAPIConfig
	client *http.Client
}

func NewProtectionCli(conf configuration.ProtectionAPIConfig) IProtectionCli {
	return &protectionCli{
		conf:   conf,
		client: &http.Client{},
	}
}

type IProtectionCli interface {
	Validate(signals Signals) (*ValidationResponse, error)
}

type Signals struct {
	UserAgent     string            `json:"userAgent" validate:"required"`
	RealAddress   string            `json:"rAddress" validate:"required"`
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

func (cli protectionCli) Validate(signals Signals) (*ValidationResponse, error) {

	body, err := json.Marshal(signals)
	if err != nil {
		log.Error().Err(err).Msg("error marshaling signals to JSON")
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, cli.conf.ProtectionEndpoint, bytes.NewBuffer(body))

	if err != nil {
		log.Error().Err(err).Msg("failed to create protection request")
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+cli.conf.ProtectionToken)

	resp, err := cli.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to do protection request")
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		bodyErr := Body.Close()
		if bodyErr != nil {
			log.Error().Err(bodyErr).Msg("unable to close validation response body")
		}
	}(resp.Body)

	var validationResponse ValidationResponse
	if err := json.NewDecoder(resp.Body).Decode(&validationResponse); err != nil {
		log.Error().Err(err).Msg("error marshaling validation response")
		return nil, err
	}
	return &validationResponse, nil
}

type ValidationResponse struct {
	Judgment string `json:"judgment"`
	Action   string `json:"action"`
}
