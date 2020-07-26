package main

import (
  "fmt"
  // "math"
  "strings"
  // "os"
  "math/rand"
   "time"
   "strconv"
)

type Carta struct {
  Nome string
  Naipe string
  Pontos int32
}

var cartasUtilizadas []string

func main() {

  var carta Carta
  testejogo := true
  var verif string

  for testejogo {
    if testejogo {
      carta = gerarCarta()
      fmt.Printf("carta %s de %s, %d pontos\n", carta.Nome, carta.Naipe, carta.Pontos)
      fmt.Scanln(&verif)
      testejogo = strings.Contains(verif, "s")
    }
  }
}

func gerarRandom()(random int32){
  tempo := time.Now()
  nanosecond := tempo.Nanosecond()
  rand.Seed(int64(nanosecond))
  random = rand.Int31n(52)

  return
}

func verificarGeracao(numeroDaCarta int32)(utilizacao bool){

  numString := strconv.Itoa(int(numeroDaCarta))

  for _, v := range cartasUtilizadas {
    jaUtilizado := strings.Contains(v, numString)

    if jaUtilizado {
      utilizacao = false
    } else {
      utilizacao = true
      cartasUtilizadas = append(cartasUtilizadas, numString)
      fmt.Println(cartasUtilizadas)
    }
  }
  return
}

func gerarCarta()(carta Carta){
  var nome = [13]string{"Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez", "Valete", "Rainha", "Rei", "As"}
  var naipe = [4]string{"Ouros", "Copas", "Espadas", "Paus"}
  var pontos = [13]int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 1}

  num := gerarRandom()

  utilizacao := verificarGeracao(num)
  if utilizacao {
    numCarta := num % 13
    numPontos := num % 13
    numNaipe := num % 4

    carta.Nome = nome[numCarta]
    carta.Naipe = naipe[numNaipe]
    carta.Pontos = pontos[numPontos]

  } else {
    num = gerarRandom()
  }

  var primeiraCcarta = true

  if primeiraCcarta {
    num = gerarRandom()

    numCarta := num % 13
    numPontos := num % 13
    numNaipe := num % 4

    carta.Nome = nome[numCarta]
    carta.Naipe = naipe[numNaipe]
    carta.Pontos = pontos[numPontos]

    primeiraCcarta = false
  }

return
}
