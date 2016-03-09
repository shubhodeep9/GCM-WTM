package main

import (
    "fmt"
    "net/http"
    "github.com/alexjlockwood/gcm"
    "log"
    "encoding/json"
)

func receiver(w http.ResponseWriter, r *http.Request) {
    var message string = ""
    if len(r.FormValue("message"))!=0 {
        message = r.FormValue("message")
    }
    data := map[string]interface{}{"message": message}
    regIDs := []string{"4", "8", "15", "16", "23", "42"}
    msg := gcm.NewMessage(data, regIDs...)

    // Create a Sender to send the message.
    sender := &gcm.Sender{ApiKey: "AIzaSyDZIaBYzRfULj5V8-QiwqN_z2HL0kPnnAQ"}

    // Send the message and receive the response after at most two retries.
    response, err := sender.Send(msg, 2)
    if err != nil {
        fmt.Println("Failed to send message:", err)
        return
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    } // send data to client side
}

func main() {
    http.HandleFunc("/", receiver) // set router
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}