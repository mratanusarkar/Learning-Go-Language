

var str = "Old Value!"

func main() {
    str, err := foo()
    if err != nil {
        // There will never be an error
        log.Println(err)
    }

    fmt.Println(str) // Prints "New Value!"
    printStr()       // Prints "Old Value!"
}

func foo() (string, error) {
    return "New Value!", nil
}

func printStr() {
    fmt.Println(str)
}