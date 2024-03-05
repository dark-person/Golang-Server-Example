# Simple Example of Server

Using `Go`, `vite`, `react` with `typescript`.

This example provide the following features:

-   `Makefile` with frequently use commands
-   Graceful shutdown of backend service by `Ctrl + C`
-   Predefined formatter for `vscode`
-   Dictionaries for `vscode` code spell extensions

## Build & Run

You are suggested to use `Makefile` for development.

If you are using `Window` OS, and would like to use `make`, consider using `Chocolatey` ([link](https://chocolatey.org/install)) command:

```
choco install make
```

### Dependencies

Install all dependencies of `npm`:

```
make install
```

OR reinstall everything:

```
make reinstall
```

### Development Server

Run Development Server:

```
make dev
```

Running with golang server:

```
make run
```

### Build

Build executable of `go`, with static web file built:

```
make build
```
