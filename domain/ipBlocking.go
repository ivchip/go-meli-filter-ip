package domain

import (
	"net/http"
	"time"
)

type IpBlocking struct {
	Ip        string    `json:"ip"gorm:"type:varchar(15); primarykey"`
	Active    bool      `json:"active"gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IpsBlocking []IpBlocking

type IpBlockingUseCases interface {
	Migrate() error
	GetAll() (IpsBlocking, error)
	GetByIp(string) (IpBlocking, error)
	Create(*IpBlocking) error
	Update(string, *IpBlocking) error
}

type IpBlockingAdapter interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByIp(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
