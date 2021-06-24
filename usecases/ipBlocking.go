package usecases

import (
	"fmt"
	"github.com/ivchip/go-meli-filter-ip/domain"
	"github.com/ivchip/go-meli-filter-ip/repository/postgres"
)

var repo domain.IpBlockingUseCases

func New() {
	repo = postgres.DAOIpBlocking()
}

func Migrate() error {
	return repo.Migrate()
}

func GetAll() (domain.IpsBlocking, error) {
	return repo.GetAll()
}

func GetByIp(ip string) (domain.IpBlocking, error) {
	if ip == "" {
		fmt.Errorf("IpBlocking ip is can not empty")
	}
	return repo.GetByIp(ip)
}

func Create(ipb *domain.IpBlocking) error {
	return repo.Create(ipb)
}

func Update(ip string, ipb *domain.IpBlocking) error {
	if ip == "" {
		fmt.Errorf("IpBlocking ip is can not empty")
	}
	return repo.Update(ip, ipb)
}
