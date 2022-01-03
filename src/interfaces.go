package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const templ = `
<h1>{{. | len}} Items</h1>
<table>
<tr style='text-align: left'>
<th>Item</th>
<th>Price</th>
</tr>
{{range $key, $value := .}}
<tr>
 <td>{{$key}}</td>
 <td>{{$value}}</td>
</tr>
{{end}}
</table>`

var myTemplate = template.Must(template.New("boo").Parse(templ))

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	//for item, price := range db {
	//	fmt.Fprintf(w, "%s: %s\n", item, price)
	//}
	myTemplate.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad price value %q\n", price)
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	db[item] = dollars(price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
