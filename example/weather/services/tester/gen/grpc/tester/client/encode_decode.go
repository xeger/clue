// Code generated by goa v3.14.4, DO NOT EDIT.
//
// tester gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/tester/design -o
// services/tester

package client

import (
	"context"

	testerpb "goa.design/clue/example/weather/services/tester/gen/grpc/tester/pb"
	tester "goa.design/clue/example/weather/services/tester/gen/tester"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildTestAllFunc builds the remote method to invoke for "tester" service
// "test_all" endpoint.
func BuildTestAllFunc(grpccli testerpb.TesterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.TestAll(ctx, reqpb.(*testerpb.TestAllRequest), opts...)
		}
		return grpccli.TestAll(ctx, &testerpb.TestAllRequest{}, opts...)
	}
}

// EncodeTestAllRequest encodes requests sent to tester test_all endpoint.
func EncodeTestAllRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*tester.TesterPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tester", "test_all", "*tester.TesterPayload", v)
	}
	return NewProtoTestAllRequest(payload), nil
}

// DecodeTestAllResponse decodes responses from the tester test_all endpoint.
func DecodeTestAllResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*testerpb.TestAllResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tester", "test_all", "*testerpb.TestAllResponse", v)
	}
	if err := ValidateTestAllResponse(message); err != nil {
		return nil, err
	}
	res := NewTestAllResult(message)
	return res, nil
}

// BuildTestSmokeFunc builds the remote method to invoke for "tester" service
// "test_smoke" endpoint.
func BuildTestSmokeFunc(grpccli testerpb.TesterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.TestSmoke(ctx, reqpb.(*testerpb.TestSmokeRequest), opts...)
		}
		return grpccli.TestSmoke(ctx, &testerpb.TestSmokeRequest{}, opts...)
	}
}

// DecodeTestSmokeResponse decodes responses from the tester test_smoke
// endpoint.
func DecodeTestSmokeResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*testerpb.TestSmokeResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tester", "test_smoke", "*testerpb.TestSmokeResponse", v)
	}
	if err := ValidateTestSmokeResponse(message); err != nil {
		return nil, err
	}
	res := NewTestSmokeResult(message)
	return res, nil
}

// BuildTestForecasterFunc builds the remote method to invoke for "tester"
// service "test_forecaster" endpoint.
func BuildTestForecasterFunc(grpccli testerpb.TesterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.TestForecaster(ctx, reqpb.(*testerpb.TestForecasterRequest), opts...)
		}
		return grpccli.TestForecaster(ctx, &testerpb.TestForecasterRequest{}, opts...)
	}
}

// DecodeTestForecasterResponse decodes responses from the tester
// test_forecaster endpoint.
func DecodeTestForecasterResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*testerpb.TestForecasterResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tester", "test_forecaster", "*testerpb.TestForecasterResponse", v)
	}
	if err := ValidateTestForecasterResponse(message); err != nil {
		return nil, err
	}
	res := NewTestForecasterResult(message)
	return res, nil
}

// BuildTestLocatorFunc builds the remote method to invoke for "tester" service
// "test_locator" endpoint.
func BuildTestLocatorFunc(grpccli testerpb.TesterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.TestLocator(ctx, reqpb.(*testerpb.TestLocatorRequest), opts...)
		}
		return grpccli.TestLocator(ctx, &testerpb.TestLocatorRequest{}, opts...)
	}
}

// DecodeTestLocatorResponse decodes responses from the tester test_locator
// endpoint.
func DecodeTestLocatorResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*testerpb.TestLocatorResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tester", "test_locator", "*testerpb.TestLocatorResponse", v)
	}
	if err := ValidateTestLocatorResponse(message); err != nil {
		return nil, err
	}
	res := NewTestLocatorResult(message)
	return res, nil
}
