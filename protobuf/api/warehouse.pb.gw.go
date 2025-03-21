// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: protobuf/api/warehouse.proto

/*
Package commerce_api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package commerce_api

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_WarehouseService_UpdateProductStock_0(ctx context.Context, marshaler runtime.Marshaler, client WarehouseServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq UpdateProductStockRequest
		metadata runtime.ServerMetadata
		err      error
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	val, ok := pathParams["warehouse_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "warehouse_id")
	}
	protoReq.WarehouseId, err = runtime.Int64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "warehouse_id", err)
	}
	val, ok = pathParams["product_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "product_id")
	}
	protoReq.ProductId, err = runtime.Int64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "product_id", err)
	}
	msg, err := client.UpdateProductStock(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_WarehouseService_UpdateProductStock_0(ctx context.Context, marshaler runtime.Marshaler, server WarehouseServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq UpdateProductStockRequest
		metadata runtime.ServerMetadata
		err      error
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	val, ok := pathParams["warehouse_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "warehouse_id")
	}
	protoReq.WarehouseId, err = runtime.Int64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "warehouse_id", err)
	}
	val, ok = pathParams["product_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "product_id")
	}
	protoReq.ProductId, err = runtime.Int64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "product_id", err)
	}
	msg, err := server.UpdateProductStock(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterWarehouseServiceHandlerServer registers the http handlers for service WarehouseService to "mux".
// UnaryRPC     :call WarehouseServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterWarehouseServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterWarehouseServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server WarehouseServiceServer) error {
	mux.Handle(http.MethodPut, pattern_WarehouseService_UpdateProductStock_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/mughieams.evermosassessment.v1.WarehouseService/UpdateProductStock", runtime.WithHTTPPathPattern("/v1/warehouses/{warehouse_id}/products/{product_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_WarehouseService_UpdateProductStock_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WarehouseService_UpdateProductStock_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterWarehouseServiceHandlerFromEndpoint is same as RegisterWarehouseServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWarehouseServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterWarehouseServiceHandler(ctx, mux, conn)
}

// RegisterWarehouseServiceHandler registers the http handlers for service WarehouseService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWarehouseServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWarehouseServiceHandlerClient(ctx, mux, NewWarehouseServiceClient(conn))
}

// RegisterWarehouseServiceHandlerClient registers the http handlers for service WarehouseService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WarehouseServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WarehouseServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WarehouseServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterWarehouseServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WarehouseServiceClient) error {
	mux.Handle(http.MethodPut, pattern_WarehouseService_UpdateProductStock_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/mughieams.evermosassessment.v1.WarehouseService/UpdateProductStock", runtime.WithHTTPPathPattern("/v1/warehouses/{warehouse_id}/products/{product_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WarehouseService_UpdateProductStock_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_WarehouseService_UpdateProductStock_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_WarehouseService_UpdateProductStock_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"v1", "warehouses", "warehouse_id", "products", "product_id"}, ""))
)

var (
	forward_WarehouseService_UpdateProductStock_0 = runtime.ForwardResponseMessage
)
