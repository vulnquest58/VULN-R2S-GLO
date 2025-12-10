package main

import (
	"VULN-R2S-GLO/config"
	"VULN-R2S-GLO/scanner"
	"VULN-R2S-GLO/utils"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	hours      = flag.Int("hours", 24, "Scan duration in hours")
	concurrent = flag.Int("concurrent", 50000, "Max concurrent goroutines")
)

func main() {
	flag.Parse()
	utils.InitLogger()

	cfg := &config.Config{
		MaxConcurrent:  *concurrent,
		ScanHours:      *hours,
		WAFBypassKB:    128,
		RequestTimeout: 8 * time.Second,
		PauseInterval:  10000,
		PauseDuration:  120 * time.Second,
	}

	utils.LogInfo("🌍 Starting VULN-R2S-GLO — Global React2Shell Scanner by @vulnquest58")
	utils.LogInfo(fmt.Sprintf("⚙️  Config: %d hours | %d concurrent | WAF Bypass: %dKB", cfg.ScanHours, cfg.MaxConcurrent, cfg.WAFBypassKB))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		utils.LogWarn("🛑 Received interrupt. Saving state and shutting down...")
		scanner.SaveState()
		cancel()
	}()

	scanner.Run(ctx, cfg)
}
