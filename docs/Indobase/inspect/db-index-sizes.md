# db-index-sizes

This command displays the size of each each index in the database. It is calculated by taking the number of pages (reported in `relpages`) and multiplying it by the page size (8192 bytes).

```
              NAME              â”‚    SIZE
  â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
    user_events_index           â”‚ 2082 MB
    job_run_details_pkey        â”‚ 3856 kB
    schema_migrations_pkey      â”‚ 16 kB
    refresh_tokens_token_unique â”‚ 8192 bytes
    users_instance_id_idx       â”‚ 0 bytes
    buckets_pkey                â”‚ 0 bytes
```