// Code generated by goa v3.14.4, DO NOT EDIT.
//
// tester gRPC server types
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/tester/design -o
// services/tester

package server

import (
	testerpb "goa.design/clue/example/weather/services/tester/gen/grpc/tester/pb"
	tester "goa.design/clue/example/weather/services/tester/gen/tester"
)

// NewTestAllPayload builds the payload of the "test_all" endpoint of the
// "tester" service from the gRPC request type.
func NewTestAllPayload(message *testerpb.TestAllRequest) *tester.TesterPayload {
	v := &tester.TesterPayload{}
	if message.Include != nil {
		v.Include = make([]string, len(message.Include))
		for i, val := range message.Include {
			v.Include[i] = val
		}
	}
	if message.Exclude != nil {
		v.Exclude = make([]string, len(message.Exclude))
		for i, val := range message.Exclude {
			v.Exclude[i] = val
		}
	}
	return v
}

// NewProtoTestAllResponse builds the gRPC response type from the result of the
// "test_all" endpoint of the "tester" service.
func NewProtoTestAllResponse(result *tester.TestResults) *testerpb.TestAllResponse {
	message := &testerpb.TestAllResponse{
		Duration:  result.Duration,
		PassCount: int32(result.PassCount),
		FailCount: int32(result.FailCount),
	}
	if result.Collections != nil {
		message.Collections = make([]*testerpb.TestCollection, len(result.Collections))
		for i, val := range result.Collections {
			message.Collections[i] = &testerpb.TestCollection{
				Name:      val.Name,
				Duration:  val.Duration,
				PassCount: int32(val.PassCount),
				FailCount: int32(val.FailCount),
			}
			if val.Results != nil {
				message.Collections[i].Results = make([]*testerpb.TestResult, len(val.Results))
				for j, val := range val.Results {
					message.Collections[i].Results[j] = &testerpb.TestResult{
						Name:     val.Name,
						Passed:   val.Passed,
						Error:    val.Error,
						Duration: val.Duration,
					}
				}
			}
		}
	}
	return message
}

// NewProtoTestSmokeResponse builds the gRPC response type from the result of
// the "test_smoke" endpoint of the "tester" service.
func NewProtoTestSmokeResponse(result *tester.TestResults) *testerpb.TestSmokeResponse {
	message := &testerpb.TestSmokeResponse{
		Duration:  result.Duration,
		PassCount: int32(result.PassCount),
		FailCount: int32(result.FailCount),
	}
	if result.Collections != nil {
		message.Collections = make([]*testerpb.TestCollection, len(result.Collections))
		for i, val := range result.Collections {
			message.Collections[i] = &testerpb.TestCollection{
				Name:      val.Name,
				Duration:  val.Duration,
				PassCount: int32(val.PassCount),
				FailCount: int32(val.FailCount),
			}
			if val.Results != nil {
				message.Collections[i].Results = make([]*testerpb.TestResult, len(val.Results))
				for j, val := range val.Results {
					message.Collections[i].Results[j] = &testerpb.TestResult{
						Name:     val.Name,
						Passed:   val.Passed,
						Error:    val.Error,
						Duration: val.Duration,
					}
				}
			}
		}
	}
	return message
}

// NewProtoTestForecasterResponse builds the gRPC response type from the result
// of the "test_forecaster" endpoint of the "tester" service.
func NewProtoTestForecasterResponse(result *tester.TestResults) *testerpb.TestForecasterResponse {
	message := &testerpb.TestForecasterResponse{
		Duration:  result.Duration,
		PassCount: int32(result.PassCount),
		FailCount: int32(result.FailCount),
	}
	if result.Collections != nil {
		message.Collections = make([]*testerpb.TestCollection, len(result.Collections))
		for i, val := range result.Collections {
			message.Collections[i] = &testerpb.TestCollection{
				Name:      val.Name,
				Duration:  val.Duration,
				PassCount: int32(val.PassCount),
				FailCount: int32(val.FailCount),
			}
			if val.Results != nil {
				message.Collections[i].Results = make([]*testerpb.TestResult, len(val.Results))
				for j, val := range val.Results {
					message.Collections[i].Results[j] = &testerpb.TestResult{
						Name:     val.Name,
						Passed:   val.Passed,
						Error:    val.Error,
						Duration: val.Duration,
					}
				}
			}
		}
	}
	return message
}

// NewProtoTestLocatorResponse builds the gRPC response type from the result of
// the "test_locator" endpoint of the "tester" service.
func NewProtoTestLocatorResponse(result *tester.TestResults) *testerpb.TestLocatorResponse {
	message := &testerpb.TestLocatorResponse{
		Duration:  result.Duration,
		PassCount: int32(result.PassCount),
		FailCount: int32(result.FailCount),
	}
	if result.Collections != nil {
		message.Collections = make([]*testerpb.TestCollection, len(result.Collections))
		for i, val := range result.Collections {
			message.Collections[i] = &testerpb.TestCollection{
				Name:      val.Name,
				Duration:  val.Duration,
				PassCount: int32(val.PassCount),
				FailCount: int32(val.FailCount),
			}
			if val.Results != nil {
				message.Collections[i].Results = make([]*testerpb.TestResult, len(val.Results))
				for j, val := range val.Results {
					message.Collections[i].Results[j] = &testerpb.TestResult{
						Name:     val.Name,
						Passed:   val.Passed,
						Error:    val.Error,
						Duration: val.Duration,
					}
				}
			}
		}
	}
	return message
}
