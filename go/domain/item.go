package domain

import "time"

type Item struct {
	Id       int64     `ksql:"id" json:"id"`
	Col_Text string    `ksql:"col_texto" json:"col_texto"`
	Col_dt   time.Time `ksql:"col_dt" json:"col_dt"`
}
