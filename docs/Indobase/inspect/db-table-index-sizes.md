# db-table-index-sizes

This command displays the total size of indexes for each table. It is calculated by using the system administration function `pg_indexes_size()`.

```
                 TABLE               â”‚ INDEX SIZE
  â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
    job_run_details                  â”‚ 10104 kB
    users                            â”‚ 128 kB
    job                              â”‚ 32 kB
    instances                        â”‚ 8192 bytes
    http_request_queue               â”‚ 0 bytes
```
