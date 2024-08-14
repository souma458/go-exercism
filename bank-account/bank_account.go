package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	closed  bool
	lock    sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, closed: false}
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.closed {
		return a.balance, false
	}

	a.lock.Lock()
	defer a.lock.Unlock()
	if a.balance+amount < 0 {
		return a.balance, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	if a.closed {
		return 0, false
	}
	balanceAtClose := a.balance
	a.balance = 0
	a.closed = true
	return balanceAtClose, a.closed
}
