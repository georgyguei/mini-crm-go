# Mini Command Line CRM

A simple command-line Customer Relationship Management (CRM) system built with Go. This project demonstrates key Go concepts including maps, structs, error handling, the comma ok idiom, and command-line interfaces.

## Features

✅ **Interactive Menu System**
- Clean, user-friendly command-line interface
- Easy navigation with numbered options

✅ **Contact Management**
- Add contacts with ID, Name, and Email
- List all contacts in a formatted table
- Search for contacts by ID
- Update existing contact information
- Delete contacts with confirmation

✅ **Command-Line Flags Support**
- Add contacts directly via command-line arguments
- Quick contact creation without interactive menu

✅ **Data Validation**
- Input validation for names and emails
- Error handling for invalid IDs
- Confirmation prompts for destructive operations

## Installation

1. **Clone the repository:**
   ```bash
   git clone <your-repo-url>
   cd mini-crm
   ```

2. **Ensure you have Go installed:**
   ```bash
   go version
   ```
   *Requires Go 1.21 or later*

3. **Run the application:**
   ```bash
   go run main.go
   ```

## Usage

### Interactive Mode

Run the application without flags to enter interactive mode:

```bash
go run main.go
```

You'll see the main menu:
```
=== Mini CRM ===
1. Add a contact
2. List all contacts
3. Search for a contact by ID
4. Update a contact
5. Delete a contact
6. Exit
Choose an option (1-6):
```

### Command-Line Mode

Add contacts directly using flags:

```bash
# Add a contact via command line
go run main.go -add -name="John Doe" -email="john@example.com"

# Add another contact
go run main.go -add -name="Jane Smith" -email="jane.smith@company.com"
```

## Examples

### Adding a Contact (Interactive)
```
Choose an option (1-6): 1
Enter contact name: Alice Johnson
Enter contact email: alice@email.com
Contact added successfully! ID: 1
```

### Listing All Contacts
```
Choose an option (1-6): 2

=== All Contacts ===
ID    Name                 Email                         
-------------------------------------------------------
1     Alice Johnson        alice@email.com               
2     Bob Wilson          bob@company.com               
3     Carol Davis         carol.davis@startup.io        
```

### Searching for a Contact
```
Choose an option (1-6): 3
Enter contact ID to search: 2

=== Contact Found ===
ID: 2
Name: Bob Wilson
Email: bob@company.com
```

### Updating a Contact
```
Choose an option (1-6): 4
Enter contact ID to update: 1

=== Current Contact Info ===
ID: 1
Name: Alice Johnson
Email: alice@email.com
Enter new name (current: Alice Johnson, press Enter to keep): Alice J. Johnson
Enter new email (current: alice@email.com, press Enter to keep): 
Contact updated successfully!
```

### Deleting a Contact
```
Choose an option (1-6): 5
Enter contact ID to delete: 3
Are you sure you want to delete contact: Carol Davis (carol.davis@startup.io)? (y/N): y
Contact deleted successfully!
```

## Go Concepts Demonstrated

### 🗺️ **Maps**
- Uses `map[int]Contact` to store contacts with ID as key
- Demonstrates map initialization, insertion, lookup, and deletion

### 🔍 **Comma OK Idiom**
```go
contact, ok := crm.contacts[id]
if !ok {
    fmt.Printf("Contact with ID %d not found.\n", id)
    return
}
```

### 🔄 **Control Structures**
- `switch` statement for menu navigation
- `for` loop for continuous menu operation
- `range` for iterating over contacts map

### ⚠️ **Error Handling**
- `if err != nil` pattern for error checking
- Input validation and user-friendly error messages

### 📊 **String Conversion**
- `strconv.Atoi()` for converting string input to integers
- Proper error handling for conversion failures

### 📥 **Input/Output**
- `bufio.Scanner` for reading user input from `os.Stdin`
- `flag` package for command-line argument parsing

### 🏗️ **Structs and Methods**
- `Contact` struct with ID, Name, Email fields
- `CRM` struct with methods for contact management
- Constructor pattern with `NewCRM()`

## Project Structure

```
mini-crm/
├── main.go          # Main application code
├── go.mod           # Go module definition
└── README.md        # This file
```

## Code Architecture

- **Contact Struct**: Represents individual contacts
- **CRM Struct**: Manages the collection of contacts and provides methods
- **Separation of Concerns**: Menu handling, input validation, and business logic are separated
- **Error Handling**: Comprehensive error checking throughout the application

## Building and Distribution

### Build executable:
```bash
go build -o mini-crm main.go
```

### Run the executable:
```bash
./mini-crm
```

### Cross-platform builds:
```bash
# For Windows
GOOS=windows GOARCH=amd64 go build -o mini-crm.exe main.go

# For Linux
GOOS=linux GOARCH=amd64 go build -o mini-crm-linux main.go

# For macOS
GOOS=darwin GOARCH=amd64 go build -o mini-crm-mac main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Learning Objectives Achieved

- ✅ Working with Go maps for data storage
- ✅ Implementing the comma ok idiom
- ✅ Using `for` loops and `switch` statements
- ✅ Error handling with `if err != nil`
- ✅ String conversion with `strconv`
- ✅ Reading from `os.Stdin` with `bufio`
- ✅ Command-line flag parsing
- ✅ Struct definition and method implementation
- ✅ Input validation and user experience design

## License

This project is open source and available under the [MIT License](LICENSE).

---

**Author**: Your Name  
**Course**: Go Programming  
**Project**: Mini Command Line CRM  
**Date**: September 2025