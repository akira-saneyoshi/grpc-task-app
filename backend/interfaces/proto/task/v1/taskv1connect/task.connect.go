// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/task/v1/task.proto

package taskv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/akira-saneyoshi/task-app/interfaces/proto/task/v1"
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
	// TaskServiceName is the fully-qualified name of the TaskService service.
	TaskServiceName = "proto.task.v1.TaskService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TaskServiceGetTaskListProcedure is the fully-qualified name of the TaskService's GetTaskList RPC.
	TaskServiceGetTaskListProcedure = "/proto.task.v1.TaskService/GetTaskList"
	// TaskServiceCreateTaskProcedure is the fully-qualified name of the TaskService's CreateTask RPC.
	TaskServiceCreateTaskProcedure = "/proto.task.v1.TaskService/CreateTask"
	// TaskServiceUpdateTaskStatusProcedure is the fully-qualified name of the TaskService's
	// UpdateTaskStatus RPC.
	TaskServiceUpdateTaskStatusProcedure = "/proto.task.v1.TaskService/UpdateTaskStatus"
	// TaskServiceUpdateTaskDueDateProcedure is the fully-qualified name of the TaskService's
	// UpdateTaskDueDate RPC.
	TaskServiceUpdateTaskDueDateProcedure = "/proto.task.v1.TaskService/UpdateTaskDueDate"
	// TaskServiceUpdateTaskDetailsProcedure is the fully-qualified name of the TaskService's
	// UpdateTaskDetails RPC.
	TaskServiceUpdateTaskDetailsProcedure = "/proto.task.v1.TaskService/UpdateTaskDetails"
	// TaskServiceDeleteTaskProcedure is the fully-qualified name of the TaskService's DeleteTask RPC.
	TaskServiceDeleteTaskProcedure = "/proto.task.v1.TaskService/DeleteTask"
)

// TaskServiceClient is a client for the proto.task.v1.TaskService service.
type TaskServiceClient interface {
	GetTaskList(context.Context, *connect.Request[v1.GetTaskListRequest]) (*connect.Response[v1.GetTaskListResponse], error)
	CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.CreateTaskResponse], error)
	UpdateTaskStatus(context.Context, *connect.Request[v1.UpdateTaskStatusRequest]) (*connect.Response[v1.UpdateTaskStatusResponse], error)
	UpdateTaskDueDate(context.Context, *connect.Request[v1.UpdateTaskDueDateRequest]) (*connect.Response[v1.UpdateTaskDueDateResponse], error)
	UpdateTaskDetails(context.Context, *connect.Request[v1.UpdateTaskDetailsRequest]) (*connect.Response[v1.UpdateTaskDetailsResponse], error)
	DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[v1.DeleteTaskResponse], error)
}

// NewTaskServiceClient constructs a client for the proto.task.v1.TaskService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTaskServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TaskServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	taskServiceMethods := v1.File_proto_task_v1_task_proto.Services().ByName("TaskService").Methods()
	return &taskServiceClient{
		getTaskList: connect.NewClient[v1.GetTaskListRequest, v1.GetTaskListResponse](
			httpClient,
			baseURL+TaskServiceGetTaskListProcedure,
			connect.WithSchema(taskServiceMethods.ByName("GetTaskList")),
			connect.WithClientOptions(opts...),
		),
		createTask: connect.NewClient[v1.CreateTaskRequest, v1.CreateTaskResponse](
			httpClient,
			baseURL+TaskServiceCreateTaskProcedure,
			connect.WithSchema(taskServiceMethods.ByName("CreateTask")),
			connect.WithClientOptions(opts...),
		),
		updateTaskStatus: connect.NewClient[v1.UpdateTaskStatusRequest, v1.UpdateTaskStatusResponse](
			httpClient,
			baseURL+TaskServiceUpdateTaskStatusProcedure,
			connect.WithSchema(taskServiceMethods.ByName("UpdateTaskStatus")),
			connect.WithClientOptions(opts...),
		),
		updateTaskDueDate: connect.NewClient[v1.UpdateTaskDueDateRequest, v1.UpdateTaskDueDateResponse](
			httpClient,
			baseURL+TaskServiceUpdateTaskDueDateProcedure,
			connect.WithSchema(taskServiceMethods.ByName("UpdateTaskDueDate")),
			connect.WithClientOptions(opts...),
		),
		updateTaskDetails: connect.NewClient[v1.UpdateTaskDetailsRequest, v1.UpdateTaskDetailsResponse](
			httpClient,
			baseURL+TaskServiceUpdateTaskDetailsProcedure,
			connect.WithSchema(taskServiceMethods.ByName("UpdateTaskDetails")),
			connect.WithClientOptions(opts...),
		),
		deleteTask: connect.NewClient[v1.DeleteTaskRequest, v1.DeleteTaskResponse](
			httpClient,
			baseURL+TaskServiceDeleteTaskProcedure,
			connect.WithSchema(taskServiceMethods.ByName("DeleteTask")),
			connect.WithClientOptions(opts...),
		),
	}
}

