package main

import (
  "fmt"
  "strings"
  "os"
  "math/rand"
  "time"
)

func main() {
  var resposta string
  var jogar bool
  var cartas [13]Carta

  cartas = gerarCartas()
  fmt.Printf("Bem vindo ao BlackJack \nGostaria de iniciar uma partida? \n")
  fmt.Scanln(&resposta)
  jogar = simOuNao(resposta)

  if jogar {
    jogo(cartas)
  } else {
    fmt.Printf("Até a próxima!\n")
    os.Exit(0)
  }
}

type Carta struct {
  Nome string
  Pontos int32
}

func jogo(cartas [13]Carta){
  var resposta string

  carta := entregarCarta(cartas)
  pontos := carta.Pontos
  fmt.Printf("A carta sorteada é %s\nSua pontuação atual é %d\n", carta.Nome, pontos)

  for pontos <= 21 {
    fmt.Printf("Você gostaria de mais uma carta? \n")
    fmt.Scanln(&resposta)
    respVerificada := simOuNao(resposta)

    if respVerificada{
      carta = entregarCarta(cartas)
      pontos = pontos + carta.Pontos
      if pontos > 21{
        fmt.Printf("A carta sorteada é %s sua pontuação foi %d\n      Game Over\n", carta.Nome, pontos)
        os.Exit(0)
      } else if pontos == 21{
        fmt.Printf("A carta sorteada é %s sua pontuação foi %d\nParabéns!\n      Game Over\n", carta.Nome, pontos)
        os.Exit(0)
      }
      fmt.Printf("A carta sorteada é %s\nSua pontuação atual é %d\n", carta.Nome, pontos)
    } else{
      fmt.Printf("Sua pontuação foi %d\nObrigado por jogar!\n", pontos)
      os.Exit(0)
    }
  }
}

func gerarCartas ()(cartas [13]Carta){
  cartas[0].Nome = "Dois"
  cartas[0].Pontos = 2

  cartas[1].Nome = "Três"
  cartas[1].Pontos = 3

  cartas[2].Nome = "Quatro"
  cartas[2].Pontos = 4

  cartas[3].Nome = "Cinco"
  cartas[3].Pontos = 5

  cartas[4].Nome = "Seis"
  cartas[4].Pontos = 6

  cartas[5].Nome = "Sete"
  cartas[5].Pontos = 7

  cartas[6].Nome = "Oito"
  cartas[6].Pontos = 8

  cartas[7].Nome = "Nove"
  cartas[7].Pontos = 9

  cartas[8].Nome = "Dez"
  cartas[8].Pontos = 10

  cartas[9].Nome = "Valete"
  cartas[9].Pontos = 10

  cartas[10].Nome = "Dama"
  cartas[10].Pontos = 10

  cartas[11].Nome = "Rei"
  cartas[11].Pontos = 10

  cartas[12].Nome = "As"
  cartas[12].Pontos = 1

  return
}

func simOuNao(resp string)(sim bool){
  strings.ToLower(resp)
  if resp == "sim" || resp == "s" {
    sim = true
  } else if resp == "não" || resp == "n" || resp == "nao"{
    sim = false
  } else {
    os.Exit(0)
  }
  return
}

func gerarRandom()(random int32){
  tempo := time.Now()
  nanosecond := tempo.Nanosecond()
  rand.Seed(int64(nanosecond))
  random = rand.Int31n(13)

  return
}

func entregarCarta (cartas [13]Carta)(carta Carta){
  num := gerarRandom()
  carta = cartas[num]
  return
}
