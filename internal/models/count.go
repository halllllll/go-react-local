package models

import "time"

// entity
type CountId int
type CountValue int
type CountCreated time.Time
type CountUpdated time.Time

type Count struct {
	Id      CountId      `json:"id"`
	Val     CountValue   `json:"value"`
	Created CountCreated `json:"created"`
	Updated CountUpdated `json:"updated"`
}

type Counts *[]Count
