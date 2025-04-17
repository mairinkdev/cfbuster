package modules

import (
    "fmt"
    "net/http"
    "strings"
    "time"
)

func FingerprintFromLine(line string) {
    parts := strings.Split(line, " -> ")
    if len(parts) != 2 {
        return
    }
    ip := strings.TrimSpace(parts[1])

    url := "http://" + ip
    client := http.Client{Timeout: 3 * time.Second}

    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("[!] Fingerprint FAIL:", ip)
        return
    }
    defer resp.Body.Close()

    server := resp.Header.Get("Server")
    powered := resp.Header.Get("X-Powered-By")
    fmt.Printf("[FP] %s | Server: %s | Powered-By: %s\n", ip, server, powered)
}
