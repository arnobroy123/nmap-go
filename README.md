# 🔍 GoPortScanner

A lightweight, fast, and concurrent port scanner built in **Go**, inspired by Nmap. This tool scans a target host for open TCP ports using Go’s built-in concurrency and networking libraries.

## 📌 Features

- Scans TCP ports 1–1024 on a specified host
- Fast concurrent scanning using goroutines
- Timeout handling for unreachable ports
- Simple and clean CLI output
- Beginner-friendly code, great for learning Go networking and concurrency

---

## 🧠 How It Works

- Spawns 100 goroutines (lightweight threads)
- Sends ports (1–1024) through a channel
- Each goroutine listens for ports and attempts a TCP connection using `net.DialTimeout`
- Results (open ports) are returned through a results channel
- Finally, open ports are displayed in sorted order

---

## 🚀 Getting Started

### ✅ Prerequisites

- Go installed (version 1.17+ recommended)  
  👉 [Install Go](https://golang.org/dl/)

### 📦 Clone the Repository
