package issue

import (
	"html/template"
	"net/http"

	"github.com/opstalent/tracker/logger"
	"golang.org/x/net/context"
)

var (
	funcs = template.FuncMap{"countTotal": countTotal}
	tmpl = template.Must(template.New("list.html").Funcs(funcs).ParseFiles("views/issue/list.html"))
)

type Data struct {
	List SortedIssues
}

func listHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	issues, err := Get(ctx, r)
	if err != nil {
		log := logger.New()
		log.Critical(ctx, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		data := &Data{Sort(issues)}
		render(w, data)
	}
}

func render(w http.ResponseWriter, args *Data) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := tmpl.Execute(w, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func countTotal(s map[string][]Issue) int {
	total := 0
	for _, p := range s {
		total += len(p)
	}

	return total
}
