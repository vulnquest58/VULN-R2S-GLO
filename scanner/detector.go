package scanner

import (
	"VULN-R2S-GLO/payloads"
	"bytes"
	"crypto/tls"
	"net/http"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: 8 * time.Second,
	Transport: &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		MaxConnsPerHost:     50,
	},
}

func IsVulnerableSafe(url string) bool {
	body, ctype := payloads.BuildSafePayload()
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(body))
	req.Header.Set("Next-Action", "x")
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

	resp, err := httpClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 500 {
		return false
	}

	buf := make([]byte, 512)
	resp.Body.Read(buf)
	bodyStr := string(buf)
	return strings.Contains(bodyStr, `E{"digest"`)
}
