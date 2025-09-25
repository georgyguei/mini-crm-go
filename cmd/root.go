package cmd

import (
	"fmt"
	"os"

	"mini-crm/internal/config"
	"mini-crm/internal/contact"
	"mini-crm/internal/storage"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cfg     *config.Config
	service contact.Service
	store   storage.Storer
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "A simple and efficient command-line contact manager",
	Long: `Mini CRM is a professional command-line contact management system built with Go.

It demonstrates best practices including:
• Decoupled package architecture
• Dependency injection via interfaces  
• Professional CLI with Cobra
• External configuration with Viper
• Multiple storage backends (Memory, JSON, SQLite)

Switch between storage types by editing config.yaml - no recompilation needed!`,
	PersistentPreRunE: initializeApp,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	defer func() {
		if store != nil {
			store.Close()
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	// Bind flags to viper
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

// initConfig reads in config file and ENV variables.
func initConfig() {
	var err error
	cfg, err = config.Load("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid configuration: %v\n", err)
		os.Exit(1)
	}
}

// initializeApp initializes the storage and service layers
func initializeApp(cmd *cobra.Command, args []string) error {
	// Use factory pattern for cleaner storage creation
	factory := storage.NewFactory()

	var err error
	store, err = factory.CreateStorage(cfg.Storage.Type, cfg.GetStorageFilePath())
	if err != nil {
		return err
	}

	// Initialize service with dependency injection
	service = contact.NewService(store)

	return nil
}
