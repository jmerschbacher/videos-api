package domain

type Data struct {
	Data any `json:"data"`
}

type Error struct {
	Rota     string `json:"rota"`
	Codigo   int    `json:"codigo"`
	Mensagem string `json:"mensagem"`
}
