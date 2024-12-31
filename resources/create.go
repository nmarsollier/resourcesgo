package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/errs"
	"github.com/nmarsollier/resourcesgo/tools/logx"
)

type CreateResourceRequest struct {
	ProjectId string            `json:"project" binding:"required"`
	Language  string            `json:"language" binding:"required"`
	SemVer    string            `json:"version" binding:"required"`
	Values    map[string]string `json:"values" binding:"required"`
}

func Create(logenv logx.Fields, request *CreateResourceRequest) (*Resource, error) {
	if err := validator.New().Struct(request); err != nil {
		return nil, err
	}

	resource, err := insert(logenv, request.ProjectId, request.Language, request.SemVer, request.Values)
	if err != nil {
		switch db.ErrorCode(err) {
		case db.ERR_FOREIGN_KEY:
			return nil, errs.ErrProjectNotExist
		case db.ERR_EXIST:
			return nil, errs.AlreadyExist
		}
		return nil, err
	}

	return resource, nil
}

func insert(
	logenv logx.Fields,
	project string,
	language string,
	semVer string,
	values map[string]string,
) (*Resource, error) {
	resource := newResource(project, language, semVer, values)

	err := db.Exec(logenv,
		`
		INSERT INTO resources (id, project, language, sem_ver, values, created, enabled)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
		`,
		resource.ID, resource.Project, resource.Language, resource.SemVer, resource.Values, resource.Created, resource.Enabled)

	if err != nil {
		return nil, err
	}
	return resource, nil
}
