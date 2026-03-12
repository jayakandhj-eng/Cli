## Extended man pages for CLI commands

### Build

Update [version string](https://github.com/Indobase/cli/blob/main/docs/main.go#L33) to match latest release.

```bash
go run docs/main.go > cli_v1_commands.yaml
```

### Release

1. Clone the [Indobase/Indobase](https://github.com/Indobase/Indobase) repo
2. Copy over the CLI reference and reformat using Indobase config

```bash
mv ../cli/cli_v1_commands.yaml specs/
npx prettier -w specs/cli_v1_commands.yaml
```

3. If there are new commands added, update [common-cli-sections.json](https://github.com/Indobase/Indobase/blob/master/spec/common-cli-sections.json) manually

