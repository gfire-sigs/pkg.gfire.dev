# pkg.gfire.dev Module Index

Go modules hosted on pkg.gfire.dev

## Add a module

To add a module to the index, fork this repository, pick a unique name for your module and add a file named `modname` to the `modules` directory. The file should contain the following content:

```yaml
# full module name
root: pkg.gfire.dev/controlplane
# the VCS used to fetch the module (default: git)
vcs: git
# the URL to fetch the module from
url: https://github.com/gfire-sigs/fire-controlplane.git
# the description of the module
description: gfire control plane 
```

then create a pull request.
