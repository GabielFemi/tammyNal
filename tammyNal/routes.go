package tammyNal

import (
	"fmt"
	"net/http"
)

func Index (w http.ResponseWriter, r *http.Request) {
	_ , _ = fmt.Fprintf(w, "Hey buddy")
}