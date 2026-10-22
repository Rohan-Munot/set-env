# set-env

Load `.env.<name>` into the **current** shell.

Works on **zsh** (and bash). The binary prints `unset`/`export` lines; a shell function `eval`s them **only if the binary succeeds**.

## Setup

Needs [Go](https://go.dev/dl/).

**1. Build the binary onto your PATH**

```sh
git clone <this-repo>
cd set-env
mkdir -p ~/.local/bin
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
  case ${1:-} in
    current|help)
      command set-env "$@"
      ;;
    *)
      local out
      out=$(command set-env "$@") || return 1
      eval "$out"
      ;;
  esac
}
```

bash — `~/.bashrc`, same function.

**3. Reload**

```sh
source ~/.zshrc    # bash: source ~/.bashrc
```

`se` is the command you type. `set-env` is the binary. After changing `main.go`, rebuild:

```sh
go build -o ~/.local/bin/set-env
```

## Use

In a directory that has the files:

```
.env.*
```

```sh
se current      # which env is active in this shell
se local        # apply .env.local to this shell
se staging      # apply .env.staging to this shell
se help
```

Any name works if `.env.<name>` exists.