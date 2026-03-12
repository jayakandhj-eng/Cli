# db-seq-scans

This command displays the number of sequential scans recorded against all tables, descending by count of sequential scans. Tables that have very high numbers of sequential scans may be underindexed, and it may be worth investigating queries that read from these tables.


```
                  NAME               â”‚ COUNT
  â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€
    emails                           â”‚ 182435
    users                            â”‚  25063
    job_run_details                  â”‚     60
    schema_migrations                â”‚      0
    migrations                       â”‚      0
```

