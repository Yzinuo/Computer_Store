package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Computer struct {
	ID          int    `gorm:"primaryKey;autoIncrement;column:id;comment:商品唯一标识" json:"id"`
	Name        string  `gorm:"not null;column:name;comment:商品名称" json:"name"`
	Image       string  `gorm:"not null;column:image;comment:商品图片" json:"imageUrl"`
	Price       float64 `gorm:"not null;column:price;comment:商品价格" json:"price"`
	Description string  `gorm:"not null;column:description;comment:电脑介绍" json:"description"`
	Brand       string  `gorm:"not null;column:brand;comment:品牌"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:created_at;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:更新时间"`
}

type ProductDetail struct {
    ID           int      `gorm:"primaryKey;autoIncrement;column:id;comment:商品详情唯一标识"`
    ProductID    int      `gorm:"unique;not null;column:product_id;comment:商品ID"`
    Processor    string    `gorm:"not null;column:processor;comment:处理器信息" json:"processor"`
    Stock        int       `gorm:"not null;column:stock;comment:库存数量"`
    Memory       string    `gorm:"not null;column:memory;comment:内存信息" json:"memory"`
    HardDisk     string    `gorm:"not null;column:hard_disk;comment:硬盘信息" json:"hard_disk"`
    GraphicsCard string    `gorm:"not null;column:graphics_card;comment:显卡信息" json:"graphics_card"`
    CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:created_at;comment:创建时间"`
    UpdatedAt    time.Time `gorm:"column:updated_at;comment:更新时间"`

    Product Computer `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}

type ProductImage struct {
    ID          int      `gorm:"primaryKey;autoIncrement;column:id;comment:图片唯一标识"`
    ProductID   int      `gorm:"not null;column:product_id;comment:商品ID"`
    ImageURL    string    `gorm:"not null;column:image_url;comment:图片URL" json:"image_url"`
    Description string    `gorm:"column:description;comment:图片描述"`
    CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:created_at;comment:创建时间"`
    UpdatedAt   time.Time `gorm:"column:updated_at;comment:更新时间"`

    Product Computer `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}

func GetComputerList(db *gorm.DB, page, size int,keyword string) ([]Computer, error) {
    var computers []Computer
	if keyword != ""{
		db = db.Where("brand LIKE ?", keyword+"%")
	}

	if page != 0 {
		db = db.Scopes(Paginate(page, size))
	}

    err := db.Find(&computers).Error
    return computers, err
}

func GetComputerDetail(db *gorm.DB, id int) (*ProductDetail, error) {
    var productDetail ProductDetail
	err := db.Where("product_id = ?", id).First(&productDetail).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("product detail for computer with id %d not found", id)
		}
		return nil, err
	}

	return &productDetail, nil
}

func GetComputerImage(db *gorm.DB, id int) ([]ProductImage, error) {
    var Images []ProductImage
	err := db.Where("product_id = ?", id).First(&Images).Error
	if err != nil {
		return nil, err
	}

	return Images, nil
}

func GetComputerInfo(db *gorm.DB,id int)(*Computer,error){
	var computer Computer
	err := db.Where("id =?", id).First(&computer).Error
	if err!= nil {
		return nil, err
	}
	return &computer, nil
}