package user

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
		Name:  "flight-master/user/UserComponent",
		Iface: reflect.TypeOf((*UserComponent)(nil)).Elem(),
		New:   func() any { return &userComponent{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return userComponent_local_stub{impl: impl.(UserComponent), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return userComponent_client_stub{stub: stub, listMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "flight-master/user/UserComponent", Method: "List"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return userComponent_server_stub{impl: impl.(UserComponent), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type userComponent_local_stub struct {
	impl   UserComponent
	tracer trace.Tracer
}

func (s userComponent_local_stub) List(ctx context.Context) (r0 []User, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "user.UserComponent.List", trace.WithSpanKind(trace.SpanKindInternal))
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

type userComponent_client_stub struct {
	stub        codegen.Stub
	listMetrics *codegen.MethodMetrics
}

func (s userComponent_client_stub) List(ctx context.Context) (r0 []User, err error) {
	// Update metrics.
	start := time.Now()
	s.listMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "user.UserComponent.List", trace.WithSpanKind(trace.SpanKindClient))
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
	r0 = serviceweaver_dec_slice_User_493ffc06(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type userComponent_server_stub struct {
	impl    UserComponent
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s userComponent_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "List":
		return s.list
	default:
		return nil
	}
}

func (s userComponent_server_stub) list(ctx context.Context, args []byte) (res []byte, err error) {
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
	serviceweaver_enc_slice_User_493ffc06(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = &User{}

func (x *User) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.ID)
	enc.String(x.FirstName)
	enc.String(x.LastName)
	enc.String(x.Email)
	enc.String(x.Phone)
}

func (x *User) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverUnmarshal: nil receiver"))
	}
	x.ID = dec.Int()
	x.FirstName = dec.String()
	x.LastName = dec.String()
	x.Email = dec.String()
	x.Phone = dec.String()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_User_493ffc06(enc *codegen.Encoder, arg []User) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_User_493ffc06(dec *codegen.Decoder) []User {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]User, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}
