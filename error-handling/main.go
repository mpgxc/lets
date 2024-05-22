package main

import (
  "fmt"
  "error/notifications"
)

type Pessoa struct {
  Nome string
  Idade int
}

func validatePessoa(pessoa Pessoa) *notifications.Notification {
  notification := notifications.NewNotification()

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

  notification := validatePessoa(pessoa)

  if notification.HasErrors() {
    fmt.Println("Erros de validação:")
    for _, error := range notification.GetErrors() {
      fmt.Println("Nome:", error.Name)
      fmt.Println("Mensagem:", error.Message)
      fmt.Println("Timestamp:", error.Timestamp)
      fmt.Println("Contexto:", error.Context)
    }
  } else {
    fmt.Println("Pessoa validada com sucesso!")
  }
}
