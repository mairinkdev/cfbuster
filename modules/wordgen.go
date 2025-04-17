package modules

import (
    "fmt"
    "strings"
)

func GenerateDynamicWordlist(domain string, fromCrt []string) []string {
    fmt.Println("[~] Gerando wordlist dinâmica...")

    base := []string{
        "www", "api", "mail", "dev", "test", "staging", "admin",
        "login", "cdn", "static", "ftp", "app", "panel",
    }

    var result []string
    for _, word := range base {
        result = append(result, word)
        result = append(result, word+"-dev")
        result = append(result, word+"-beta")
        result = append(result, "br-"+word)
        result = append(result, word+"-2024")
    }

    for _, sub := range fromCrt {
        parts := strings.Split(sub, ".")
        if len(parts) > 0 {
            result = append(result, parts[0])
        }
    }

    return unique(result)
}

func unique(input []string) []string {
    m := make(map[string]bool)
    var result []string
    for _, item := range input {
        if _, exists := m[item]; !exists {
            m[item] = true
            result = append(result, item)
        }
    }
    return result
}
