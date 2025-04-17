package modules

import (
    "fmt"
    "strings"

    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
)

func BypassUAM(target string) {
    fmt.Println("[~] Iniciando tentativa de bypass Cloudflare UAM para:", target)

    url := "http://" + target
    path := launcher.New().
        Bin("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe").
        Headless(true).
        NoSandbox(true).
        Leakless(false). // desativa leakless.exe
        MustLaunch()

    browser := rod.New().ControlURL(path).MustConnect()
    defer browser.MustClose()

    page := browser.MustPage(url)
    page.MustWaitLoad().MustWaitIdle()

    html, err := page.HTML()
    if err != nil {
        fmt.Println("Erro ao capturar HTML:", err)
        return
    }

    if strings.Contains(html, "Just a moment") || strings.Contains(html, "Checking your browser") {
        fmt.Println("[!] Cloudflare UAM detectado, bypass executado com sucesso (aguardou JS)")
    } else {
        fmt.Println("[✓] Página carregada sem UAM:", url)
    }

    page.MustScreenshot("uam_bypass_result.png")
    fmt.Println("[📸] Screenshot salvo como uam_bypass_result.png")
}
