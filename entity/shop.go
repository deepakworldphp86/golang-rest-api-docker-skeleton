package entity

import "time"


type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Users struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Categories struct {
	ID            uint     `json:"id" gorm:"primary_key"`
	CategoryName  string   `json:"category_name"`
	Description   string   `json:"description"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

type Customers struct {
	ID            uint     `json:"id" gorm:"primary_key"`
	CustomerEmail string   `json:"customer_email" gorm:"unique"`
	CustomerName  string   `json:"customer_name"`
	ContactName   string   `json:"contact_name"`
    Address   	  string   `json:"address"`
	City          string   `json:"city"`
	PostalCode    string   `json:"postal_code"`
	Country       string   `json:"country"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

type Products struct {
	ID            uint     `json:"id" gorm:"primary_key"`
	ProductName  string    `json:"product_name" gorm:"unique"`
	CategoryId   string    `json:"category_id"`
    Unit   	      string   `json:"unit"`
	Price         string   `json:"price"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}


type Orders struct {
	ID           uint         `json:"id" gorm:"primary_key"`
	CustomerId   string       `json:"customer_id"`
	order_date   time.Time    `json:"order_date"`
	CreatedAt    time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}


type OrderDetails struct {
	ID           uint         `json:"id" gorm:"primary_key"`
	order_id     uint         `json:"order_id"`
	product_id   uint         `json:"product_id"`
	quantity     uint         `json:"quantity"`
	CreatedAt    time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}
