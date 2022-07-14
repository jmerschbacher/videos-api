package domain

type Video struct {
	Id        int64  `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	URL       string `json:"url"`
}
