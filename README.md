# set-env

Load `.env.<name>` into the **current** shell.

Works on **zsh** (and bash). The binary prints `export` lines; a shell function `eval`s them.

## Setup

Needs [Go](https://go.dev/dl/).

**1. Build the binary onto your PATH**

```sh
git clone <this-repo>
cd set-env
go build -o ~/.local/bin/set-env
```

`~/.local/bin` must be on `PATH`. If it is not, add this to `~/.zshrc` (or `~/.bashrc`):

```sh
export PATH="$HOME/.local/bin:$PATH"
```

**2. Add the function** so `se` applies vars in *this* shell.

zsh — `~/.zshrc`:

```zsh
se() {
  eval "$(command set-env "$@")"
}
```

bash — `~/.bashrc`, same function.

**3. Reload**

```sh
source ~/.zshrc    # bash: source ~/.bashrc
```

## Use

In a directory that has the files:

```
.env.*
```

```sh
se              # which env is active (or "ENV not set yet!")
se local        # apply .env.local to this shell
se staging      # apply .env.staging to this shell
```

Any name works if `.env.<name>` exists.