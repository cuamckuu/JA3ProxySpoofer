# JA3 Proxy Spoofer

This is prototype of HTTP Forward Proxy Server with JA3 Spoofing to bypass TLS Fingerprint check from CloudFlare and other alike service

```bash
# Cloudflare uses TLS fingerprint to block curl request
$ curl -L --silent 'http://api.mangalib.me/api/manga/14034--super-god-gene/chapters' | head -c100
# <!DOCTYPE html><html lang="en-US"><head><title>Just a moment...</title><meta http-equiv="Content-Typ

# Using proxy to spoof TLS fingerprint and bypass Cloudflare check
$ curl -L --silent --proxy 'http://localhost:8080' 'http://api.mangalib.me/api/manga/14034--super-god-gene/chapters' | head -c100
{"data":[{"id":301971,"index":1,"item_number":1,"volume":"1","number":"1","number_secondary":"1","na
```

# Why Proxy?

I want to scrape websites using python, but there is no easy way to spoof JA3, so here is a simple proxy server to do this.

# References

- https://blog.foxio.io/ja4%2B-network-fingerprinting
- https://eli.thegreenplace.net/2022/go-and-proxy-servers-part-1-http-proxies/
- https://github.com/Danny-Dasilva/CycleTLS
- https://github.com/Skyuzii/CycleTLS
- https://github.com/cucyber/JA3Transport
- https://github.com/lwthiker/curl-impersonate
- https://github.com/refraction-networking/utls
- https://habr.com/ru/articles/596411/
- https://kovardin.ru/articles/go/https-proxy-golang/
- https://reintech.io/blog/creating-simple-proxy-server-with-go
- https://scrapfly.io/blog/how-to-avoid-web-scraping-blocking-tls/
- https://scrapfly.io/web-scraping-tools/ja3-fingerprint

# TODO

- [X] Proof of concept
- [X] Bypass check at ranobelib
- [ ] Add support for proxy chaining
- [ ] Check bypass works for betfair and paddypower
- [ ] Handle proxy related headers, like `X-Forward` and others...
- [ ] Somehow fix `undexpected EOF` for https://tools.scrapfly.io/api/tls
- [ ] Find a way to remove replacing `http` for `https` and to avoid related SSL errors
- [ ] Randomize fingerprints on each request
- [ ] Add socks5 support
- [ ] Add benchmarcs and stress test
- [ ] Add support for proxy credentials
- [ ] Add proxy chain mode with predefined proxies
