package domain

type Video struct {
	Id        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	URL       string `json:"url"`
}

func (v *Video) IsValidoEdicao() bool {
	if v.Id != 0 {
		if v.Titulo == "" && v.Descricao == "" && v.URL == "" {
			return false
		}
		return true
	}
	return false
}

func (v *Video) IsValido() bool {
	if v.Titulo == "" || v.Descricao == "" || v.URL == "" {
		return false
	}
	return true
}
