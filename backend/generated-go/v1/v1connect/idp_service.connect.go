// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/idp_service.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/bytebase/bytebase/backend/generated-go/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// IdentityProviderServiceName is the fully-qualified name of the IdentityProviderService service.
	IdentityProviderServiceName = "bytebase.v1.IdentityProviderService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IdentityProviderServiceGetIdentityProviderProcedure is the fully-qualified name of the
	// IdentityProviderService's GetIdentityProvider RPC.
	IdentityProviderServiceGetIdentityProviderProcedure = "/bytebase.v1.IdentityProviderService/GetIdentityProvider"
	// IdentityProviderServiceListIdentityProvidersProcedure is the fully-qualified name of the
	// IdentityProviderService's ListIdentityProviders RPC.
	IdentityProviderServiceListIdentityProvidersProcedure = "/bytebase.v1.IdentityProviderService/ListIdentityProviders"
	// IdentityProviderServiceCreateIdentityProviderProcedure is the fully-qualified name of the
	// IdentityProviderService's CreateIdentityProvider RPC.
	IdentityProviderServiceCreateIdentityProviderProcedure = "/bytebase.v1.IdentityProviderService/CreateIdentityProvider"
	// IdentityProviderServiceUpdateIdentityProviderProcedure is the fully-qualified name of the
	// IdentityProviderService's UpdateIdentityProvider RPC.
	IdentityProviderServiceUpdateIdentityProviderProcedure = "/bytebase.v1.IdentityProviderService/UpdateIdentityProvider"
	// IdentityProviderServiceDeleteIdentityProviderProcedure is the fully-qualified name of the
	// IdentityProviderService's DeleteIdentityProvider RPC.
	IdentityProviderServiceDeleteIdentityProviderProcedure = "/bytebase.v1.IdentityProviderService/DeleteIdentityProvider"
	// IdentityProviderServiceTestIdentityProviderProcedure is the fully-qualified name of the
	// IdentityProviderService's TestIdentityProvider RPC.
	IdentityProviderServiceTestIdentityProviderProcedure = "/bytebase.v1.IdentityProviderService/TestIdentityProvider"
)

