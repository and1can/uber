package main

type Uber struct {
    Id int              `json:"id"`
    Start int           `json:"start"`
    Distance int        `json:"distance"`
    Reqid string        `json:"reqid"`
    TotalCharged int    `json:"totalcharged"`
}