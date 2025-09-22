package main

import (
	"testing"
)

func TestNewCRM(t *testing.T) {
	crm := NewCRM()

	if crm == nil {
		t.Fatal("NewCRM() returned nil")
	}

	if crm.contacts == nil {
		t.Error("contacts map should be initialized")
	}

	if crm.nextID != 1 {
		t.Errorf("Expected nextID to be 1, got %d", crm.nextID)
	}

	if len(crm.contacts) != 0 {
		t.Errorf("Expected empty contacts map, got %d contacts", len(crm.contacts))
	}
}

func TestContactStorage(t *testing.T) {
	crm := NewCRM()

	// Test adding a contact
	contact := Contact{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	crm.contacts[contact.ID] = contact

	// Test comma ok idiom
	retrievedContact, ok := crm.contacts[1]
	if !ok {
		t.Error("Contact should exist in map")
	}

	if retrievedContact.Name != "Test User" {
		t.Errorf("Expected name 'Test User', got '%s'", retrievedContact.Name)
	}

	if retrievedContact.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", retrievedContact.Email)
	}

	// Test non-existent contact using comma ok idiom
	_, ok = crm.contacts[999]
	if ok {
		t.Error("Non-existent contact should not be found")
	}
}

func TestMapOperations(t *testing.T) {
	crm := NewCRM()

	// Add multiple contacts
	contacts := []Contact{
		{ID: 1, Name: "Alice", Email: "alice@test.com"},
		{ID: 2, Name: "Bob", Email: "bob@test.com"},
		{ID: 3, Name: "Charlie", Email: "charlie@test.com"},
	}

	for _, contact := range contacts {
		crm.contacts[contact.ID] = contact
	}

	// Test map length
	if len(crm.contacts) != 3 {
		t.Errorf("Expected 3 contacts, got %d", len(crm.contacts))
	}

	// Test range over map
	count := 0
	for id, contact := range crm.contacts {
		count++
		if contact.ID != id {
			t.Errorf("Contact ID mismatch: map key %d, contact ID %d", id, contact.ID)
		}
	}

	if count != 3 {
		t.Errorf("Expected to iterate over 3 contacts, got %d", count)
	}

	// Test delete operation
	delete(crm.contacts, 2)

	if len(crm.contacts) != 2 {
		t.Errorf("Expected 2 contacts after deletion, got %d", len(crm.contacts))
	}

	_, ok := crm.contacts[2]
	if ok {
		t.Error("Deleted contact should not exist")
	}
}
