package cyoa

import (
	handler "cyoa/template"
	"io/ioutil"
	"net/http"
	"text/template"
)

func main() {

    contents, err := ioutil.ReadFile();
    if err!= nil{
        panic(err)
    }

    template.New("story").Parse()
    
    http.HandleFunc("intro",handler.IntroHandler)
    http.HandleFunc("new-york",)
    http.HandleFunc("debate",)
    http.HandleFunc("sean-kelly",)
    http.HandleFunc("mark-bates",)
    http.HandleFunc("denver",)
    http.HandleFunc("home",)
	http.ListenAndServe(":8080", nil)
}


