package pointer

import(
     "fmt";
     "errors"
)

type String interface{
    String() string
}

func (b Bitcoin) String()string{
    return fmt.Sprintf("%d BTC", b)
}

type Bitcoin int

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet)WithDraw(coin Bitcoin)error{
    if coin > w.balance{
        return errors.New("more than the amount you have")
    }
    w.balance -= coin
    return nil
}

func (w *Wallet)Balance()String{
    fmt.Printf("address of balance is %v\n",&w)
    return String(w.balance)
}

func (w *Wallet)Deposit(amount Bitcoin){
    fmt.Printf("address of deposit is %v\n",&w)
    w.balance += amount
}
