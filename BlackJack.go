package main

import (
  "fmt"
  "strings"
  "os"
  "math/rand"
  "time"
)

var cartasUtilizadas = []int32{-1}
var pontosJ, pontosPC int32
var cartaJ, cartaPC CartaType
var respVerificada, jogadorContinua, dealerContinua bool

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

type CartaType struct {
  Nome string
  Naipe string
  Pontos int32
}

func jogo(){
  var resposta string
  iniciante := gerarRandom(2)
  jogadorContinua = true
  dealerContinua = true

  cartaJogador(); cartaJogador()
  cartaDealer(); cartaDealer()

  for pontosJ <= 21 {
    if jogadorContinua {
      fmt.Printf("Você gostaria de mais uma carta? \n")
      fmt.Scanln(&resposta)
      respVerificada = simOuNao(resposta)
    }
    if iniciante == 0 {
      jogadorInicia()
    } else if iniciante == 1 {
      dealerInicia()
    }
    verificarPontuacao(cartaJ, cartaPC, pontosJ, pontosPC, jogadorContinua, dealerContinua)
  }
}

func verificarPontuacao(cartaJ CartaType, cartaPC CartaType, pontosJ int32, pontosPC int32, jogadorContinua bool, dealerContinua bool){

  if pontosJ > 21 && pontosPC < 21{
    fmt.Printf("\nA carta sorteada é %s de %s sua pontuação foi %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
    fmt.Printf("A carta do dealer é %s de %s, ele venceu com %d pontos!\n      Game Over\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
    os.Exit(0)
  } else if pontosJ == 21{
    fmt.Printf("\nA carta sorteada é %s de %s sua pontuação foi %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
    fmt.Printf("A carta do dealer é %s de %s, ele perdeu com %d pontos\nVocê ganhou, parabéns!\n  Game Over\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
    os.Exit(0)
  } else if pontosPC == 21{
    fmt.Printf("\nA carta sorteada é %s de %s sua pontuação foi %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
    fmt.Printf("A carta do dealer é %s de %s, ele venceu com %d pontos!\n      Game Over\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
    os.Exit(0)
  } else if pontosPC > 21 && pontosJ < 21{
    fmt.Printf("\nA carta sorteada é %s de %s sua pontuação foi %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
    fmt.Printf("A carta do dealer é %s de %s, ele perdeu com %d pontos\nVocê ganhou, parabéns!\n  Game Over\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
    os.Exit(0)
  } else if !dealerContinua && !jogadorContinua {
      if pontosJ > pontosPC {
        fmt.Printf("\nVocê venceu com %d Pontos!\nParabéns!\n      Game Over\n", pontosJ)
        os.Exit(0)
      } else if pontosJ < pontosPC {
        fmt.Printf("\nO dealer venceu com %d Pontos!\n      Game Over\n", pontosPC)
        os.Exit(0)
      } else {
        fmt.Printf("\nHouve um empate!\n Dealer:%d\n Você:%d\n      Game Over\n", pontosPC, pontosJ)
        os.Exit(0)
      }
  }
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

func jogadorInicia (){
  if respVerificada && jogadorContinua{
    cartaJogador()

  } else if !respVerificada && jogadorContinua{
    jogadorContinua = false
    respVerificada = false
    fmt.Printf("Sua pontuação é %d\n", pontosJ)
  }

  if pontosPC >= 19 && pontosJ < pontosPC && dealerContinua{
    fmt.Printf("O dealer terminou com %d pontos\n", pontosPC)
    dealerContinua = false

  } else if pontosPC <= 19 && dealerContinua{
    cartaDealer()
  }
}

func dealerInicia(){
  if pontosPC >= 19 && pontosJ < pontosPC && dealerContinua{
    fmt.Printf("O dealer terminou com %d pontos\n", pontosPC)
    dealerContinua = false

  } else if pontosPC <= 19 && dealerContinua{
    cartaDealer()
  }
  if respVerificada && jogadorContinua{
    cartaJogador()

  } else if !respVerificada && jogadorContinua{
    jogadorContinua = false
    respVerificada = false
    fmt.Printf("Sua pontuação é %d\n", pontosJ)
  }
}

func cartaJogador(){
  cartaJ = gerarCarta()
  pontosJ = pontosJ + cartaJ.Pontos
  fmt.Printf("\nA sua carta sorteada é %s de %s\nSua pontuação atual é %d\n", cartaJ.Nome, cartaJ.Naipe, pontosJ)
}

func cartaDealer(){
  cartaPC = gerarCarta()
  pontosPC = pontosPC + cartaPC.Pontos
  fmt.Printf("\nA carta do dealer é %s de %s\n E sua pontuação atual é %d\n", cartaPC.Nome, cartaPC.Naipe, pontosPC)
}

func gerarCarta()(carta CartaType){
  var nome = [13]string{"Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez", "Valete", "Rainha", "Rei", "As"}
  var naipe = [4]string{"Ouros", "Copas", "Espadas", "Paus"}
  var pontos = [13]int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 1}
  var num int32
  utilizada := true

  for utilizada {
    num = gerarRandom(52)
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

func gerarRandom(num int32)(random int32){
  tempo := time.Now()
  nanosecond := tempo.Nanosecond()
  rand.Seed(int64(nanosecond))
  random = rand.Int31n(num)

  return
}
