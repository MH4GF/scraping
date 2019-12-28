package main

type totalTable struct {
	Income   string `json:"income"`
	Expenses string `json:"expenses"`
	Balance  string `json:"balance"`
}

func newTotalTable(income string, expenses string, balance string) *totalTable {
	return &totalTable{
		Income:   income,
		Expenses: expenses,
		Balance:  balance,
	}
}
