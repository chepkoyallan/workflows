package main

import (
	"github.com/golang/glog"
)

// ClientManager Manages all the clients used on the platform
type ClientManager struct {
}

// newClientManager a factory function that creates and initializes new instance of the ClientManager
// returns a ClientManager
func newClientManager() ClientManager {
	clientManager := ClientManager{}
	clientManager.init()
	return clientManager
}

// This will run first when the package is referenced
func (c *ClientManager) init() {
	glog.Info("Initializing the client manager")
	glog.Info("Client manager initialized successfully")
}
