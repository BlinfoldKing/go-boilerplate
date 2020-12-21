package console

import (
	"fmt"
	"go-boilerplate/entity"
	"os"
	"reflect"

	"github.com/dave/jennifer/jen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
)

var crudCmd = &cobra.Command{
	Use:   "crud",
	Short: "crud by entity",
	Long:  `This subcommand used to creating module files`,
	Run:   moduleGenerator,
}

func init() {
	crudCmd.PersistentFlags().String("name", "example", "module name")
	crudCmd.PersistentFlags().String("directory", "modules", "modules directory")
	Root.AddCommand(crudCmd)
}

func moduleGenerator(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	directory := cmd.Flag("directory").Value.String()

	currentDir, err := os.Getwd()
	if err != nil {
		logrus.Info(fmt.Sprintf("Error: %s", err))
	}

	base := currentDir + string(os.PathSeparator) + directory + string(os.PathSeparator)

	// create directory
	err = os.Mkdir(base+"/"+name, 0755)
	if err != nil {
		logrus.Info(fmt.Sprintf("Error: %s", err))
		os.Remove(base + "/" + name)
	}

	err = generateRepository(name, base)
	if err != nil {
		logrus.Info(fmt.Sprintf("Error: %s", err))
	}

	err = generateValidation(name, base)
	if err != nil {
		logrus.Info(fmt.Sprintf("Error: %s", err))
	}

	err = generateService(name, base)
	if err != nil {
		logrus.Info(fmt.Sprintf("Error: %s", err))
	}

	err = generateRoutes(name, base)
	if err != nil {
		logrus.Info(fmt.Sprintf("Error: %s", err))
	}

	err = generatePostgresRepository(name, base)
	if err != nil {
		logrus.Info(fmt.Sprintf("Error: %s", err))
	}

	err = generateHandler(name, base)
	if err != nil {
		logrus.Info(fmt.Sprintf("Error: %s", err))
	}
}

func generateRepository(pkg, dest string) error {
	file := jen.NewFilePathName(dest, pkg)

	upperPkg := strcase.UpperCamelCase(pkg)

	file.Comment("//go:generate mockgen -package " + pkg + " -source=repository.go -destination repository_mock.go")
	file.Add(jen.Id("import").Parens(jen.Id(`"go-boilerplate/entity"`)))

	file.Comment("Repository abstraction for storage")
	file.Type().Id("Repository").Interface(
		jen.Id("Save").Params(jen.Id("entity."+upperPkg)).Error(),
		jen.Id("DeleteByID").Params(jen.Id("id").Id("string")).Error(),
		jen.Id("FindByID").Params(jen.Id("id").Id("string")).Params(jen.Id("entity."+upperPkg), jen.Id("error")),

		jen.Id("Update").Params(
			jen.Id("id").Id("string"),
			jen.Id("changeset").Id("entity."+upperPkg+"ChangeSet"),
		).Error(),

		jen.Id("GetList").Params(
			jen.Id("limit"),
			jen.Id("offset").Id("int"),
		).Params(jen.Id("[]entity."+upperPkg), jen.Id("error")),
	)

	err := file.Save(dest + pkg + "/repository.go")
	return err
}

func generateValidation(pkg, dest string) error {
	file := jen.NewFilePathName(dest, pkg)

	file.Type().Id("CreateRequest").Struct()
	file.Type().Id("UpdateRequest").Struct()

	err := file.Save(dest + pkg + "/validation.go")
	return err
}

