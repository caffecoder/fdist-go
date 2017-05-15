Simple GO file distribution library.
========================================

## Description:

Simple library that allows organize distribution of files within hex based tree.

## Install package:

```shell
go get -u github.com/xcdr/fdist-go
```

## Example of usage:

```go
import fdist github.com/xcdr/fdist-go

var fd = fdist.NewFileDistribution("/tmp/storage")

// default extensions
databaseID := 1
fd.HexPath(databaseID)
fd.RenameFrom("/tmp/upload/file1.txt")
path := fd.GetPath() // saved file path

// set all extensions to .pdf
databaseID = 256
fd.SetExtension(".pdf")
fd.HexPath(databaseID)
fd.RenameFrom("/tmp/upload/file2.txt")
path = fd.GetPath() // saved file path
```

Files should be stored in /tmp/storage/01.dat and /tmp/storage/01/00.pdf
