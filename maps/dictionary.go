package dictionary

import "errors"

var errorNotFound = errors.New("cannot find the word you searched")
var wordExists = errors.New("cannot add word because it already exists")

type Dictionary map[string]string

func (d Dictionary)Search(target string)(string,error){
	definition,ok := d[target]
	if !ok {
	    return "",errorNotFound
	}
	return definition,nil
}

func (d Dictionary)Add(key, value string)error{
    d[key] = value
    return nil
}
