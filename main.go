package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:lenslocked@gmail.com\">@gmail.com</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<details><summary>Is there a free version?</summary>Yesm there is.</details>
	<details><summary>What are your support hours?</summary>We have support staff answering emails 24/7.</details>
	<details><summary>How do I contact support?</summary>Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>.</details>`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// switch r.URL.Path {
// case "/":
// 	homeHandler(w, r)
// case "/contact":
// 	contactHandler(w, r)
// default:
// 	//TODO handle the page not found error
// 	http.Error(w, "Page not found", http.StatusNotFound)
// }
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
