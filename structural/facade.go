package structural

import "fmt"

// Facade is a structural design pattern that provides a simplified (but limited) interface
// to a complex system of classes, library or framework.

// Take a pizza ordering system where you just select the pizza and enter card details for payment.

// The facade is the GUI/operator taking the order hiding the complexities of the underlying infrastructure/subsystem
// Subsystems are : Check Card details, Credit/Debit , Balance check , notification , ledger entry

type Account struct {
	card string
	code int
}

func (acc *Account) checkDetails(card string, code int) error {
	if acc.card == card && acc.code == code {
		fmt.Println("Verified account.")
		return nil
	}
	return fmt.Errorf("Account details are incorrect")
}

type Wallet struct {
	balance float32
}

func (w *Wallet) addMoney(amount float32) error {
	w.balance += amount
	fmt.Println("Balance added successfully.")
	return nil
}

func (w *Wallet) debitMoney(amount float32) error {
	if w.balance < amount {
		return fmt.Errorf("Insufficient balance")
	}
	fmt.Println("Sufficient balance for debit")
	w.balance -= amount
	return nil
}

type Notification struct {
}

func (n *Notification) sendCreditNotification() {
	fmt.Println("Sending credit notification")
}

func (n *Notification) sendDebitNotification() {
	fmt.Println("Sending debit notification")
}

type Ledger struct {
}

func (l *Ledger) makeEntry(acc string, t string, amount float32) {
	fmt.Printf("Making entry for account %s , txn type %s with amount %.2f\n", acc, t, amount)
}

type OrderFacade struct {
	account      *Account
	wallet       *Wallet
	notification *Notification
	ledger       *Ledger
}

func newOrderFacade(card string, code int) *OrderFacade {
	return &OrderFacade{
		account:      &Account{card, code},
		wallet:       &Wallet{0},
		notification: &Notification{},
		ledger:       &Ledger{},
	}
}

func (f *OrderFacade) DebitFromWallet(card string, code int, amount float32) error {
	// Check Card details first
	err := f.account.checkDetails(card, code)
	if err != nil {
		return err
	}
	// Add to wallet
	err = f.wallet.debitMoney(amount)
	if err != nil {
		return err
	}

	// Make ledger entry
	f.ledger.makeEntry(card, "debit", amount)

	// Send notification
	f.notification.sendDebitNotification()

	return nil
}

func (f *OrderFacade) AddMoneyToWallet(card string, code int, amount float32) error {
	// Check Card details first
	err := f.account.checkDetails(card, code)
	if err != nil {
		return err
	}
	// Add to wallet
	err = f.wallet.addMoney(amount)
	if err != nil {
		return err
	}

	// Make ledger entry
	f.ledger.makeEntry(card, "credit", amount)

	// Send notification
	f.notification.sendCreditNotification()

	return nil

}

func RunFacadeDemo() {
	// client makes payment through facade
	// As soon as he enters card details, we make a facade object for interacting with all subsystems
	f := newOrderFacade("abcd", 231)

	// If he is adding money
	err := f.AddMoneyToWallet("abcd", 231, 20)
	if err != nil {
		panic(err)
	}

	// If he is making payment
	err = f.DebitFromWallet("abcd", 231, 10)
	if err != nil {
		panic(err)
	}

}
