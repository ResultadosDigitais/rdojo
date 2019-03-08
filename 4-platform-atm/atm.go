package atm

import "errors"

type ATM struct {
	billStocks []*billStock
}

func NewATM() *ATM {
	return &ATM{
		billStocks: startBillStocks(),
	}
}

type billStock struct {
	bill      int
	available int
}

func (atm *ATM) Withdraw(moneyAmount int) (map[int]int, error) {
	billCounts := make(map[int]int)

	for _, stock := range atm.billStocks {
		bill := stock.bill

		for stock.available > 0 && moneyAmount >= bill {
			moneyAmount -= bill
			stock.available--
			billCounts[bill]++
			if moneyAmount == 0 {
				return billCounts, nil
			}
		}
	}
	return map[int]int{}, errors.New("valor invalido")

}

func startBillStocks() []*billStock {
	return []*billStock{
		{
			bill:      100,
			available: 10,
		},
		{
			bill:      50,
			available: 20,
		},
		{
			bill:      20,
			available: 50,
		},
		{
			bill:      10,
			available: 100,
		},
	}
}
