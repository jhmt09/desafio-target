# Índice 

* [Título e Imagem de capa](#Título-e-Imagem-de-capa)
* [Tecnologias utilizadas](#badges)
* [Índice](#índice)
* [Descrição do Projeto](#descrição-do-projeto)
* [Status do Projeto](#status-do-Projeto)
* [Conclusão](#conclusão)

# Desafio da Target Sistemas!
![Desenho de uma cabeça com várias engrenagens, simbolizando o pensamento de uma ideia](https://blogs.iadb.org/conocimiento-abierto/wp-content/uploads/sites/10/2017/03/idea-concept-banner.jpg)

Tecnologia Utilizadas

![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual_Studio_Code-0078D4?style=for-the-badge&logo=visual%20studio%20code&logoColor=white)


# Descrição do projeto.

    Desafio 1

Fibonacci

A primeira descrição será do desáfio 1, onde foi feito um programa para  onde, informado um número, ele calculou a sequência de Fibonacci e retornou uma mensagem avisando se o número informado pertence ou não a sequência.

Explicação do código:

A função fibonacci(n int) recebe um número inteiro n como entrada e retorna um slice contendo os n primeiros números da sequência de Fibonacci.

A função contains(slice []int, num int) recebe um slice de inteiros e um número inteiro como entrada, e retorna true se o número estiver presente no slice, e false caso contrário.
Na função main(), o programa solicita que o usuário informe um número inteiro.

O programa chama a função fibonacci(n) para calcular a sequência de Fibonacci com n elementos.

O programa chama a função contains(fib, n) para verificar se o número informado está presente na sequência de Fibonacci.

O programa exibe uma mensagem informando se o número pertence ou não à sequência de Fibonacci.

    Desafio 2

O desafio numero 2 consiste na criação de um programa que calcula e retorna o maior e menor valor de um faturamento de uma distribuidora fictícia. Ainda calcula a média mensal e o numeros de dias onde o faturamento foi maior do que a média mensal

Explicação do código:

O vetor faturamentoDiario guarda os valores de faturamento diário da distribuidora. O programa usa um laço for para calcular o menor e o maior valor de faturamento, somar o total de faturamento e contar o número de dias em que o faturamento diário foi superior à média mensal. Em seguida, os resultados são exibidos na tela usando a função Printf do pacote fmt.

Desafio 3

Partindo para o desafio de numero 3, criei um programa na qual calcula o percentual de representação que cada estado teve dentro do valor total mensal da distribuidora.

Explicação do código:

Este código em Go tem como objetivo calcular o percentual de representação de cada estado em relação ao faturamento mensal total de uma distribuidora. Ele faz uso de um map chamado faturamento que contém os valores de faturamento de cada estado.

Inicialmente, é declarada uma variável total do tipo float64, que será utilizada para armazenar o valor total de faturamento. Em seguida, é feita uma iteração sobre os valores do map faturamento usando o _ para ignorar a chave e utilizar somente o valor. Dentro do loop, é realizado o somatório dos valores na variável total.

Por fim, é feita uma segunda iteração sobre o map faturamento, desta vez utilizando a chave e o valor. Dentro do loop, é calculado o percentual de representação do valor em relação ao total, e o resultado é exibido na tela usando a função fmt.Printf. O resultado é uma lista contendo o percentual de representação de cada estado no faturamento total da distribuidora.


    Desafio 4

Por ultimo o desafio de numero 4, onde fiz um programa que inverte uma String.

Explicação do código:

Neste programa, a string a ser invertida é definida na variável original. Em seguida, a string é convertida em um slice de bytes na variável bytes, que será manipulada para inverter os caracteres.

O início e o fim do slice a ser invertido são definidos nas variáveis inicio e fim, respectivamente. A partir daí, é feita uma iteração para trocar os caracteres das extremidades do slice, avançando o início e retrocedendo o fim a cada iteração, até que todo o slice seja invertido.

Por fim, o slice de bytes resultante é convertido novamente em uma string na variável invertida e é impressa na tela.


# Status do projeto

    ✅ Projeto feito ✅


# Conclusão

Escreva uma breve conclusão de um projeto. O projeto foi feito somente por mim
Com base no projeto desenvolvido exclusivamente por mim, posso concluir que os objetivos foram alcançados com sucesso. Ao longo do processo, foi necessário o uso de habilidades técnicas, gerenciais e criativas, além de muita dedicação e disciplina para concluir as tarefas dentro do prazo estabelecido.

Durante o desenvolvimento do projeto, foram identificadas algumas dificuldades e desafios, mas minha habilidade e comprometimento foram fundamentais para superá-los de forma eficiente e satisfatória.

Além disso, o projeto contribuiu para o meu crescimento pessoal e profissional, consolidando meu conhecimento e experiência em determinado assunto.

Em resumo, o projeto foi um sucesso e atingiu os objetivos estabelecidos, contribuindo para o meu aprimoramento pessoal e profissional. Estou satisfeito com o resultado obtido e confiante de que o trabalho desenvolvido poderá contribuir positivamente para a organização ou para outras áreas da minha vida.

Desde já eu agradeço imensamente aos colaboradores da Target sistemas pela atenção!



