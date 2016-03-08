package main

import (
    "fmt"

    "github.com/alexjlockwood/gcm"
)

func main() {
    // Create the message to be sent.
    data := map[string]interface{}{"score": "5x1", "time": "15:10"}
    regIDs := []string{"4", "8", "15", "16", "23", "42"}
    msg := gcm.NewMessage(data, regIDs...)

    // Create a Sender to send the message.
    sender := &gcm.Sender{ApiKey: "sample_api_key"}

    // Send the message and receive the response after at most two retries.
    response, err := sender.Send(msg, 2)
    if err != nil {
        fmt.Println("Failed to send message:", err)
        return
    }

    /* ... */
}