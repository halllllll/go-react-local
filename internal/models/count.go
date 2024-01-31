package models

import "time"

// entity
type CountId int
type CountValue int
type CountCreated time.Time
type CountUpdated time.Time

type Count struct {
	Id      CountId      `json:"id" db:"count_id"`
	Val     CountValue   `json:"value" db:"count_value"`
	Created CountCreated `json:"created" db:"created_at"`
	Updated CountUpdated `json:"updated" db:"updated_at"`
}

type Counts *[]Count
