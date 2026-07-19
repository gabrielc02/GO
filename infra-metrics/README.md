# Project: infra-metrics
**Job Ticket ID:** #002  
**Role:** Infrastructure Architect  
**Supervising:** Head of SRE  

## 1. Objective
Build a professional CLI tool to parse raw server logs and aggregate error counts per service. This tool will serve as a foundation for our automated incident alerting.

## 2. Technical Requirements 
* **CLI Entry Point:** The tool must be executable via terminal.
* **Input:** A slice of strings representing raw server logs.
* **Filtering Logic:** 
    * Extract and keep only logs containing the substring `"[ERROR]"`.
* **Aggregation:** 
    * Count occurrences of distinct services triggering an error.
* **Output:** Output a cleanup report mapped as `map[string]int` displaying `{Service: ErrorCount}`.

## 3. Architectural Constraints
* **Package Separation:** String parsing must be decoupled from map aggregation logic.
* **Professional Structure:** 
    * Use `cmd/` for the application entry point.
    * Use `internal/` for core logic to prevent external importing.
* **Error Handling:** Graceful handling of malformed log lines.

## 4. Sample Data
```text
"2026-04-24 10:00:00 [ERROR] Connection failed on srv-01"
"2026-04-24 10:01:00 [INFO] Heartbeat detected on srv-02"
"2026-04-24 10:02:00 [ERROR] Disk full on srv-02"
"2026-04-24 10:03:00 [ERROR] Timeout on srv-01"
```
