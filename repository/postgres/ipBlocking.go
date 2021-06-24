package postgres

import (
	"github.com/ivchip/go-meli-filter-ip/domain"
	"gorm.io/gorm"
)

type PsqlIpBlocking struct {
	db *gorm.DB
}

func newPsqlIpBlocking(db *gorm.DB) *PsqlIpBlocking {
	return &PsqlIpBlocking{db}
}

func (p *PsqlIpBlocking) Migrate() error {
	err := p.db.AutoMigrate(&domain.IpBlocking{})
	if err != nil {
		return err
	}
	return nil
}

func (p *PsqlIpBlocking) GetAll() (domain.IpsBlocking, error) {
	ipsBlocking := domain.IpsBlocking{}
	result := p.db.Find(&ipsBlocking)
	if result.Error != nil {
		return domain.IpsBlocking{}, result.Error
	}
	return ipsBlocking, nil
}

func (p *PsqlIpBlocking) GetByIp(ip string) (domain.IpBlocking, error) {
	ipBlocking := domain.IpBlocking{}
	result := p.db.Where("ip = ?", ip).First(&ipBlocking)
	if result.Error != nil {
		return domain.IpBlocking{}, result.Error
	}
	return ipBlocking, nil
}

func (p *PsqlIpBlocking) Create(blocking *domain.IpBlocking) error {
	result := p.db.Create(&blocking)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PsqlIpBlocking) Update(ip string, blocking *domain.IpBlocking) error {
	var idDB domain.IpBlocking
	result := p.db.Where("ip = ?", ip).First(&idDB)
	if result.Error != nil {
		return result.Error
	}
	idDB.Active = blocking.Active
	result = p.db.Save(&idDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
