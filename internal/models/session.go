package models

import "time"

type Session struct {
	Id     string    `json:"id"`
	Issued time.Time `json:"issued"`
	Expiry time.Time `json:"expiry"`
}
