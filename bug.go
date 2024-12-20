func main() {
    a := make([]int, 0, 5)
    for i := 0; i < 5; i++ {
        a = append(a, i)
    }

    for _, v := range a {
        fmt.Println(v)
    }

    b := a[:2]
    fmt.Println("b:", b)

    b = append(b, 100)
    fmt.Println("b:", b)
    fmt.Println("a:", a)
}