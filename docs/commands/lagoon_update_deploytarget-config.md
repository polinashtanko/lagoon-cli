## lagoon update deploytarget-config

Update a deploytarget config

### Synopsis

Update a deploytarget config

```
lagoon update deploytarget-config [flags]
```

### Options

```
  -b, --branches string       Branches regex
  -d, --deploytarget uint     Deploytarget id
  -h, --help                  help for deploytarget-config
  -I, --id uint               Deploytarget config id
  -P, --pullrequests string   Pullrequests title regex
  -w, --weight uint           Deploytarget config weighting (default:1) (default 1)
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

* [lagoon update](lagoon_update.md)	 - Update a resource

