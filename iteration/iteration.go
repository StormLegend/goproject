package iteration

func Repeat(a string)string{
    repeat := ""
    for i:=0;i<5;i++ {
        repeat = repeat + a
    }
    return repeat
}
