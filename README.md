# Notify - Distributed Notification Delivery Service (Go + gRPC + GraphQL)

This is a simple, beginner-friendly project built in **Go (Golang)** that simulates how platforms like Instagram or Twitter notify users when someone they follow posts new content. It uses **gRPC for publishing posts**, a **background queue with retry logic** for dispatching notifications, and a **GraphQL API** to retrieve those notifications.

---

## 🧠 What You’ll Learn

- ✅ How distributed notification systems work
- ✅ gRPC for efficient communication
- ✅ Queues and background workers (using Go routines)
- ✅ Retrying failed notifications
- ✅ In-memory data storage and retrieval
- ✅ GraphQL API (via [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)) for querying data

---

## ⚙️ How It Works

### ✅ Step 1: Start the Servers

```bash
make run
```

- gRPC Server starts on port `:5050`
- GraphQL Server starts on port `:8080`

### ✅ Step 2: Simulate a Post (via gRPC)

Use [grpcurl](https://github.com/fullstorydev/grpcurl) to simulate a post:

```bash
grpcurl -plaintext -d '{"postId":"p1","authorId":"user1","content":"Hello world!"}' localhost:5050 post.PostService/PublishPost
```

- The server finds followers of `user1`
- Queues notification for each follower.
- Background workers send notifications.
- Notifications are saved in memory.

### ✅ Step 3: View Notifications (via GraphQL)

Open Postman (select POST method and then Body -> GraphQL) and hit:
```
http://localhost:8080/query
```

Then send this GraphQL query:

```graphql
{
  notifications(UserId: "user2")
}
```

If `user2` is a follower of `user1`, they will see:
```json
{
  "data": {
    "notifications": ["New post from user1: Hello world!"]
  }
}
```

---

## 🛠 Under the Hood (Concepts Explained Simply)

### gRPC (Post Publishing)

We use gRPC (faster alternative to REST) to send posts:
```go
rpc PublishPost(Post) returns (NotificationResponse);
```

### Queue + Dispatcher (Notification Engine)

- All notifications go into a queue.
- Background **workers** pick from this queue.
- If sending fails, it retries 3 times with delay.

### In-Memory Store (Temporary Data)

Notifications are stored in memory using a map like:
```go
map[userID][]string
```

This helps us retrieve them later via GraphQL.

### GraphQL (Query Layer)

Want to get all notifications for a user? Use GraphQL:
```graphql
{
  notifications(userId: "user3")
}
```

It only returns the notifications you care about—fast and clean!


## 🔥 Want to Try More?

- Add persistent storage (e.g., SQLite or Redis)
- Show notification status (delivered, failed)
- Add real-time WebSocket for live updates.
