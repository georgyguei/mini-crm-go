package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update an existing contact",
	Long: `Update an existing contact by ID.
	
You can provide new values via flags. Only provided fields will be updated.
Example: mini-crm update 1 --name "Jane Doe" --email "jane@newdomain.com"`,
	Args: cobra.ExactArgs(1),
	RunE: runUpdateContact,
}

var (
	updateName  string
	updateEmail string
	updatePhone string
)

func init() {
	rootCmd.AddCommand(updateCmd)

	// Flags for update command
	updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "New contact name")
	updateCmd.Flags().StringVarP(&updateEmail, "email", "e", "", "New contact email")
	updateCmd.Flags().StringVarP(&updatePhone, "phone", "p", "", "New contact phone")
}

// runUpdateContact handles the update contact command
func runUpdateContact(cmd *cobra.Command, args []string) error {
	// Parse contact ID
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid contact ID: %s", args[0])
	}

	// Get current contact to preserve unchanged fields
	currentContact, err := service.GetContact(uint(id))
	if err != nil {
		return fmt.Errorf("contact not found: %w", err)
	}

	// Use current values if flags not provided
	name := updateName
	if name == "" {
		name = currentContact.Name
	}

	email := updateEmail
	if email == "" {
		email = currentContact.Email
	}

	phone := updatePhone
	if !cmd.Flags().Changed("phone") {
		phone = currentContact.Phone
	}

	// Update the contact
	updatedContact, err := service.UpdateContact(uint(id), name, email, phone)
	if err != nil {
		return fmt.Errorf("failed to update contact: %w", err)
	}

	fmt.Printf("✅ Contact updated successfully!\n")
	fmt.Printf("ID: %d\n", updatedContact.ID)
	fmt.Printf("Name: %s\n", updatedContact.Name)
	fmt.Printf("Email: %s\n", updatedContact.Email)
	if updatedContact.Phone != "" {
		fmt.Printf("Phone: %s\n", updatedContact.Phone)
	}
	fmt.Printf("Updated: %s\n", updatedContact.UpdatedAt.Format("2006-01-02 15:04:05"))

	return nil
}
