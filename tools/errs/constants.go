package errs

var ErrID = NewValidationField("id", "Invalid")
var NotFound = NewCustom(400, "Document not found")
var AlreadyExist = NewCustom(400, "Already exist")
var Internal = NewCustom(500, "Internal server error")
var ErrResourceExist = NewValidationField("resource", "exist")
var ErrProjectNotExist = NewCustom(404, "project")
