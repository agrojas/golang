package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func GetCuit(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    gender := vars["gender"]
    dni := vars["dni"]

    var cuit = Cuit{
        genderLabel: gender,
        dni:         dni,
    }
    // CompleteCuit(&cuit)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(cuit)
    // if err := json.NewEncoder(w).Encode(cuit); err != nil {
    //     panic(err)
    // }
}

func CuitVerificationCode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    gender := vars["gender"]
    dni := vars["dni"]
    fmt.Fprintln(w, "Todo show:", gender, dni)
}
