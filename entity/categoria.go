package entity

type Categoria struct {
	Id     uint   `gorm:"primarykey;auto_increment"`
	Titulo string `gorm:"not null"`
	Cor    string `gorm:"not null"`
}
