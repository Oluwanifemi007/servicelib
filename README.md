# ServiceLib 🚀

![ServiceLib](https://img.shields.io/badge/ServiceLib-Go%20Library-brightgreen.svg)  
[![Release Version](https://img.shields.io/github/v/release/Oluwanifemi007/servicelib.svg)](https://github.com/Oluwanifemi007/servicelib/releases)  
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Welcome to **ServiceLib**, a comprehensive Go library designed to accelerate the development of robust, production-ready microservices. Whether you are building a new service from scratch or enhancing an existing one, ServiceLib provides the tools and patterns you need to create scalable, maintainable applications.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Core Concepts](#core-concepts)
  - [Adapter Pattern](#adapter-pattern)
  - [Dependency Injection](#dependency-injection)
  - [Middleware](#middleware)
  - [Error Handling](#error-handling)
  - [Logging](#logging)
  - [Telemetry](#telemetry)
- [Usage Examples](#usage-examples)
- [Health Checks](#health-checks)
- [Configuration](#configuration)
- [Database Integration](#database-integration)
- [GraphQL Support](#graphql-support)
- [Prometheus and Grafana](#prometheus-and-grafana)
- [Releases](#releases)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Microservice Patterns**: Implement common patterns like Adapter, Factory, and Repository.
- **Authentication**: Secure your services with built-in authentication mechanisms.
- **Configuration Management**: Simplify configuration with structured approaches.
- **Database Support**: Integrate with various databases easily.
- **Error Handling**: Standardize error management across services.
- **Health Checks**: Monitor service health with built-in checks.
- **Telemetry**: Collect metrics for performance monitoring.

## Installation

To install ServiceLib, use the following command:

```bash
go get github.com/Oluwanifemi007/servicelib
```

## Getting Started

To get started with ServiceLib, check the [Releases](https://github.com/Oluwanifemi007/servicelib/releases) section for the latest version. Download and execute the binaries to integrate ServiceLib into your project.

### Basic Setup

Here is a simple example to set up a microservice using ServiceLib:

```go
package main

import (
    "github.com/Oluwanifemi007/servicelib"
)

func main() {
    service := servicelib.NewService("MyService")
    service.Start()
}
```

## Core Concepts

### Adapter Pattern

The Adapter Pattern allows you to create a bridge between two incompatible interfaces. ServiceLib includes a flexible adapter system to help you integrate various components without changing their existing code.

### Dependency Injection

Dependency Injection (DI) is a technique to achieve Inversion of Control (IoC) between classes and their dependencies. ServiceLib provides a DI container to manage your service dependencies efficiently.

### Middleware

Middleware functions are essential for handling requests and responses. ServiceLib allows you to define custom middleware for logging, authentication, and more.

### Error Handling

Error handling is crucial in any application. ServiceLib provides a standardized way to manage errors across your microservices, making it easier to track and resolve issues.

### Logging

Logging is vital for monitoring and debugging. ServiceLib offers a simple logging interface that integrates with various logging backends.

### Telemetry

Collecting telemetry data helps you monitor the performance of your services. ServiceLib supports various telemetry libraries, enabling you to gather and analyze metrics easily.

## Usage Examples

Here are some usage examples to help you understand how to implement different features of ServiceLib.

### Example: Creating a Service with Middleware

```go
package main

import (
    "github.com/Oluwanifemi007/servicelib"
)

func loggingMiddleware(next servicelib.Handler) servicelib.Handler {
    return func(req servicelib.Request) servicelib.Response {
        // Log the request
        log.Printf("Received request: %v", req)
        return next(req)
    }
}

func main() {
    service := servicelib.NewService("MyService")
    service.Use(loggingMiddleware)
    service.Start()
}
```

### Example: Implementing Health Checks

```go
package main

import (
    "github.com/Oluwanifemi007/servicelib"
)

func main() {
    service := servicelib.NewService("MyService")
    service.AddHealthCheck("/health", func() bool {
        return true // Replace with actual health check logic
    })
    service.Start()
}
```

## Health Checks

Health checks are crucial for ensuring that your services are running correctly. ServiceLib provides a simple way to implement health checks. You can define endpoints that return the health status of your service.

## Configuration

ServiceLib supports various configuration formats, including JSON and YAML. You can easily load configurations from files or environment variables.

### Example: Loading Configuration

```go
package main

import (
    "github.com/Oluwanifemi007/servicelib"
)

func main() {
    config := servicelib.LoadConfig("config.yaml")
    service := servicelib.NewService(config.ServiceName)
    service.Start()
}
```

## Database Integration

ServiceLib supports integration with multiple databases, including SQL and NoSQL databases. You can define repositories for data access, ensuring clean separation of concerns.

### Example: Using a Repository

```go
package main

import (
    "github.com/Oluwanifemi007/servicelib"
)

type User struct {
    ID   int
    Name string
}

type UserRepository struct {
    db servicelib.Database
}

func (r *UserRepository) GetUser(id int) (*User, error) {
    // Implement database logic here
}

func main() {
    service := servicelib.NewService("MyService")
    userRepo := &UserRepository{db: service.Database}
    user, err := userRepo.GetUser(1)
    if err != nil {
        log.Println(err)
    }
}
```

## GraphQL Support

ServiceLib provides built-in support for GraphQL, making it easy to create APIs that allow clients to request only the data they need.

### Example: Setting Up a GraphQL Server

```go
package main

import (
    "github.com/Oluwanifemi007/servicelib"
)

func main() {
    service := servicelib.NewService("MyService")
    service.UseGraphQL("/graphql")
    service.Start()
}
```

## Prometheus and Grafana

Integrate with Prometheus and Grafana for monitoring and visualization. ServiceLib supports exporting metrics to Prometheus, allowing you to visualize service performance in Grafana.

### Example: Exposing Metrics

```go
package main

import (
    "github.com/Oluwanifemi007/servicelib"
)

func main() {
    service := servicelib.NewService("MyService")
    service.ExposeMetrics("/metrics")
    service.Start()
}
```

## Releases

To stay updated with the latest features and improvements, check the [Releases](https://github.com/Oluwanifemi007/servicelib/releases) section. Download and execute the binaries to keep your project up to date.

## Contributing

Contributions are welcome! If you would like to contribute to ServiceLib, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/YourFeature`).
6. Open a pull request.

## License

ServiceLib is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.