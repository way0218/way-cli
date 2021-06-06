package server

import (
	"github.com/google/wire"
	"{{.Appname}}/internal/server/repo"
	"{{.Appname}}/internal/server/service"
	"{{.Appname}}/internal/server/usecase"
)

// Set for di
var Set = wire.NewSet(
	service.NewArticleService,
	usecase.NewArticleUsecase,
	repo.NewArticleRepo,
)