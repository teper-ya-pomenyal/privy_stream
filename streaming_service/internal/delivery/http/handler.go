package http

type HTTPHandlers struct {
	streamer Streamer
	catalog  Catalog
}

func NewHTTPHandler(streamer Streamer, catalog Catalog) *HTTPHandlers {
	return &HTTPHandlers{streamer: streamer, catalog: catalog}
}
