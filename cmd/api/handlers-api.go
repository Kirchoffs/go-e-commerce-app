package main

import (
    "encoding/json"
    "net/http"
)

type stripePayload struct {
    Currency string `json:"currency"`
    Amount   string `json:"amount"`
}

type jsonResponse struct {
    OK      bool   `json:"ok"`
    Message string `json:"message"`
    Content string `json:"content"`
    ID      int    `json:"id"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
    res := jsonResponse{
        OK: true,
    }

    out, err := json.MarshalIndent(res, "", "  ")
    if err != nil {
        app.errorLog.Println(err)
    }

    w.Header().Set("Content-Type", "json/application")
    w.Write(out)
}
