package simpro

import (
	"errors"
	"fmt"
)

// NewSimPROSDK errors
var (
	ErrorEmptyAccessToken = errors.New(
		"NewSimPROSDK error: `apiAccessToken` cannot be empty")
	ErrorEmptyDomain = errors.New(
		"NewSimPROSDK error: `simPRODomain` cannot be empty")
)

// GET errors
var (
	ErrorFailedReadingBody = func(err string) error {
		return fmt.Errorf(
			"GetCompanies error: Failed to read the response body: %s", err)
	}
	ErrorFailedJSONUnmarshal = func(err string) error {
		return fmt.Errorf(
			"GetCompanies error: Failed to unmarshal the response body: %s", err)
	}
)

// makeHTTPRequest errors
var (
	ErrorFailedCreatingRequest = func(err string) error {
		return fmt.Errorf(
			"makeHTTPRequest error: Failed while creating new request: %s", err)
	}
	ErrorFailedMakingRequest = func(err string) error {
		return fmt.Errorf(
			"makeHTTPRequest error: Failed to make request: %s", err)
	}
	ErrorUnexpectedResponse = func(statusCode int) error {
		return fmt.Errorf(
			"makeHTTPRequest error: Status code %d", statusCode)
	}
)
