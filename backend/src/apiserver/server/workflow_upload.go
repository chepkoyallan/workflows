package server

import (
	"fmt"
	"github.com/chepkoyallan/workflows/backend/src/apiserver/resource"
	"github.com/golang/glog"
	"net/http"
)

// WorkflowUploadServerOptions workflow upload server options
type WorkflowUploadServerOptions struct {
	CollectMetrics bool `json:"collect_metrics,omitempty"`
}

// WorkflowUploadServer uploads workflows
type WorkflowUploadServer struct {
	resourceManager *resource.Manager
	options         *WorkflowUploadServerOptions
}

// UploadWorkflow workflow upload api
func (s *WorkflowUploadServer) UploadWorkflow(w http.ResponseWriter, r *http.Request) {
	if s.options.CollectMetrics {

	}
	s.uploadWorkflow("v2beta1", w, r)
}

// uploadWorkflow creates a workflow and a workflow version
func (s *WorkflowUploadServer) uploadWorkflow(apiVersion string, w http.ResponseWriter, r *http.Request) {
	glog.Infof("Upload workflow called")
	fmt.Printf("api_version:%v \t w: %v, r: %v", apiVersion, w, r)
}

// NewWorkflowUploadServer constructor
func NewWorkflowUploadServer(resourceManager *resource.Manager, options *WorkflowUploadServerOptions) *WorkflowUploadServer {
	return &WorkflowUploadServer{resourceManager: resourceManager, options: options}
}
