package main

import "fmt"

func main() {
    // Definindo a string a ser invertida
    original := "Aqui vai ser invertido, igual no filme Tenet"

    // Convertendo a string em um slice de bytes para facilitar a manipulação
    bytes := []byte(original)

    // Definindo o início e o fim do slice a ser invertido
    inicio := 0
    fim := len(bytes) - 1

    // Invertendo os caracteres do slice
    for inicio < fim {
        // Troca os caracteres das extremidades
        bytes[inicio], bytes[fim] = bytes[fim], bytes[inicio]

        // Avança o início e retrocede o fim
        inicio++
        fim--
    }

    // Convertendo o slice de bytes resultante em uma string novamente
    invertida := string(bytes)

    // Imprimindo a string invertida
    fmt.Println(invertida)
}
