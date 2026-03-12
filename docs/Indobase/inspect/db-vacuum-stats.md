# db-vacuum-stats

This shows you stats about the vacuum activities for each table. Due to Postgres' [MVCC](https://www.postgresql.org/docs/current/mvcc.html) when data is updated or deleted new rows are created and old rows are made invisible and marked as "dead tuples". Usually the [autovaccum](https://Indobase.com/docs/guides/platform/database-size#vacuum-operations) process will aysnchronously clean the dead tuples.

The command lists when the last vacuum and last auto vacuum took place, the row count on the table as well as the count of dead rows and whether autovacuum is expected to run or not. If the number of dead rows is much higher than the row count, or if an autovacuum is expected but has not been performed for some time, this can indicate that autovacuum is not able to keep up and that your vacuum settings need to be tweaked or that you require more compute or disk IOPS to allow autovaccum to complete.


```
        SCHEMA        â”‚              TABLE               â”‚ LAST VACUUM â”‚ LAST AUTO VACUUM â”‚      ROW COUNT       â”‚ DEAD ROW COUNT â”‚ EXPECT AUTOVACUUM?
â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
 auth                 â”‚ users                            â”‚             â”‚ 2023-06-26 12:34 â”‚               18,030 â”‚              0 â”‚ no
 public               â”‚ profiles                         â”‚             â”‚ 2023-06-26 23:45 â”‚               13,420 â”‚             28 â”‚ no
 public               â”‚ logs                             â”‚             â”‚ 2023-06-26 01:23 â”‚            1,313,033 â”‚      3,318,228 â”‚ yes
 storage              â”‚ objects                          â”‚             â”‚                  â”‚             No stats â”‚              0 â”‚ no
 storage              â”‚ buckets                          â”‚             â”‚                  â”‚             No stats â”‚              0 â”‚ no
 Indobase_migrations  â”‚ schema_migrations                â”‚             â”‚                  â”‚             No stats â”‚              0 â”‚ no

```

