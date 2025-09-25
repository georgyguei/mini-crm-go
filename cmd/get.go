package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get a contact by ID",
	Long: `Retrieve and display detailed information about a specific contact by ID.
	
Example: mini-crm get 1`,
	Args: cobra.ExactArgs(1),
	RunE: runGetContact,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

// runGetContact handles the get contact command
func runGetContact(cmd *cobra.Command, args []string) error {
	// Parse contact ID
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid contact ID: %s", args[0])
	}

	// Get the contact
	contact, err := service.GetContact(uint(id))
	if err != nil {
		return fmt.Errorf("contact not found: %w", err)
	}

	// Display contact details
	fmt.Printf("📇 Contact Details\n")
	fmt.Printf("==================\n")
	fmt.Printf("ID: %d\n", contact.ID)
	fmt.Printf("Name: %s\n", contact.Name)
	fmt.Printf("Email: %s\n", contact.Email)
	if contact.Phone != "" {
		fmt.Printf("Phone: %s\n", contact.Phone)
	} else {
		fmt.Printf("Phone: N/A\n")
	}
	fmt.Printf("Created: %s\n", contact.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated: %s\n", contact.UpdatedAt.Format("2006-01-02 15:04:05"))

	return nil
}
