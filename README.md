# Service Sales Order

A high-performance and scalable service for managing sales orders, developed in Go (Golang) with the Echo framework and
MariaDB.

## 📋 Requirements

To set up and run this service, ensure the following dependencies are installed:

- [Golang](https://go.dev/): Version 1.20 or higher.
- [Echo Framework](https://echo.labstack.com/): For building the web application.
- [MariaDB](https://mariadb.org/download/): Version 10.x for database operations.

## ⚙️ Development Setup

### Step 1: Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/SAMBPLG/service-sales-order.git
cd service-sales-order

```

### Step 2: Install Dependencies

Download and install the required dependencies:

```bash
go mod tidy
```

### Step 3: Configure Environment Variables

Create a .env file and set the required environment variables. A template is provided for convenience:

```bash
cp .env.examples .env
```

Edit the .env file to configure your application.

### Step 4: Run the Application

Run the service locally using one of the following commands:

```bash
go run cmd/web/main.go
```

### Step 5 (Optional): Run the worker

Run the worker locally using one of the following commands:

```bash
go run cmd/worker/main.go
```

## 🌍 Environment Variables

Below is a list of environment variables that can be configured for the application:

### General Configuration

| Variable                  | Description                                | Default      |
|---------------------------|--------------------------------------------|--------------|
| PORT                      | Application port                           | 8080         |
| TIMEZONE                  | Application timezone                       | Asia/Jakarta |
| TIMEOUT_GRACEFUL_SHUTDOWN | Timeout for graceful shutdown (in seconds) | 15           |

### Database Configuration (Read)

| Variable           | Description                             | Default |
|--------------------|-----------------------------------------|---------|
| MYSQL_READ_HOST    | MySQL host for read operations          |         |
| MYSQL_READ_USER    | MySQL username for read operations      |         |
| MYSQL_READ_PASS    | MySQL password for read operations      |         |
| MYSQL_READ_PORT    | MySQL port for read operations          | 3306    |
| MYSQL_READ_DBNAME  | MySQL database name for read operations |         |
| MYSQL_READ_USE_TLS | Enable TLS for MySQL read connections   | false   |

### Database Configuration (Write)

| Variable            | Description                              | Default |
|---------------------|------------------------------------------|---------|
| MYSQL_WRITE_HOST    | MySQL host for write operations          |         |
| MYSQL_WRITE_USER    | MySQL username for write operations      |         |
| MYSQL_WRITE_PASS    | MySQL password for write operations      |         |
| MYSQL_WRITE_PORT    | MySQL port for write operations          | 3306    |
| MYSQL_WRITE_DBNAME  | MySQL database name for write operations |         |
| MYSQL_WRITE_USE_TLS | Enable TLS for MySQL write connections   | false   |

### Database Connection Pooling

| Variable             | Description                              | Default |
|----------------------|------------------------------------------|---------|
| DB_MAX_OPEN_CONNS    | Maximum number of open connections       | 10      |
| DB_MAX_IDLE_CONNS    | Maximum number of idle connections       | 5       |
| DB_CONN_MAX_LIFETIME | Maximum connection lifetime (in seconds) | 300     |
| DB_LOG_LEVEL         | Log level for database operations        | info    |

### RabbitMQ Configuration

| Variable                        | Description                                     | Default |
|---------------------------------|-------------------------------------------------|---------|
| RABBITMQ_HOST                   | RabbitMQ server host                            |         |
| RABBITMQ_PORT                   | RabbitMQ server port                            | 5672    |
| RABBITMQ_USER                   | RabbitMQ username                               |         |
| RABBITMQ_PASSWORD               | RabbitMQ password                               |         |
| RABBITMQ_VHOST                  | RabbitMQ virtual host                           | /       |
| RABBITMQ_HEARTBEAT_INTERVAL     | Heartbeat interval (in seconds)                 | 60      |
| RABBITMQ_RETRY_CONNECT_INTERVAL | Retry interval for reconnect attempts (seconds) | 5       |
| RABBITMQ_MAX_RETRY_CONNECT      | Maximum retry attempts for RabbitMQ connection  | 10      |

### Table
