package modules

import (
    "fmt"
    "net"
    "strings"
    "time"
    "net/http"
)

func ScanCIDRRange(cidr string, originalServer string) {
    fmt.Println("[~] Iniciando varredura no CIDR:", cidr)

    ip, ipnet, err := net.ParseCIDR(cidr)
    if err != nil {
        fmt.Println("Erro ao interpretar CIDR:", err)
        return
    }

    for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
        target := ip.String()
        checkHTTPFingerprint(target, originalServer)
    }
}

func inc(ip net.IP) {
    for j := len(ip) - 1; j >= 0; j-- {
        ip[j]++
        if ip[j] > 0 {
            break
        }
    }
}

func checkHTTPFingerprint(ip string, originalServer string) {
    url := "http://" + ip
    client := http.Client{Timeout: 3 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        return
    }
    defer resp.Body.Close()

    server := resp.Header.Get("Server")
    if strings.Contains(server, originalServer) && server != "" {
        fmt.Printf("[MATCH] %s | Server: %s\n", ip, server)
    }
}
