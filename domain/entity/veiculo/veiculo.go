package veiculo

type Veiculo struct {
	Marca         string  `json:"marca"`
	Modelo        string  `json:"modelo"`
	Ano           string  `json:"ano"`
	Cor           string  `json:"cor"`
	Preco         float32 `json:"preco"`
	Flagvendido   bool    `json:"flagvendido"`
	Id            int     `json:"id"`
	Cpf           string  `json:"cpf"`
	Datavenda     string  `json:"datavenda"`
	Pagamento     bool    `json:"pagamento"`
	Pagamentodesc string  `json:"pagamentodesc"`
}
