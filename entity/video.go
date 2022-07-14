package entity

type Video struct {
	Id        int64  `gorm:"primarykey;type:integer(7);AUTO_INCREMENT"`
	Titulo    string `gorm:"type:varchar(100); not null"`
	Descricao string `gorm:"type:text; not null"`
	URL       string `gorm:"type:varchar(255); not null"`
}
