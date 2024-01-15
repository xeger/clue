// Code generated by goa v3.14.4, DO NOT EDIT.
//
// tester endpoints
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/tester/design -o
// services/tester

package tester

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "tester" service endpoints.
type Endpoints struct {
	TestAll        goa.Endpoint
	TestSmoke      goa.Endpoint
	TestForecaster goa.Endpoint
	TestLocator    goa.Endpoint
}

// NewEndpoints wraps the methods of the "tester" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		TestAll:        NewTestAllEndpoint(s),
		TestSmoke:      NewTestSmokeEndpoint(s),
		TestForecaster: NewTestForecasterEndpoint(s),
		TestLocator:    NewTestLocatorEndpoint(s),
	}
}

// Use applies the given middleware to all the "tester" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.TestAll = m(e.TestAll)
	e.TestSmoke = m(e.TestSmoke)
	e.TestForecaster = m(e.TestForecaster)
	e.TestLocator = m(e.TestLocator)
}

// NewTestAllEndpoint returns an endpoint function that calls the method
// "test_all" of service "tester".
func NewTestAllEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*TesterPayload)
		return s.TestAll(ctx, p)
	}
}

// NewTestSmokeEndpoint returns an endpoint function that calls the method
// "test_smoke" of service "tester".
func NewTestSmokeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.TestSmoke(ctx)
	}
}

// NewTestForecasterEndpoint returns an endpoint function that calls the method
// "test_forecaster" of service "tester".
func NewTestForecasterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.TestForecaster(ctx)
	}
}

// NewTestLocatorEndpoint returns an endpoint function that calls the method
// "test_locator" of service "tester".
func NewTestLocatorEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.TestLocator(ctx)
	}
}