package models

type MigrationFunc func() error

// Model is an interface that all models must implement.
type Model interface {
	RunMigration() error
}
