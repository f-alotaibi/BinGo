package api

import (
	"bingo/utils"
	"bingo/web"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type PasteService struct {
	ServiceName string
	Endpoint    string
}

func (paste PasteService) Init(router *httprouter.Router) {
	router.GET(fmt.Sprintf("/%s/:id", paste.ServiceName), paste.pasteServiceHandler)
}

func (paste PasteService) pasteServiceHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.Endpoint, id))
	if err != nil {
		http.Error(w, "Paste Not Found", http.StatusNotFound)
		println(fmt.Sprintf("%s: Paste not found %s", paste.ServiceName, id))
		return
	}
	queries := r.URL.Query()
	// If user is asking for raw paste
	if rawQuery, err := strconv.Atoi(queries.Get("raw")); err == nil && rawQuery == 1 {
		println(fmt.Sprintf("%s: Paste raw %s OK", paste.ServiceName, id))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "text/plain")
		_, err := w.Write([]byte(content))
		if err != nil {
			http.Error(w, "Could not write data", http.StatusInternalServerError)
			return
		}
		return
	}
	// If user is asking for downloading
	if downloadQuery, err := strconv.Atoi(queries.Get("dl")); err == nil && downloadQuery == 1 {
		println(fmt.Sprintf("%s: Paste download %s OK", paste.ServiceName, id))
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", id+".txt"))
		w.Header().Set("Content-type", "text/plain")
		w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		_, err := w.Write([]byte(content))
		if err != nil {
			http.Error(w, "Could not write data", http.StatusInternalServerError)
			return
		}
		return
	}
	// Otherwise, get the paste page
	println(fmt.Sprintf("%s: Paste web %s OK", paste.ServiceName, id))
	component := web.PasteComponent(paste.ServiceName, id, content)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Could not render component", http.StatusInternalServerError)
		return
	}
}
