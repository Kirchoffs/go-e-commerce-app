package main

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
    app.infoLog.Println("Hit the handler")
    if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
        app.errorLog.Println(err)
    }
}
