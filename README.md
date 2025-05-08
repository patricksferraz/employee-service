# Employee Service

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/employee-service)](https://goreportcard.com/report/github.com/patricksferraz/employee-service)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/patricksferraz/employee-service?status.svg)](https://godoc.org/github.com/patricksferraz/employee-service)

A modern, scalable microservice for employee management built with Go, featuring both REST and gRPC APIs, event-driven architecture, and comprehensive monitoring.

## 🚀 Features

- **Dual API Support**: REST and gRPC endpoints for maximum flexibility
- **Event-Driven Architecture**: Kafka integration for real-time employee updates
- **Database Management**: PostgreSQL with GORM for robust data persistence
- **API Documentation**: Swagger/OpenAPI integration
- **Monitoring**: Elastic APM integration for performance monitoring
- **Containerized**: Docker and Docker Compose support
- **Kubernetes Ready**: Includes K8s deployment configurations
- **Environment Configuration**: Flexible configuration management with Viper
- **Input Validation**: Comprehensive validation using govalidator
- **CORS Support**: Built-in CORS middleware
- **Database Administration**: PGAdmin included for database management
- **Kafka Management**: Confluent Control Center for Kafka monitoring

## 🛠️ Tech Stack

- **Language**: Go 1.16+
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Message Broker**: Apache Kafka
- **API Documentation**: Swagger/OpenAPI
- **Monitoring**: Elastic APM
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Configuration**: Viper
- **Validation**: govalidator

## 📋 Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Make (for using Makefile commands)
- PostgreSQL client (optional)
- Kafka client (optional)

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/patricksferraz/employee-service.git
   cd employee-service
   ```

2. Copy the environment file and configure it:
   ```bash
   cp .env.example .env
   ```

3. Start the services using Docker Compose:
   ```bash
   make up
   ```

4. The service will be available at:
   - REST API: http://localhost:8080
   - gRPC: localhost:50051
   - PGAdmin: http://localhost:9000
   - Kafka Control Center: http://localhost:9021

## 🏗️ Project Structure

```
.
├── application/     # Application layer (use cases)
├── cmd/            # Command line interface
├── domain/         # Domain models and interfaces
├── infrastructure/ # Infrastructure implementations
├── k8s/           # Kubernetes configurations
├── utils/         # Utility functions
└── .docker/       # Docker related files
```

## 🔧 Configuration

The service can be configured using environment variables or a configuration file. See `.env.example` for all available options.

## 📚 API Documentation

Once the service is running, you can access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

## 🧪 Testing

Run the tests using:
```bash
make test
```

## 🐳 Docker

Build and run the service using Docker:
```bash
make build
make up
```

## ☸️ Kubernetes

Deploy to Kubernetes using the configurations in the `k8s/` directory:
```bash
kubectl apply -f k8s/
```

## 📈 Monitoring

The service integrates with Elastic APM for monitoring. Access the APM dashboard to view:
- Request latency
- Error rates
- Transaction traces
- Service maps

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Patrick Ferraz** - *Initial work* - [patricksferraz](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Confluent Kafka](https://www.confluent.io/)
- [Elastic APM](https://www.elastic.co/apm)
