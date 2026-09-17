package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/teper-ya-pomenyal/privy_stream/gateway/internal/clients"
)

const defaultPageLimit = 20

type CatalogHandler struct {
	catalogClient *clients.CatalogClient
}

func NewCatalogHandler(catalogClient *clients.CatalogClient) *CatalogHandler {
	return &CatalogHandler{catalogClient: catalogClient}
}

func parsePageParams(r *http.Request) (int32, int32) {
	limit := int32(defaultPageLimit)
	offset := int32(0)
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = int32(n)
		}
	}
	if v := r.URL.Query().Get("offset"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			offset = int32(n)
		}
	}
	return limit, offset
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *CatalogHandler) GetTrackByID(w http.ResponseWriter, r *http.Request) {
	res, err := h.catalogClient.GetTrackByID(r.Context(), chi.URLParam(r, "track_uuid"))
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *CatalogHandler) TrackExists(w http.ResponseWriter, r *http.Request) {
	exists, err := h.catalogClient.TrackExists(r.Context(), chi.URLParam(r, "track_uuid"))
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"exists": exists})
}

func (h *CatalogHandler) SearchTrack(w http.ResponseWriter, r *http.Request) {
	trackName := r.URL.Query().Get("track_name")
	if trackName == "" {
		http.Error(w, "track_name is required", http.StatusBadRequest)
		return
	}
	limit, offset := parsePageParams(r)
	res, err := h.catalogClient.SearchTrack(r.Context(), trackName, limit, offset)
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *CatalogHandler) GetArtistByID(w http.ResponseWriter, r *http.Request) {
	res, err := h.catalogClient.GetArtistByID(r.Context(), chi.URLParam(r, "artist_uuid"))
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *CatalogHandler) SearchArtist(w http.ResponseWriter, r *http.Request) {
	artistName := r.URL.Query().Get("artist_name")
	if artistName == "" {
		http.Error(w, "artist_name is required", http.StatusBadRequest)
		return
	}
	limit, offset := parsePageParams(r)
	res, err := h.catalogClient.SearchArtist(r.Context(), artistName, limit, offset)
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *CatalogHandler) GetArtistAlbums(w http.ResponseWriter, r *http.Request) {
	limit, offset := parsePageParams(r)
	res, err := h.catalogClient.GetArtistAlbums(r.Context(), chi.URLParam(r, "artist_uuid"), limit, offset)
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *CatalogHandler) GetArtistTracks(w http.ResponseWriter, r *http.Request) {
	limit, offset := parsePageParams(r)
	res, err := h.catalogClient.GetArtistTracks(r.Context(), chi.URLParam(r, "artist_uuid"), limit, offset)
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *CatalogHandler) GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	res, err := h.catalogClient.GetAlbumByID(r.Context(), chi.URLParam(r, "album_uuid"))
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (h *CatalogHandler) GetAlbumTracks(w http.ResponseWriter, r *http.Request) {
	res, err := h.catalogClient.GetAlbumTracks(r.Context(), chi.URLParam(r, "album_uuid"))
	if err != nil {
		mapGRPCError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, res)
}
