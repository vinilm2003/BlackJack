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
  Naipe string
  Pontos int32
}

func jogo(cartas [13]Carta){
  var resposta string

  naipePC, cartaPC := cartasDealer(cartas)
  pontosPC := cartaPC.Pontos
  fmt.Printf("A carta do dealer é %s %s\n E sua pontuação atual é %d\n", cartaPC.Nome, naipePC.Naipe, pontosPC)

  naipe, carta := cartasJogador(cartas)
  pontos := carta.Pontos
  fmt.Printf("\nA sua carta sorteada é %s %s\nSua pontuação atual é %d\n", carta.Nome, naipe.Naipe, pontos)

  for pontos <= 21 {
    fmt.Printf("Você gostaria de mais uma carta? \n")
    fmt.Scanln(&resposta)
    respVerificada := simOuNao(resposta)

    if respVerificada{
      naipePC, cartaPC = cartasDealer(cartas)
      pontosPC = pontosPC + cartaPC.Pontos

      naipe, carta = cartasJogador(cartas)
      pontos = pontos + carta.Pontos

      verificacaoJogador(naipe, carta, naipePC, cartaPC, pontos, pontosPC)
    } else{
      fmt.Printf("O dealer terminou com %d pontos\n", pontosPC)
      fmt.Printf("Sua pontuação foi %d\nObrigado por jogar!\n", pontos)
      os.Exit(0)
    }
  }
}

func gerarCartas ()(cartas [13]Carta){
  cartas[0].Naipe = "de Paus"
  cartas[0].Nome = "Dois"
  cartas[0].Pontos = 2

  cartas[1].Naipe = "de Ouros"
  cartas[1].Nome = "Três"
  cartas[1].Pontos = 3

  cartas[2].Naipe = "de Copas"
  cartas[2].Nome = "Quatro"
  cartas[2].Pontos = 4

  cartas[3].Naipe = "de Espadas"
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

func verificacaoJogador(naipe Carta, carta Carta, naipePC Carta, cartaPC Carta, pontos int32, pontosPC int32){
  if pontos > 21 && pontosPC < 21{
    fmt.Printf("A carta sorteada é %s %s sua pontuação foi %d\n      Game Over\n", carta.Nome, naipe.Naipe, pontos)
    fmt.Printf("A carta do dealer é %s %s, ele venceu com %d pontos!\n      Game Over\n", cartaPC.Nome, naipePC.Naipe, pontosPC)
    os.Exit(0)
  } else if pontos == 21{
    fmt.Printf("A carta sorteada é %s %s sua pontuação foi %d\n", carta.Nome, naipe.Naipe, pontos)
    fmt.Printf("A carta do dealer é %s %s, ele perdeu com %d pontos\nParabéns!\n      Game Over\n", cartaPC.Nome, naipePC.Naipe, pontosPC)
    os.Exit(0)
  } else {
  fmt.Printf("A carta do dealer é %s %s\n E sua pontuação atual é %d\n", cartaPC.Nome, naipePC.Naipe, pontosPC)
  fmt.Printf("\nA sua carta sorteada é %s %s\nSua pontuação atual é %d\n", carta.Nome, naipe.Naipe, pontos)
  }
  return
}

func cartasDealer (cartas [13]Carta) (naipe Carta, carta Carta){
  naipe = entregarNaipe(cartas)
  carta = entregarCarta(cartas)
  return
}

func cartasJogador (cartas [13]Carta) (naipe Carta, carta Carta){
  naipe = entregarNaipe(cartas)
  carta = entregarCarta(cartas)
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

func gerarCartaRandom()(random int32){
  return gerarRandom(13)
}

func entregarCarta (cartas [13]Carta)(carta Carta){
  num := gerarCartaRandom()
  carta = cartas[num]
  return
}

func gerarNaipeRandom()(random int32){
  return gerarRandom(4)
}

func gerarRandom(limite int32)(random int32){
  tempo := time.Now()
  nanosecond := tempo.Nanosecond()
  rand.Seed(int64(nanosecond))
  random = rand.Int31n(limite)

  return
}

func entregarNaipe (cartas [13]Carta)(naipeCarta Carta){
  num := gerarNaipeRandom()
  naipeCarta = cartas[num]
  return
}
