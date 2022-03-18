package ns2

type Bank struct {
	Balance []int64
	Num     int
}

func Constructor(balance []int64) Bank {
	b := Bank{}
	b.Balance = make([]int64, len(balance)+1)
	for i := 0; i < len(balance); i++ {
		b.Balance[i+1] = balance[i]
	}
	b.Num = len(balance) + 1
	return b
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 <= 0 || account1 > this.Num {
		return false
	}
	if account2 <= 0 || account2 > this.Num {
		return false
	}
	if this.Balance[account1] < money {
		return false
	}
	this.Balance[account1] -= money
	this.Balance[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account <= 0 || account > this.Num {
		return false
	}
	this.Balance[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account <= 0 || account > this.Num {
		return false
	}
	if this.Balance[account] < money {
		return false
	}
	this.Balance[account] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
