package modelos

import (
	"api/src/seguranca"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (Usuario *Usuario) Preparar(etapa string) error {

	if erro := Usuario.Validar(etapa); erro != nil {
		return erro
	}

	if erro := Usuario.Formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) Validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("Nome não informado")
	}

	if usuario.Nick == "" {
		return errors.New("Nick não informado")
	}

	if usuario.Email == "" {
		return errors.New("Email não informado")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("Formato de email inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("Senha não informada")
	}

	return nil
}
func (usuario *Usuario) Formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
