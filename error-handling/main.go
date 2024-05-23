package main

import (
	"fmt"

	. "mpgxc/notification"
	. "mpgxc/result"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func validatePessoa(pessoa Pessoa) *Notification {
	notification := NewNotification()

	if len(pessoa.Nome) < 3 {
		notification.AddError("NomeInvalido", "Nome precisa ter pelo menos 3 caracteres", struct{}{})
	}

	if pessoa.Idade < 18 {
		notification.AddError("IdadeInvalida", "Idade precisa ser maior ou igual a 18", struct{}{})
	}

	return notification
}

func main() {
	pessoa := Pessoa{Nome: "João", Idade: 16}

	notification := NewNotification()

	notification.AddNotification(validatePessoa(pessoa))

	result := Result[int](divide(10, 0))

	if !result.Ok() {
		notification.AddError("DivisaoInvalida", result.Err().Error(), nil)
	}

	if notification.HasErrors() {

		fmt.Println("Erros de validação:")

		for _, error := range notification.GetErrors() {
			fmt.Println("Nome:", error.Name)
			fmt.Println("Mensagem:", error.Message)
			fmt.Println("Timestamp:", error.Timestamp)
			fmt.Println("Contexto:", error.Context)
		}
	}
}

func divide(numerator, denominator int) Result[int] {
	if denominator == 0 {
		return Err[int](fmt.Errorf("division by zero"))
	}

	return Ok(numerator / denominator)
}
