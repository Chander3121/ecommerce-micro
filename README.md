# Ecommerce Polyglot Microservices System 🚀

A polyglot microservices-based ecommerce backend system built using multiple languages and frameworks to explore distributed architecture, inter-service communication, Docker networking, and service orchestration.

This project demonstrates how different services written in different languages can communicate with each other using REST APIs inside a Dockerized microservices environment.

---

# 🏗️ Architecture

```text
                ┌────────────────────┐
                │   Order Service    │
                │     Golang/Gin     │
                └───────┬─────┬──────┘
                        │     │
         HTTP REST      │     │ HTTP REST
                        │     │
                        ▼     ▼
              ┌────────────┐  ┌──────────────┐
              │Auth Service│  │ProductService│
              │   Rails    │  │ FastAPI/Py   │
              └─────┬──────┘  └──────┬───────┘
                    │                │
                    ▼                ▼
              PostgreSQL       PostgreSQL


                ┌────────────────────┐
                │   Future Plans     │
                │ RabbitMQ / gRPC    │
                └────────────────────┘
```

# 🧩 Services

| Service         | Tech Stack     | Responsibility                       |
| --------------- | -------------- | ------------------------------------ |
| Auth Service    | Ruby on Rails  | JWT Authentication & User Management |
| Product Service | Python FastAPI | Product & Inventory Management       |
| Order Service   | Golang + Gin   | Order Orchestration                  |
| Databases       | PostgreSQL     | Separate DB per service              |
| Infrastructure  | Docker Compose | Container Orchestration              |

# ✨ Features

✅ Authentication Service
 - User Registration
 - User Login
 - JWT Token Generation
 - Token Validation
 - Protected APIs

✅ Product Service
 - Create Products
 - Product Listing
 - Product Details
 - Stock Management

✅ Order Service
 - Order Creation
 - Product Validation
 - JWT Validation via Auth Service
 - Distributed Business Logic
 - Total Price Calculation

✅ Microservices Concepts
 - Polyglot architecture
 - Service-to-service communication
 - Separate database per service
 - Docker networking
 - REST-based orchestration
 - Distributed authentication

# 🛠️ Tech Stack
## Backend
 - Ruby on Rails 8
 - Python FastAPI
 - Golang Gin
## Database
 - PostgreSQL
## DevOps
 - Docker
 - Docker Compose
## Communication
 - REST APIs
 - Internal Docker Networking

# 📂 Project Structure
```text
ecommerce-micro/
│
├── auth-service/
│
├── product-service/
│
├── order-service/
│
└── docker-compose.yml
```

# ⚙️ Setup Instructions
## 1. Clone Repository
```bash
git clone https://github.com/Chander3121/ecommerce-micro.git

cd ecommerce-micro
```

## 2. Start All Services
```bash
docker compose up --build
```

# 🌐 Services
| Service         | URL                                            |
| --------------- | ---------------------------------------------- |
| Auth Service    | [http://localhost:3001](http://localhost:3001) |
| Product Service | [http://localhost:3002](http://localhost:3002) |
| Order Service   | [http://localhost:3003](http://localhost:3003) |

# 🔐 Authentication APIs
## Register User
```curl
curl -X POST http://localhost:3001/register \
-H "Content-Type: application/json" \
-d '{
  "name":"Test",
  "email":"test@example.com",
  "password":"password",
  "password_confirmation":"password"
}'
```
## Login User
```curl
curl -X POST http://localhost:3001/login \
-H "Content-Type: application/json" \
-d '{
  "email":"test@example.com",
  "password":"password"
}'
```

# 📦 Product APIs
## Create Product
```curl
curl -X POST http://localhost:3002/products \
-H "Content-Type: application/json" \
-d '{
  "name":"Chocolate Cake",
  "description":"Dark chocolate cake",
  "price":499,
  "stock":10
}'
```

## Get Products
```curl
curl http://localhost:3002/products
```

# 🛒 Order APIs
## Create Order
```curl
curl -X POST http://localhost:3003/orders \
-H "Authorization: Bearer YOUR_TOKEN" \
-H "Content-Type: application/json" \
-d '{
  "product_id":1,
  "quantity":2
}'
```
# 🔄 Order Flow
## Distributed Flow
```text
Client
  ↓
Order Service
  ↓
Validate JWT from Auth Service
  ↓
Fetch Product from Product Service
  ↓
Calculate Total Price
  ↓
Create Order
```
# 🐳 Dockerized Architecture
Each service runs independently inside Docker containers.

 - Independent databases
 - Internal Docker networking
 - Service discovery using container names

Example:
```text
http://auth-service:3001
http://product-service:3002
```

# 🚀 Future Improvements
 - RabbitMQ Event-Driven Architecture
 - Notification Service (Node.js)
 - gRPC Communication
 - API Gateway
 - Kubernetes Deployment
 - Distributed Logging
 - Monitoring & Observability
 - CI/CD Pipelines

# 📚 Learning Goals
This project was built to gain hands-on experience with:
 - Microservices Architecture
 - Distributed Systems
 - Polyglot Backend Development
 - Docker Networking
 - Service Orchestration
 - REST Communication
 - JWT Authentication
 - Database Per Service Pattern

# 👨‍💻 Author
## Chander Prakash

Ruby on Rails Developer exploring:
 - Golang
 - Python
 - Microservices
 - Distributed Architecture
 - Event-Driven Systems

# ⭐ Repository

If you found this project useful or interesting, feel free to star the repository 🚀
