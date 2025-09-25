// Package contact provides the core domain models and interfaces for contact management
package contact

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Contact represents a contact in our CRM system
// It follows the domain model pattern with validation
type Contact struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Validate performs business logic validation on the contact
func (c *Contact) Validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return errors.New("name cannot be empty")
	}

	if strings.TrimSpace(c.Email) == "" {
		return errors.New("email cannot be empty")
	}

	if !strings.Contains(c.Email, "@") {
		return errors.New("invalid email format")
	}

	// Validate phone format if provided (must start with 06 or 07 for French mobile)
	if c.Phone != "" && !strings.HasPrefix(c.Phone, "06") && !strings.HasPrefix(c.Phone, "07") {
		return errors.New("phone number must start with '06' or '07' if provided")
	}

	return nil
}

// BeforeCreate is a GORM hook that runs before creating a record
func (c *Contact) BeforeCreate(tx *gorm.DB) error {
	c.Name = strings.TrimSpace(c.Name)
	c.Email = strings.TrimSpace(strings.ToLower(c.Email))
	c.Phone = strings.TrimSpace(c.Phone)
	return c.Validate()
}

// BeforeUpdate is a GORM hook that runs before updating a record
func (c *Contact) BeforeUpdate(tx *gorm.DB) error {
	c.Name = strings.TrimSpace(c.Name)
	c.Email = strings.TrimSpace(strings.ToLower(c.Email))
	c.Phone = strings.TrimSpace(c.Phone)
	return c.Validate()
}
