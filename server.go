package main

import (
    "fmt"
    "net/http"
    "github.com/alexjlockwood/gcm"
    "log"
    "encoding/json"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

func receiver(w http.ResponseWriter, r *http.Request) {
    var message string = ""
    
    if len(r.FormValue("message"))!=0 {
        message = r.FormValue("message")
    }
    o := orm.NewOrm()
    var registrations []*Registrations
    _, err := o.QueryTable("registrations").All(&registrations)
    if err == orm.ErrNoRows {
        fmt.Fprintf(w, "No registrations")
        return
    }
    var regIDs []string
    for _, val := range registrations {
        regIDs = append(regIDs,val.Regid)
    }
    fmt.Println(regIDs)
    
    data := map[string]interface{}{"message": message}
    msg := gcm.NewMessage(data, regIDs...)

    // Create a Sender to send the message.
    sender := &gcm.Sender{ApiKey: "AIzaSyAmDb9Gv7rY8dWvEUbwyU0y3hQTz2eoatU"}

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

type Registrations struct{
    Id int
    Regid string
} 

func (a *Registrations) TableName() string {
    return "registrations"
}

func register(w http.ResponseWriter, r *http.Request){
    var regID string = ""
    if len(r.FormValue("regID"))!=0 {
        regID = r.FormValue("regID")
    }
    
    o := orm.NewOrm()
    reg := Registrations{Regid: regID}
    _, err := o.Insert(&reg)
    if err == nil {
        fmt.Fprintf(w,"Success")
    }
    
}

func main() {
    orm.RegisterDriver("sqlite", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "data.db")
    orm.RegisterModel(new(Registrations))
    http.HandleFunc("/send/", receiver)
    http.HandleFunc("/register/", register)
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}