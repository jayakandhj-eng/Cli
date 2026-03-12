## Indobase-link

Link your local development project to a hosted Indobase project.

PostgREST configurations are fetched from the Indobase platform and validated against your local configuration file.

Optionally, database settings can be validated if you provide a password. Your database password is saved in native credentials storage if available.

> If you do not want to be prompted for the database password, such as in a CI environment, you may specify it explicitly via the `Indobase_DB_PASSWORD` environment variable.

Some commands like `db dump`, `db push`, and `db pull` require your project to be linked first.

