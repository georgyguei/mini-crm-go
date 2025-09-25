package contact

// Repository defines the interface for contact storage operations
// This follows the Repository pattern for clean architecture
type Repository interface {
	// Create adds a new contact to storage
	Create(contact *Contact) error

	// GetByID retrieves a contact by its ID
	GetByID(id uint) (*Contact, error)

	// GetAll retrieves all contacts from storage
	GetAll() ([]*Contact, error)

	// Update modifies an existing contact
	Update(contact *Contact) error

	// Delete removes a contact by ID
	Delete(id uint) error

	// GetByEmail finds a contact by email address
	GetByEmail(email string) (*Contact, error)
}

// Service defines the business logic operations for contact management
// This layer contains business rules and orchestrates repository calls
type Service interface {
	// CreateContact creates a new contact with validation
	CreateContact(name, email, phone string) (*Contact, error)

	// ListContacts retrieves all contacts
	ListContacts() ([]*Contact, error)

	// GetContact retrieves a contact by ID
	GetContact(id uint) (*Contact, error)

	// UpdateContact updates an existing contact
	UpdateContact(id uint, name, email, phone string) (*Contact, error)

	// DeleteContact removes a contact by ID
	DeleteContact(id uint) error

	// SearchByEmail finds a contact by email
	SearchByEmail(email string) (*Contact, error)
}
