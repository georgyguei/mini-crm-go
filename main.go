package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Contact represents a contact in our CRM
type Contact struct {
	ID    int
	Name  string
	Email string
}

// CRM manages our contacts using a map
type CRM struct {
	contacts map[int]Contact
	nextID   int
}

// NewCRM creates a new CRM instance
func NewCRM() *CRM {
	return &CRM{
		contacts: make(map[int]Contact),
		nextID:   1,
	}
}

func main() {
	// Parse command line flags
	var (
		addFlag   = flag.Bool("add", false, "Add a contact via command line")
		nameFlag  = flag.String("name", "", "Contact name")
		emailFlag = flag.String("email", "", "Contact email")
	)
	flag.Parse()

	crm := NewCRM()

	// Handle command line flags for adding contacts
	if *addFlag {
		if *nameFlag == "" || *emailFlag == "" {
			fmt.Println("Error: Both -name and -email flags are required when using -add")
			fmt.Println("Usage: go run main.go -add -name=\"John Doe\" -email=\"john@example.com\"")
			return
		}

		contact := Contact{
			ID:    crm.nextID,
			Name:  *nameFlag,
			Email: *emailFlag,
		}

		crm.contacts[crm.nextID] = contact
		crm.nextID++

		fmt.Printf("Contact added successfully: %+v\n", contact)
		return
	}

	// Start interactive menu
	crm.runMenu()
}

// runMenu displays the interactive menu and handles user input
func (crm *CRM) runMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Mini CRM ===")
		fmt.Println("1. Add a contact")
		fmt.Println("2. List all contacts")
		fmt.Println("3. Search for a contact by ID")
		fmt.Println("4. Update a contact")
		fmt.Println("5. Delete a contact")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option (1-6): ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %v\n", err)
			}
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			crm.addContact(scanner)
		case "2":
			crm.listContacts()
		case "3":
			crm.searchContact(scanner)
		case "4":
			crm.updateContact(scanner)
		case "5":
			crm.deleteContact(scanner)
		case "6":
			fmt.Println("Thank you for using Mini CRM!")
			return
		default:
			fmt.Println("Invalid option. Please choose 1-6.")
		}
	}
}

// addContact adds a new contact
func (crm *CRM) addContact(scanner *bufio.Scanner) {
	fmt.Print("Enter contact name: ")
	if !scanner.Scan() {
		fmt.Println("Error reading name")
		return
	}
	name := strings.TrimSpace(scanner.Text())

	if name == "" {
		fmt.Println("Error: Name cannot be empty")
		return
	}

	fmt.Print("Enter contact email: ")
	if !scanner.Scan() {
		fmt.Println("Error reading email")
		return
	}
	email := strings.TrimSpace(scanner.Text())

	if email == "" {
		fmt.Println("Error: Email cannot be empty")
		return
	}

	// Simple email validation
	if !strings.Contains(email, "@") {
		fmt.Println("Error: Please enter a valid email address")
		return
	}

	contact := Contact{
		ID:    crm.nextID,
		Name:  name,
		Email: email,
	}

	crm.contacts[crm.nextID] = contact
	crm.nextID++

	fmt.Printf("Contact added successfully! ID: %d\n", contact.ID)
}

// listContacts displays all contacts
func (crm *CRM) listContacts() {
	if len(crm.contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("\n=== All Contacts ===")
	fmt.Printf("%-5s %-20s %-30s\n", "ID", "Name", "Email")
	fmt.Println(strings.Repeat("-", 55))

	for id, contact := range crm.contacts {
		fmt.Printf("%-5d %-20s %-30s\n", id, contact.Name, contact.Email)
	}
}

// searchContact searches for a contact by ID using the comma ok idiom
func (crm *CRM) searchContact(scanner *bufio.Scanner) {
	fmt.Print("Enter contact ID to search: ")
	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}

	idStr := strings.TrimSpace(scanner.Text())
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Error: Invalid ID '%s'. Please enter a valid number.\n", idStr)
		return
	}

	// Using comma ok idiom to check if contact exists
	contact, ok := crm.contacts[id]
	if !ok {
		fmt.Printf("Contact with ID %d not found.\n", id)
		return
	}

	fmt.Println("\n=== Contact Found ===")
	fmt.Printf("ID: %d\n", contact.ID)
	fmt.Printf("Name: %s\n", contact.Name)
	fmt.Printf("Email: %s\n", contact.Email)
}

// updateContact updates an existing contact
func (crm *CRM) updateContact(scanner *bufio.Scanner) {
	fmt.Print("Enter contact ID to update: ")
	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}

	idStr := strings.TrimSpace(scanner.Text())
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Error: Invalid ID '%s'. Please enter a valid number.\n", idStr)
		return
	}

	// Using comma ok idiom to check if contact exists
	contact, ok := crm.contacts[id]
	if !ok {
		fmt.Printf("Contact with ID %d not found.\n", id)
		return
	}

	// Display current contact info
	fmt.Println("\n=== Current Contact Info ===")
	fmt.Printf("ID: %d\n", contact.ID)
	fmt.Printf("Name: %s\n", contact.Name)
	fmt.Printf("Email: %s\n", contact.Email)

	// Get new name
	fmt.Printf("Enter new name (current: %s, press Enter to keep): ", contact.Name)
	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}
	newName := strings.TrimSpace(scanner.Text())
	if newName == "" {
		newName = contact.Name
	}

	// Get new email
	fmt.Printf("Enter new email (current: %s, press Enter to keep): ", contact.Email)
	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}
	newEmail := strings.TrimSpace(scanner.Text())
	if newEmail == "" {
		newEmail = contact.Email
	} else if !strings.Contains(newEmail, "@") {
		fmt.Println("Error: Please enter a valid email address")
		return
	}

	// Update the contact
	updatedContact := Contact{
		ID:    id,
		Name:  newName,
		Email: newEmail,
	}

	crm.contacts[id] = updatedContact
	fmt.Println("Contact updated successfully!")
}

// deleteContact removes a contact by ID
func (crm *CRM) deleteContact(scanner *bufio.Scanner) {
	fmt.Print("Enter contact ID to delete: ")
	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}

	idStr := strings.TrimSpace(scanner.Text())
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Error: Invalid ID '%s'. Please enter a valid number.\n", idStr)
		return
	}

	// Using comma ok idiom to check if contact exists
	contact, ok := crm.contacts[id]
	if !ok {
		fmt.Printf("Contact with ID %d not found.\n", id)
		return
	}

	// Show contact to be deleted and ask for confirmation
	fmt.Printf("Are you sure you want to delete contact: %s (%s)? (y/N): ", contact.Name, contact.Email)
	if !scanner.Scan() {
		fmt.Println("Error reading input")
		return
	}

	confirmation := strings.TrimSpace(strings.ToLower(scanner.Text()))
	if confirmation == "y" || confirmation == "yes" {
		delete(crm.contacts, id)
		fmt.Println("Contact deleted successfully!")
	} else {
		fmt.Println("Delete operation cancelled.")
	}
}
