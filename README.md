# AnimeImageAnalyser-v2

This project is a rebuild version of `AnimeImageAnalyser`.

## Difference with original version

This project will aims improve flaws in previous version:

-   better documentation for both frontend & backend code
-   better version control with github flow
-   implement complete testing for all codes
-   use `typescript` instead of `javascript`, provide better type checking
-   use web server instead of original `wails` GUI

## Feature

_to be implemented_

## Development

### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

### Building

To build a re-distributable, production mode package, use `wails build`.
