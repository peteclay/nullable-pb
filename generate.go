package nullablepb

// Update buf build modules
//go:generate buf dep update
// Generate files
//go:generate buf generate
// Generate self-contained descriptor / protoset
//go:generate buf build -o nullable.protoset
// Ensure we've got dependencies added to the go.mod file so mockery can run
//go:generate go mod tidy
// Generate mocks
//go:generate mockery
// Ensure we've got mockery's dependencies in the go.mod file too
//go:generate go mod tidy
