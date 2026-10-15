# set-env

Load `.env.<name>` into the **current** shell.

## Commands

```sh
se           # which env is active in this shell
se local     # apply .env.local to current shell
se staging   # apply .env.staging to current shell
```

Works if `.env.<name>` exists. New terminal starts unset.