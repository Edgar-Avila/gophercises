package website

import (
	"choose-your-own-adventure/story"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type RootHandler struct {
}

func (rh RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var st story.Story
	st.Load("./gopher.json")
	tmpl, err := template.ParseFiles("./website/arc.html")
	if err != nil {
		log.Fatal(err)
	}
    arcname := strings.Trim(r.URL.Path, " /")
	arc, ok := st.GetArc(arcname)
    if !ok {
        arc, _ = st.GetArc("intro")
    }
    tmpl.Execute(w, arc)
}

func Start() error {
	http.Handle("/", new(RootHandler))
    fmt.Println("Server started at port 8080")
	return http.ListenAndServe(":8080", nil)
}
