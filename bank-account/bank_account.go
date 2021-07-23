// package account implements routines that manages
// a bank account
package account

import "sync"

type BankAccountT struct {
	mux     sync.Mutex
	balance int
	closed  bool
}

// Open creates a new bank account with
// the input amount as initial balance
func Open(amount int) *BankAccountT {
	if amount < 0 {
		return nil
	}
	return &BankAccountT{balance: amount}
}

// Balance returns the available balance
// of active account
func (acc *BankAccountT) Balance() (int, bool) {
	acc.mux.Lock();
	defer acc.mux.Unlock()

	if acc.closed {
		return 0, false
	}
	return acc.balance, true
}

// Close returns the available balance of
// an active account and then deactivates it
func (acc *BankAccountT) Close() (int, bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()

	if acc.closed {
		return 0, false
	}
    balance := acc.balance
	acc.balance = 0
	acc.closed = true
	return balance, true
}

// Deposit adds deposit to the balance of an
// active account
func (acc *BankAccountT) Deposit(deposit int) (int, bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()

	if acc.closed {
		return deposit, false
	}

	if acc.balance + deposit < 0 {
		return acc.balance, false
	}

	acc.balance += deposit
	return acc.balance, true
}
