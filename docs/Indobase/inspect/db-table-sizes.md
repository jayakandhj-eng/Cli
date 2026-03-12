# db-table-sizes

This command displays the size of each table in the database. It is calculated by using the system administration function `pg_table_size()`, which includes the size of the main data fork, free space map, visibility map and TOAST data. It does not include the size of the table's indexes.


```
                  NAME               â”‚    SIZE
  â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
    job_run_details                  â”‚ 385 MB
    emails                           â”‚ 584 kB
    job                              â”‚ 40 kB
    sessions                         â”‚ 0 bytes
    prod_resource_notifications_meta â”‚ 0 bytes
```