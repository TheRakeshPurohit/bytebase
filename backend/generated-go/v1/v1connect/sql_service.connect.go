// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/sql_service.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/bytebase/bytebase/backend/generated-go/v1"
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
	// SQLServiceName is the fully-qualified name of the SQLService service.
	SQLServiceName = "bytebase.v1.SQLService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SQLServiceQueryProcedure is the fully-qualified name of the SQLService's Query RPC.
	SQLServiceQueryProcedure = "/bytebase.v1.SQLService/Query"
	// SQLServiceAdminExecuteProcedure is the fully-qualified name of the SQLService's AdminExecute RPC.
	SQLServiceAdminExecuteProcedure = "/bytebase.v1.SQLService/AdminExecute"
	// SQLServiceSearchQueryHistoriesProcedure is the fully-qualified name of the SQLService's
	// SearchQueryHistories RPC.
	SQLServiceSearchQueryHistoriesProcedure = "/bytebase.v1.SQLService/SearchQueryHistories"
	// SQLServiceExportProcedure is the fully-qualified name of the SQLService's Export RPC.
	SQLServiceExportProcedure = "/bytebase.v1.SQLService/Export"
	// SQLServiceCheckProcedure is the fully-qualified name of the SQLService's Check RPC.
	SQLServiceCheckProcedure = "/bytebase.v1.SQLService/Check"
	// SQLServicePrettyProcedure is the fully-qualified name of the SQLService's Pretty RPC.
	SQLServicePrettyProcedure = "/bytebase.v1.SQLService/Pretty"
	// SQLServiceDiffMetadataProcedure is the fully-qualified name of the SQLService's DiffMetadata RPC.
	SQLServiceDiffMetadataProcedure = "/bytebase.v1.SQLService/DiffMetadata"
	// SQLServiceAICompletionProcedure is the fully-qualified name of the SQLService's AICompletion RPC.
	SQLServiceAICompletionProcedure = "/bytebase.v1.SQLService/AICompletion"
)

// SQLServiceClient is a client for the bytebase.v1.SQLService service.
type SQLServiceClient interface {
	// Permissions required: bb.databases.get
	Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error)
	// Permissions required: bb.sql.admin
	AdminExecute(context.Context) *connect.BidiStreamForClient[v1.AdminExecuteRequest, v1.AdminExecuteResponse]
	// SearchQueryHistories searches query histories for the caller.
	// Permissions required: None
	SearchQueryHistories(context.Context, *connect.Request[v1.SearchQueryHistoriesRequest]) (*connect.Response[v1.SearchQueryHistoriesResponse], error)
	// Permissions required: bb.databases.get
	Export(context.Context, *connect.Request[v1.ExportRequest]) (*connect.Response[v1.ExportResponse], error)
	// Permissions required: bb.databases.check
	Check(context.Context, *connect.Request[v1.CheckRequest]) (*connect.Response[v1.CheckResponse], error)
	// Permissions required: None
	Pretty(context.Context, *connect.Request[v1.PrettyRequest]) (*connect.Response[v1.PrettyResponse], error)
	// Permissions required: None
	DiffMetadata(context.Context, *connect.Request[v1.DiffMetadataRequest]) (*connect.Response[v1.DiffMetadataResponse], error)
	// Permissions required: None
	AICompletion(context.Context, *connect.Request[v1.AICompletionRequest]) (*connect.Response[v1.AICompletionResponse], error)
}

