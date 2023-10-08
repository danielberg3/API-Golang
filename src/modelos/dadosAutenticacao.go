package modelos

//DadosAutenticacao são os dados obtidos ao fazer login
type DadosAutenticacao struct {
	ID    string `json:"ID"`
	Token string `json:"Token"`
}
