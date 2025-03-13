package domain

import "time"

type Shipping struct {
	ID uint `gorm:"primaryKey" json:"id"`
	// ต้องตรงกับ order
	OrderID uint `gorm:"unique;not null" json:"order_id"`
	// หมายเลขพัสดุ
	TrackingNumber string    `gorm:"unique" json:"tracking_number"`
	Carrier        string    `json:"carrier"`
	Status         string    `gorm:"default:PENDING" json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