// NewSQLServiceClient constructs a client for the bytebase.v1.SQLService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSQLServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SQLServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	sQLServiceMethods := v1.File_v1_sql_service_proto.Services().ByName("SQLService").Methods()
	return &sQLServiceClient{
		query: connect.NewClient[v1.QueryRequest, v1.QueryResponse](
			httpClient,
			baseURL+SQLServiceQueryProcedure,
			connect.WithSchema(sQLServiceMethods.ByName("Query")),
			connect.WithClientOptions(opts...),
		),
		adminExecute: connect.NewClient[v1.AdminExecuteRequest, v1.AdminExecuteResponse](
			httpClient,
			baseURL+SQLServiceAdminExecuteProcedure,
			connect.WithSchema(sQLServiceMethods.ByName("AdminExecute")),
			connect.WithClientOptions(opts...),
		),
		searchQueryHistories: connect.NewClient[v1.SearchQueryHistoriesRequest, v1.SearchQueryHistoriesResponse](
			httpClient,
			baseURL+SQLServiceSearchQueryHistoriesProcedure,
			connect.WithSchema(sQLServiceMethods.ByName("SearchQueryHistories")),
			connect.WithClientOptions(opts...),
		),
		export: connect.NewClient[v1.ExportRequest, v1.ExportResponse](
			httpClient,
			baseURL+SQLServiceExportProcedure,
			connect.WithSchema(sQLServiceMethods.ByName("Export")),
			connect.WithClientOptions(opts...),
		),
		check: connect.NewClient[v1.CheckRequest, v1.CheckResponse](
			httpClient,
			baseURL+SQLServiceCheckProcedure,
			connect.WithSchema(sQLServiceMethods.ByName("Check")),
			connect.WithClientOptions(opts...),
		),
		pretty: connect.NewClient[v1.PrettyRequest, v1.PrettyResponse](
			httpClient,
			baseURL+SQLServicePrettyProcedure,
			connect.WithSchema(sQLServiceMethods.ByName("Pretty")),
			connect.WithClientOptions(opts...),
		),
		diffMetadata: connect.NewClient[v1.DiffMetadataRequest, v1.DiffMetadataResponse](
			httpClient,
			baseURL+SQLServiceDiffMetadataProcedure,
			connect.WithSchema(sQLServiceMethods.ByName("DiffMetadata")),
			connect.WithClientOptions(opts...),
		),
		aICompletion: connect.NewClient[v1.AICompletionRequest, v1.AICompletionResponse](
			httpClient,
			baseURL+SQLServiceAICompletionProcedure,
			connect.WithSchema(sQLServiceMethods.ByName("AICompletion")),
			connect.WithClientOptions(opts...),
		),
	}
}

// sQLServiceClient implements SQLServiceClient.
type sQLServiceClient struct {
	query                *connect.Client[v1.QueryRequest, v1.QueryResponse]
	adminExecute         *connect.Client[v1.AdminExecuteRequest, v1.AdminExecuteResponse]
	searchQueryHistories *connect.Client[v1.SearchQueryHistoriesRequest, v1.SearchQueryHistoriesResponse]
	export               *connect.Client[v1.ExportRequest, v1.ExportResponse]
	check                *connect.Client[v1.CheckRequest, v1.CheckResponse]
	pretty               *connect.Client[v1.PrettyRequest, v1.PrettyResponse]
	diffMetadata         *connect.Client[v1.DiffMetadataRequest, v1.DiffMetadataResponse]
	aICompletion         *connect.Client[v1.AICompletionRequest, v1.AICompletionResponse]
}

// Query calls bytebase.v1.SQLService.Query.
func (c *sQLServiceClient) Query(ctx context.Context, req *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error) {
	return c.query.CallUnary(ctx, req)
}

// AdminExecute calls bytebase.v1.SQLService.AdminExecute.
func (c *sQLServiceClient) AdminExecute(ctx context.Context) *connect.BidiStreamForClient[v1.AdminExecuteRequest, v1.AdminExecuteResponse] {
	return c.adminExecute.CallBidiStream(ctx)
}

// SearchQueryHistories calls bytebase.v1.SQLService.SearchQueryHistories.
func (c *sQLServiceClient) SearchQueryHistories(ctx context.Context, req *connect.Request[v1.SearchQueryHistoriesRequest]) (*connect.Response[v1.SearchQueryHistoriesResponse], error) {
	return c.searchQueryHistories.CallUnary(ctx, req)
}

