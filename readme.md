# 🛡️ cfbuster v1 

**cfbuster** é uma ferramenta de reconhecimento ofensivo e evasão inteligente de infraestrutura web, projetada para **descobrir subdomínios reais, identificar IPs de origem expostos mesmo atrás da Cloudflare e outras CDNs**, e **testar proteção anti-bot (UAM)** em modo totalmente automatizado.

---

## Funcionalidades

| Tipo     | Recurso |
|----------|---------|
|  Inteligência Passiva | Subdomínio discovery via [crt.sh](https://crt.sh) |
|  Wordlist Dinâmica | Geração de nomes com padrões combinatórios e bruteforce inteligente |
|  DNS Resolver | Resolução paralela de subdomínios com filtro automático de IPs Cloudflare |
|  HTTP Fingerprint | Coleta automática de headers HTTP (Server, X-Powered-By) |
|  Consulta ASN | Identifica a faixa CIDR do IP de origem |
|  Scanner CIDR | Varredura de IPs vizinhos no ASN buscando fingerprint semelhante |
|  Shodan Lookup | Integração via API para extração de subdomínios históricos e metadados |
|  Bypass Cloudflare UAM | Executa browser headless real (Chrome via Rod) e burla proteções JS |
|  Screenshot | Salva imagem da tela após renderização UAM com JS |

---

## 📦 Como usar

### Execução simples:
```bash
go run main.go --target example.com
```

### Execução com chave da Shodan:
```bash
go run main.go --target example.com --shodan-key SUA_API_KEY
```

---

## 🧰 Requisitos

- Go 1.18 ou superior
- Navegador Google Chrome instalado (para bypass UAM)
- Conexão com internet
- (Opcional) Conta no [shodan.io](https://shodan.io) para usar a API

---

## 📂 Exemplo de execução

```bash
go run main.go --target ???.com
```

```
🚀 cfbuster v3.5 - Varredura ofensiva iniciada.
[✓] Nenhum wildcard DNS detectado.
[✓] 18 subdomínios encontrados via crt.sh
[FOUND] ???.com -> 00.000.00.000
[FP] 00.000.00.000 | Server: envoy | Powered-By:
[ASN] Organização: AMAZON-02
[ASN] Faixa CIDR: 00.000.0.0/00
[MATCH] 00.000.00.000 | Server: envoy
[✓] Página carregada sem UAM: http://???.com
```

---

## ⚠️ Aviso Legal

> Esta ferramenta foi criada exclusivamente para **estudos, treinamentos, auditorias autorizadas e testes em ambientes controlados**.
> O uso não autorizado contra terceiros **pode configurar crime previsto em lei**.

---

## 👨‍💻 Autor

**Arthur Mairink**
https://github.com/mairinkdev
