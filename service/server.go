package service

import (
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/{op}", testHandler(formatter)).Methods("GET")
    mx.HandleFunc("/{op}/{id}", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
		op := vars["op"]
        id := vars["id"]
		switch {
			case op == "hello":
				formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello, " + id})
			case op == "bye":
				formatter.JSON(w, http.StatusOK, struct{ Test string }{"Bye, " + id})
			default:
				formatter.JSON(w, http.StatusOK, struct{ Test string }{"Wrong url"})
		}
    }
}
