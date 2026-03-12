# db-locks

This command displays queries that have taken out an exclusive lock on a relation. Exclusive locks typically prevent other operations on that relation from taking place, and can be a cause of "hung" queries that are waiting for a lock to be granted.

If you see a query that is hanging for a very long time or causing blocking issues you may consider killing the query by connecting to the database and running `SELECT pg_cancel_backend(PID);` to cancel the query. If the query still does not stop you can force a hard stop by running `SELECT pg_terminate_backend(PID);`

```
     PID   â”‚ RELNAME â”‚ TRANSACTION ID â”‚ GRANTED â”‚                  QUERY                  â”‚   AGE
  â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
    328112 â”‚ null    â”‚              0 â”‚ t       â”‚ SELECT * FROM logs;                     â”‚ 00:04:20
```
