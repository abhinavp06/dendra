package routing

import (
	"fmt"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
