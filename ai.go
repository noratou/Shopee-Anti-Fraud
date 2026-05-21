package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// AnalyzeTraffic sends the logs to Groq API and returns malicious IPs
func AnalyzeTraffic(logsJSON []byte) []interface{} {
	fmt.Println("🧠 AI Agent is analyzing traffic patterns...")

	apiKey := os.Getenv("GROQ_API_KEY")
	
	prompt := fmt.Sprintf(`You are a Cyber Security Agent for Shopee Cross-Border Trade. 
Analyze the following traffic logs and identify potential Fraud or Bot Attacks.
Output a strict JSON object with a key "malicious_ips" containing an array of flagged IPs, risk level (High/Medium), and the reason.
Logs: %s`, string(logsJSON))

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "llama-3.3-70b-versatile",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"response_format": map[string]string{"type": "json_object"},
	})

	req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ API Error:", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		// 取得 AI 回傳的純文字 (JSON 字串)
		messageStr := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
		
		// 👉 加了這兩行！讓終端機印出你最喜歡的完美 JSON 格式
		fmt.Println("=== AI Security Agent Analysis Result (Powered by Groq + Go) ===")
		fmt.Println(messageStr)
		
		var aiResponse map[string]interface{}
		json.Unmarshal([]byte(messageStr), &aiResponse)
		
		if maliciousIPs, exists := aiResponse["malicious_ips"].([]interface{}); exists {
			return maliciousIPs
		}
	}
	return nil
}