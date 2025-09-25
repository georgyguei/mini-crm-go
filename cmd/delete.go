package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a contact",
	Long: `Delete a contact by ID.
	
This action requires confirmation unless --force flag is used.
Example: mini-crm delete 1`,
	Args: cobra.ExactArgs(1),
	RunE: runDeleteContact,
}

var forceDelete bool

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Flags for delete command
	deleteCmd.Flags().BoolVarP(&forceDelete, "force", "f", false, "Skip confirmation prompt")
}

// runDeleteContact handles the delete contact command
func runDeleteContact(cmd *cobra.Command, args []string) error {
	// Parse contact ID
	id, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid contact ID: %s", args[0])
	}

	// Get contact details for confirmation
	contact, err := service.GetContact(uint(id))
	if err != nil {
		return fmt.Errorf("contact not found: %w", err)
	}

	// Confirm deletion unless force flag is used
	if !forceDelete {
		fmt.Printf("⚠️  Are you sure you want to delete this contact?\n")
		fmt.Printf("ID: %d\n", contact.ID)
		fmt.Printf("Name: %s\n", contact.Name)
		fmt.Printf("Email: %s\n", contact.Email)
		if contact.Phone != "" {
			fmt.Printf("Phone: %s\n", contact.Phone)
		}
		fmt.Print("\nType 'yes' to confirm: ")

		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read confirmation: %w", err)
		}

		response = strings.TrimSpace(strings.ToLower(response))
		if response != "yes" {
			fmt.Println("❌ Delete cancelled.")
			return nil
		}
	}

	// Delete the contact
	if err := service.DeleteContact(uint(id)); err != nil {
		return fmt.Errorf("failed to delete contact: %w", err)
	}

	fmt.Printf("✅ Contact deleted successfully! (ID: %d, Name: %s)\n", contact.ID, contact.Name)
	return nil
}
