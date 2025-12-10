#!/bin/bash
echo "🚀 Building VULN-R2S-GLO..."

GOOS=linux GOARCH=amd64 go build -o VULN-R2S-GLO-linux main.go
GOOS=windows GOARCH=amd64 go build -o VULN-R2S-GLO.exe main.go
GOOS=darwin GOARCH=amd64 go build -o VULN-R2S-GLO-mac main.go

echo "✅ Build completed:"
ls -lh VULN-R2S-GLO* 2>/dev/null || echo "No binaries built (check Go installation)"