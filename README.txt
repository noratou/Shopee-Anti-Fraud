# 🛡️ Shopee Anti-Fraud AI Copilot

An autonomous, cloud-native cybersecurity agent designed for high-concurrency cross-border e-commerce (e.g., Shopee). This system ingests traffic logs and leverages AI to perform context-aware reasoning, instantly detecting and blocking sophisticated multi-border promotional abuse, credit card testing, and VPN fraud.

## ✨ Key Features
* **High-Performance Backend**: Built with Go to efficiently handle high-QPS traffic logs without bottlenecks.
* **Context-Aware AI Reasoning**: Utilizes Llama-3.3 (via Groq API) to go beyond traditional static WAF rules. It analyzes transaction contexts (e.g., rapid cross-border checkouts, mass promo abuse) to catch dynamic fraud patterns.
* **Database-Driven Pre-Filtering**: Implements a cost-saving architectural pipeline that checks MongoDB before calling the AI. If an IP is already flagged, it triggers an instant block and skips the LLM queue entirely.
* **Idempotent Storage (Upsert)**: Leverages MongoDB update operations with upsert configurations to maintain a clean, unique blacklist collection and prevent data duplication.
* **Threat Injection Testing**: Includes a dynamic synthetic data generator seeded with system time that simulates Southeast Asian cross-border traffic, randomly injecting varied targeted attack profiles to validate AI accuracy.

## 🛠️ Tech Stack
* **Language**: Go (Golang)
* **AI Model**: Llama-3.3-70b-versatile (Groq API)
* **Database**: MongoDB Atlas / Local MongoDB
* **Architecture**: Flat Modular Design (Data Generation, Contextual Analysis, Idempotent Persistence)

## 📁 Repository Structure
The project uses a clean, production-grade flat modular architecture for straightforward service division:
* main.go - The system orchestrator controlling the pipeline flow.
* generator.go - Synthetic log generator with time-seeded threat injection.
* ai.go - Handles structured JSON payload exchanges with the Groq Llama-3 API.
* db.go - Manages low-latency MongoDB connections, state checks, and upserts.

## 🚀 How to Run Locally

1. Clone the repository:
   git clone https://github.com/YourUsername/shopee-antifraud-go.git
   cd shopee-antifraud-go

2. Initialize your dependencies:
   go mod tidy

3. Create a .env file in the root directory and add your credentials:
   GROQ_API_KEY=your_groq_api_key_here
   MONGO_URI=your_mongodb_atlas_connection_string

4. Run the complete autonomous pipeline:
   go run .

## 📊 Sample System Execution Flow
=== 🚀 Starting Cross-Border Anti-Fraud Copilot ===
✅ Successfully connected to MongoDB!
🔄 Generating synthetic cross-border traffic logs...
🛑 [Instant Block] IP 185.220.101.5 is already in local database blacklist. Dropped from AI queue.
🧠 AI Agent is analyzing traffic patterns...
=== AI Security Agent Analysis Result (Powered by Groq + Go) ===
{
  "malicious_ips": [
     {
       "ip": "45.134.144.12",
       "risk_level": "High",
       "reason": "Mass promo-code brute-force attempt from known Datacenter IP"
     }
   ]
}
⚠️  AI Detected 1 new unique threats. Saving to blocklist...
🛡️  Processed security blacklist. Successfully synced 1 unique threats to MongoDB.
=== Process Completed ===