func generateService(pkg, dest string) error {
	file := jen.NewFilePathName(dest, pkg)

	upperPkg := strcase.UpperCamelCase(pkg)

	file.Add(jen.Id("import").Params(
		jen.Id(`
		"errors"
		"go-boilerplate/entity"`),
	))

	file.Comment("Service contains business logic")
	file.Type().Id("Service").Struct(
		jen.Id("repository").Id("Repository"),
	)

	file.Comment("CreateService init service")
	file.Func().Id("CreateService").Params(jen.Id("repo").Id("Repository")).Id("Service").Block(
		jen.Return(jen.Id("Service{repo}")),
	)

	e := reflect.ValueOf(&entity.User{}).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}

	file.Comment("Create" + upperPkg + " create new " + pkg)
	file.Func().Params(jen.Id("service").Id("Service")).Id("Create"+upperPkg).Params(jen.Id("dataset").Id("entity."+upperPkg)).Params(
		jen.Id(pkg).Id("entity."+upperPkg),
		jen.Id("err").Id("error"),
	).Block(
		jen.List(jen.Id(pkg), jen.Err()).Op(":=").Id("entity.New"+upperPkg+"(dataset)"),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(),
		),

		jen.Id("err").Op("=").Id("service.repository.Save("+pkg+")"),

		jen.Return(),
	)

	file.Comment("GetList get list of " + pkg)
	file.Func().Params(jen.Id("service").Id("Service")).Id("GetList").Params(jen.Id("limit"), jen.Id("offset").Id("int")).Params(
		jen.Id(pkg).Id("[]entity."+upperPkg),
		jen.Id("err").Id("error"),
	).Block(
		jen.List(jen.Id(pkg), jen.Err()).Op(":=").Id("service.repository.GetList(limit, offset)"),

		jen.Return(),
	)

	file.Comment("Update update " + pkg)
	file.Func().Params(jen.Id("service").Id("Service")).Id("Update").Params(jen.Id("id").Id("string"), jen.Id("changeset").Id("entity."+upperPkg+"ChangeSet")).Params(
		jen.Id(pkg).Id("entity."+upperPkg),
		jen.Id("err").Id("error"),
	).Block(
		jen.List(jen.Err()).Op(":=").Id("service.repository.Update(id, changeset)"),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("entity."+upperPkg+"{}"), jen.Id("err")),
		),
		jen.Return(jen.Id("service.GetByID(id)")),
	)

	file.Comment("GetByID find " + pkg + "by id")
	file.Func().Params(jen.Id("service").Id("Service")).Id("GetByID").Params(jen.Id("id").Id("string")).Params(
		jen.Id(pkg).Id("entity."+upperPkg),
		jen.Id("err").Id("error"),
	).Block(
		jen.Return(jen.Id("service.repository.FindByID(id)")),
	)

	file.Comment("DeleteByID delete " + pkg + "by id")
	file.Func().Params(jen.Id("service").Id("Service")).Id("DeleteByID").Params(jen.Id("id").Id("string")).Params(
		jen.Id("err").Id("error"),
	).Block(
		jen.Return(jen.Id("service.repository.DeleteByID(id)")),
	)

	err := file.Save(dest + pkg + "/service.go")
	return err
}

func generateRoutes(pkg, dest string) error {
	file := jen.NewFilePathName(dest, pkg)

	file.Add(jen.Id("import").Params(
		jen.Id(`
		"go-boilerplate/adapters"
		"go-boilerplate/middlewares"

		"github.com/kataras/iris/v12"`),
	))

	// route
	file.Const().Id("name").Op("=").Lit("/" + pkg)

	file.Comment("Routes init " + pkg)
	file.Func().Id("Routes").Params(
		jen.Id("app").Id("*iris.Application"),
		jen.Id("adapters").Id("adapters.Adapters"),
	).Block(
		jen.Id("repository").Op(":=").Id("CreatePosgresRepository(adapters.Postgres)"),
		jen.Id("service").Op(":=").Id("CreateService(repository)"),
		jen.Id("handler").Op(":=").Id("handler{service, adapters}"),

		jen.Id(pkg).Op(":=").Id("app.Party(name)"),

		jen.Id(pkg+`.Get("/", handler.GetList)`),
		jen.Id(pkg+`.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)`),
		jen.Id(pkg+`.Get("/{id:string}", handler.GetByID)`),
		jen.Id(pkg+`.Delete("/{id:string}", handler.DeleteByID)`),
		jen.Id(pkg+`.Put("/", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)`),
	)

	err := file.Save(dest + pkg + "/routes.go")
	return err
}

