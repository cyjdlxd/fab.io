package fab

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/kooinam/fabio/actors"
	"github.com/markbates/pkger"
)

func serveStats() {
	f, _ := pkger.Open("github.com/kooinam/fabio:/stats/index.html")
	b, _ := ioutil.ReadAll(f)
	content := string(b)

	tmpl := template.Must(template.New("stats").Parse(content))

	http.HandleFunc("/stats/index.html", func(w http.ResponseWriter, r *http.Request) {
		data := &struct {
			ActorInfos []*actors.ActorInfo
		}{
			ActorInfos: ActorManager().GetActorInfos(),
		}

		tmpl.Execute(w, data)
	})
}