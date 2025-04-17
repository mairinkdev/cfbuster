package modules

import (
    "fmt"
    "math/rand"
    "net"
    "time"
)

func CheckWildcard(domain string) {
    rand.Seed(time.Now().UnixNano())
    testSub := fmt.Sprintf("%d-test.%s", rand.Intn(99999), domain)
    ips, err := net.LookupHost(testSub)
    if err == nil && len(ips) > 0 {
        fmt.Println("[!] Wildcard DNS detectado:", testSub, "->", ips[0])
    } else {
        fmt.Println("[✓] Nenhum wildcard DNS detectado.")
    }
}
