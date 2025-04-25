# Notify - Distributed Notification Delivery Service (Go + gRPC + GraphQL + Prometheus)

This project simulates a **miniature version of how social media platforms (like Instagram or Twitter)** notify users when someone they follow creates a new post.

It uses:

- ✅ **gRPC** to simulate users publishing posts  
- ✅ **Background worker queue** to dispatch notifications  
- ✅ **Retry logic** for failed deliveries  
- ✅ **GraphQL API** to fetch a user’s notifications  
- ✅ **Prometheus metrics** to monitor performance and failures  

---

## 📦 Project Features Overview

| Feature                        | Description |
|-------------------------------|-------------|
| gRPC API                      | Simulates publishing a post (`PublishPost`) |
| Notification Queue            | In-memory delivery queue for notifications |
| Worker Pool                   | Concurrent workers simulate delivery |
| Retry Logic                   | Retry 3 times for failed notifications (10% simulated failures) |
| GraphQL API                   | Fetch recent notifications for a user |
| Prometheus Metrics            | Monitor sent, failed notifications, and delivery times |
| In-Memory Store               | All data (posts, followers, notifications) is stored temporarily in Go maps |

---

## ⚙️ How to Run (Locally Without Docker)

### ✅ Step 1: Start the System

```bash
make run
```

- gRPC Server → `localhost:5050`  
- GraphQL Server → `localhost:8081`  
- Prometheus Metrics → `localhost:8081/metrics`

---

### ✅ Step 2: Publish a Post (via gRPC)

```bash
grpcurl -plaintext -d '{"postId":"p1","authorId":"user1","content":"Hello world!"}' localhost:5050 post.PostService/PublishPost
```

🔄 Internally:

- Finds all followers of `user1` (from mock data)
- Creates and queues notifications
- Background workers pick from the queue and deliver them
- Logs show success or failure

---

### ✅ Step 3: Query Notifications (via GraphQL)

URL:  
```
http://localhost:8081/query
```

Example query:
```graphql
{
  notifications(userId: "user2")
}
```

✅ Output:
```json
{
  "data": {
    "notifications": ["New post from user1: Hello world!"]
  }
}
```

---

## 📊 Step 4: Monitor with Prometheus (Manual)

Visit raw metrics endpoint:
```
http://localhost:8081/metrics
```

Metrics:

- `notifications_sent_total`
- `notifications_failed_total`
- `notification_delivery_duration_seconds`

---

## 🖥️ Optional: Prometheus UI (Manual Mode)

### 📥 Download Prometheus

From: [https://prometheus.io/download/](https://prometheus.io/download/)

### 📁 Create `prometheus.yml`

```yaml
global:
  scrape_interval: 5s

scrape_configs:
  - job_name: 'notify_service'
    static_configs:
      - targets: ['localhost:8081']
```

### ▶️ Start Prometheus

```bash
./prometheus --config.file=prometheus.yml
```

Open Prometheus UI:  
```
http://localhost:9090
```

---

## 🐳 Docker + Prometheus Monitoring

Run the whole system using Docker Compose.

---

### ✅ Step 1: Build and Run with Docker Compose

```bash
docker-compose up --build
```

Runs:

- 🟣 `notify` (Go app):
  - gRPC: `localhost:5050`
  - GraphQL + Metrics: `localhost:8081`
- 🟡 `prometheus`:
  - UI: `localhost:9090`

---

### ✅ Step 2: gRPC Call (Inside Docker)

```bash
grpcurl -plaintext -d '{"postId":"p1","authorId":"user1","content":"Hello from Docker!"}' localhost:5050 post.PostService/PublishPost
```

---

### ✅ Step 3: Query Notifications (GraphQL)

URL:
```
http://localhost:8081/query
```

Query:
```graphql
{
  notifications(userId: "user2")
}
```

---

### 📊 Step 4: Prometheus Metrics

Open Prometheus UI:

```
http://localhost:9090
```

Query metrics like:

- `notifications_sent_total`
- `notifications_failed_total`
- `notification_delivery_duration_seconds`

> 💡 Scraping happens every 5s.

---

### 🧱 Docker File Structure

#### `Dockerfile`

```dockerfile
FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o notify main.go

EXPOSE 5050
EXPOSE 8081

CMD ["./notify"]
```

#### `docker-compose.yml`

```yaml
services:
  notify:
    build: .
    ports:
      - "5050:5050"
      - "8081:8081"

  prometheus:
    image: prom/prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./prome_ui/prometheus.yml:/etc/prometheus/prometheus.yml
```

#### `prometheus.yml`

```yaml
global:
  scrape_interval: 5s

scrape_configs:
  - job_name: 'notify-app'
    static_configs:
      - targets: ['notify:8081']
```

---

## 🧰 Under the Hood – Explained Simply

### 🟣 gRPC (PublishPost)

```protobuf
rpc PublishPost(Post) returns (NotificationResponse);
```

Simulates user post publishing and returns how many notifications were queued.

---

### 🔁 Queue + Worker Dispatcher

- Adds notifications to a channel (queue)
- Worker goroutines consume the queue
- Each "delivers" (logs) a notification
- 10% chance of simulated failure → retried 3 times with backoff

---

### 🧠 In-Memory Data

Used maps to simulate DBs:

```go
map[string][]string        // followers
map[string][]Notification  // user notifications
```

---

### 📈 Prometheus Metrics Tracked

| Metric                                | Description |
|---------------------------------------|-------------|
| `notifications_sent_total`           | # delivered |
| `notifications_failed_total`         | # failed    |
| `notification_delivery_duration_seconds` | Timing histogram |

---


![alt text](image.png)

![alt text](image-1.png)