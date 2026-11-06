package main

import (
	"time"

	"github.com/google/uuid"
)

type ClaimFile struct {
	FileID         string    `json:"file_id"` // this indicates a file_id that the source system maintains.. gateway then creates it's own file_id which dendra uses for it's internal tracking
	FilePointer    string    `json:"file_pointer"`
	AltFilePointer string    `json:"alt_file_pointer"`
	DFileID        uuid.UUID `json:"-"` // gateway's own file_id
}

type Claim struct {
	// expected from partner
	ClaimID        string      `json:"claim_id"`
	ClaimCreatedAt time.Time   `json:"claim_created_at"`
	FileMetadata   []ClaimFile `json:"file_metadata"`

	// gateway's own fields it injects
	TxnID      uuid.UUID `json:"-"`
	ReceivedAt time.Time `json:"-"`
}
