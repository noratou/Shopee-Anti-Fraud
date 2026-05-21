package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("=== 🚀 Starting Cross-Border Anti-Fraud Copilot ===")

	// 1. Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// 2. Setup MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" 
	}
	dbClient := ConnectDB(mongoURI)

	// 3. Generate Test Data
	rawLogs := GenerateTestLogs(10) 
	
	// 4. Pre-Filter: Remove IPs that are ALREADY blacklisted in MongoDB
	filteredLogs := []TrafficLog{}
	for _, logLine := range rawLogs {
		if IsIPBlacklisted(dbClient, logLine.IP) {
			fmt.Printf("🛑 [Instant Block] IP %s is already in local database blacklist. Dropped from AI queue.\n", logLine.IP)
		} else {
			filteredLogs = append(filteredLogs, logLine)
		}
	}

	// 5. If everything was already blocked, we can exit early!
	if len(filteredLogs) == 0 {
		fmt.Println("✅ No new traffic needs AI analysis. All threats already mitigated.")
		fmt.Println("=== Process Completed ===")
		return
	}

	// 6. AI Analysis on NEW/UNKNOWN traffic only
	logsJSON, _ := json.MarshalIndent(filteredLogs, "", "  ")
	maliciousData := AnalyzeTraffic(logsJSON)

	// 7. Save New Threats to Database
	if maliciousData != nil && len(maliciousData) > 0 {
		fmt.Printf("⚠️  AI Detected %d new unique threats. Saving to blocklist...\n", len(maliciousData))
		SaveBlacklist(dbClient, maliciousData)
	} else {
		fmt.Println("✅ AI Analysis complete. No new threats detected.")
	}

	fmt.Println("=== Process Completed ===")
}