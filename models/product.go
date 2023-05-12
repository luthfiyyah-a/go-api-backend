package models

type Product struct {
    ID          int64   `gorm:"primaryKey" json:"id"`
    NamaProduct string  `gorm:"type:varchar(300)" json:"nama_product"`
    Deskripsi   string  `gorm:"type:text" json:"deskripsi"`
    Harga       float32 `gorm:"type:decimal(10,2)" json:"harga"`
    Rating      float32 `gorm:"type:decimal(2,1)" json:"rating"`
}
