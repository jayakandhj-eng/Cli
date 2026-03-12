## Indobase-stop

Stops the Indobase local development stack.

Requires `Indobase/config.toml` to be created in your current working directory by running `Indobase init`.

All Docker resources are maintained across restarts.  Use `--no-backup` flag to reset your local development data between restarts.

Use the `--all` flag to stop all local Indobase projects instances on the machine. Use with caution with `--no-backup` as it will delete all Indobase local projects data.
