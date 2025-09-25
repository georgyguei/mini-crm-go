// Package storage provides different storage implementations for contact persistence
package storage

import "mini-crm/internal/contact"

// Storer defines the interface for different storage backends
// This allows for dependency injection and easy swapping of storage mechanisms
type Storer interface {
	// Embed the contact repository interface
	contact.Repository
	// Close closes the storage connection if applicable
	Close() error
}
