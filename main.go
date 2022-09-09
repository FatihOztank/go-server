package main // klasik go kodu böyle başlıyor

import(
	"fmt"
	"log"
	"net/http" // most of the functionality comes from here
)

func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"parseform error %v", err)
		return
	}

	fmt.Fprintf(w,"post request succeded")
	name := r.FormValue("name")

	fmt.Fprintf(w, "name = %s\n", name)


	// if r.URL.Path != "/hello" {
	// 	http.Error(w,"404 not found", http.StatusNotFound)
	// 	return
	// }

	// if r.Method != "GET" {
	// 	http.Error(w, "method is not supported", http.StatusNotFound)
	// 	return
	// }
}


func helloHandler(w http.ResponseWriter, r *http.Request){ // * is pointer here too
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

