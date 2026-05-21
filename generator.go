package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TrafficLog struct {
	OrderID       string  `json:"order_id"`
	IP            string  `json:"ip"`
	TargetCountry string  `json:"target_country"`
	Amount        float64 `json:"amount"`
	Note          string  `json:"note"`
}

// GenerateTestLogs creates a batch of normal and suspicious traffic logs
func GenerateTestLogs(count int) []TrafficLog {
	fmt.Println("🔄 Generating synthetic cross-border traffic logs...")
	
	// 1. Seed the random generator so it changes every time you run the app
	rand.Seed(time.Now().UnixNano())

	countries := []string{"ID", "TH", "VN", "MY", "SG"}
	
	// 2. A pool of different enterprise-level attack scenarios for the portfolio
	maliciousScenarios := []TrafficLog{
		{IP: "185.220.101.5", Amount: 1500.00, Note: "High-frequency cross-border checkout within 2 seconds using VPN"},
		{IP: "45.134.144.12", Amount: 3000.00, Note: "Mass promo-code brute-force attempt from known Datacenter IP"},
		{IP: "193.142.146.35", Amount: 2.50, Note: "Testing multiple stolen credit cards with micro-transactions"},
		{IP: "8.242.12.99", Amount: 0.00, Note: "Account takeover attempt with rapid password resets"},
	}

	logs := []TrafficLog{}

	for i := 1; i <= count; i++ {
		// Inject a random malicious pattern every 5 logs
		if i%5 == 0 {
			// Randomly pick one of the attack scenarios
			badGuy := maliciousScenarios[rand.Intn(len(maliciousScenarios))]
			badGuy.OrderID = fmt.Sprintf("TXN_%04d", i)
			badGuy.TargetCountry = countries[rand.Intn(len(countries))]
			
			logs = append(logs, badGuy)
		} else {
			// Normal traffic
			logs = append(logs, TrafficLog{
				OrderID:       fmt.Sprintf("TXN_%04d", i),
				IP:            fmt.Sprintf("103.24.140.%d", rand.Intn(255)),
				TargetCountry: countries[rand.Intn(len(countries))],
				Amount:        float64(rand.Intn(100) + 10),
				Note:          "Normal user checkout behavior",
			})
		}
	}
	return logs
}