package demographics

import (
	"time"

	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
)

type Country struct {
	ID        string    `gorm:"type:text;not null;" json:"id"`
	Name      string    `gorm:"type:varchar(150);not null;" json:"name"`
	Continent string    `gorm:"type:varchar(150);not null;" json:"continent"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

type Region struct {
	ID        string    `gorm:"type:text;not null;" json:"id"`
	Name      string    `gorm:"type:varchar(150);not null;" json:"name"`
	CountryId string    `gorm:"not null;" json:"countryId"`
	Country   Country   `json:"country"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

type District struct {
	ID        string    `gorm:"type:text;not null;" json:"id"`
	Name      string    `gorm:"type:varchar(150);not null;" json:"name"`
	RegionId  string    `gorm:"not null;" json:"regionId"`
	Region    Region    `json:"region"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (country *Country) BeforeCreate(_ *gorm.DB) error {
	country.ID = ulids.GenerateUUID().String()
	return nil
}

func (country *Country) BeforeUpdate(_ *gorm.DB) error {
	country.UpdatedAt = time.Now()
	return nil
}

func (country *Country) Prepare() {
	country.CreatedAt = time.Now()
	country.UpdatedAt = time.Now()
}

func (region *Region) BeforeCreate(_ *gorm.DB) error {
	region.ID = ulids.GenerateUUID().String()
	return nil
}

func (region *Region) BeforeUpdate(_ *gorm.DB) error {
	region.UpdatedAt = time.Now()
	return nil
}

func (region *Region) Prepare() {
	region.CreatedAt = time.Now()
	region.UpdatedAt = time.Now()
}

func (district *District) BeforeCreate(_ *gorm.DB) error {
	district.ID = ulids.GenerateUUID().String()
	return nil
}

func (district *District) BeforeUpdate(_ *gorm.DB) error {
	district.UpdatedAt = time.Now()
	return nil
}

func (district *District) Prepare() {
	district.CreatedAt = time.Now()
	district.UpdatedAt = time.Now()
}
