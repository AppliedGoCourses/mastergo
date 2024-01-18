module github.com/AppliedGoCourses/mastergo/2_LanguageBasics/2-11_Packages/2-11-2_create_custom_packages/diceclient

go 1.16

// The dice module exists in different versions inside this repository. (Typically, an incomplete "todo" version and a "solution" version. To avoid an error message, we use a nonexisting import path here. Otherewise, we would get "conflicting replacements for the same module" error message. The "replace" directive ensures that the module is imported from a local directory. It also ensures that we can edit the module locally, and the client can pick up the changes immediately.
// THIS IS A HACK. Do not use it with your repositories. A typical Go code repository does not have this problem.
require github.com/appliedgocourses/doesnotexist/2-11-2/dice v0.3.0

replace github.com/appliedgocourses/doesnotexist/2-11-2/dice v0.3.0 => ../dice
