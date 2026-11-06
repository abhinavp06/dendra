package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/abhinavp06/dendra/gateway/utils"
	"github.com/google/uuid"
)

const maxClaimBodyBytes = 1 << 20 // 1 MiB

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", healthHandler)
	mux.HandleFunc("POST /claims/register", registerClaimHandler)

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func registerClaimHandler(w http.ResponseWriter, r *http.Request) {
	var claim Claim

	claim.ReceivedAt = time.Now()

	r.Body = http.MaxBytesReader(w, r.Body, maxClaimBodyBytes)
	if err := json.NewDecoder(r.Body).Decode(&claim); err != nil {
		if _, tooLarge := errors.AsType[*http.MaxBytesError](err); tooLarge {
			utils.WriteJSONError(w, http.StatusRequestEntityTooLarge, "Request body too large")
			return
		}
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if claim.ClaimID == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Claim ID is required")
		return
	}

	if claim.ClaimCreatedAt.IsZero() {
		utils.WriteJSONError(w, http.StatusBadRequest, "Expected a claim creation timestamp")
		return
	}

	if len(claim.FileMetadata) == 0 {
		utils.WriteJSONError(w, http.StatusBadRequest, "Please provide file metadata")
		return
	}
	for i := range claim.FileMetadata {
		file := &claim.FileMetadata[i]
		if file.FilePointer == "" {
			utils.WriteJSONError(w, http.StatusBadRequest, "Expected a file URL")
			return
		}

		file.DFileID = uuid.New()
	}

	claim.TxnID = uuid.New()

	utils.WriteJSON(w, http.StatusAccepted, map[string]string{"txn_id": claim.TxnID.String()})
}
