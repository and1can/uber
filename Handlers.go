package main

import (
    "encoding/json"
    "net/http"
    "strings"
    "log"
    _ "github.com/lib/pq"
    "database/sql"
    
)


func IntervalHistory(w http.ResponseWriter, r *http.Request) {
    
    u := r.URL.String()
    usplit := strings.Split(u, "/")

    db2, err := sql.Open("postgres","user=and1can dbname=uber sslmode=disable")
   
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



