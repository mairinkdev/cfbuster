package modules

import (
    "fmt"
    "net"
    "strings"
    "sync"
)

var cfRanges = []string{
    "104.", "172.", "188.", "198.41", "190.93", "141.101", "108.162", "103.21", "103.22",
}

func isCloudflare(ip string) bool {
    for _, prefix := range cfRanges {
        if strings.HasPrefix(ip, prefix) {
            return true
        }
    }
    return false
}

func ResolveAndValidate(domain string, subdomains []string) []string {
    var wg sync.WaitGroup
    var mu sync.Mutex
    var results []string

    for _, sub := range subdomains {
        full := sub + "." + domain
        wg.Add(1)
        go func(s string) {
            defer wg.Done()
            ips, err := net.LookupHost(s)
            if err == nil {
                for _, ip := range ips {
                    if !isCloudflare(ip) {
                        result := fmt.Sprintf("[FOUND] %s -> %s", s, ip)
                        fmt.Println(result)
                        mu.Lock()
                        results = append(results, result)
                        mu.Unlock()
                    }
                }
            }
        }(full)
    }

    wg.Wait()
    return results
}
