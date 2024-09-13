package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// ResultData holds the data to be displayed in the template
type ResultData struct {
	Result string
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	var data ResultData

	if r.Method == http.MethodPost {
		r.ParseForm()

		characterCopiesStr := r.FormValue("characterCopies")
		lightconeCopiesStr := r.FormValue("lightconeCopies")
		currentWarpsStr := r.FormValue("currentWarps")

		// Log the inputs for debugging
		log.Printf("characterCopiesStr: %s, lightconeCopiesStr: %s, currentWarpsStr: %s", characterCopiesStr, lightconeCopiesStr, currentWarpsStr)

		characterCopies, err := strconv.Atoi(characterCopiesStr)
		if err != nil {
			http.Error(w, "Invalid input for character copies", http.StatusBadRequest)
			return
		}

		lightconeCopies, err := strconv.Atoi(lightconeCopiesStr)
		if err != nil {
			http.Error(w, "Invalid input for lightcone copies", http.StatusBadRequest)
			return
		}

		currentWarps, err := strconv.Atoi(currentWarpsStr)
		if err != nil {
			http.Error(w, "Invalid input for current warps", http.StatusBadRequest)
			return
		}

		// Calculate guaranteed characters and lightcones
		guarantee := calculateGuaranteed(characterCopies, lightconeCopies, currentWarps)

		// Log the calculated guarantee for debugging
		log.Printf("Calculated guarantee: %d", guarantee)

		if characterCopiesStr != "" && lightconeCopiesStr != "" {
			if guarantee < (lightconeCopies + characterCopies) {
				data.Result = "You are not likely to get your desired copies."
			} else {
				data.Result = "You are likely to get your desired copies."
			}
		} else {
			data.Result = ""
		}

		// Log the result before rendering the template
		log.Printf("Result: %s", data.Result)
	}

	// Handle template execution error
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Unable to render the page", http.StatusInternalServerError)
	}
}

func calculateGuaranteed(characterCopies int, lightconeCopies int, warps int) int {
	charCount := 0
	lcCount := 0
	for i := 0; i < characterCopies; i++ {
		if warps > 160 {
			charCount += 1
			warps -= 160
		}
	}
	if warps > 0 {
		for j := 0; j < lightconeCopies; j++ {
			if warps > 140 {
				lcCount += 1
				warps -= 140
			}
		}
	}

	return (charCount + lcCount)
}
