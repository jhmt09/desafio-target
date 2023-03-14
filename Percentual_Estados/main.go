package main

import "fmt"

func main() {
	// Valores de faturamento por estado
	faturamento := map[string]float64{
		"SP":     67836.43,
		"RJ":     36678.66,
		"MG":     29229.88,
		"ES":     27165.48,
		"Outros": 19849.53,
	}

	// Variável para armazenar o total de faturamento de cada estado
	var total float64

	// Calcular o total de faturamento
	for _, valor := range faturamento {
		total += valor
	}

	// Calcular e exibir o percentual de representação de cada estado
	for estado, valor := range faturamento {
		percentual := (valor / total) * 100
		fmt.Printf("%s: %.2f%%\n", estado, percentual)
	}
}
