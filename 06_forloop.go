package main
import "fmt"
func main() {

    students := map[int]string{
        101: "Alice",
        102: "Bob",
        103: "Charlie",
    }

    for key, value := range students {

        fmt.Printf("ID: %d, Name: %s\n", key, value)
    }
}