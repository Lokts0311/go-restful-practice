package model

import (
	"time"
)

type HttpSuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Command struct {
	Uuid     string    `json:"uuid"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Url      string    `json:"url"`
}

type Request struct {
	Uuid     string    `json:"uuid"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Name     string    `json:"name"`
}
