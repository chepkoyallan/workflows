package main

import (
	"context"
	"flag"
	"github.com/chepkoyallan/workflows/backend/src/apiserver/resource"
	"github.com/chepkoyallan/workflows/backend/src/apiserver/server"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
)

var (
	rpcPortFlag        = flag.String("rpcPortFlag", ":8887", "RPC Port")
	httpPortFlag       = flag.String("httpPortFlag", ":8888", "Http Proxy Port")
	configPath         = flag.String("config", "", "Path to JSON configuration file")
	collectMetricsFlag = flag.Bool("collectMetricsFlag", true, "Whether to collect Prometheus metrics in the API Server")
	defaultNamespace   = flag.String("defaultNamespace", "", "Default namespace")
)

func main() {
	flag.Parse()
	clientManager := newClientManager()
	resourceManager := resource.NewResourceManager(&clientManager, *defaultNamespace)
	//fmt.Printf("client manager: %v, Value: %T", resourceManager, resourceManager)
	//go startRpcServer(resourceManager)
	startHttpProxy(resourceManager)
}

func startRpcServer(resourceManager *resource.Manager) {
	glog.Info("Starting GRPC Server")

}

func startHttpProxy(resourceManager *resource.Manager) {
	glog.Info("Starting Http Proxy")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create gRPC HTTP MUX and register services for v1beta1 api.
	//runtimeMux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(grpcCustomMatcher))
	runtimeMux := runtime.NewServeMux()

	// create a top level mux to include both workflow upload server and the grpc servers
	topMux := mux.NewRouter()

	sharedWorkflowUploadServer := server.NewWorkflowUploadServer(resourceManager, &server.WorkflowUploadServerOptions{
		*collectMetricsFlag,
	})

	//
	topMux.HandleFunc("/apis/v2beta1/workflows/upload", sharedWorkflowUploadServer.UploadWorkflow)
	topMux.PathPrefix("/apis/").Handler(runtimeMux)

	http.ListenAndServe(*httpPortFlag, topMux)
	glog.Info("Http Proxy started")
}
