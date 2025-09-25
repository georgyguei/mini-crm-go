package storage

import (
	"fmt"
)

// Factory provides a clean way to create storage instances
// Implements the Factory Pattern for better separation of concerns
type Factory struct{}

// NewFactory creates a new storage factory
func NewFactory() *Factory {
	return &Factory{}
}

// CreateStorage creates a storage instance based on the type and configuration
// This centralizes storage creation logic and makes it easy to add new storage types
func (f *Factory) CreateStorage(storageType string, filePath string) (Storer, error) {
	switch storageType {
	case "memory":
		return NewMemoryStore(), nil
	case "json":
		return NewJSONStore(filePath)
	case "gorm":
		return NewGORMStore(filePath)
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", storageType)
	}
}

// GetSupportedTypes returns a list of supported storage types
func (f *Factory) GetSupportedTypes() []string {
	return []string{"memory", "json", "gorm"}
}
