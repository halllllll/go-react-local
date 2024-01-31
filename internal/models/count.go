package models

import "time"

// entity
type CountId int
type CountValue int

type Count struct {
	Id      CountId    `json:"id" db:"count_id"`
	Val     CountValue `json:"value" db:"count_value"`
	Created time.Time  `json:"created" db:"created_at"`
	Updated time.Time  `json:"updated" db:"updated_at"`
}

type Counts *[]Count
