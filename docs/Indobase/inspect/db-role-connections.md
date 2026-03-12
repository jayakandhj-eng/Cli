# db-role-connections

This command shows the number of active connections for each database roles to see which specific role might be consuming more connections than expected.

This is a Indobase specific command. You can see this breakdown on the dashboard as well:
https://app.Indobase.com/project/_/database/roles

The maximum number of active connections depends [on your instance size](https://Indobase.com/docs/guides/platform/compute-add-ons). You can [manually overwrite](https://Indobase.com/docs/guides/platform/performance#allowing-higher-number-of-connections) the allowed number of connection but it is not advised.

```


            ROLE NAME         â”‚ ACTIVE CONNCTION
  â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
    authenticator             â”‚                5
    postgres                  â”‚                5
    Indobase_admin            â”‚                1
    pgbouncer                 â”‚                1
    anon                      â”‚                0
    authenticated             â”‚                0
    service_role              â”‚                0
    dashboard_user            â”‚                0
    Indobase_auth_admin       â”‚                0
    Indobase_storage_admin    â”‚                0
    Indobase_functions_admin  â”‚                0
    pgsodium_keyholder        â”‚                0
    pg_read_all_data          â”‚                0
    pg_write_all_data         â”‚                0
    pg_monitor                â”‚                0

Active connections 12/90

```

