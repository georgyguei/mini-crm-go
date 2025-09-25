package storage

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"mini-crm/internal/contact"
)

// MemoryStore provides in-memory storage for testing and development
// Implements the Single Responsibility Principle by focusing only on memory operations
type MemoryStore struct {
	contacts map[uint]*contact.Contact
	nextID   uint
	mu       sync.RWMutex
}

// NewMemoryStore creates a new in-memory storage instance
func NewMemoryStore() Storer {
	return &MemoryStore{
		contacts: make(map[uint]*contact.Contact),
		nextID:   1,
	}
}

// Create adds a new contact to memory storage
func (m *MemoryStore) Create(c *contact.Contact) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Validate before storing
	if err := c.Validate(); err != nil {
		return err
	}

	c.ID = m.nextID
	// Set timestamps for memory storage
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now

	m.contacts[c.ID] = c
	m.nextID++
	return nil
}

// GetByID retrieves a contact by its ID from memory
func (m *MemoryStore) GetByID(id uint) (*contact.Contact, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	c, exists := m.contacts[id]
	if !exists {
		return nil, fmt.Errorf("there is no contact with ID %d", id)
	}
	return c, nil
}

// GetAll retrieves all contacts from memory
func (m *MemoryStore) GetAll() ([]*contact.Contact, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	contacts := make([]*contact.Contact, 0, len(m.contacts))
	for _, c := range m.contacts {
		contacts = append(contacts, c)
	}
	return contacts, nil
}

// Update modifies an existing contact in memory
func (m *MemoryStore) Update(c *contact.Contact) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.contacts[c.ID]; !exists {
		return errors.New("contact not found")
	}

	if err := c.Validate(); err != nil {
		return err
	}

	// Update timestamp
	c.UpdatedAt = time.Now()

	m.contacts[c.ID] = c
	return nil
}

// Delete removes a contact by ID from memory
func (m *MemoryStore) Delete(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.contacts[id]; !exists {
		return errors.New("contact not found")
	}

	delete(m.contacts, id)
	return nil
}

// GetByEmail finds a contact by email address in memory
func (m *MemoryStore) GetByEmail(email string) (*contact.Contact, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, c := range m.contacts {
		if c.Email == email {
			return c, nil
		}
	}
	return nil, errors.New("contact not found")
}

// Close closes the memory store (no-op for memory)
func (m *MemoryStore) Close() error {
	return nil
}
