package main // klasik go kodu böyle başlıyor

import (
	"fmt"
	"log"
	"net/http" // most of the functionality comes from here
	"os"
	"path/filepath"
	"regexp"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform error %v", err)
		return
	}

	fmt.Fprintf(w, "post request succeded")
	name := r.FormValue("name")

	fmt.Fprintf(w, "name = %s\n", name)

}

func helloHandler(w http.ResponseWriter, r *http.Request) { // * is pointer here too
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform error %v", err)
		return
	}

	fmt.Fprintf(w, "post request succeded\n")
	filename := r.FormValue("filename")
	if filename == ""{
		fmt.Fprintf(w, "No filename specified")
		return

	}
	username := os.Getenv("USER")
	root := fmt.Sprintf("/home/%s/Downloads", username) 
	
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() {
			return nil
		}
		reg, err2 := regexp.Compile(filename)

		if err2 != nil {

			return err2
		}

		if reg.MatchString(info.Name()) {

			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(w, "fileprocessing error %v\n", err)
		return
	}

	if len(files) == 0 {
		fmt.Fprintf(w, "No file found...\n")
	}
	for _, file := range files {
        fmt.Fprintf(w, "%s\n", file)
    }
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/filesystem", fileHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}