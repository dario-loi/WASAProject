package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// applyCORSHandler applies a CORS policy to the router. CORS stands for Cross-Origin Resource Sharing: it's a security
// feature present in web browsers that blocks JavaScript requests going across different domains if not specified in a
// policy. This function sends the policy of this API server.
func applyCORSHandler(h http.Handler) http.Handler {
	// Some of the allowed headers are legacy.
	return handlers.CORS(
		handlers.AllowedHeaders([]string{
			"x-example-header",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.ExposedHeaders([]string{"Authorization"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "searcher_id",
			"identifier", "requester_ID", "user_id", "user_name", "commenter_name"}))(h)
}
