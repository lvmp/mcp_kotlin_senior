# Kotlin Senior MCP Server

An advanced Model Context Protocol (MCP) server designed to act as a **Senior Backend Kotlin Developer**. 

This tool provides AI agents with specialized capabilities to analyze, generate, and optimize Kotlin codebases, focusing on enterprise-grade architecture, best practices, and performant design.

## Features

### 🏗️ Architecture Analysis
- **`analyze_architecture`**: Analyzes project structures and suggests improvements aligning with:
  - **Clean Architecture**
  - **Hexagonal Architecture (Ports & Adapters)**
  - **Modular Monoliths**
  - **Microservices**

### 🎨 Design Patterns
- **`generate_design_pattern`**: Generates idiomatic Kotlin implementations of classic GoF patterns (Singleton, Strategy, Observer, Factory, etc.), ensuring thread safety and leveraging Kotlin features (e.g., `object`, `sealed class`, Coroutines).

### ✅ Best Practices & Code Quality
- **`check_best_practices`**: Scans code snippets for anti-patterns (e.g., `!!`, `GlobalScope`), suggests SOLID improvements, and enforces optimization techniques.

### 🧪 Testing
- **`generate_test_template`**: Creates robust test templates using **JUnit 5** and **MockK**.
  - Supports **Unit Tests** (fast, mocked).
  - Supports **Integration Tests** (Spring Boot context loading).

### ☁️ Cloud & DevOps
- **`suggest_cloud_solution`**: tailored recommendations for **Google Cloud Platform (GCP)**.
  - Compute: Cloud Run vs GKE.
  - Database: Cloud SQL (PostgreSQL) vs Firestore.
  - CI/CD pipelines and container strategies.

## Technology Stack

- **Language**: Go (Golang) - Migrated from TypeScript for superior performance, binary distribution, and minimal footprint.
- **SDK**: `github.com/mark3labs/mcp-go`
- **Distribution**: Docker (Extreme Scratch/UPX Optimization)

## 🚀 Docker Optimization

The Go implementation is built for extreme efficiency:
- **Build Stage**: Multi-stage build with static binary compilation.
- **Compression**: UPX extreme compression (-9).
- **Final Stage**: `FROM scratch` (0 bytes base image).
- **Result**: Final image size is approximately **3.6MB** (Reduced from ~227MB in the TS version).

## Usage

### Docker (Recommended)

1. **Pull the Image**
   ```bash
   docker pull lvmp7/mcp-kotlin-senior:latest
   ```

2. **Run via MCP**
   Add to your MCP client configuration (e.g., Claude Desktop, IDE extensions):
   ```json
   {
     "mcpServers": {
       "kotlin-senior": {
         "command": "docker",
         "args": ["run", "-i", "--rm", "lvmp7/mcp-kotlin-senior:latest"]
       }
     }
   }
   ```

## Development

1. **Prerequisites**
   - Go 1.23+

2. **Initialize & Build**
   ```bash
   go mod download
   go build -o mcp-server main.go
   ```

3. **Run Locally**
   ```bash
   ./mcp-server
   ```

4. **Deploy (Docker)**
   Use the provided PowerShell script:
   ```powershell
   .\push_docker.ps1
   ```

## Versions
- **v1.0.0**: Original TypeScript implementation.
- **v2.0.0**: Migrated to Go with extreme Docker optimization (Current).
- **latest**: Points to the latest Go implementation.

## License

MIT

