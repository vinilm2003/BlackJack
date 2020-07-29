package main

import (
  "fmt"
  "strings"
  "math/rand"
   "time"
)

type Carta struct {
  Nome string
  Naipe string
  Pontos int32
}

var cartasUtilizadas = []int32{-1}

func main() {
  var carta Carta
  var verif string
  testejogo := true

  for testejogo {
    carta = gerarCarta()
    fmt.Printf("carta %s de %s, %d pontos\n", carta.Nome, carta.Naipe, carta.Pontos)
    fmt.Scanln(&verif)
    testejogo = strings.Contains(verif, "")
  }
}

func gerarRandom()(random int32){
  tempo := time.Now()
  nanosecond := tempo.Nanosecond()
  rand.Seed(int64(nanosecond))
  random = rand.Int31n(52)

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
