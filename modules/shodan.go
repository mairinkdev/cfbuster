package modules

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func ShodanLookup(domain, key string) {
    fmt.Println("[~] Buscando hostnames no Shodan para:", domain)
    url := fmt.Sprintf("https://api.shodan.io/dns/domain/%s?key=%s", domain, key)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("[Shodan] Erro:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("[Shodan] Resultado bruto:\n", string(body))

    var data map[string]interface{}
    _ = json.Unmarshal(body, &data)

    if subs, ok := data["subdomains"]; ok {
        fmt.Println("[Shodan] Subdomínios encontrados:", subs)
    }
}
