package pointers

import "fmt"

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

//type Stringer interface {
//	String() string
//}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return fmt.Errorf("insufficient funds: have %s, want %s", w.balance, amt)
	}
	w.balance -= amt
	return nil
}

func newPointer() {
	var ptr *int
	var a int = 10
	ptr = &a
	fmt.Println(a)
	fmt.Println("Memory address", ptr)

	fmt.Println("Value at the memory address", *ptr)
}
