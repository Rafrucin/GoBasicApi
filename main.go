package main

// run an test @ http://localhost:3000/?name=Raf

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
		names := r.URL.Query()["name"]
		var name string
		if len(names)==1{
			name = names[0]
		}
		w.Write([]byte("Hello "+name+"\n"))

		m := map[string]string{"name": name} 
		enc := json.NewEncoder(w)
		enc.Encode(m)
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}	
	
}