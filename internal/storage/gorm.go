package storage

import (
	"errors"
	"fmt"

	"mini-crm/internal/contact"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GORMStore provides GORM/SQLite-based storage
// Implements the Single Responsibility Principle by focusing only on database operations
type GORMStore struct {
	db *gorm.DB
}

// NewGORMStore creates a new GORM storage instance with SQLite
func NewGORMStore(dbPath string) (Storer, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto migrate the contact schema
	if err := db.AutoMigrate(&contact.Contact{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &GORMStore{db: db}, nil
}

// Create adds a new contact to GORM storage
func (g *GORMStore) Create(c *contact.Contact) error {
	if err := g.db.Create(c).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a contact by its ID from GORM storage
func (g *GORMStore) GetByID(id uint) (*contact.Contact, error) {
	var c contact.Contact
	if err := g.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact not found")
		}
		return nil, err
	}
	return &c, nil
}

// GetAll retrieves all contacts from GORM storage
func (g *GORMStore) GetAll() ([]*contact.Contact, error) {
	var contacts []*contact.Contact
	if err := g.db.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

// Update modifies an existing contact in GORM storage
func (g *GORMStore) Update(c *contact.Contact) error {
	return g.db.Save(c).Error
}

// Delete removes a contact by ID from GORM storage
func (g *GORMStore) Delete(id uint) error {
	return g.db.Delete(&contact.Contact{}, id).Error
}

// GetByEmail finds a contact by email address in GORM storage
func (g *GORMStore) GetByEmail(email string) (*contact.Contact, error) {
	var c contact.Contact
	if err := g.db.Where("email = ?", email).First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact not found")
		}
		return nil, err
	}
	return &c, nil
}

// Close closes the GORM database connection
func (g *GORMStore) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
