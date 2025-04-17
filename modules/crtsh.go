package modules

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func GetSubdomains(domain string) []string {
    fmt.Println("[~] Consultando crt.sh...")
    url := "https://crt.sh/?q=%25." + domain + "&output=json"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Erro:", err)
        return []string{}
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    var data []map[string]interface{}
    _ = json.Unmarshal(body, &data)

    found := make(map[string]bool)
    for _, entry := range data {
        if name, ok := entry["name_value"].(string); ok {
            for _, sub := range strings.Split(name, "\n") {
                sub = strings.TrimSpace(sub)
                if strings.HasSuffix(sub, domain) {
                    found[sub] = true
                }
            }
        }
    }

    var subs []string
    for sub := range found {
        subs = append(subs, strings.TrimPrefix(sub, domain+"."))
    }

    fmt.Printf("[✓] %d subdomínios do crt.sh\n", len(subs))
    return subs
}