func generatePostgresRepository(pkg, dest string) error {
	file := jen.NewFilePathName(dest, pkg)

	upperPkg := strcase.UpperCamelCase(pkg)
	pkgWithS := pkg + "s"

	file.Add(jen.Id("import").Params(
		jen.Id(`"go-boilerplate/adapters/postgres"
		"go-boilerplate/entity"`),
	))

	file.Comment("PostgresRepository repository implementation on postgres")
	file.Type().Id("PostgresRepository").Struct(
		jen.Id("db").Id("*postgres.Postgres"),
	)

	file.Comment("CreatePosgresRepository init PostgresRepository")
	file.Func().Id("CreatePosgresRepository").Params(jen.Id("db").Id("*postgres.Postgres")).Id("Repository").Block(
		jen.Return(jen.Id("PostgresRepository{db}")),
	)

	file.Comment("Save save " + pkg + " to db")
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("Save").Params(jen.Id(pkg).Id("entity."+upperPkg)).Id("error").Block(
		jen.List(jen.Id("_"), jen.Err()).Op(":=").Id(`repo.db.Table("`+pkgWithS+`").Insert(&`+pkg+`)`),
		jen.Return(jen.Id("err")),
	)

	file.Comment("GetList get list of " + pkg)
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("GetList").Params(jen.Id("limit"), jen.Id("offset").Id("int")).Params(jen.Id(pkgWithS).Id("[]entity."+upperPkg), jen.Id("err").Id("error")).Block(
		jen.Id("err").Op("=").Id(`repo.db.
		Paginate("`+pkgWithS+`", &`+pkgWithS+`, postgres.PaginationOpt{
			Limit:  &limit,
			Offset: &offset,
		})`),
		jen.Return(),
	)

	file.Comment("Update update " + pkg)
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("Update").Params(jen.Id("id").Id("string"), jen.Id("changeset").Id("entity."+upperPkg+"ChangeSet")).Params(jen.Id("error")).Block(
		jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(`repo.db.Table("`+pkgWithS+`").Where("id = ?", id).Update(&changeset)`),
		jen.Return().Id("err"),
	)

	file.Comment("FindByID find " + pkg + " by id")
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("FindByID").Params(jen.Id("id").Id("string")).Params(jen.Id(pkg).Id("entity."+upperPkg), jen.Id("err").Id("error")).Block(
		jen.List(jen.Id("_"), jen.Id("err")).Op("=").Id(`repo.db.SQL("SELECT * FROM `+pkgWithS+` WHERE id = ?", id).Get(&`+pkg+`)`),
		jen.Return(),
	)

	file.Comment("DeleteByID delete " + pkg + " by id")
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("DeleteByID").Params(jen.Id("id").Id("string")).Params(jen.Id("error")).Block(
		jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(`repo.db.Exec("DELETE FROM `+pkgWithS+` WHERE id = ?", id)`),
		jen.Return().Id("err"),
	)

	err := file.Save(dest + pkg + "/postgres_repository.go")
	return err
}

func generateHandler(pkg, dest string) error {
	file := jen.NewFilePathName(dest, pkg)

	upperPkg := strcase.UpperCamelCase(pkg)
	pkgWithS := pkg + "s"
	errResponse := `helper.
	CreateErrorResponse(ctx, err).
	InternalServer().
	JSON()`
	successResponse := `helper.CreateResponse(ctx).Ok().WithData(" + pkg + ").JSON()`

	file.Add(jen.Id("import").Params(
		jen.Id(`"fmt"
		"go-boilerplate/adapters"
		"go-boilerplate/entity"
		"go-boilerplate/helper"
	
		"github.com/kataras/iris/v12"`),
	))

	file.Type().Id("handler").Struct(
		jen.Id(pkgWithS).Id("Service"),
		jen.Id("adapters").Id("adapters.Adapters"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("GetList").Params(jen.Id("ctx").Id("iris.Contenxt")).Block(
		jen.Id("limit").Op(":=").Id(`ctx.URLParamIntDefault("limit", 10)`),
		jen.Id("offset").Op(":=").Id(`ctx.URLParamIntDefault("offset", 10)`),

		jen.List(jen.Id(pkg), jen.Id("err")).Op(":=").Id("h." + pkgWithS + ".GetList(limit, offset)"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id(successResponse),
		jen.Id("ctx.Next()"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("GetByID").Params(jen.Id("ctx").Id("iris.Contenxt")).Block(
		jen.Id("id").Op(":=").Id(`ctx.Params().GetString("id")`),

		jen.List(jen.Id(pkg), jen.Id("err")).Op(":=").Id("h." + pkgWithS + ".GetByID(id)"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id(successResponse),
		jen.Id("ctx.Next()"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("DeleteByID").Params(jen.Id("ctx").Id("iris.Contenxt")).Block(
		jen.Id("id").Op(":=").Id(`ctx.Params().GetString("id")`),

		jen.Id("err").Op(":=").Id("h." + pkgWithS + ".DeleteByID(id)"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id(`helper.CreateResponse(ctx).Ok().WithMessage(fmt.Sprintf("%s deleted", id)).JSON()`),
		jen.Id("ctx.Next()"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("Update").Params(jen.Id("ctx").Id("iris.Contenxt")).Block(
		jen.Id("request").Op(":=").Id(`ctx.Values().Get("body").(*UpdateRequest)`),

		jen.List(jen.Id(pkg), jen.Id("err")).Op(":=").Id("h." + pkgWithS + ".GetByID(id)"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.List(jen.Id(pkg), jen.Id("err")).Op("=").Id("h." + pkgWithS + ".Update(request.ID, entity." + upperPkg + "ChangeSet{})"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id(successResponse),
		jen.Id("ctx.Next()"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("Create").Params(jen.Id("ctx").Id("iris.Contenxt")).Block(
		jen.Id("request").Op(":=").Id(`ctx.Values().Get("body").(*CreateRequest)`),

		jen.List(jen.Id(pkg), jen.Id("err")).Op(":=").Id("h." + pkgWithS + ".Create(entity." + upperPkg + "{})"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id(successResponse),
		jen.Id("ctx.Next()"),
	)

	err := file.Save(dest + pkg + "/handler.go")
	return err
}