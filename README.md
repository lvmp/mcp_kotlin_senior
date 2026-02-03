# Kotlin Senior MCP Server

An advanced Model Context Protocol (MCP) server designed to act as a **Senior Backend Kotlin Developer**. 

This tool provides AI agents with specialized capabilities to analyze, generate, and optimize Kotlin codebases, focusing on enterprise-grade architecture, best practices, and performant design.

## Features

### 🏗️ Architecture Analysis
- **`analyze_architecture`**: Analyzes project structures and suggests improvements aligning with:
  - **Clean Architecture**
  - **Hexagonal Architecture (Ports & Adapters)**
  - **Modular Monoliths**

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

- **Language**: TypeScript (Node.js) - Chosen for high I/O performance, low latency, and lightweight resource footprint compared to a JVM-based MCP server.
- **Runtime**: Node.js 20+
- **Distribution**: Docker (Alpine/Distroless)

## Usage

### Docker (Recommended)

1. **Build the Image**
   ```bash
   docker build -t mcp-kotlin-senior .
   ```

2. **Run via MCP**
   Add to your MCP client configuration (e.g., Claude Desktop, IDE extensions):
   ```json
   {
     "mcpServers": {
       "kotlin-senior": {
         "command": "docker",
         "args": ["run", "-i", "--rm", "mcp-kotlin-senior"]
       }
     }
   }
   ```

## Development

1. **Install Dependencies**
   ```bash
   npm install
   ```

2. **Build**
   ```bash
   npm run build
   ```

3. **Run Locally**
   ```bash
   npm start
   ```

## License

MIT
