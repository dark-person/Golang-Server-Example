# Simple Example of Server

Using `Go`, `create-react-app`, `react` with `typescript`.

Although `create-react-app` is not suggested to use anymore (use `vite`) instead, this repository can still view as example of server folder structure.

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

