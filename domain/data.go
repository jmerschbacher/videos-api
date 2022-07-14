package domain

type Data struct {
	Data any `json:"data"`
}

type Error struct {
	Method   string `json:"method"`
	Rota     string `json:"rota"`
	Codigo   int    `json:"codigo"`
	Mensagem string `json:"mensagem"`
}
