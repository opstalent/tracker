package issue

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/opstalent/tracker/env"
	"github.com/opstalent/tracker/user"
	"golang.org/x/net/context"
)

type element struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

var (
	funcs = template.FuncMap{
		"countTotal": countTotal,
		"getDonut": func(users map[string]map[string][]Issue) template.JS {
			var data []element
			for _, status := range users {
				for pName, issues := range status {
					data = append(data, element{pName, len(issues)})
				}
			}
			b, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
				return ""
			}
			return template.JS(b)
		},
		"userSlug": func(name string) string {
			return strings.Replace(name, " ", "_", -1)
		},
	}
	tmpl = template.Must(template.New("list.html").Funcs(funcs).ParseFiles("views/issue/list.html"))
)

type Data struct {
	List SortedIssues
}

func listHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var issues = &Issues{}
	users, err := user.Get(ctx, r)
	for _, user := range users.Resources {
		if is, err := Get(ctx, r, &user); err == nil {
			issues.Resources = append(issues.Resources, is.Resources...)
		}
	}
	if err != nil {
		env.Config.Log.Critical(ctx, "%s", err)
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
