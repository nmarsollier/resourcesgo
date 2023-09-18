package resources

import "github.com/go-playground/validator/v10"

type CreateResourceRequest struct {
	ProjectId string            `json:"project" binding:"required"`
	Language  string            `json:"language" binding:"required"`
	SemVer    string            `json:"version" binding:"required"`
	Values    map[string]string `json:"values" binding:"required"`
}

func Create(request *CreateResourceRequest) (*Resource, error) {
	if err := validator.New().Struct(request); err != nil {
		return nil, err
	}
	// validar semver y projectId

	if exist, _ := findBy(request.ProjectId, request.Language, request.SemVer); exist != nil {
		return nil, ErrResourceExist
	}

	resource, err := insert(request.ProjectId, request.Language, request.SemVer, request.Values)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
