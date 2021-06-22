package main

import (
	"fmt"
	"github.com/demonyangyue/gopl/src/ch07/eval"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cal", calculate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calculate(w http.ResponseWriter , r * http.Request)  {
	raw_expr := r.FormValue("expr")
	if raw_expr == "" {
		http.Error(w, "no expression", http.StatusBadRequest)
		return
	}
	expr, err := eval.Parse(raw_expr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result:= expr.Eval(eval.Env{})
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "result is : %g", result)
}

