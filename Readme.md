## 🧱 Architecture Diagram

The bot follows a **simple, clean, and scalable pipeline**, where only the **independent I/O operations** (notifications) are concurrent.

```
                ┌────────────────────┐
                │  GitHub Actions     │
                │  (Cron / Manual)    │
                └─────────┬──────────┘
                          │
                          ▼
                ┌────────────────────┐
                │   main.go (Bot)     │
                │ cmd/bot/main.go     │
                └─────────┬──────────┘
                          │
        ┌─────────────────┴─────────────────┐
        │                                   │
        ▼                                   ▼
┌────────────────────┐            ┌────────────────────┐
│ LeetCode Fetcher   │            │ Gemini AI           │
│ Daily Problem      │──────────▶ │ Explanation         │
└────────────────────┘            └────────────────────┘
                          │
                          ▼
                ┌────────────────────┐
                │ Message Formatter  │
                └─────────┬──────────┘
                          │
          ┌───────────────┴────────────────┐
          │                                │
          ▼                                ▼
┌────────────────────┐        ┌────────────────────┐
│ SMS Sender         │        │ WhatsApp Sender    │
│ (Twilio)           │        │ (Twilio Sandbox)   │
└─────────▲──────────┘        └─────────▲──────────┘
          │                                │
          └─────────── Concurrent ────────┘
                    (goroutines + WaitGroup)
```

### Key Architecture Decisions

* **Sequential**

  * LeetCode fetch
  * AI explanation generation
    (data-dependent steps)

* **Concurrent**

  * SMS notification
  * WhatsApp notification
    (independent I/O-bound operations)

* **Concurrency Tools**

  * Goroutines
  * `sync.WaitGroup`
  * Buffered error channel

---

## ▶️ Local Run Instructions

You can run the bot **locally** without GitHub Actions.

---

### 1️⃣ Prerequisites

* Go **1.22+**
* Twilio account (trial is fine)
* Google Gemini API key

---

### 2️⃣ Clone the repository

```bash
git clone https://github.com/<your-username>/Leetcode-daily-bot.git
cd Leetcode-daily-bot
```

---

### 3️⃣ Create `.env` file (local only)

```env
TWILIO_AUTH_TOKEN=
FromTwilioNumber=
ToTwilioNumber=
TWILIO_ACCOUNT_SID=
SMS_MODE=
#  SMS_MODE=
GEMINI_API_KEY=
# WAPP_MODE=
WAPP_MODE=
```

> ⚠️ `.env` is ignored by Git and should never be committed.

---

### 4️⃣ Download dependencies

```bash
go mod download
```

---

### 5️⃣ Run the bot

```bash
go run ./cmd/bot
```

---

### 6️⃣ Expected behavior

* Fetches the **LeetCode Daily Challenge**
* Generates AI explanation
* Sends message via:

  * SMS
  * WhatsApp
    (depending on `MESSAGE_CHANNEL`)

---

### 🧪 Fake Mode (recommended for testing)

To avoid burning Twilio credits:

```env
MESSAGE_CHANNEL=fake
```

This runs the entire pipeline **without sending messages**.

---

## 🧠 Why `go run ./cmd/bot`?

* Keeps `main.go` isolated
* Allows multiple entry points in future
* Follows **standard Go project layout**
* Works cleanly in **CI/CD and local runs**

---

## 🟢 Summary

* Architecture is **simple, modular, and scalable**
* Concurrency is used **only where it adds value**
* Local and CI environments behave consistently

This makes the project **production-ready and interview-ready**.
