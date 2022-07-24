package domain

type Categoria struct {
	Id     uint   `json:"id"`
	Titulo string `json:"titulo"`
	Cor    string `json:"cor"`
}

func (c *Categoria) IsPatchValid() bool {
	if c.Id != 0 {
		if c.Titulo == "" && c.Cor == "" {
			return false
		}
		return true
	}
	return false
}

func (c *Categoria) IsValid() bool {
	if c.Id != 0 || c.Titulo == "" || c.Cor == "" {
		return false
	}
	return true
}
