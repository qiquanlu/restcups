package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

type Printer struct {
	Name string `json:"name"`
}

type PrintResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"Message": e.Message,
	})
}

func main() {

	port := flag.Int("port", 8080, "The port to listen on")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	// GET route to list all printers
	r.Get("/printers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cmd := exec.Command("lpstat", "-p")
		output, err := cmd.Output()

		if err != nil {
			errorResponse := &ErrorResponse{Message: "Error listing printers"}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		printers := strings.Split(string(output), "\n")

		var printerList []Printer
		for _, p := range printers {
			outputStrings := strings.Split(string(p), " ")

			if len(outputStrings) >= 2 {
				printerList = append(printerList, Printer{Name: outputStrings[1]})
			}
		}

		json.NewEncoder(w).Encode(printerList)
	})

	// POST route to print to a printer
	r.Post("/print", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		printer := r.FormValue("printer")
		if printer == "" {

			errorResponse := &ErrorResponse{Message: "Printer not specified"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse)

			return
		}

		copies, err := strconv.Atoi(r.FormValue("copies"))
		if err != nil {
			copies = 1
		}
		// Parse the "raw" form value as a boolean
		raw, err := strconv.ParseBool(r.FormValue("raw"))
		if err != nil {
			raw = false
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// Print the binary data to the specified printer
		rawCmd := ""
		if raw {
			rawCmd = "raw"
		}
		cmd := exec.Command("lp", "-d", printer, "-n", strconv.Itoa(copies), "-o", rawCmd)
		cmd.Stdin = bytes.NewReader(data)

		if err := cmd.Run(); err != nil {
			http.Error(w, "Error printing file", http.StatusInternalServerError)
			return
		}

		response := PrintResponse{Message: "Print job submitted successfully"}
		json.NewEncoder(w).Encode(response)
	})
	addr := fmt.Sprintf(":%d", *port)

	log.Fatal(http.ListenAndServe(addr, r))
}
