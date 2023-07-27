package controllers

import (
	services "atm/Services"
	validations "atm/Validations"
	"html/template"
	"net/http"
	"strconv"
)

type TemplateData struct {
	IsError      bool
	ErrorMessage string
	Amount       int
	Data         map[int]int
}

var view = template.Must(template.ParseGlob("Views/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	view.ExecuteTemplate(w, "Index", nil)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {

	amount := r.FormValue("amount")
	amountFormatted, err := strconv.Atoi(amount)

	var data = TemplateData{
		IsError: false,
	}

	if err != nil {
		data = TemplateData{
			IsError:      true,
			ErrorMessage: "Valor inválido para saque! Por favor, não utilize vírgulas nem pontos (exemplo: 120).",
		}

		view.ExecuteTemplate(w, "Withdrawal", data)
	}

	if validations.Validate(amountFormatted) {

		data = TemplateData{
			IsError:      true,
			ErrorMessage: "Valor inválido para saque! Por favor, digite um valor entre R$ 2,00 e R$ 1.000.000.000,00 de reais, exceto R$ 3,00, pois não é possível efetuar o saque nesse valor.",
		}

		view.ExecuteTemplate(w, "Withdrawal", data)
	}

	banknotes := []int{200, 100, 50, 20, 10, 5, 2}

	response, err := services.CalculateBills(amountFormatted, banknotes)

	data = TemplateData{
		Amount: amountFormatted,
		Data:   response,
	}

	if err != nil {
		data = TemplateData{
			IsError:      true,
			ErrorMessage: err.Error(),
		}
	}

	view.ExecuteTemplate(w, "Withdrawal", data)
}
