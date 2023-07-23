package contas

import (
	"projeto/clientes"
)

type ContaPoupanca struct {
	Titular                                  clientes.Titular
	NumeroDaAgencia, NumeroDaConta, Operacao int
	saldo                                    float64
}

func (c *ContaPoupanca) Sacar(ValorDoSaque float64) string {
	PodeSacar := ValorDoSaque > 0 && ValorDoSaque <= c.saldo
	if PodeSacar {
		c.saldo -= ValorDoSaque
		return "Saque realizado com sucesso: "
	} else {
		return "Saque insuficiente..."
	}
}

func (c *ContaPoupanca) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Depósito realizado com sucesso: ", c.saldo
	} else {
	}
	return "Depósito não foi realizado...", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
