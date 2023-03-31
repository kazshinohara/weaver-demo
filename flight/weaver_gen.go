package flight

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "flight-master/flight/FlightComponent",
		Iface: reflect.TypeOf((*FlightComponent)(nil)).Elem(),
		New:   func() any { return &flightComponent{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return flightComponent_local_stub{impl: impl.(FlightComponent), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return flightComponent_client_stub{stub: stub, listMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "flight-master/flight/FlightComponent", Method: "List"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return flightComponent_server_stub{impl: impl.(FlightComponent), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type flightComponent_local_stub struct {
	impl   FlightComponent
	tracer trace.Tracer
}

func (s flightComponent_local_stub) List(ctx context.Context) (r0 []Flight, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "flight.FlightComponent.List", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.List(ctx)
}

// Client stub implementations.

type flightComponent_client_stub struct {
	stub        codegen.Stub
	listMetrics *codegen.MethodMetrics
}

func (s flightComponent_client_stub) List(ctx context.Context) (r0 []Flight, err error) {
	// Update metrics.
	start := time.Now()
	s.listMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "flight.FlightComponent.List", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.listMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.listMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.listMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	if err != nil {
		return
	}
	s.listMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_Flight_ab6b58dc(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type flightComponent_server_stub struct {
	impl    FlightComponent
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s flightComponent_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "List":
		return s.list
	default:
		return nil
	}
}

func (s flightComponent_server_stub) list(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.List(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_Flight_ab6b58dc(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = &Flight{}

func (x *Flight) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Flight.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.ID)
	enc.String(x.Origin)
	enc.String(x.Destination)
	enc.Int(x.AirlineID)
	enc.String(x.Departure)
	enc.String(x.Arrival)
	enc.Int(x.AvailableSeats)
}

func (x *Flight) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Flight.WeaverUnmarshal: nil receiver"))
	}
	x.ID = dec.Int()
	x.Origin = dec.String()
	x.Destination = dec.String()
	x.AirlineID = dec.Int()
	x.Departure = dec.String()
	x.Arrival = dec.String()
	x.AvailableSeats = dec.Int()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_Flight_ab6b58dc(enc *codegen.Encoder, arg []Flight) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Flight_ab6b58dc(dec *codegen.Decoder) []Flight {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Flight, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}
