## lagoon delete user

Delete a user from Lagoon

### Synopsis

Delete a user from Lagoon

```
lagoon delete user [flags]
```

### Options

```
  -E, --email string   Email address of the user
  -h, --help           help for user
```

### Options inherited from parent commands

```
      --config-file string   Path to the config file to use (must be *.yml or *.yaml)
      --debug                Enable debugging output (if supported)
  -e, --environment string   Specify an environment to use
      --force                Force yes on prompts (if supported)
  -l, --lagoon string        The Lagoon instance to interact with
      --no-header            No header on table (if supported)
      --output-csv           Output as CSV (if supported)
      --output-json          Output as JSON (if supported)
      --pretty               Make JSON pretty (if supported)
  -p, --project string       Specify a project to use
      --skip-update-check    Skip checking for updates
  -i, --ssh-key string       Specify path to a specific SSH key to use for lagoon authentication
```

### SEE ALSO

* [lagoon delete](lagoon_delete.md)	 - Delete a project, or delete notifications and variables from projects or environments

