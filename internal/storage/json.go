package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"mini-crm/internal/contact"
)

// JSONStore provides JSON file-based storage
// Implements the Single Responsibility Principle by focusing only on JSON file operations
type JSONStore struct {
	filename string
	contacts map[uint]*contact.Contact
	nextID   uint
	mu       sync.RWMutex
}

// NewJSONStore creates a new JSON file storage instance
func NewJSONStore(filename string) (Storer, error) {
	store := &JSONStore{
		filename: filename,
		contacts: make(map[uint]*contact.Contact),
		nextID:   1,
	}

	if err := store.load(); err != nil {
		return nil, fmt.Errorf("failed to load JSON store: %w", err)
	}

	return store, nil
}

// load reads contacts from the JSON file
func (j *JSONStore) load() error {
	if _, err := os.Stat(j.filename); os.IsNotExist(err) {
		// File doesn't exist, start fresh
		return nil
	}

	data, err := os.ReadFile(j.filename)
	if err != nil {
		return err
	}

	var contacts []*contact.Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return err
	}

	for _, c := range contacts {
		j.contacts[c.ID] = c
		if c.ID >= j.nextID {
			j.nextID = c.ID + 1
		}
	}

	return nil
}

// save writes contacts to the JSON file
func (j *JSONStore) save() error {
	contacts := make([]*contact.Contact, 0, len(j.contacts))
	for _, c := range j.contacts {
		contacts = append(contacts, c)
	}

	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(j.filename, data, 0644)
}

// Create adds a new contact to JSON storage
func (j *JSONStore) Create(c *contact.Contact) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	if err := c.Validate(); err != nil {
		return err
	}

	c.ID = j.nextID
	// Set timestamps for JSON storage
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now

	j.contacts[c.ID] = c
	j.nextID++

	return j.save()
}

// GetByID retrieves a contact by its ID from JSON storage
func (j *JSONStore) GetByID(id uint) (*contact.Contact, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()

	c, exists := j.contacts[id]
	if !exists {
		return nil, fmt.Errorf("there is no contact with ID %d", id)
	}
	return c, nil
}

// GetAll retrieves all contacts from JSON storage
func (j *JSONStore) GetAll() ([]*contact.Contact, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()

	contacts := make([]*contact.Contact, 0, len(j.contacts))
	for _, c := range j.contacts {
		contacts = append(contacts, c)
	}
	return contacts, nil
}

// Update modifies an existing contact in JSON storage
func (j *JSONStore) Update(c *contact.Contact) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	if _, exists := j.contacts[c.ID]; !exists {
		return errors.New("contact not found")
	}

	if err := c.Validate(); err != nil {
		return err
	}

	// Update timestamp
	c.UpdatedAt = time.Now()

	j.contacts[c.ID] = c
	return j.save()
}

// Delete removes a contact by ID from JSON storage
func (j *JSONStore) Delete(id uint) error {
	j.mu.Lock()
	defer j.mu.Unlock()

	if _, exists := j.contacts[id]; !exists {
		return errors.New("contact not found")
	}

	delete(j.contacts, id)
	return j.save()
}

// GetByEmail finds a contact by email address in JSON storage
func (j *JSONStore) GetByEmail(email string) (*contact.Contact, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()

	for _, c := range j.contacts {
		if c.Email == email {
			return c, nil
		}
	}
	return nil, errors.New("contact not found")
}

// Close closes the JSON store (no-op for file)
func (j *JSONStore) Close() error {
	return nil
}
