import datetime
import random

levels = ["INFO", "ERROR", "WARN", "DEBUG"]
messages = [
    "Heartbeat detected", "Connection failed", "Disk full", "Timeout",
    "High memory usage", "CPU usage high", "Disk nearing capacity",
    "Service restarted", "Health check passed", "Cache cleared",
    "Cache warmed", "Metrics flushed"
]

start_time = datetime.datetime(2026, 4, 24, 12, 0, 0)

with open("server_status.log", "w") as f:
    for i in range(10000):
        timestamp = (start_time + datetime.timedelta(seconds=i * 30)).strftime("%Y-%m-%d %H:%M:%S")
        level = random.choice(levels)
        message = random.choice(messages)
        server = f"srv-{random.randint(1, 10):02d}"
        
        f.write(f"{timestamp} [{level}] {message} on {server}\n")

print("Arquivo logs.txt com 10.000 linhas gerado.")
