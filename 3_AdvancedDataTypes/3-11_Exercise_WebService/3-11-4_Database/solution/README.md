# A note on the package locations

This code imports the `quotes` package from a remote import path. In the `go.mod` file, however, this import is permanently replaced by a local copy of this package, right in the `quotes` folder in this directory. 

This allows you to modify the `quotes` package if you wish, and the Go compiler will pick up the changes immediately. No need for changing the package version, publishing the package, or downloading it again.
