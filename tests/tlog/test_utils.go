package tlog

import (
	"context"

	"github.com/nmarsollier/resourcesgo/internal/tools/logx"
)

var TestContext = logx.CtxWithFields(context.Background(), logx.NewFields())
