package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	_ "embed"
)

//go:embed go.mod.tpl
var mod string

//go:embed v1.proto.tpl
var v1 string

var (
	//go:embed cmd.cron.wire.go.tpl
	cmdCronWire string

	//go:embed cmd.server.wire.go.tpl
	cmdServerWire string

	//go:embed cmd.server.main.go.tpl
	cmdServerMain string

	//go:embed cmd.server.server.go.tpl
	cmdServerServer string
)

var (
	//go:embed initializer.gorm.go.tpl
	initializerGorm string

	//go:embed initializer.wire_set.go.tpl
	initializerWireSet string
)

var (
	//go:embed internal.cron.go.tpl
	internalCron string
	//go:embed internal.domain.article.go.tpl
	internalDomainArticle string
	//go:embed internal.domain.model.go.tpl
	internalDomainModel string
	//go:embed internal.domain.tag.go.tpl
	internalDomainTag string
	//go:embed internal.pkg.copier.go.tpl
	internalPkgCopier string
	//go:embed internal.server.pepo.go.tpl
	internalServerPepo string
	//go:embed internal.server.service.go.tpl
	internalServerService string
	//go:embed internal.server.usecase.go.tpl
	internalServerUsecase string

	//go:embed internal.wire.set.go.tpl
	internalWireSet string
)

var (
	isWrite *bool
	appPath *string
	appName *string
)

func main() {
	isWrite = flag.Bool("write", false, "help message for write")
	appPath = flag.String("path", "test", "help message for test")
	appName = flag.String("name", "test", "help message for test")

	flag.Parse()

	CreatedApp(*appPath, *appName)
}

func CreatedApp(appPath, appName string) {
	log.Println("Creating application...")
	os.MkdirAll(appName, 0755)
	os.MkdirAll(path.Join(appName, "api/product/app/v1"), 0755)
	os.MkdirAll(path.Join(appName, "/cmd/server"), 0755)
	os.MkdirAll(path.Join(appName, "/cmd/cron"), 0755)

	os.MkdirAll(path.Join(appName, "/config/initializer"), 0755)

	os.MkdirAll(path.Join(appName, "/internal/cron"), 0755)
	os.MkdirAll(path.Join(appName, "/internal/domain"), 0755)
	os.MkdirAll(path.Join(appName, "/internal/pkg/copier"), 0755)
	os.MkdirAll(path.Join(appName, "/internal/server/repo"), 0755)
	os.MkdirAll(path.Join(appName, "/internal/server/service"), 0755)
	os.MkdirAll(path.Join(appName, "/internal/server/usecase"), 0755)

	if *isWrite {

		WriteToFile(path.Join(appName, "go.mod"), strings.Replace(mod, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "api/product/app/v1", "v1.proto"), strings.Replace(v1, "{{.Appname}}", appName, -1))

		WriteToFile(path.Join(appName, "/cmd/server", "main.go"), strings.Replace(cmdServerMain, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/cmd/server", "server.go"), strings.Replace(cmdServerServer, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/cmd/server", "wire.go"), strings.Replace(cmdServerWire, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/cmd/cron", "wire.go"), strings.Replace(cmdCronWire, "{{.Appname}}", appName, -1))

		WriteToFile(path.Join(appName, "/config/initializer", "gorm.go"), strings.Replace(initializerGorm, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/config/initializer", "wire_set.go"), strings.Replace(initializerWireSet, "{{.Appname}}", appName, -1))

		WriteToFile(path.Join(appName, "/internal/cron", "wire_set.go"), strings.Replace(internalCron, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/internal/domain", "article.go"), strings.Replace(internalDomainArticle, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/internal/domain", "model.go"), strings.Replace(internalDomainModel, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/internal/domain", "tag.go"), strings.Replace(internalDomainTag, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/internal/pkg/copier", "copier.go"), strings.Replace(internalPkgCopier, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/internal/server/repo", "repo.go"), strings.Replace(internalServerPepo, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/internal/server/service", "article.go"), strings.Replace(internalServerService, "{{.Appname}}", appName, -1))
		WriteToFile(path.Join(appName, "/internal/server/usecase", "usecase.go"), strings.Replace(internalServerUsecase, "{{.Appname}}", appName, -1))

		WriteToFile(path.Join(appName, "/internal/server", "wire_set.go"), strings.Replace(internalWireSet, "{{.Appname}}", appName, -1))

	}

	log.Println("new application successfully created!")
}

func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
