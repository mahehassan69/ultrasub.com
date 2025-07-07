package main

import (
    "html/template"
    "net/http"
    "ultrafinder/internal/sources"
    "ultrafinder/internal/resolver"
)

type Result struct {
    Domain     string
    Subdomains []string
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        domain := r.FormValue("domain")
        subs := sources.Collect(domain)
        resolved := resolver.ResolveSubdomains(subs)

        tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
        tmpl.Execute(w, Result{Domain: domain, Subdomains: resolved})
    } else {
        tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
        tmpl.Execute(w, nil)
    }
}

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
    http.HandleFunc("/", handler)

    println("🌍 Visit: http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
