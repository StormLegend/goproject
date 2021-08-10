package dictionary

//import "errors"

const(
    ErrNotFound = DictionaryErr("cannot find the word you searched")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)


type DictionaryErr string
type Dictionary map[string]string


func (e DictionaryErr) Error() string{
    return string(e)
}

func (d Dictionary)Search(target string)(string,error){
	definition,ok := d[target]
	if !ok {
	    return "",ErrNotFound
	}
	return definition,nil
}

func (d Dictionary)Add(key, value string)error{
    _, err := d.Search(key)
    switch err {
        case ErrNotFound:
		d[key] = value
        case nil:
		return ErrWordExists
	default:
		return err
    }
    //d[key] = value
    return nil
}

func (d Dictionary)Update(key, value string){
    d[key] = value
}
