package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	resp := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if value, exists := body["message"]; exists && reflect.TypeOf(value).Kind() == reflect.String {
		fmt.Println(value)
		resp["message"] = "Данные успешно приняты"
		resp["status"] = "success"
	} else {
		resp["message"] = "Некорректное JSON-сообщение"
		resp["status"] = "404"
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/", mainPage)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
