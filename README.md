
# go-api-gateway: A Simple API Gateway in Go 
This project implements a basic **API Gateway service** written in **Go**. It acts as a reverse proxy, forwarding requests to backend services based on configuration.

### Features 

- **Routing:** Routes incoming requests to backend services based on URL paths.
- **Service Registration:** Allows configuration of backend services with details like base URL and routes.
- **Authentication (placeholder):** Includes a placeholder function for implementing authentication checks.
- **Configuration Management:** Loads configuration from a YAML file using the `viper` library.
- **Rate Limiting:** Prevents excessive requests from a client IP to a particular service within a configurable window.
- **Logging**: Logs messages for service registration, request proxying, and errors.

### Getting Started

### Prerequisites:

Go installed (version 1.17 or higher recommended)
**Running the API Gateway:**

Clone the repository:

    git clone https://github.com/your-username/go-api-gateway.git

Navigate to the project directory:

    cd go-api-gateway

Build and run the service:

    go run .


This will start the API Gateway on port 8080 by default (configurable in the code).

**Configuration:**

The API Gateway loads configuration from a YAML file named `config.yaml` located in the current directory (you can adjust the file path using viper options). The configuration file should define the following settings:
Backend Services:

    rateLimitWindow: 5m  # Rate limiting window (e.g., 5 minutes)
    rateLimitCount: 10   # Maximum allowed requests per window
    
    services:
      # Service name 1
      service1:
        base_url: https://your-backend-service1.com
        routes:
          - path: /api/v1/data  # Route path within the API Gateway
      # Service name 2
      service2:
        base_url: https://your-backend-service2.com
        routes:
          - path: /api/v2/users  # Route path within the API Gateway



**Explanation:**

-   `rateLimitWindow`: Defines the duration for the rate limiting window (e.g., 5 minutes).
-   `rateLimitCount`: Sets the maximum number of requests allowed from a client IP within the rate limit window.
-   `services`: A map containing service configurations. Each service has a name and details like:
    -   `base_url`: The base URL of the backend service.
    -   `routes`: A list of route definitions for the service within the API Gateway. Each route has a `path` specifying the path segment within the API Gateway that should be forwarded to the corresponding path on the backend service.

**Backend Services:**

Backend services are registered in the configuration with their base URL and a list of routes. The routes define path mappings within the API Gateway that should be forwarded to the corresponding backend service paths. ➡️

### Development ️

This is a basic example, and further development can be done to enhance functionalities. Here are some potential areas for improvement:

-   **API Versioning:** Implement support for different API versions.
-   **Metrics and Monitoring:** Collect and expose metrics on API usage and backend service performance.
-   **Error Handling:** Implement more robust error handling for informative responses in different error scenarios.
-   **Testing:** Add unit and integration tests for better code coverage.

### License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).