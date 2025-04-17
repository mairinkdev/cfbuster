package main

import (
    "fmt"
    "os"
    "cfbuster/modules"
)

func main() {
    if len(os.Args) < 3 || os.Args[1] != "--target" {
        fmt.Println("Uso: ./cfbuster --target <domínio> [--shodan-key SUA_KEY]")
        return
    }

    domain := os.Args[2]
    var shodanKey string

    if len(os.Args) == 5 && os.Args[3] == "--shodan-key" {
        shodanKey = os.Args[4]
    }

    fmt.Println("🦈 cfbuster v1 - Varredura ofensiva iniciada.")
    modules.CheckWildcard(domain)

    subs := modules.GetSubdomains(domain)
    wordlist := modules.GenerateDynamicWordlist(domain, subs)
    wordlist = append(wordlist, modules.LoadWordlist("wordlist.txt")...)

    results := modules.ResolveAndValidate(domain, wordlist)

    for _, r := range results {
        modules.FingerprintFromLine(r)
    }

    if shodanKey != "" {
        modules.ShodanLookup(domain, shodanKey)
    }

    cidr := modules.LookupASN("185.76.43.122")
    modules.ScanCIDRRange(cidr, "nginx")

    modules.BypassUAM("email.gamesow.com")

    fmt.Println("\nhttps://github.com/mairinkdev")
	fmt.Println("Desenvolvida para fins educacionais 🐧")
}
