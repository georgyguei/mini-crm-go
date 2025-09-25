package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contacts",
	Long: `List all contacts in the CRM system.
	
Displays contacts in a formatted table with ID, name, email, phone, and creation date.`,
	RunE: runListContacts,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// runListContacts handles the list contacts command
func runListContacts(cmd *cobra.Command, args []string) error {
	contacts, err := service.ListContacts()
	if err != nil {
		return fmt.Errorf("failed to retrieve contacts: %w", err)
	}

	if len(contacts) == 0 {
		fmt.Println("📭 No contacts found.")
		return nil
	}

	// Create tabwriter for formatted output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	// Print header
	fmt.Fprintf(w, "ID\tName\tEmail\tPhone\tCreated\n")
	fmt.Fprintf(w, "--\t----\t-----\t-----\t-------\n")

	// Print each contact
	for _, contact := range contacts {
		phone := contact.Phone
		if phone == "" {
			phone = "N/A"
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
			contact.ID,
			contact.Name,
			contact.Email,
			phone,
			contact.CreatedAt.Format("2006-01-02 15:04"))
	}

	fmt.Printf("\n📊 Total contacts: %d\n", len(contacts))
	return nil
}
