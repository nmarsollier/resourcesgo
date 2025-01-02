package errs

var NotFound = NewCustom(400, "Document not found")
var AlreadyExist = NewCustom(400, "Already exist")
var Internal = NewCustom(500, "Internal server error")
var ErrProjectNotExist = NewCustom(404, "projectId")
var ErrLanguageNotExist = NewCustom(404, "languageId")
