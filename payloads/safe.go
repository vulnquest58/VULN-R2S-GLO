package payloads

import "fmt"

func BuildSafePayload() (string, string) {
	boundary := "----WebKitFormBoundaryR2SDetect1337"
	body := fmt.Sprintf(
		"------%s\r\n"+
			"Content-Disposition: form-data; name=\"1\"\r\n\r\n{}\r\n"+
			"------%s\r\n"+
			"Content-Disposition: form-data; name=\"0\"\r\n\r\n[\"$1:aa:aa\"]\r\n"+
			"------%s--",
		boundary, boundary, boundary,
	)
	return body, "multipart/form-data; boundary=" + boundary
}
