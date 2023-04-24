package resource

// ClientManagerInterface client management interface
type ClientManagerInterface interface {
}

// Manager resource manager
type Manager struct {
}

// NewResourceManager  creates an instance of a resource manager
func NewResourceManager(clientManager ClientManagerInterface, defaultNamespace string) *Manager {
	return &Manager{}
}