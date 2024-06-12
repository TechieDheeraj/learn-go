# learn-go
1. https://opensource.google/projects/go?hl=en

# NOTES:

## Enviornment variables:
1. GOROOT -> Used to find GO binaries and Libraries in the system
2. GOPATH -> Workspace directory (where go code will reside)

## GOROOT:
1. src -> source files for standard libraries
2. pkg -> compiled packages used by go tools
3. bin -> executable commands 

## GOPATH:
1. src -> GO code source files (which we are developing)
2. pkg -> third party binaries and libraries 
3. bin -> executables

## Module:
1. A Module is collection of go packages. We can add a module to go project

## Package:
1. A Package is directory of .go files.

# Notes:
1. go mod tidy -> This command will remove any unused dependencies in the file
2. go.mod -> This file contains module name, go version, direct/indirect dependency to the code
3. go.sum -> This file contains all the modules imported in the project with versioning and checksum
4. go clean -modcache -> This command cleans the module cache and allows to pull the modules again
5. go mod vendor -v -> This command creates vendor folder in local and stores all the modules in the dir which are pulled from internet

# Interface:
1. In gccgo, an interface value is really a struct with three fields (libgo/runtime/interface.h): a pointer to the type descriptor for the type of the value currently stored in the interface value; a pointer to a table of functions implementing the methods for the current value; a pointer to the current value itself. 
    ## refs:
    1. https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
    2. https://research.swtch.com/interfaces
    3. https://www.airs.com/blog/archives/277

## Blogs:
1. https://swtch.com/~rsc/

