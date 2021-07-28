package dictionary

import "testing"

func TestDictionary(t *testing.T){
    t.Run("known word",func(t *testing.T){
        dictionary := Dictionary{"test":"this is just a test"}
        got,_ := dictionary.Search("test")
        want := "this is just a test"
        assertStrings(t,got,want)
    })
    t.Run("unknown word", func(t *testing.T){
        dictionary := Dictionary{"test":"this is just a test"}
        _,got := dictionary.Search("unknown")
        //want := "cannot find the word you searched"
	assertError(t,got,errorNotFound)
})
}

func TestAdd(t *testing.T){
    t.Run("add new word",func(t *testing.T){
        dic := Dictionary{}
        key := "key1"
        value := "this is the first key"
	err := dic.Add(key,value)
        want := "this is the first key"
	assertError(t,err,nil)
        assertDefinition(t, dic, key, value, want)
    })
    t.Run("add exist word",func(t *testing.T){
        key := "key1"
        value := "this is the first key"
        want := "this is the first key"
	dic := Dictionary{key:value}
	err := dic.Add(key, "new one")
	assertError(t, err, wordExists)
        assertDefinition(t, dic, key, value, want)
    })
}

func assertDefinition(t *testing.T, d Dictionary, key, value, target string){
    t.Helper()
    got,err := d.Search(key)
    if err != nil{
        t.Fatal("should find added word:",err)
    }
    if got != target {
        t.Errorf("want %s but got %s",target,got)
    }
}

func assertError(t *testing.T, got, want error){
    if got != want{
        t.Errorf("got %v but want %v",got,want)
    }
}

func assertStrings(t *testing.T, got, want string){
    t.Helper()
    if got != want {
	    t.Errorf("got '%s' but want '%s', ",got,want)
    }
}
