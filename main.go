// Desenvolvimento de um site de vendas, conectado a um banco de dados.
// Utilização do PostgreSQL e das linguagens Go e HTML.

package main

import (
	"net/http"

	"github.com/rodrigorissettoterra/Site-em-Go-e-HTML-com-conexao-PostgreSQL/routes"
)

// Definindo a função principal
func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
