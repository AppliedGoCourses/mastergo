module github.com/AppliedGoCourses/mastergo/3_AdvancedDataTypes/3-11_Exercise_WebService/3-11-4_Database/webservice/solution/webservice

go 1.16

// The quotes module exists in different versions inside this repository. (Typically, an incomplete "todo" version and a "solution" version. To avoid an error message, we use a nonexisting import path here. Otherewise, we would get "conflicting replacements for the same module" error message. The "replace" directive ensures that the module is imported from a local directory.
require (
	github.com/appliedgocourses/doesnotexist/3-11-4/quotes v0.1.0
	github.com/pkg/errors v0.9.1
)

replace github.com/appliedgocourses/doesnotexist/3-11-4/quotes v0.1.0 => ./quotes
