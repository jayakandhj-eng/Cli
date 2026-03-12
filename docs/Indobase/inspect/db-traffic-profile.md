# db-traffic-profile

This command analyzes table I/O patterns to show read/write activity ratios based on block-level operations. It combines data from PostgreSQL's `pg_stat_user_tables` (for tuple operations) and `pg_statio_user_tables` (for block I/O) to categorize each table's workload profile.


The command classifies tables into categories:
- **Read-Heavy** - Read operations are more than 5x write operations (e.g., 1:10, 1:50)
- **Write-Heavy** - Write operations are more than 20% of read operations (e.g., 1:2, 1:4, 2:1, 10:1)
- **Balanced** - Mixed workload where writes are between 20% and 500% of reads
- **Read-Only** - Only read operations detected
- **Write-Only** - Only write operations detected

```
SCHEMA â”‚ TABLE        â”‚ BLOCKS READ â”‚ WRITE TUPLES â”‚ BLOCKS WRITE â”‚ ACTIVITY RATIO
â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
public â”‚ user_events  â”‚     450,234 â”‚     9,004,680â”‚       23,450 â”‚ 20:1 (Write-Heavy)
public â”‚ users        â”‚      89,203 â”‚        12,451â”‚        1,203 â”‚ 7.2:1 (Read-Heavy)
public â”‚ sessions     â”‚      15,402 â”‚        14,823â”‚        2,341 â”‚ â‰ˆ1:1 (Balanced)
public â”‚ cache_data   â”‚     123,456 â”‚             0â”‚            0 â”‚ Read-Only
auth   â”‚ audit_logs   â”‚           0 â”‚        98,234â”‚       12,341 â”‚ Write-Only
```

**Note:** This command only displays tables that have had both read and write activity. Tables with no I/O operations are not shown. The classification ratio threshold (default: 5:1) determines when a table is considered "heavy" in one direction versus balanced.

