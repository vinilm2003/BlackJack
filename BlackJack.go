package main

import (
  "fmt"
  "strings"
  "os"
  "math/rand"
  "time"
)

var cartasUtilizadas = []int32{-1}

func main() {
  var resposta string
  var jogar bool

  fmt.Printf("Bem vindo ao BlackJack \nGostaria de iniciar uma partida? \n")
  fmt.Scanln(&resposta)
  jogar = simOuNao(resposta)

  if jogar {
    jogo()
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

func jogo(){
  var resposta string

  cartaPC := gerarCarta()
  pontosPC := cartaPC.Pontos
  fmt.Printf("A carta do dealer é %s de %s\n E sua pontuação atual é %d\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)

  cartaJ := gerarCarta()
  pontosJ := cartaJ.Pontos
  fmt.Printf("\nA sua carta sorteada é %s de %s\nSua pontuação atual é %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)

  for pontosJ <= 21 {
    fmt.Printf("Você gostaria de mais uma carta? \n")
    fmt.Scanln(&resposta)
    respVerificada := simOuNao(resposta)

    if respVerificada{
      cartaPC = gerarCarta()
      pontosPC = pontosPC + cartaPC.Pontos

      cartaJ = gerarCarta()
      pontosJ = pontosJ + cartaJ.Pontos

      verificacaoJogador(cartaJ, cartaPC, pontosJ, pontosPC)
    } else{
      fmt.Printf("O dealer terminou com %d pontos\n", pontosPC)
      fmt.Printf("Sua pontuação foi %d\nObrigado por jogar!\n", pontosJ)
      os.Exit(0)
    }
  }
}

func verificacaoJogador(cartaJ Carta, cartaPC Carta, pontosJ int32, pontosPC int32){

  if pontosJ > 21 && pontosPC < 21{
    fmt.Printf("A carta sorteada é %s de %s sua pontuação foi %d\n      Game Over\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
    fmt.Printf("A carta do dealer é %s de %s, ele venceu com %d pontos!\n      Game Over\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
    os.Exit(0)
  } else if pontosJ == 21{
    fmt.Printf("A carta sorteada é %s de %s sua pontuação foi %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
    fmt.Printf("A carta do dealer é %s de %s, ele perdeu com %d pontos\nParabéns!\n      Game Over\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
    os.Exit(0)
  } else if pontosPC == 21{
    fmt.Printf("A carta sorteada é %s de %s sua pontuação foi %d\n      Game Over\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
    fmt.Printf("A carta do dealer é %s de %s, ele venceu com %d pontos!\n      Game Over\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
    os.Exit(0)
  } else if pontosPC > 21 && pontosJ < 21{
    fmt.Printf("A carta sorteada é %s de %s sua pontuação foi %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
    fmt.Printf("A carta do dealer é %s de %s, ele perdeu com %d pontos\nParabéns!\n      Game Over\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
    os.Exit(0)
  } else {
  fmt.Printf("A carta do dealer é %s de %s\n E sua pontuação atual é %d\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
  fmt.Printf("\nA sua carta sorteada é %s de %s\nSua pontuação atual é %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
  }
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

func gerarCarta()(carta Carta){
  var nome = [13]string{"Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez", "Valete", "Rainha", "Rei", "As"}
  var naipe = [4]string{"Ouros", "Copas", "Espadas", "Paus"}
  var pontos = [13]int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 1}
  var num int32
  utilizada := true

  for utilizada {
    num = gerarRandom()
    for _, v := range cartasUtilizadas {
      utilizada = v == num
      if utilizada {
        break
      }
    }
  }

  cartasUtilizadas = append(cartasUtilizadas, num)

  numCarta := num % 13
  numPontos := num % 13
  numNaipe := num % 4

  carta.Nome = nome[numCarta]
  carta.Naipe = naipe[numNaipe]
  carta.Pontos = pontos[numPontos]

return
}

func gerarRandom()(random int32){
  tempo := time.Now()
  nanosecond := tempo.Nanosecond()
  rand.Seed(int64(nanosecond))
  random = rand.Int31n(52)

  return
}
