package main

import (
	"fmt"
	"html/template"// After adding this package(html/template), we won't be using fmt anymore.
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

// save ...
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //filenameというp.Bodyという内容のfileを作る。最後のやつがわかんない。
}

// loadPage loads pages
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// render rendering
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page)  {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// viewHandler HandleFunc needs this Handler
func viewHandler(w http.ResponseWriter, r *http.Request)  {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)// if there does not exist page:title, make a new page.
		return		
	}
	renderTemplate(w, "view", p)
	// t, _ := template.ParseFiles("view.html")// return *template.Template, error
	// t.Execute(w,p)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// editHandler 
func editHandler(w http.ResponseWriter, r *http.Request)  {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
	// t, _ := template.ParseFiles("edit.html")// return *template.Template, error
	// t.Execute(w,p)
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+// postするやつ？
	// 	"<textarea name=\"body\">%s</textarea><br>"+// writing area to edit
	// 	"<input type=\"submit\" value=\"Save\">"+// make a button to submit
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)
}

// saveHandler ...
func saveHandler(w http.ResponseWriter, r *http.Request)  {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body") //"body"にdataが格納されていて、FormValue で読み出している、と思う
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}


func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// http://localhost:8080/view/test
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
}
