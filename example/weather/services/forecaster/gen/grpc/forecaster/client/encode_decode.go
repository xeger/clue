// Code generated by goa v3.14.4, DO NOT EDIT.
//
// Forecaster gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/forecaster/design -o
// services/forecaster

package client

import (
	"context"

	forecaster "goa.design/clue/example/weather/services/forecaster/gen/forecaster"
	forecasterpb "goa.design/clue/example/weather/services/forecaster/gen/grpc/forecaster/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildForecastFunc builds the remote method to invoke for "Forecaster"
// service "forecast" endpoint.
func BuildForecastFunc(grpccli forecasterpb.ForecasterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Forecast(ctx, reqpb.(*forecasterpb.ForecastRequest), opts...)
		}
		return grpccli.Forecast(ctx, &forecasterpb.ForecastRequest{}, opts...)
	}
}

// EncodeForecastRequest encodes requests sent to Forecaster forecast endpoint.
func EncodeForecastRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*forecaster.ForecastPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Forecaster", "forecast", "*forecaster.ForecastPayload", v)
	}
	return NewProtoForecastRequest(payload), nil
}

// DecodeForecastResponse decodes responses from the Forecaster forecast
// endpoint.
func DecodeForecastResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*forecasterpb.ForecastResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Forecaster", "forecast", "*forecasterpb.ForecastResponse", v)
	}
	if err := ValidateForecastResponse(message); err != nil {
		return nil, err
	}
	res := NewForecastResult(message)
	return res, nil
}
