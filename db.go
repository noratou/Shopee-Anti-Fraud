package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB establishes a connection to MongoDB
func ConnectDB(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")
	return client
}

// SaveBlacklist upserts malicious IPs into MongoDB to prevent duplicate rows
func SaveBlacklist(client *mongo.Client, maliciousData []interface{}) {
	collection := client.Database("ShopeeSecurity").Collection("BlacklistIPs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	blockedCount := 0

	for _, item := range maliciousData {
		// Convert the interface item to a clean map
		doc, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		ip, exists := doc["ip"]
		if !exists {
			continue
		}

		// Filter to check if the IP already exists
		filter := bson.M{"ip": ip}
		
		// Update details, set updated timestamp
		update := bson.M{
			"$set": bson.M{
				"risk_level": doc["risk_level"],
				"reason":     doc["reason"],
				"updated_at": time.Now(),
			},
		}

		// Enable Upsert option (Insert if missing, Update if exists)
		opts := options.Update().SetUpsert(true)

		_, err := collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Printf("❌ Failed to upsert IP %v: %v", ip, err)
			continue
		}
		blockedCount++
	}

	fmt.Printf("🛡️  Processed security blacklist. Successfully synced %d unique threats to MongoDB.\n", blockedCount)
}

// IsIPBlacklisted checks if an IP already exists in the database
func IsIPBlacklisted(client *mongo.Client, ip string) bool {
	collection := client.Database("ShopeeSecurity").Collection("BlacklistIPs")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var result bson.M
	err := collection.FindOne(ctx, bson.M{"ip": ip}).Decode(&result)
	
	// If no document is found, err will be mongo.ErrNoDocuments (meaning it's safe/not blacklisted)
	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		return false // Fallback if DB query fails
	}

	return true // IP found in blacklist!
}