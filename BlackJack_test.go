package main

import (
	"testing"
)

func TestGerarRandom(t *testing.T) {
	for numDeCartas, i := gerarRandom(52), 0; i <= 100; i++ {
		if numDeCartas < 0 || numDeCartas > 51 {
			t.Errorf("Geração de carta inválida - %d", numDeCartas)
		}
	}
}

func TestJogo(t *testing.T) {
	for numDeCartas, i := gerarRandom(2), 0; i <= 10; i++ {
		if numDeCartas < 0 || numDeCartas > 1 {
			t.Errorf("Geração de número para início inválido - %d", numDeCartas)
		}
	}
}

func TestJogadorInicia(t *testing.T) {
	var respostasPossiveis = [5]string{"sim", "s", "não", "nao", "n"}

	for pontosPC := 0; pontosPC <= 21; pontosPC++ {

		if pontosPC < 21 {
			for _, v := range respostasPossiveis {
				switch v {
				case "sim", "s":

				case "não", "nao", "n":

				default:
					t.Error("Valor inválido")
				}
			}
		} else if pontosPC == 21 {

		} else {
			t.Errorf("Valor inválido")
		}
	}
}

func TestGerarCarta(t *testing.T) {
	var cartasUtilizadas = []int32{-1}
	var num, numCarta, ultimoNumero int32
	utilizada := true

	for i := 0; i < 13; i++ {

		for utilizada {
			num = gerarRandom(13)
			for _, v := range cartasUtilizadas {
				utilizada = v == num
				if utilizada {
					break
				}
			}
		}

		utilizada = true
		cartasUtilizadas = append(cartasUtilizadas, num)

		numCarta = num % 13

		if numCarta != ultimoNumero {

		} else {
			t.Errorf("Nome de carta repetido")
		}

		ultimoNumero = num % 13
	}
}

func TestJogadorContinua(t *testing.T) {
	for continua, pontuacao := true, 0; pontuacao <= 21; pontuacao++ {
		if continua && pontuacao == 21 {
			continua = false
		} else if continua && (pontuacao > 21 || pontuacao < 0) {
			t.Errorf("Pontuação inválida %d", pontuacao)
		}
	}
}

func TestVerificarPontuacao(t *testing.T) {
	var pontosJ, pontosPC int32
	for i := 0; i < 105; i++ {
		pontosJ = gerarRandom(22)
		pontosPC = gerarRandom(22)

		if pontosJ > 21 || pontosPC > 21 {
			t.Errorf("Pontuação inválida: JOGADOR %d / PC %d", pontosJ, pontosPC)
		} else if pontosJ < 0 || pontosPC < 0 {
			t.Errorf("Pontuação inválida: JOGADOR %d / PC %d", pontosJ, pontosPC)
		}
	}
}
