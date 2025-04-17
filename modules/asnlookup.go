package modules

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func LookupASN(ip string) string {
    fmt.Println("[~] Consultando ASN para IP:", ip)
    url := fmt.Sprintf("https://api.iptoasn.com/v1/as/ip/%s", ip)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Erro na consulta ASN:", err)
        return ""
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    var data map[string]interface{}
    json.Unmarshal(body, &data)

    if val, ok := data["as_description"]; ok {
        fmt.Println("[ASN] Organização:", val)
    }

    if cidr, ok := data["announced"]; ok {
        fmt.Println("[ASN] Faixa CIDR:", cidr)
        return cidr.(string)
    }

    return ""
}
