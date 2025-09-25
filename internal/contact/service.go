package contact

import (
	"errors"
	"fmt"
)

// service implements the Service interface with business logic
type service struct {
	repo Repository
}

// NewService creates a new contact service with dependency injection
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// CreateContact creates a new contact with validation
func (s *service) CreateContact(name, email, phone string) (*Contact, error) {
	// Check if email already exists
	existing, _ := s.repo.GetByEmail(email)
	if existing != nil {
		return nil, errors.New("contact with this email already exists")
	}

	contact := &Contact{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	if err := s.repo.Create(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

// ListContacts retrieves all contacts
func (s *service) ListContacts() ([]*Contact, error) {
	contacts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

// GetContact retrieves a contact by ID
func (s *service) GetContact(id uint) (*Contact, error) {
	contact, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

// UpdateContact updates an existing contact
func (s *service) UpdateContact(id uint, name, email, phone string) (*Contact, error) {
	// Get existing contact
	contact, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("contact not found: %w", err)
	}

	// Check if new email conflicts with another contact
	if email != contact.Email {
		existing, _ := s.repo.GetByEmail(email)
		if existing != nil && existing.ID != id {
			return nil, errors.New("another contact with this email already exists")
		}
	}

	// Update fields
	contact.Name = name
	contact.Email = email
	contact.Phone = phone

	if err := s.repo.Update(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

// DeleteContact removes a contact by ID
func (s *service) DeleteContact(id uint) error {
	// Check if contact exists
	_, err := s.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("contact not found: %w", err)
	}

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

// SearchByEmail finds a contact by email
func (s *service) SearchByEmail(email string) (*Contact, error) {
	contact, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("contact not found: %w", err)
	}
	return contact, nil
}
