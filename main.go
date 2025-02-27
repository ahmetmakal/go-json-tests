package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Log yapısı
type Log struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	IP        string    `json:"ip"`
	Timestamp time.Time `json:"timestamp"`
}

// JSON formatında örnek veri
var JsonData string = `[
	{"id": 1, "type": "click", "ip": "181.20.12.4", "timestamp": "2024-02-20T10:15:30Z"},
	{"id": 2, "type": "visit", "ip": "184.123.0.1", "timestamp": "2024-02-20T10:16:00Z"},
	{"id": 3, "type": "click", "ip": "203.0.113.5", "timestamp": "2024-02-20T10:18:30Z"},
	{"id": 4, "type": "purchase", "ip": "192.168.0.1", "timestamp": "2024-02-20T10:20:00Z"},
	{"id": 5, "type": "visit", "ip": "181.20.12.4", "timestamp": "2024-02-20T10:25:30Z"},
	{"id": 6, "type": "click", "ip": "192.168.0.2", "timestamp": "2024-02-20T10:30:00Z"},
	{"id": 7, "type": "visit", "ip": "203.0.113.5", "timestamp": "2024-02-20T10:32:30Z"},
	{"id": 8, "type": "click", "ip": "192.168.0.1", "timestamp": "2024-02-20T10:35:00Z"},
	{"id": 9, "type": "purchase", "ip": "203.0.113.5", "timestamp": "2024-02-20T10:40:30Z"},
	{"id": 10, "type": "visit", "ip": "192.168.0.1", "timestamp": "2024-02-20T10:42:00Z"},
	{"id": 11, "type": "click", "ip": "203.0.113.5", "timestamp": "2024-02-20T10:45:30Z"},
	{"id": 12, "type": "purchase", "ip": "192.168.0.1", "timestamp": "2024-02-20T10:50:00Z"},
	{"id": 13, "type": "visit", "ip": "181.20.12.4", "timestamp": "2024-02-20T10:52:30Z"},
	{"id": 14, "type": "click", "ip": "192.168.0.2", "timestamp": "2024-02-20T10:55:00Z"},
	{"id": 15, "type": "visit", "ip": "203.0.113.5", "timestamp": "2024-02-20T10:57:30Z"},
	{"id": 16, "type": "click", "ip": "192.168.0.1", "timestamp": "2024-02-20T11:00:00Z"},
	{"id": 17, "type": "purchase", "ip": "203.0.113.5", "timestamp": "2024-02-20T11:05:30Z"},
	{"id": 18, "type": "visit", "ip": "192.168.0.2", "timestamp": "2024-02-20T11:10:00Z"},
	{"id": 19, "type": "click", "ip": "181.20.12.4", "timestamp": "2024-02-20T11:15:30Z"},
	{"id": 20, "type": "visit", "ip": "192.168.0.1", "timestamp": "2024-02-20T11:17:00Z"},
	{"id": 21, "type": "click", "ip": "203.0.113.5", "timestamp": "2024-02-20T11:18:30Z"},
	{"id": 22, "type": "click", "ip": "192.168.0.2", "timestamp": "2024-02-20T11:20:00Z"},
	{"id": 23, "type": "purchase", "ip": "192.168.0.1", "timestamp": "2024-02-20T11:25:30Z"},
	{"id": 24, "type": "click", "ip": "203.0.113.5", "timestamp": "2024-02-20T11:27:00Z"},
	{"id": 25, "type": "click", "ip": "203.0.113.5", "timestamp": "2024-02-20T11:30:00Z"}
]`

func main() {

	printConsole()
}

func printConsole() {
	// JSON verisini struct dizisine dönüştürme
	var logs []Log
	err := json.Unmarshal([]byte(JsonData), &logs)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Logları ekrana yazdır
	for _, log := range logs {
		fmt.Printf("ID: %d, Type: %s, IP: %s, Timestamp: %s\n",
			log.ID, log.Type, log.IP, log.Timestamp)
	}
}
