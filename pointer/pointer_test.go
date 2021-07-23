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

    assertNoError := func(t *testing.T, got error){
        if got != nil{
	    t.Fatal("got an error but didnt expected one")
	}
    }

    assertError := func(t *testing.T,got error, want error){
        if got == nil{
	    t.Fatal("want an error but got none")
	}
	if got != want {
	    t.Errorf("got %q but want %q", got, want)
	}
    }

    t.Run("rebuild struct Wallet",func(t *testing.T){
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(100))
        want := Bitcoin(100)
	assertWallet(t, wallet, want)
    })

    t.Run("withdraw from wallet with fund",func(t *testing.T){
	startBalance := Bitcoin(20)
	wallet := Wallet{balance:startBalance}
	err := wallet.WithDraw(Bitcoin(10))
//	want := Bitcoin(20)
	assertWallet(t, wallet, Bitcoin(10))
	assertNoError(t, err)
    })

    t.Run("withdraw without enough fund",func(t *testing.T){
        wallet := Wallet{Bitcoin(20)}
	err := wallet.WithDraw(Bitcoin(100))
	assertWallet(t, wallet, Bitcoin(20))
	assertError(t,err,insufficientFundErr)
    })

}
