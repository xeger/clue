// Code generated by goa v3.14.4, DO NOT EDIT.
//
// front HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/front/design -o
// services/front

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	front "goa.design/clue/example/weather/services/front/gen/front"
	goahttp "goa.design/goa/v3/http"
)

// BuildForecastRequest instantiates a HTTP request object with method and path
// set to call the "front" service "forecast" endpoint
func (c *Client) BuildForecastRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		ip string
	)
	{
		p, ok := v.(string)
		if !ok {
			return nil, goahttp.ErrInvalidType("front", "forecast", "string", v)
		}
		ip = p
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ForecastFrontPath(ip)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("front", "forecast", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeForecastResponse returns a decoder for responses returned by the front
// forecast endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeForecastResponse may return the following errors:
//   - "not_usa" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeForecastResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ForecastResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("front", "forecast", err)
			}
			err = ValidateForecastResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("front", "forecast", err)
			}
			res := NewForecast2OK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ForecastNotUsaResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("front", "forecast", err)
			}
			err = ValidateForecastNotUsaResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("front", "forecast", err)
			}
			return nil, NewForecastNotUsa(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("front", "forecast", resp.StatusCode, string(body))
		}
	}
}

// BuildTestAllRequest instantiates a HTTP request object with method and path
// set to call the "front" service "test_all" endpoint
func (c *Client) BuildTestAllRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: TestAllFrontPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("front", "test_all", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeTestAllRequest returns an encoder for requests sent to the front
// test_all server.
func EncodeTestAllRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*front.TestAllPayload)
		if !ok {
			return goahttp.ErrInvalidType("front", "test_all", "*front.TestAllPayload", v)
		}
		body := NewTestAllRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("front", "test_all", err)
		}
		return nil
	}
}

// DecodeTestAllResponse returns a decoder for responses returned by the front
// test_all endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeTestAllResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body TestAllResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("front", "test_all", err)
			}
			err = ValidateTestAllResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("front", "test_all", err)
			}
			res := NewTestAllTestResultsOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("front", "test_all", resp.StatusCode, string(body))
		}
	}
}

// BuildTestSmokeRequest instantiates a HTTP request object with method and
// path set to call the "front" service "test_smoke" endpoint
func (c *Client) BuildTestSmokeRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: TestSmokeFrontPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("front", "test_smoke", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeTestSmokeResponse returns a decoder for responses returned by the
// front test_smoke endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeTestSmokeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body TestSmokeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("front", "test_smoke", err)
			}
			err = ValidateTestSmokeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("front", "test_smoke", err)
			}
			res := NewTestSmokeTestResultsOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("front", "test_smoke", resp.StatusCode, string(body))
		}
	}
}

// unmarshalLocationResponseBodyToFrontLocation builds a value of type
// *front.Location from a value of type *LocationResponseBody.
func unmarshalLocationResponseBodyToFrontLocation(v *LocationResponseBody) *front.Location {
	res := &front.Location{
		Lat:   *v.Lat,
		Long:  *v.Long,
		City:  *v.City,
		State: *v.State,
	}

	return res
}

// unmarshalPeriodResponseBodyToFrontPeriod builds a value of type
// *front.Period from a value of type *PeriodResponseBody.
func unmarshalPeriodResponseBodyToFrontPeriod(v *PeriodResponseBody) *front.Period {
	res := &front.Period{
		Name:            *v.Name,
		StartTime:       *v.StartTime,
		EndTime:         *v.EndTime,
		Temperature:     *v.Temperature,
		TemperatureUnit: *v.TemperatureUnit,
		Summary:         *v.Summary,
	}

	return res
}

// unmarshalTestCollectionResponseBodyToFrontTestCollection builds a value of
// type *front.TestCollection from a value of type *TestCollectionResponseBody.
func unmarshalTestCollectionResponseBodyToFrontTestCollection(v *TestCollectionResponseBody) *front.TestCollection {
	res := &front.TestCollection{
		Name:      *v.Name,
		Duration:  *v.Duration,
		PassCount: *v.PassCount,
		FailCount: *v.FailCount,
	}
	if v.Results != nil {
		res.Results = make([]*front.TestResult, len(v.Results))
		for i, val := range v.Results {
			res.Results[i] = unmarshalTestResultResponseBodyToFrontTestResult(val)
		}
	}

	return res
}

// unmarshalTestResultResponseBodyToFrontTestResult builds a value of type
// *front.TestResult from a value of type *TestResultResponseBody.
func unmarshalTestResultResponseBodyToFrontTestResult(v *TestResultResponseBody) *front.TestResult {
	if v == nil {
		return nil
	}
	res := &front.TestResult{
		Name:     *v.Name,
		Passed:   *v.Passed,
		Error:    v.Error,
		Duration: *v.Duration,
	}

	return res
}
