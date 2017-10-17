package main

import (
	"github.com/go-calculator-api/pkg/calculator"
	"html/template"
	"net/http"
)

var T *template.Template

func init() {
	t, err := template.ParseFiles("./../views/index.html", "./../views/calculator.html")
	if err != nil {
		panic("Template parsing error")
	}
	T = t
}

func Index(w http.ResponseWriter, r *http.Request) {
	// validate the method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// return with result if expression was set
	T.Execute(w, "index.html")
}

func Calculate(w http.ResponseWriter, r *http.Request) {

	// validate the method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// parse form data
	r.ParseForm()

	// prepare the data
	data := struct {
		Expr   string
		Result string
	}{
		Expr:   r.FormValue("expr"),
		Result: "",
	}

	// calculate expression if set
	if data.Expr != "" {
		result := calculator.CalculateParsedExpression(calculator.ParseExpression(data.Expr))
		data.Result = result
	}

	// return with result if expression was set
	T.ExecuteTemplate(w, "calculator.html", data)
}
