package agent

import (
	protov1 "github.com/bryk-io/did-method/proto/v1"
	"go.bryk.io/x/ccg/did"
)

// Storage defines an abstract component that provides and manage
// persistent data requirements for DID documents.
type Storage interface {
	// Setup the instance and prepare for usage.
	Open(info string) error

	// Free resources and finish processing.
	Close() error

	// Returns a brief information summary for the storage instance.
	Description() string

	// Check if a record exists for the specified DID.
	Exists(id *did.Identifier) bool

	// Return a previously stored DID instance.
	Get(req *protov1.QueryRequest) (*did.Identifier, error)

	// Create or update the record for the given DID instance.
	Save(id *did.Identifier) error

	// Remove any existing records for the given DID instance.
	Delete(id *did.Identifier) error
}
