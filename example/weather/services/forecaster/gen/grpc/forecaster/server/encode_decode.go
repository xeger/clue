// Code generated by goa v3.14.4, DO NOT EDIT.
//
// Forecaster gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/forecaster/design -o
// services/forecaster

package server

import (
	"context"

	forecaster "goa.design/clue/example/weather/services/forecaster/gen/forecaster"
	forecasterpb "goa.design/clue/example/weather/services/forecaster/gen/grpc/forecaster/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeForecastResponse encodes responses from the "Forecaster" service
// "forecast" endpoint.
func EncodeForecastResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(*forecaster.Forecast2)
	if !ok {
		return nil, goagrpc.ErrInvalidType("Forecaster", "forecast", "*forecaster.Forecast2", v)
	}
	resp := NewProtoForecastResponse(result)
	return resp, nil
}

// DecodeForecastRequest decodes requests sent to "Forecaster" service
// "forecast" endpoint.
func DecodeForecastRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *forecasterpb.ForecastRequest
		ok      bool
	)
	{
		if message, ok = v.(*forecasterpb.ForecastRequest); !ok {
			return nil, goagrpc.ErrInvalidType("Forecaster", "forecast", "*forecasterpb.ForecastRequest", v)
		}
	}
	var payload *forecaster.ForecastPayload
	{
		payload = NewForecastPayload(message)
	}
	return payload, nil
}
