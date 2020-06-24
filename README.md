# lua-go-pdf

This repository contains a simple go module that renders a pdf file using chromedp (webkit engine). 
The go module gets compiled to a shared library, which can be called from Lua using the foreign function interface.

## Run
Take a look at the Makefile or run the following commands (linux only):

```shell script
make
make run
```

If nothing went wrong you should be able to find a new `invoice.pdf` at the root of your repository.

Requires a working Go installation and a Docker binary to be available on your system.


