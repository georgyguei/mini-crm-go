# 🚀 Mini CRM CLI

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/doc/install)

A command-line Customer Relationship Management system built in Go, demonstrating clean architecture and professional development practices.

## ✨ Features

### 🎯 Core Functionality

- 📋 **Complete CRUD Operations**: Create, Read, Update, and Delete contacts
- 🎨 **Beautiful CLI Interface**: Professional command-line experience with Cobra
- ⚡ **Zero-Downtime Configuration**: Switch storage backends without recompiling
- 🛡️ **Smart Validation**: French mobile number validation and business rules
- 🔍 **Intuitive Commands**: Flag-based interface with comprehensive help

### 🗄️ Multiple Storage Backends

- 🧠 **Memory Storage**: Lightning-fast in-memory storage for testing
- 📄 **JSON Storage**: Human-readable file storage for small datasets
- 🗃️ **SQLite Database**: Production-ready database with GORM ORM

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/georgyguei/mini-crm-go
cd mini-crm

# Build the application
go build -o mini-crm

# Run it!
./mini-crm --help
```

_Requires Go 1.21 or later_

### Basic Usage

```bash
# Add a contact
./mini-crm add --name "John Doe" --email "john@example.com" --phone "0612345678"

# List all contacts
./mini-crm list

# Get specific contact
./mini-crm get 1

# Update a contact
./mini-crm update 1 --name "John Smith"

# Delete a contact (with confirmation)
./mini-crm delete 1
```

## ⚙️ Configuration

The magic ✨ of Mini CRM lies in its **zero-downtime configuration switching**. Simply edit `config.yaml` to change storage backends without recompiling!

```yaml
# config.yaml
app:
  name: "Mini CRM"
  version: "2.0.0"

storage:
  type: "gorm" # Switch between: memory, json, gorm
  filepath: "contacts.db" # Auto-adapts: contacts.json for JSON, contacts.db for SQLite
```

### Storage Options

| Type     | Description               | Use Case                  | Persistence    |
| -------- | ------------------------- | ------------------------- | -------------- |
| `memory` | In-memory HashMap         | Testing, demos            | ❌ Per-command |
| `json`   | Human-readable JSON       | Small datasets, debugging | ✅ File-based  |
| `gorm`   | SQLite database with GORM | Production use            | ✅ Database    |

## 🏛️ Architecture

```
mini-crm/
├── cmd/                    # 🎯 CLI Commands (Cobra)
│   ├── root.go            # Root command & initialization
│   ├── add.go             # Add contact command
│   ├── list.go            # List contacts command
│   ├── get.go             # Get contact command
│   ├── update.go          # Update contact command
│   └── delete.go          # Delete contact command
├── internal/               # 🔒 Private application code
│   ├── contact/           # 📋 Domain Layer
│   │   ├── contact.go     # Contact model & validation
│   │   └── service.go     # Business logic service
│   ├── storage/           # 💾 Data Access Layer
│   │   ├── interface.go   # Storage contract
│   │   ├── factory.go     # Storage factory pattern
│   │   ├── memory.go      # In-memory implementation
│   │   ├── json.go        # JSON file implementation
│   │   └── gorm.go        # SQLite/GORM implementation
│   └── config/            # ⚙️ Configuration Layer
│       └── config.go      # Viper configuration handling
├── config.yaml            # 📝 Application configuration
├── main.go                # 🚪 Application entry point
└── go.mod                 # 📦 Go module definition
```

## 💡 Code Examples

### Adding a New Storage Backend

Want to add Redis storage? Just implement the `Storer` interface:

```go
type RedisStore struct {
    client *redis.Client
}

func (r *RedisStore) Create(contact *contact.Contact) error {
    // Implement Redis storage logic
    return nil
}

func (r *RedisStore) GetAll() ([]*contact.Contact, error) {
    // Implement Redis retrieval
    return contacts, nil
}

// Implement other interface methods...
```

Then add it to the factory:

```go
case StorageTypeRedis:
    return NewRedisStore(filepath)
```

### Custom Validation Rules

Extend validation by modifying the Contact model:

```go
func (c *Contact) Validate() error {
    if err := c.basicValidation(); err != nil {
        return err
    }

    // Add custom business rules
    if strings.Contains(c.Name, "@") {
        return errors.New("name cannot contain @ symbol")
    }

    return nil
}
```

## 🎯 Advanced Usage

### Custom Configuration Paths

```bash
# Use custom config file
./mini-crm --config /path/to/custom/config.yaml list

# Environment-specific configs
./mini-crm --config ./config/production.yaml list
```

### Batch Operations

```bash
# Add multiple contacts
./mini-crm add --name "John Doe" --email "john@example.com"
./mini-crm add --name "Jane Smith" --email "jane@example.com"
./mini-crm add --name "Bob Wilson" --email "bob@example.com"

# Force delete without confirmation
./mini-crm delete 1 --force
```

### Storage Switching Examples

```bash
# Test with memory storage
echo 'storage: {type: "memory"}' > test-config.yaml
./mini-crm --config test-config.yaml add --name "Test User" --email "test@example.com"

# Switch to JSON for persistence
echo 'storage: {type: "json", filepath: "backup.json"}' > prod-config.yaml
./mini-crm --config prod-config.yaml list
```

## 🔬 Development

### Building

```bash
# Development build
go build -o mini-crm

# Production build with optimizations
go build -ldflags="-s -w" -o mini-crm

# Cross-platform builds
GOOS=windows GOARCH=amd64 go build -o mini-crm.exe
GOOS=linux GOARCH=amd64 go build -o mini-crm-linux
```

## 🔧 Troubleshooting

### Common Issues

**❓ "Contact not found" when using memory storage**  
💡 Memory storage doesn't persist between commands. Use JSON or GORM for persistence.

**❓ Database file permissions error**  
💡 Ensure directory is writable: `chmod 755 .`

**❓ Configuration file not found**  
💡 App searches: `./config.yaml` → `$HOME/.mini-crm/config.yaml` → `/etc/mini-crm/config.yaml`

**❓ Phone validation failing**  
💡 French mobile numbers must start with `06` or `07`

## 🤝 Contributing

We welcome contributions! This project is perfect for learning Go best practices.

### 💡 Contribution Ideas

- 📊 Export/import functionality (CSV, XML)
- 🔍 Advanced search and filtering
- 📱 International phone number validation

### Getting Started

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Follow existing code patterns and SOLID principles
4. Add tests for new functionality
5. Submit a pull request with clear description

## 📚 Learning Outcomes

This project demonstrates essential Go concepts and best practices:

- ✅ **Database Integration** - GORM/SQLite with auto-migration
- ✅ **Professional CLI** - Cobra & Viper integration
- ✅ **SOLID Architecture** - Maintainable "Lego brick" design
- ✅ **Multiple Storage Backends** - Seamless switching without recompilation
- ✅ **Production Ready** - Error handling, validation, and professional UX

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

<div align="center">

**Built for learning Go development best practices**

_Author_: **Georgy Guei** | _Course_: **Go Programming** | _Version_: **2.0.0** | _Date_: **September 2025**

</div>
