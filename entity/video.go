package entity

type Video struct {
	Id          uint   `gorm:"primarykey;auto_increment"`
	Titulo      string `gorm:"not null"`
	Descricao   string `gorm:"not null"`
	URL         string `gorm:"not null"`
	CategoriaId uint
	Categoria   Categoria `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
