# Description
This is a simple project where I simulate ticket demand to apply my Go knowledge to a real-world scenario. It represents the second stage of my language-learning strategy."

# Project: infra-audit
**Job Ticket ID:** #001  
**Role:** Lead Architect  
**Supervising:** Head of SRE  

## 1. Objective
Build a professional CLI tool to identify high-risk production assets from a raw infrastructure inventory. This tool will serve as a foundation for our automated health auditing.

## 2. Technical Requirements
- **CLI Entry Point:** The tool must be executable via terminal.
- **Data Source:** Read a local `inventory.json` file.
- **Filtering Logic:** 
    - Keep only assets where `status == "production"`.
    - AND where (`cpu_usage > 80` OR `memory_usage > 80`).
- **Data Enrichment:** Calculate a `severity_level` for each filtered asset:
    - **Critical:** If both CPU and Memory usage are > 90.
    - **High:** Otherwise (but still meeting the initial filter).
- **Output:** Generate a file named `audit_report.json` containing only the filtered and enriched assets.

## 3. Architectural Constraints
- **Package Separation:** Business logic (filtering/scoring) must be decoupled from I/O logic (file reading/writing).
- **Professional Structure:** 
    - Use `cmd/` for the application entry point.
    - Use `internal/` for core logic to prevent external importing.
- **Error Handling:** Graceful handling of missing files or malformed JSON.

## 4. Sample Data (`inventory.json`)
```json
[
  {
    "id": "srv-prod-01",
    "hostname": "api-gateway-01",
    "status": "production",
    "cpu_usage": 85,
    "memory_usage": 40,
    "region": "us-east-1"
  },
  {
    "id": "srv-dev-02",
    "hostname": "test-runner-02",
    "status": "development",
    "cpu_usage": 95,
    "memory_usage": 90,
    "region": "us-west-2"
  },
  {
    "id": "srv-prod-03",
    "hostname": "db-primary-01",
    "status": "production",
    "cpu_usage": 92,
    "memory_usage": 94,
    "region": "eu-central-1"
  },
  {
    "id": "srv-prod-04",
    "hostname": "cache-node-01",
    "status": "production",
    "cpu_usage": 15,
    "memory_usage": 20,
    "region": "sa-east-1"
  }
]

REF https://github.com/golang-standards/project-layout

We got a inventory file, the main object is, create a sistem that can read from this file and apply a audit on those itens, generating a output with valid assets and invalida assets, saying why those are invalidate
/report contains the model for assets output, and audit.go that apply the filter
/io contains the reader function
/inventory contains model.go for assets input
/main is the clue of all those