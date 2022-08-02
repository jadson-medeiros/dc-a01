package domain

import "time"

type Item struct {
	Id       int64     `json:"id"`
	Col_Text string    `json:"col_texto"`
	Col_dt   time.Time `json:"col_dt"`
}
