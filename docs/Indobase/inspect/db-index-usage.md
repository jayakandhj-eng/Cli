# db-index-usage

This command provides information on the efficiency of indexes, represented as what percentage of total scans were index scans. A low percentage can indicate under indexing, or wrong data being indexed.

```
       TABLE NAME     â”‚ PERCENTAGE OF TIMES INDEX USED â”‚ ROWS IN TABLE
  â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
    user_events       â”‚                             99 â”‚       4225318 
    user_feed         â”‚                             99 â”‚       3581573
    unindexed_table   â”‚                              0 â”‚        322911
    job               â”‚                            100 â”‚         33242
    schema_migrations â”‚                             97 â”‚             0
    migrations        â”‚ Insufficient data              â”‚             0
```