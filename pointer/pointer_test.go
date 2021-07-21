package pointer

import (
    "testing"
)


func TestWallet(t *testing.T){
    assertWallet := func(t testing.TB, wallet Wallet, want Bitcoin){
        t.Helper()
	if wallet.Balance() != String(want){
	    t.Errorf("got %s but want %s", wallet.Balance(), String(want))
	}
    }
    assertError :=func(t *testing.T,got error, want string){
        if got == nil{
	    t.Errorf("want an error but got none")
	}
    }
    /*t.Run("first Test of wallet", func(t *testing.T){
        wallet := Wallet{}
        wallet.Deposit(10)
        want := Bitcoin(10)
	assertWallet(t, wallet, want)
    })*/

    t.Run("rebuild struct Wallet",func(t *testing.T){
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(100))
        want := Bitcoin(100)
	assertWallet(t, wallet, want)
    })
    t.Run("withdraw from wallet",func(t *testing.T){
	startBalance := Bitcoin(20)
	wallet := Wallet{balance:startBalance}
	err := wallet.WithDraw(Bitcoin(100))
//	want := Bitcoin(20)
	assertError(t, err)
	assertWallet(t, wallet, startBalance)
    })

}
