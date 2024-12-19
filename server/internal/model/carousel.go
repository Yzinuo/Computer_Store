package model

import (
	"time"

	"gorm.io/gorm"
)

type Carousel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;column:id;comment:轮播图唯一标识"`
	ImageURL  string    `gorm:"not null;column:image_url;comment:图片URL"`
	LinkURL   string    `gorm:"column:link_url;comment:跳转链接"`
	IsActive  bool      `gorm:"not null;default:true;column:is_active;comment:是否启用（1：启用，0：禁用）"`
	Order     int       `gorm:"not null;default:0;column:order;comment:显示顺序"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:created_at;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:更新时间"`
}

type CarouselVO struct{
	ImageURL string `json:"picUrl"`
	LinkURL  string `json:"linkUrl"`
}

func GetCarousel(db *gorm.DB) ([]CarouselVO, error) {
	var carouselList []CarouselVO
	err := db.Model(&Carousel{}).Select("image_url", "link_url").Where("is_active = ?", true).Order("`order`").Find(&carouselList).Error
	if err != nil {
		return nil, err
	}
	return carouselList, nil
}

