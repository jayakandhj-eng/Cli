# db-calls

This command is much like the `Indobase inspect db outliers` command, but ordered by the number of times a statement has been called.

You can use this information to see which queries are called most often, which can potentially be good candidates for optimisation.

```

                        QUERY                      â”‚ TOTAL EXECUTION TIME â”‚ PROPORTION OF TOTAL EXEC TIME â”‚ NUMBER CALLS â”‚  SYNC IO TIME
  â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¼â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€
    SELECT * FROM users WHERE id = $1              â”‚ 14:50:11.828939      â”‚ 89.8%                         â”‚  183,389,757 â”‚ 00:00:00.002018
    SELECT * FROM user_events                      â”‚ 01:20:23.466633      â”‚ 1.4%                          â”‚       78,325 â”‚ 00:00:00
    INSERT INTO users (email, name) VALUES ($1, $2)â”‚ 00:40:11.616882      â”‚ 0.8%                          â”‚       54,003 â”‚ 00:00:00.000322

```

