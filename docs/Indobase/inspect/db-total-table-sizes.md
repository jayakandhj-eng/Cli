# db-total-table-sizes

This command displays the total size of each table in the database. It is the sum of the values that `pg_table_size()` and `pg_indexes_size()` gives for each table. System tables inside `pg_catalog` and `information_schema` are not included.

```
                NAME               â”‚    SIZE
â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
  job_run_details                  â”‚ 395 MB
  slack_msgs                       â”‚ 648 kB
  emails                           â”‚ 640 kB
```