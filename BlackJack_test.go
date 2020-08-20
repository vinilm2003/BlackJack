package main

import(
  "testing"
)

func TestGerarRandom (t *testing.T){
  for numDeCartas, i := gerarRandom(52), 0; i <= 100; i++{
    if numDeCartas < 0 || numDeCartas > 51{
      t.Errorf("Geração de carta inválida - %d", numDeCartas)
    }
  }
}
