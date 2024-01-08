// Code generated by goa v3.14.4, DO NOT EDIT.
//
// tester client
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/tester/design -o
// services/tester

package tester

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "tester" service client.
type Client struct {
	TestAllEndpoint        goa.Endpoint
	TestSmokeEndpoint      goa.Endpoint
	TestForecasterEndpoint goa.Endpoint
	TestLocatorEndpoint    goa.Endpoint
}

// NewClient initializes a "tester" service client given the endpoints.
func NewClient(testAll, testSmoke, testForecaster, testLocator goa.Endpoint) *Client {
	return &Client{
		TestAllEndpoint:        testAll,
		TestSmokeEndpoint:      testSmoke,
		TestForecasterEndpoint: testForecaster,
		TestLocatorEndpoint:    testLocator,
	}
}

// TestAll calls the "test_all" endpoint of the "tester" service.
// TestAll may return the following errors:
//   - "include_exclude_both" (type *goa.ServiceError): Cannot specify both include and exclude
//   - error: internal error
func (c *Client) TestAll(ctx context.Context, p *TesterPayload) (res *TestResults, err error) {
	var ires any
	ires, err = c.TestAllEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TestResults), nil
}

// TestSmoke calls the "test_smoke" endpoint of the "tester" service.
// TestSmoke may return the following errors:
//   - "include_exclude_both" (type *goa.ServiceError): Cannot specify both include and exclude
//   - error: internal error
func (c *Client) TestSmoke(ctx context.Context) (res *TestResults, err error) {
	var ires any
	ires, err = c.TestSmokeEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*TestResults), nil
}

// TestForecaster calls the "test_forecaster" endpoint of the "tester" service.
// TestForecaster may return the following errors:
//   - "include_exclude_both" (type *goa.ServiceError): Cannot specify both include and exclude
//   - error: internal error
func (c *Client) TestForecaster(ctx context.Context) (res *TestResults, err error) {
	var ires any
	ires, err = c.TestForecasterEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*TestResults), nil
}

// TestLocator calls the "test_locator" endpoint of the "tester" service.
// TestLocator may return the following errors:
//   - "include_exclude_both" (type *goa.ServiceError): Cannot specify both include and exclude
//   - error: internal error
func (c *Client) TestLocator(ctx context.Context) (res *TestResults, err error) {
	var ires any
	ires, err = c.TestLocatorEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*TestResults), nil
}
