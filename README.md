# Eliminação de Gauss

Este repositório contém implementações de um algoritmo de eliminação gaussiana para resolver sistemas de equações lineares, utilizando várias linguagens de programação, como C, Go, Rust, e um script em Python para gerar gráficos dos testes realizados. O código pode ser utilizado para avaliar o desempenho de cada implementação e comparar a eficiência entre elas.

# Integrantes 
* Marina de Souza Brum

# Descrição

O algoritmo implementado é a Eliminação Gaussiana, que é um método comum para resolver sistemas de equações lineares. O código cria uma matriz A de tamanho N x N e um vetor B de tamanho N, gerando números aleatórios para os elementos dessas estruturas. O objetivo é encontrar o vetor solução X, que satisfaz o sistema de equações representado por:

      𝐴 × 𝑋 = 𝐵

O código de cada linguagem implementa a solução serialmente e mede o tempo de execução total, tanto em termos de tempo real quanto de tempo de CPU.

# C

Para rodar o programa em C, use o comando abaixo:
```bash
 & .\'gauss.exe' 
```

Exemplos:

* & .\'gauss.exe' 5 para rodar com N = 5 e uma seed aleatória.
* & .\'gauss.exe' 5 12345 para rodar com N = 5 e uma seed fixa.
* & .\'gauss.exe' para rodar com o valor padrão de N = 10 e uma seed aleatória.

# Go

Para rodar o programa em Go, use o comando abaixo:
```bash
go run gauss.go 
```

Exemplos:

* go run gauss.go 5 para rodar com N = 5 e uma seed aleatória.
* go run gauss.go 5 12345 para rodar com N = 5 e uma seed fixa.
* go run gauss.go para rodar com o valor padrão de N = 10 e uma seed aleatória.

# RUST

Para rodar o programa em Rust, use o comando abaixo:
```bash
cargo run 
```

Exemplos:

* cargo run 5 para rodar com N = 5 e uma seed aleatória.
* cargo run 5 12345 para rodar com N = 5 e uma seed fixa.
* cargo run para rodar com o valor padrão de N = 10 e uma seed aleatória.

# Python (Para os Gráficos)
O script em Python é utilizado para gerar gráficos de desempenho. Ele utiliza matplotlib.pyplot e pandas para visualizar o tempo de execução das diferentes implementações. Certifique-se de ter as bibliotecas necessárias instaladas!

```bash
python grafico.py
```

# Resultado Esperado
Cada execução do programa calculará o vetor X e imprimirá o tempo total de execução (em milissegundos) e o tempo de CPU utilizado. A saída também inclui os valores das matrizes A e B se N for pequeno o suficiente (menor que 10).

Além disso, ao rodar o script Python, você verá gráficos que comparam o tempo de execução entre as implementações em C, Go e Rust.
