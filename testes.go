package main

import (
  "fmt"
  "math/rand"
  "time"
)

var resposta string
var teste bool

func main() {
  fmt.Printf("digite: \n")
  fmt.Scanln(&resposta)
  teste = simOuNao(resposta)
  if teste {
    quemComeca()
  }
}

func quemComeca(){
  var num int32
  iniciar := true

  jogadorOuDealer := gerarRandom(2)

  for iniciar {
    fmt.Printf("digite: \n")
    fmt.Scanln(&resposta)
    teste = simOuNao(resposta)

    if teste{
      if jogadorOuDealer == 0{
        num = gerarRandom(11)
        fmt.Printf("jogador %d\n", num)
        fmt.Printf("digite: \n")
        fmt.Scanln(&resposta)

      }else if jogadorOuDealer == 1{
        num = gerarRandom(11)
        fmt.Printf("Dealer %d\n", num)
        fmt.Printf("digite: \n")
        fmt.Scanln(&resposta)

      }
    }
  }
}

func simOuNao(resp string)(sim bool){
  if resp == ""{
    sim = true
  }
  return
}

func gerarRandom(num int32)(random int32){
  tempo := time.Now()
  nanosecond := tempo.Nanosecond()
  rand.Seed(int64(nanosecond))
  random = rand.Int31n(num)

  return
}
