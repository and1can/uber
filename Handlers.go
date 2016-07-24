package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strings"
    "log"
    _ "github.com/lib/pq"
    "database/sql"
    //"net/url"
  

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048567))
    if err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := RepoCreateTodo(todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

type Uber struct {
    Id int              `json:"id"`
    Start int           `json:"start"`
    Distance int        `json:"distance"`
    Reqid string        `json:"reqid"`
    TotalCharged int    `json:"totalcharged"`
}

type Cost struct {
    Distance int        `json:"distance"`
    Totalcharged int    `json:"totalcharged"`
}

func IntervalHistory(w http.ResponseWriter, r *http.Request) {
    
    u := r.URL.String()
    //uparse := url.Parse(u)
    usplit := strings.Split(u, "/")
    //fmt.Fprintln(w, "In IntervalHistory with interval: ", usplit[2], usplit[3])
    
    /*db, err := sql.Open("postgres","user=and1can dbname=uber sslmode=disable")
    //db, err := sql.Open("postgres", "user=and1can dbname=restful sslmode=disable") 
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM history where start between $1 and $2", usplit[2], usplit[3])
    if err != nil {
        log.Fatal(err)
    }

    for rows.Next() {
        cost := new(Cost)
        err := rows.Scan(&cost.Distance, &cost.Totalcharged)
        // without the scan, all values will be 0 in json :O
        if err != nil {
            log.Fatal(err)
        }
        uberquery := &Uber{
            Id: trip.Id,
            Start: trip.Start,
            Distance: trip.Distance,
            Reqid: trip.Reqid,
            TotalCharged: trip.TotalCharged}

        uberjson, _ := json.Marshal(uberquery)
        w.Write(uberjson)
    }

    defer rows.Close() */

    db2, err := sql.Open("postgres","user=and1can dbname=uber sslmode=disable")
    //db, err := sql.Open("postgres", "user=and1can dbname=restful sslmode=disable") 
    if err != nil {
        log.Fatal(err)
    }

    cost, err := db2.Query("SELECT SUM(distance) as distance, SUM(totalcharged) as totalcharged from history where start between $1 and $2", usplit[2], usplit[3])
    if err != nil {
        log.Fatal(err)
    }

    for cost.Next() {
        c := new(Cost)
        err := cost.Scan(&c.Distance, &c.Totalcharged)
        // without the scan, all values will be 0 in json :O
        if err != nil {
            log.Fatal(err)
        }
        uberquery := &Cost{
            Distance: c.Distance,
            Totalcharged: c.Totalcharged}
        uberjson, _ := json.Marshal(uberquery)
        w.Write(uberjson)
    }

    defer cost.Close()
}



