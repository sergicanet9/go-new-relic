# go-new-relic

Backend web application built with the **Go** programming language. The application exposes two simple RESTful endpoints and is instrumented with **New Relic's APM and Infrastructure agents** to provide comprehensive monitoring.

---

## 🚀 Features

* **Go & HTTP**: A modern backend build with Go and its `net/http` standard library.
* **Docker & Docker Compose**: Containerization of the application and the New Relic Infrastructure Agent for easy deployment.
* **New Relic APM Agent**: Manual instrumentation with the Go Agent to monitor transaction performance, logs, errors, and latency.
* **New Relic Infrastructure Agent**: The Infrastructure Agent monitors the health of the container and host, including CPU, memory, and disk usage.
* **Unit Tests**: The project includes unit tests for the API endpoints.
* **Makefile**: Automation of common tasks like building and running the project.

---

## 🛠️ Prerequisites

Make sure you have the following software installed on your system:

* **Go (1.21+)**
* **Docker & Docker Compose**
* **Git**

---

## 🏁 Getting Started

Follow these steps to get the project up and running on your local machine.

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/sergicanet9/go-new-relic.git
    cd go-new-relic
    ```

2.  **Configure New Relic:**
    Create a file named `.env` in the root of the project and add your New Relic license key.

    ```bash
    # New Relic License Key for Go Agent (APM)
    NEW_RELIC_LICENSE_KEY="your_license_key_here"

    # New Relic License Key for Infrastructure Agent
    NRIA_LICENSE_KEY="your_license_key_here"
    ```

3.  **Build and run the application:**
    The `Makefile` will handle building the Docker image, running `docker-compose`, and starting the Go server along with both New Relic agents.

    ```bash
    make run
    ```

    The application will be available at `http://localhost:8080`.

---

## ✅ Running Tests

You can run the unit tests for the application to ensure that the endpoints are working as expected.

```bash
make test
```

---

## 📦 API Endpoints

| Method | Endpoint | Description |
| :----- | :------- | :---------- |
| `GET` | `/` | Returns a simple message to verify the server is running. |
| `GET` | `/report` | A simple route that simulates a long-running transaction. Great for testing the New Relic instrumentation. |

---

## 📈 Monitoring with New Relic

This project is fully instrumented with two New Relic Agents:

  * **New Relic Go Agent (APM)**: Monitors the application itself, collecting data on web transactions, external service calls, logs and errors.
  * **New Relic Infrastructure Agent**: Monitors the underlying host and Docker container, providing visibility into resource usage like CPU and memory.

After running `make run` and making a few requests to the API endpoints, you'll be able to see the data in your New Relic account.

---

