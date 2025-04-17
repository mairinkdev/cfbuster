package modules

import (
    "bufio"
    "fmt"
    "os"
)

func LoadWordlist(path string) []string {
    var words []string
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Erro ao abrir wordlist:", err)
        return words
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    return unique(words)
}
