package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// applyCORSHandler applies a CORS policy to the router. CORS stands for Cross-Origin Resource Sharing: it's a security
// feature present in web browsers that blocks JavaScript requests going across different domains if not specified in a
// policy. This function sends the policy of this API server.
func applyCORSHandler(h http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{
			"x-example-header",
			"*",
		}),
		handlers.AllowedHeaders(([]string{"Content-Type", "Authorization", "username"})),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*", "http://127.0.0.1:5173"}),
		handlers.ExposedHeaders([]string{"Authorization"}),
	)(h)
}
