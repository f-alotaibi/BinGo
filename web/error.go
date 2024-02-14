package web

import "net/http"

func GiveErrorPage(w http.ResponseWriter, r *http.Request, code int, errorString string) {
	errComponent := ErrorComponent(code, errorString)
	componentError := errComponent.Render(r.Context(), w)
	if componentError != nil {
		http.Error(w, "Could not render error component", http.StatusInternalServerError)
		return
	}
}
