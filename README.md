# Online Shop

## Routers

### Health Check
- **Endpoint:** `GET /health-check`
    - **Response:** `OK`

### Swagger
- **Endpoint:** `GET /swagger/index.html`
- **Response:** Swagger UI with all the available endpoints


## Models Structure

```sql
users {
    id: int,
    username: varchar(50),
    email: varchar(50),
    address: varchar(50),
    registration_date: timestamp default current_timestamp,
    role: varchar(50),
}
products {
    id: int,
    name: varchar(50),
    description: text,
    price: numeric,
    category: varchar(50),
    quantity: int,
}
orders {
    id: int,
    user_id: int,
    total_price: numeric,
    order_date: timestamp default current_timestamp,
    status: varchar(50),
}
payments {
    id: int,
    order_id: int,
    user_id: int,
    payment_date: timestamp default current_timestamp,
    payment_status: varchar(50),
    amount: numeric,
}
```

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/FIFSAK/OnlineStore
   cd OnlineStore
   ```
2. **Set up .env file like .env.example:**


3. **Build the Docker images:**
   ```bash
   make build
   ```
4. **Start the Docker containers:**
   ```bash
   make up
   ```
5. **Check the health of the server:**
   Open your browser and go to http://localhost:8080/health-check to ensure the server is running properly.


6. **Stop the Docker containers:**
   ```bash
   make down
   ```

**LINK: https://projectmanagementservice.onrender.com**