// Export calls bytebase.v1.SQLService.Export.
func (c *sQLServiceClient) Export(ctx context.Context, req *connect.Request[v1.ExportRequest]) (*connect.Response[v1.ExportResponse], error) {
	return c.export.CallUnary(ctx, req)
}

// Check calls bytebase.v1.SQLService.Check.
func (c *sQLServiceClient) Check(ctx context.Context, req *connect.Request[v1.CheckRequest]) (*connect.Response[v1.CheckResponse], error) {
	return c.check.CallUnary(ctx, req)
}

// Pretty calls bytebase.v1.SQLService.Pretty.
func (c *sQLServiceClient) Pretty(ctx context.Context, req *connect.Request[v1.PrettyRequest]) (*connect.Response[v1.PrettyResponse], error) {
	return c.pretty.CallUnary(ctx, req)
}

// DiffMetadata calls bytebase.v1.SQLService.DiffMetadata.
func (c *sQLServiceClient) DiffMetadata(ctx context.Context, req *connect.Request[v1.DiffMetadataRequest]) (*connect.Response[v1.DiffMetadataResponse], error) {
	return c.diffMetadata.CallUnary(ctx, req)
}

// AICompletion calls bytebase.v1.SQLService.AICompletion.
func (c *sQLServiceClient) AICompletion(ctx context.Context, req *connect.Request[v1.AICompletionRequest]) (*connect.Response[v1.AICompletionResponse], error) {
	return c.aICompletion.CallUnary(ctx, req)
}

// SQLServiceHandler is an implementation of the bytebase.v1.SQLService service.
type SQLServiceHandler interface {
	// Permissions required: bb.databases.get
	Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error)
	// Permissions required: bb.sql.admin
	AdminExecute(context.Context, *connect.BidiStream[v1.AdminExecuteRequest, v1.AdminExecuteResponse]) error
	// SearchQueryHistories searches query histories for the caller.
	// Permissions required: None
	SearchQueryHistories(context.Context, *connect.Request[v1.SearchQueryHistoriesRequest]) (*connect.Response[v1.SearchQueryHistoriesResponse], error)
	// Permissions required: bb.databases.get
	Export(context.Context, *connect.Request[v1.ExportRequest]) (*connect.Response[v1.ExportResponse], error)
	// Permissions required: bb.databases.check
	Check(context.Context, *connect.Request[v1.CheckRequest]) (*connect.Response[v1.CheckResponse], error)
	// Permissions required: None
	Pretty(context.Context, *connect.Request[v1.PrettyRequest]) (*connect.Response[v1.PrettyResponse], error)
	// Permissions required: None
	DiffMetadata(context.Context, *connect.Request[v1.DiffMetadataRequest]) (*connect.Response[v1.DiffMetadataResponse], error)
	// Permissions required: None
	AICompletion(context.Context, *connect.Request[v1.AICompletionRequest]) (*connect.Response[v1.AICompletionResponse], error)
}

// NewSQLServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSQLServiceHandler(svc SQLServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	sQLServiceMethods := v1.File_v1_sql_service_proto.Services().ByName("SQLService").Methods()
	sQLServiceQueryHandler := connect.NewUnaryHandler(
		SQLServiceQueryProcedure,
		svc.Query,
		connect.WithSchema(sQLServiceMethods.ByName("Query")),
		connect.WithHandlerOptions(opts...),
	)
	sQLServiceAdminExecuteHandler := connect.NewBidiStreamHandler(
		SQLServiceAdminExecuteProcedure,
		svc.AdminExecute,
		connect.WithSchema(sQLServiceMethods.ByName("AdminExecute")),
		connect.WithHandlerOptions(opts...),
	)
	sQLServiceSearchQueryHistoriesHandler := connect.NewUnaryHandler(
		SQLServiceSearchQueryHistoriesProcedure,
		svc.SearchQueryHistories,
		connect.WithSchema(sQLServiceMethods.ByName("SearchQueryHistories")),
		connect.WithHandlerOptions(opts...),
	)
	sQLServiceExportHandler := connect.NewUnaryHandler(
		SQLServiceExportProcedure,
		svc.Export,
		connect.WithSchema(sQLServiceMethods.ByName("Export")),
		connect.WithHandlerOptions(opts...),
	)
	sQLServiceCheckHandler := connect.NewUnaryHandler(
		SQLServiceCheckProcedure,
		svc.Check,
		connect.WithSchema(sQLServiceMethods.ByName("Check")),
		connect.WithHandlerOptions(opts...),
	)
	sQLServicePrettyHandler := connect.NewUnaryHandler(
		SQLServicePrettyProcedure,
		svc.Pretty,
		connect.WithSchema(sQLServiceMethods.ByName("Pretty")),
		connect.WithHandlerOptions(opts...),
	)
	sQLServiceDiffMetadataHandler := connect.NewUnaryHandler(
		SQLServiceDiffMetadataProcedure,
		svc.DiffMetadata,
		connect.WithSchema(sQLServiceMethods.ByName("DiffMetadata")),
		connect.WithHandlerOptions(opts...),
	)
	sQLServiceAICompletionHandler := connect.NewUnaryHandler(
		SQLServiceAICompletionProcedure,
		svc.AICompletion,
		connect.WithSchema(sQLServiceMethods.ByName("AICompletion")),
		connect.WithHandlerOptions(opts...),
	)
	return "/bytebase.v1.SQLService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SQLServiceQueryProcedure:
			sQLServiceQueryHandler.ServeHTTP(w, r)
		case SQLServiceAdminExecuteProcedure:
			sQLServiceAdminExecuteHandler.ServeHTTP(w, r)
		case SQLServiceSearchQueryHistoriesProcedure:
			sQLServiceSearchQueryHistoriesHandler.ServeHTTP(w, r)
		case SQLServiceExportProcedure:
			sQLServiceExportHandler.ServeHTTP(w, r)
		case SQLServiceCheckProcedure:
			sQLServiceCheckHandler.ServeHTTP(w, r)
		case SQLServicePrettyProcedure:
			sQLServicePrettyHandler.ServeHTTP(w, r)
		case SQLServiceDiffMetadataProcedure:
			sQLServiceDiffMetadataHandler.ServeHTTP(w, r)
		case SQLServiceAICompletionProcedure:
			sQLServiceAICompletionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSQLServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSQLServiceHandler struct{}

func (UnimplementedSQLServiceHandler) Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.SQLService.Query is not implemented"))
}

func (UnimplementedSQLServiceHandler) AdminExecute(context.Context, *connect.BidiStream[v1.AdminExecuteRequest, v1.AdminExecuteResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.SQLService.AdminExecute is not implemented"))
}

func (UnimplementedSQLServiceHandler) SearchQueryHistories(context.Context, *connect.Request[v1.SearchQueryHistoriesRequest]) (*connect.Response[v1.SearchQueryHistoriesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.SQLService.SearchQueryHistories is not implemented"))
}

func (UnimplementedSQLServiceHandler) Export(context.Context, *connect.Request[v1.ExportRequest]) (*connect.Response[v1.ExportResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.SQLService.Export is not implemented"))
}

func (UnimplementedSQLServiceHandler) Check(context.Context, *connect.Request[v1.CheckRequest]) (*connect.Response[v1.CheckResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.SQLService.Check is not implemented"))
}

func (UnimplementedSQLServiceHandler) Pretty(context.Context, *connect.Request[v1.PrettyRequest]) (*connect.Response[v1.PrettyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.SQLService.Pretty is not implemented"))
}

func (UnimplementedSQLServiceHandler) DiffMetadata(context.Context, *connect.Request[v1.DiffMetadataRequest]) (*connect.Response[v1.DiffMetadataResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.SQLService.DiffMetadata is not implemented"))
}

func (UnimplementedSQLServiceHandler) AICompletion(context.Context, *connect.Request[v1.AICompletionRequest]) (*connect.Response[v1.AICompletionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bytebase.v1.SQLService.AICompletion is not implemented"))
}
