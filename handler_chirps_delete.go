package main

import (
	"net/http"

	"github.com/Bgoodwin24/chirpy/internal/auth"
	"github.com/Bgoodwin24/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirps(w http.ResponseWriter, r *http.Request) {
	// Extract chirpID from URL
	path := r.URL.Path
	chirpID := path[len("/api/chirps/"):]

	if chirpID == "" {
		respondWithError(w, http.StatusBadRequest, "Missing chirpID in path", nil)
		return
	}

	// Convert chirpID to uuid.UUID
	uuidChirp, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirpID format", err)
		return
	}

	// Validate the token and extract userID from it
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get token", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	// Retrieve the chirp from the database
	chirp, err := cfg.db.GetChirpByID(r.Context(), uuidChirp)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "Access Denied", nil)
		return
	}

	err = cfg.db.DeleteChirpByID(r.Context(), database.DeleteChirpByIDParams{
		ID:     chirp.ID,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete chirp", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
