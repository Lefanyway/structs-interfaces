package contas

import (
	"projeto/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(ValorDoSaque float64) string {
	PodeSacar := ValorDoSaque > 0 && ValorDoSaque <= c.saldo
	if PodeSacar {
		c.saldo -= ValorDoSaque
		return "Saque realizado com sucesso: "
	} else {
		return "Saque insuficiente..."
	}
}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Depósito realizado com sucesso: ", c.saldo
	} else {
	}
	return "Depósito não foi realizado...", c.saldo
}

func (c *ContaCorrente) Transferir(ValorDaTransferencia float64, ContaDestino *ContaCorrente) bool {
	if ValorDaTransferencia < c.saldo {
		c.saldo -= ValorDaTransferencia
		ContaDestino.Depositar(ValorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
