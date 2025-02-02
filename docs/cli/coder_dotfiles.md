## coder dotfiles

Checkout and install a dotfiles repository from a Git URL

```
coder dotfiles [git_repo_url] [flags]
```

### Examples

```
  - Check out and install a dotfiles repository without prompts:

      $ coder dotfiles --yes git@github.com:example/dotfiles.git
```

### Options

```
  -h, --help                 help for dotfiles
      --symlink-dir string   Specifies the directory for the dotfiles symlink destinations. If empty will use $HOME.
                             Consumes $CODER_SYMLINK_DIR
  -y, --yes                  Bypass prompts
```

### Options inherited from parent commands

```
      --global-config coder   Path to the global coder config directory.
                              Consumes $CODER_CONFIG_DIR (default "~/.config/coderv2")
      --header stringArray    HTTP headers added to all requests. Provide as "Key=Value".
                              Consumes $CODER_HEADER
      --no-feature-warning    Suppress warnings about unlicensed features.
                              Consumes $CODER_NO_FEATURE_WARNING
      --no-version-warning    Suppress warning when client and server versions do not match.
                              Consumes $CODER_NO_VERSION_WARNING
      --token string          Specify an authentication token. For security reasons setting CODER_SESSION_TOKEN is preferred.
                              Consumes $CODER_SESSION_TOKEN
      --url string            URL to a deployment.
                              Consumes $CODER_URL
  -v, --verbose               Enable verbose output.
                              Consumes $CODER_VERBOSE
```

### SEE ALSO

- [coder](coder.md) -
