## db-bloat

This command displays an estimation of table "bloat" - Due to Postgres' [MVCC](https://www.postgresql.org/docs/current/mvcc.html) when data is updated or deleted new rows are created and old rows are made invisible and marked as "dead tuples". Usually the [autovaccum](https://Indobase.com/docs/guides/platform/database-size#vacuum-operations) process will asynchronously clean the dead tuples. Sometimes the autovaccum is unable to work fast enough to reduce or prevent tables from becoming bloated. High bloat can slow down queries, cause excessive IOPS and waste space in your database.

Tables with a high bloat ratio should be investigated to see if there are vacuuming is not quick enough or there are other issues.

```
    TYPE  â”‚ SCHEMA NAME â”‚        OBJECT NAME         â”‚ BLOAT â”‚ WASTE
  â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
    table â”‚ public      â”‚ very_bloated_table         â”‚  41.0 â”‚ 700 MB
    table â”‚ public      â”‚ my_table                   â”‚   4.0 â”‚ 76 MB
    table â”‚ public      â”‚ happy_table                â”‚   1.0 â”‚ 1472 kB
    index â”‚ public      â”‚ happy_table::my_nice_index â”‚   0.7 â”‚ 880 kB
```