// taskServiceClient implements TaskServiceClient.
type taskServiceClient struct {
	getTaskList       *connect.Client[v1.GetTaskListRequest, v1.GetTaskListResponse]
	createTask        *connect.Client[v1.CreateTaskRequest, v1.CreateTaskResponse]
	updateTaskStatus  *connect.Client[v1.UpdateTaskStatusRequest, v1.UpdateTaskStatusResponse]
	updateTaskDueDate *connect.Client[v1.UpdateTaskDueDateRequest, v1.UpdateTaskDueDateResponse]
	updateTaskDetails *connect.Client[v1.UpdateTaskDetailsRequest, v1.UpdateTaskDetailsResponse]
	deleteTask        *connect.Client[v1.DeleteTaskRequest, v1.DeleteTaskResponse]
}

// GetTaskList calls proto.task.v1.TaskService.GetTaskList.
func (c *taskServiceClient) GetTaskList(ctx context.Context, req *connect.Request[v1.GetTaskListRequest]) (*connect.Response[v1.GetTaskListResponse], error) {
	return c.getTaskList.CallUnary(ctx, req)
}

// CreateTask calls proto.task.v1.TaskService.CreateTask.
func (c *taskServiceClient) CreateTask(ctx context.Context, req *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.CreateTaskResponse], error) {
	return c.createTask.CallUnary(ctx, req)
}

// UpdateTaskStatus calls proto.task.v1.TaskService.UpdateTaskStatus.
func (c *taskServiceClient) UpdateTaskStatus(ctx context.Context, req *connect.Request[v1.UpdateTaskStatusRequest]) (*connect.Response[v1.UpdateTaskStatusResponse], error) {
	return c.updateTaskStatus.CallUnary(ctx, req)
}

// UpdateTaskDueDate calls proto.task.v1.TaskService.UpdateTaskDueDate.
func (c *taskServiceClient) UpdateTaskDueDate(ctx context.Context, req *connect.Request[v1.UpdateTaskDueDateRequest]) (*connect.Response[v1.UpdateTaskDueDateResponse], error) {
	return c.updateTaskDueDate.CallUnary(ctx, req)
}

// UpdateTaskDetails calls proto.task.v1.TaskService.UpdateTaskDetails.
func (c *taskServiceClient) UpdateTaskDetails(ctx context.Context, req *connect.Request[v1.UpdateTaskDetailsRequest]) (*connect.Response[v1.UpdateTaskDetailsResponse], error) {
	return c.updateTaskDetails.CallUnary(ctx, req)
}

// DeleteTask calls proto.task.v1.TaskService.DeleteTask.
func (c *taskServiceClient) DeleteTask(ctx context.Context, req *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[v1.DeleteTaskResponse], error) {
	return c.deleteTask.CallUnary(ctx, req)
}

// TaskServiceHandler is an implementation of the proto.task.v1.TaskService service.
type TaskServiceHandler interface {
	GetTaskList(context.Context, *connect.Request[v1.GetTaskListRequest]) (*connect.Response[v1.GetTaskListResponse], error)
	CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.CreateTaskResponse], error)
	UpdateTaskStatus(context.Context, *connect.Request[v1.UpdateTaskStatusRequest]) (*connect.Response[v1.UpdateTaskStatusResponse], error)
	UpdateTaskDueDate(context.Context, *connect.Request[v1.UpdateTaskDueDateRequest]) (*connect.Response[v1.UpdateTaskDueDateResponse], error)
	UpdateTaskDetails(context.Context, *connect.Request[v1.UpdateTaskDetailsRequest]) (*connect.Response[v1.UpdateTaskDetailsResponse], error)
	DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[v1.DeleteTaskResponse], error)
}

// NewTaskServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTaskServiceHandler(svc TaskServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	taskServiceMethods := v1.File_proto_task_v1_task_proto.Services().ByName("TaskService").Methods()
	taskServiceGetTaskListHandler := connect.NewUnaryHandler(
		TaskServiceGetTaskListProcedure,
		svc.GetTaskList,
		connect.WithSchema(taskServiceMethods.ByName("GetTaskList")),
		connect.WithHandlerOptions(opts...),
	)
	taskServiceCreateTaskHandler := connect.NewUnaryHandler(
		TaskServiceCreateTaskProcedure,
		svc.CreateTask,
		connect.WithSchema(taskServiceMethods.ByName("CreateTask")),
		connect.WithHandlerOptions(opts...),
	)
	taskServiceUpdateTaskStatusHandler := connect.NewUnaryHandler(
		TaskServiceUpdateTaskStatusProcedure,
		svc.UpdateTaskStatus,
		connect.WithSchema(taskServiceMethods.ByName("UpdateTaskStatus")),
		connect.WithHandlerOptions(opts...),
	)
	taskServiceUpdateTaskDueDateHandler := connect.NewUnaryHandler(
		TaskServiceUpdateTaskDueDateProcedure,
		svc.UpdateTaskDueDate,
		connect.WithSchema(taskServiceMethods.ByName("UpdateTaskDueDate")),
		connect.WithHandlerOptions(opts...),
	)
	taskServiceUpdateTaskDetailsHandler := connect.NewUnaryHandler(
		TaskServiceUpdateTaskDetailsProcedure,
		svc.UpdateTaskDetails,
		connect.WithSchema(taskServiceMethods.ByName("UpdateTaskDetails")),
		connect.WithHandlerOptions(opts...),
	)
	taskServiceDeleteTaskHandler := connect.NewUnaryHandler(
		TaskServiceDeleteTaskProcedure,
		svc.DeleteTask,
		connect.WithSchema(taskServiceMethods.ByName("DeleteTask")),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.task.v1.TaskService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TaskServiceGetTaskListProcedure:
			taskServiceGetTaskListHandler.ServeHTTP(w, r)
		case TaskServiceCreateTaskProcedure:
			taskServiceCreateTaskHandler.ServeHTTP(w, r)
		case TaskServiceUpdateTaskStatusProcedure:
			taskServiceUpdateTaskStatusHandler.ServeHTTP(w, r)
		case TaskServiceUpdateTaskDueDateProcedure:
			taskServiceUpdateTaskDueDateHandler.ServeHTTP(w, r)
		case TaskServiceUpdateTaskDetailsProcedure:
			taskServiceUpdateTaskDetailsHandler.ServeHTTP(w, r)
		case TaskServiceDeleteTaskProcedure:
			taskServiceDeleteTaskHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTaskServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTaskServiceHandler struct{}

func (UnimplementedTaskServiceHandler) GetTaskList(context.Context, *connect.Request[v1.GetTaskListRequest]) (*connect.Response[v1.GetTaskListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.task.v1.TaskService.GetTaskList is not implemented"))
}

func (UnimplementedTaskServiceHandler) CreateTask(context.Context, *connect.Request[v1.CreateTaskRequest]) (*connect.Response[v1.CreateTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.task.v1.TaskService.CreateTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) UpdateTaskStatus(context.Context, *connect.Request[v1.UpdateTaskStatusRequest]) (*connect.Response[v1.UpdateTaskStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.task.v1.TaskService.UpdateTaskStatus is not implemented"))
}

func (UnimplementedTaskServiceHandler) UpdateTaskDueDate(context.Context, *connect.Request[v1.UpdateTaskDueDateRequest]) (*connect.Response[v1.UpdateTaskDueDateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.task.v1.TaskService.UpdateTaskDueDate is not implemented"))
}

func (UnimplementedTaskServiceHandler) UpdateTaskDetails(context.Context, *connect.Request[v1.UpdateTaskDetailsRequest]) (*connect.Response[v1.UpdateTaskDetailsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.task.v1.TaskService.UpdateTaskDetails is not implemented"))
}

func (UnimplementedTaskServiceHandler) DeleteTask(context.Context, *connect.Request[v1.DeleteTaskRequest]) (*connect.Response[v1.DeleteTaskResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.task.v1.TaskService.DeleteTask is not implemented"))
}
