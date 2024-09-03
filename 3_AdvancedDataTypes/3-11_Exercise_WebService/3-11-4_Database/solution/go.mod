module github.com/AppliedGoCourses/mastergo/3_AdvancedDataTypes/3-11_Exercise_WebService/3-11-4_Database/webservice/solution/webservice

go 1.23.0

require (
	github.com/appliedgocourses/quotes v0.0.0-20240903134310-ce9321a5c65b
	// The quotes module exists in different versions inside this repository. (Typically, an incomplete "todo" version and a "solution" version. To avoid an error message, we use a nonexisting import path here. Otherewise, we would get "conflicting replacements for the same module" error message. The "replace" directive ensures that the module is imported from a local directory.
	github.com/pkg/errors v0.9.1
)

require (
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	go.etcd.io/bbolt v1.3.11 // indirect
	go.etcd.io/gofail v0.1.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/appliedgocourses/doesnotexist/3-11-4/quotes v0.1.0 => ./quotes
