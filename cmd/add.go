package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact",
	Long: `Add a new contact to the CRM system.
	
You can provide contact information via flags or interactively.
Example: mini-crm add --name "John Doe" --email "john@example.com" --phone "0612345678"`,
	RunE: runAddContact,
}

var (
	addName  string
	addEmail string
	addPhone string
)

func init() {
	rootCmd.AddCommand(addCmd)

	// Flags for add command
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "Contact name (required)")
	addCmd.Flags().StringVarP(&addEmail, "email", "e", "", "Contact email (required)")
	addCmd.Flags().StringVarP(&addPhone, "phone", "p", "", "Contact phone (optional)")

	// Mark required flags
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}

// runAddContact handles the add contact command
func runAddContact(cmd *cobra.Command, args []string) error {
	contact, err := service.CreateContact(addName, addEmail, addPhone)
	if err != nil {
		return fmt.Errorf("failed to create contact: %w", err)
	}

	fmt.Printf("✅ Contact added successfully!\n")
	fmt.Printf("ID: %d\n", contact.ID)
	fmt.Printf("Name: %s\n", contact.Name)
	fmt.Printf("Email: %s\n", contact.Email)
	if contact.Phone != "" {
		fmt.Printf("Phone: %s\n", contact.Phone)
	}
	fmt.Printf("Created: %s\n", contact.CreatedAt.Format("2006-01-02 15:04:05"))

	return nil
}
