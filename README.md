# 🌍 VULN-R2S-GLO — Global React2Shell (R2S) Scanner

> **First-in-class, high-performance scanner for CVE-2025-55182 & CVE-2025-66478**  
> Built in Go for unmatched speed, resilience, and global coverage.  
> By [@vulnquest58](https://github.com/vulnquest58)

---

## 🔥 Why VULN-R2S-GLO?

While other tools scan slowly in Python, **VULN-R2S-GLO** leverages **Go’s concurrency model** to:
- Scan **100,000+ targets/hour** with minimal memory
- Bypass WAFs using **128KB junk padding**
- **Resume automatically** after crashes or interruptions
- **Pause every 10K requests** to avoid network throttling
- Export **only verified vulnerable targets** in structured JSON

This is the **first public scanner** purpose-built for **React2Shell (R2S)** at internet scale.

---

## 🛠️ Features

| Feature | Description |
|--------|-------------|
| ⚡ **Blazing Fast** | Goroutines enable 50,000+ concurrent scans |
| 🛡️ **WAF Bypass** | Automatic junk data injection (128KB) |
| 🔁 **Resume on Crash** | State saved every 1,000 scans |
| ⏸️ **Smart Pausing** | 2-minute break every 10,000 requests |
| 📦 **Zero Dependencies** | Pure Go — no external libraries |
| 🌐 **Global Coverage** | Random public IP generation (no private/reserved ranges) |
| 📁 **Structured Output** | Clean JSON reports per vulnerable host |

---

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Git (optional)

### Build from Source
```bash
git clone https://github.com/vulnquest58/VULN-R2S-GLO.git
cd VULN-R2S-GLO
go mod init VULN-R2S-GLO
go build -o VULN-R2S-GLO main.go
```

### Run Scan (Windows)
```bash
.\VULN-R2S-GLO.exe --hours 24 --concurrent 50000
```
### Run Scan (Linux/macOS)
```bash
chmod +x VULN-R2S-GLO
./VULN-R2S-GLO --hours 6 --concurrent 10000
```
### Run Scan (Linux/macOS)
```bash
Usage of ./VULN-R2S-GLO:
  -concurrent int
        Max concurrent goroutines (default 50000)
  -hours int
        Scan duration in hours (default 24)
```
### Examples
```bash
# 24-hour full scan (default)
./VULN-R2S-GLO

# 6-hour targeted scan
./VULN-R2S-GLO --hours 6 --concurrent 20000

# Graceful shutdown: Press Ctrl+C → state auto-saved
```
### 📂 Output Structure
After a successful run, you’ll find:

```bash
VULN-R2S-GLO/
├── output/
│   └── vulnerable/
│       ├── r2s_203.45.67.89_1712345678.json
│       └── r2s_185.12.34.56_1712345699.json
└── state/
    └── scan_state.json
```
Each JSON file contains:

- Confirmed command output (id, uname, etc.)
- Server fingerprint (OS, Node.js version, user)
- Full metadata and timestamp

### ⚠️ Ethical Use Only
This tool is intended strictly for authorized penetration testing and educational research.
Unauthorized scanning of systems you do not own or have explicit permission to test is illegal.

By using VULN-R2S-GLO, you agree to comply with all applicable laws and ethical guidelines.