// IdentityProviderServiceClient is a client for the bytebase.v1.IdentityProviderService service.
type IdentityProviderServiceClient interface {
	// Permissions required: bb.identityProviders.get
	GetIdentityProvider(context.Context, *connect.Request[v1.GetIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error)
	// Permissions required: None
	ListIdentityProviders(context.Context, *connect.Request[v1.ListIdentityProvidersRequest]) (*connect.Response[v1.ListIdentityProvidersResponse], error)
	// Permissions required: bb.identityProviders.create
	CreateIdentityProvider(context.Context, *connect.Request[v1.CreateIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error)
	// Permissions required: bb.identityProviders.update
	UpdateIdentityProvider(context.Context, *connect.Request[v1.UpdateIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error)
	// Permissions required: bb.identityProviders.delete
	DeleteIdentityProvider(context.Context, *connect.Request[v1.DeleteIdentityProviderRequest]) (*connect.Response[emptypb.Empty], error)
	// Permissions required: bb.identityProviders.update
	TestIdentityProvider(context.Context, *connect.Request[v1.TestIdentityProviderRequest]) (*connect.Response[v1.TestIdentityProviderResponse], error)
}

// NewIdentityProviderServiceClient constructs a client for the bytebase.v1.IdentityProviderService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIdentityProviderServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IdentityProviderServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	identityProviderServiceMethods := v1.File_v1_idp_service_proto.Services().ByName("IdentityProviderService").Methods()
	return &identityProviderServiceClient{
		getIdentityProvider: connect.NewClient[v1.GetIdentityProviderRequest, v1.IdentityProvider](
			httpClient,
			baseURL+IdentityProviderServiceGetIdentityProviderProcedure,
			connect.WithSchema(identityProviderServiceMethods.ByName("GetIdentityProvider")),
			connect.WithClientOptions(opts...),
		),
		listIdentityProviders: connect.NewClient[v1.ListIdentityProvidersRequest, v1.ListIdentityProvidersResponse](
			httpClient,
			baseURL+IdentityProviderServiceListIdentityProvidersProcedure,
			connect.WithSchema(identityProviderServiceMethods.ByName("ListIdentityProviders")),
			connect.WithClientOptions(opts...),
		),
		createIdentityProvider: connect.NewClient[v1.CreateIdentityProviderRequest, v1.IdentityProvider](
			httpClient,
			baseURL+IdentityProviderServiceCreateIdentityProviderProcedure,
			connect.WithSchema(identityProviderServiceMethods.ByName("CreateIdentityProvider")),
			connect.WithClientOptions(opts...),
		),
		updateIdentityProvider: connect.NewClient[v1.UpdateIdentityProviderRequest, v1.IdentityProvider](
			httpClient,
			baseURL+IdentityProviderServiceUpdateIdentityProviderProcedure,
			connect.WithSchema(identityProviderServiceMethods.ByName("UpdateIdentityProvider")),
			connect.WithClientOptions(opts...),
		),
		deleteIdentityProvider: connect.NewClient[v1.DeleteIdentityProviderRequest, emptypb.Empty](
			httpClient,
			baseURL+IdentityProviderServiceDeleteIdentityProviderProcedure,
			connect.WithSchema(identityProviderServiceMethods.ByName("DeleteIdentityProvider")),
			connect.WithClientOptions(opts...),
		),
		testIdentityProvider: connect.NewClient[v1.TestIdentityProviderRequest, v1.TestIdentityProviderResponse](
			httpClient,
			baseURL+IdentityProviderServiceTestIdentityProviderProcedure,
			connect.WithSchema(identityProviderServiceMethods.ByName("TestIdentityProvider")),
			connect.WithClientOptions(opts...),
		),
	}
}

// identityProviderServiceClient implements IdentityProviderServiceClient.
type identityProviderServiceClient struct {
	getIdentityProvider    *connect.Client[v1.GetIdentityProviderRequest, v1.IdentityProvider]
	listIdentityProviders  *connect.Client[v1.ListIdentityProvidersRequest, v1.ListIdentityProvidersResponse]
	createIdentityProvider *connect.Client[v1.CreateIdentityProviderRequest, v1.IdentityProvider]
	updateIdentityProvider *connect.Client[v1.UpdateIdentityProviderRequest, v1.IdentityProvider]
	deleteIdentityProvider *connect.Client[v1.DeleteIdentityProviderRequest, emptypb.Empty]
	testIdentityProvider   *connect.Client[v1.TestIdentityProviderRequest, v1.TestIdentityProviderResponse]
}

// GetIdentityProvider calls bytebase.v1.IdentityProviderService.GetIdentityProvider.
func (c *identityProviderServiceClient) GetIdentityProvider(ctx context.Context, req *connect.Request[v1.GetIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error) {
	return c.getIdentityProvider.CallUnary(ctx, req)
}

// ListIdentityProviders calls bytebase.v1.IdentityProviderService.ListIdentityProviders.
func (c *identityProviderServiceClient) ListIdentityProviders(ctx context.Context, req *connect.Request[v1.ListIdentityProvidersRequest]) (*connect.Response[v1.ListIdentityProvidersResponse], error) {
	return c.listIdentityProviders.CallUnary(ctx, req)
}

// CreateIdentityProvider calls bytebase.v1.IdentityProviderService.CreateIdentityProvider.
func (c *identityProviderServiceClient) CreateIdentityProvider(ctx context.Context, req *connect.Request[v1.CreateIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error) {
	return c.createIdentityProvider.CallUnary(ctx, req)
}

// UpdateIdentityProvider calls bytebase.v1.IdentityProviderService.UpdateIdentityProvider.
func (c *identityProviderServiceClient) UpdateIdentityProvider(ctx context.Context, req *connect.Request[v1.UpdateIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error) {
	return c.updateIdentityProvider.CallUnary(ctx, req)
}

// DeleteIdentityProvider calls bytebase.v1.IdentityProviderService.DeleteIdentityProvider.
func (c *identityProviderServiceClient) DeleteIdentityProvider(ctx context.Context, req *connect.Request[v1.DeleteIdentityProviderRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteIdentityProvider.CallUnary(ctx, req)
}

// TestIdentityProvider calls bytebase.v1.IdentityProviderService.TestIdentityProvider.
func (c *identityProviderServiceClient) TestIdentityProvider(ctx context.Context, req *connect.Request[v1.TestIdentityProviderRequest]) (*connect.Response[v1.TestIdentityProviderResponse], error) {
	return c.testIdentityProvider.CallUnary(ctx, req)
}

// IdentityProviderServiceHandler is an implementation of the bytebase.v1.IdentityProviderService
// service.
type IdentityProviderServiceHandler interface {
	// Permissions required: bb.identityProviders.get
	GetIdentityProvider(context.Context, *connect.Request[v1.GetIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error)
	// Permissions required: None
	ListIdentityProviders(context.Context, *connect.Request[v1.ListIdentityProvidersRequest]) (*connect.Response[v1.ListIdentityProvidersResponse], error)
	// Permissions required: bb.identityProviders.create
	CreateIdentityProvider(context.Context, *connect.Request[v1.CreateIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error)
	// Permissions required: bb.identityProviders.update
	UpdateIdentityProvider(context.Context, *connect.Request[v1.UpdateIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error)
	// Permissions required: bb.identityProviders.delete
	DeleteIdentityProvider(context.Context, *connect.Request[v1.DeleteIdentityProviderRequest]) (*connect.Response[emptypb.Empty], error)
	// Permissions required: bb.identityProviders.update
	TestIdentityProvider(context.Context, *connect.Request[v1.TestIdentityProviderRequest]) (*connect.Response[v1.TestIdentityProviderResponse], error)
}

// NewIdentityProviderServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIdentityProviderServiceHandler(svc IdentityProviderServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	identityProviderServiceMethods := v1.File_v1_idp_service_proto.Services().ByName("IdentityProviderService").Methods()
	identityProviderServiceGetIdentityProviderHandler := connect.NewUnaryHandler(
		IdentityProviderServiceGetIdentityProviderProcedure,
		svc.GetIdentityProvider,
		connect.WithSchema(identityProviderServiceMethods.ByName("GetIdentityProvider")),
		connect.WithHandlerOptions(opts...),
	)
	identityProviderServiceListIdentityProvidersHandler := connect.NewUnaryHandler(
		IdentityProviderServiceListIdentityProvidersProcedure,
		svc.ListIdentityProviders,
		connect.WithSchema(identityProviderServiceMethods.ByName("ListIdentityProviders")),
		connect.WithHandlerOptions(opts...),
	)
	identityProviderServiceCreateIdentityProviderHandler := connect.NewUnaryHandler(
		IdentityProviderServiceCreateIdentityProviderProcedure,
		svc.CreateIdentityProvider,
		connect.WithSchema(identityProviderServiceMethods.ByName("CreateIdentityProvider")),
		connect.WithHandlerOptions(opts...),
	)
	identityProviderServiceUpdateIdentityProviderHandler := connect.NewUnaryHandler(
		IdentityProviderServiceUpdateIdentityProviderProcedure,
		svc.UpdateIdentityProvider,
		connect.WithSchema(identityProviderServiceMethods.ByName("UpdateIdentityProvider")),
		connect.WithHandlerOptions(opts...),
	)
	identityProviderServiceDeleteIdentityProviderHandler := connect.NewUnaryHandler(
		IdentityProviderServiceDeleteIdentityProviderProcedure,
		svc.DeleteIdentityProvider,
		connect.WithSchema(identityProviderServiceMethods.ByName("DeleteIdentityProvider")),
		connect.WithHandlerOptions(opts...),
	)
	identityProviderServiceTestIdentityProviderHandler := connect.NewUnaryHandler(
		IdentityProviderServiceTestIdentityProviderProcedure,
		svc.TestIdentityProvider,
		connect.WithSchema(identityProviderServiceMethods.ByName("TestIdentityProvider")),
		connect.WithHandlerOptions(opts...),
	)
	return "/bytebase.v1.IdentityProviderService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IdentityProviderServiceGetIdentityProviderProcedure:
			identityProviderServiceGetIdentityProviderHandler.ServeHTTP(w, r)
		case IdentityProviderServiceListIdentityProvidersProcedure:
			identityProviderServiceListIdentityProvidersHandler.ServeHTTP(w, r)
		case IdentityProviderServiceCreateIdentityProviderProcedure:
			identityProviderServiceCreateIdentityProviderHandler.ServeHTTP(w, r)
		case IdentityProviderServiceUpdateIdentityProviderProcedure:
			identityProviderServiceUpdateIdentityProviderHandler.ServeHTTP(w, r)
		case IdentityProviderServiceDeleteIdentityProviderProcedure:
			identityProviderServiceDeleteIdentityProviderHandler.ServeHTTP(w, r)
		case IdentityProviderServiceTestIdentityProviderProcedure:
			identityProviderServiceTestIdentityProviderHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIdentityProviderServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIdentityProviderServiceHandler struct{}

func (UnimplementedIdentityProviderServiceHandler) GetIdentityProvider(context.Context, *connect.Request[v1.GetIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.IdentityProviderService.GetIdentityProvider is not implemented"))
}

func (UnimplementedIdentityProviderServiceHandler) ListIdentityProviders(context.Context, *connect.Request[v1.ListIdentityProvidersRequest]) (*connect.Response[v1.ListIdentityProvidersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.IdentityProviderService.ListIdentityProviders is not implemented"))
}

func (UnimplementedIdentityProviderServiceHandler) CreateIdentityProvider(context.Context, *connect.Request[v1.CreateIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.IdentityProviderService.CreateIdentityProvider is not implemented"))
}

func (UnimplementedIdentityProviderServiceHandler) UpdateIdentityProvider(context.Context, *connect.Request[v1.UpdateIdentityProviderRequest]) (*connect.Response[v1.IdentityProvider], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.IdentityProviderService.UpdateIdentityProvider is not implemented"))
}

func (UnimplementedIdentityProviderServiceHandler) DeleteIdentityProvider(context.Context, *connect.Request[v1.DeleteIdentityProviderRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.IdentityProviderService.DeleteIdentityProvider is not implemented"))
}

func (UnimplementedIdentityProviderServiceHandler) TestIdentityProvider(context.Context, *connect.Request[v1.TestIdentityProviderRequest]) (*connect.Response[v1.TestIdentityProviderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.IdentityProviderService.TestIdentityProvider is not implemented"))
}
