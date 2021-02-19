package console

import (
	"fmt"
	"os"
	"strings"

	"go-boilerplate/helper"

	"github.com/dave/jennifer/jen"
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
	crudCmd.PersistentFlags().String("fields", "title:string,content:string", "module name")
	Root.AddCommand(crudCmd)
}

func parseFields(args string) (res map[string]string) {
	res = make(map[string]string)
	fields := strings.Split(args, ",")
	for _, field := range fields {
		content := strings.Split(field, ":")
		name := ""
		fieldType := "string"

		if len(content) > 0 {
			name = content[0]
			if len(content) > 1 {
				fieldType = content[1]
			}

			res[name] = fieldType
		}
	}
	return
}

func moduleGenerator(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	directory := "modules"

	currentDir, err := os.Getwd()
	if err != nil {
		helper.Logger.Info(fmt.Sprintf("Error: %s", err))
	}

	base := currentDir + string(os.PathSeparator) + directory + string(os.PathSeparator)

	fields := parseFields(cmd.Flag("fields").Value.String())

	// create entity
	err = generateEntity(name, currentDir+string(os.PathSeparator)+"entity"+string(os.PathSeparator), fields)
	if err != nil {
		helper.Logger.Error(fmt.Sprintf("Error: %s", err))
	} else {
		helper.Logger.Println("entity created")
	}

	// create directory
	err = os.Mkdir(base+"/"+name, 0755)
	if err != nil {
		helper.Logger.Error(fmt.Sprintf("Error: %s", err))
		os.Exit(0)
	}

	err = generateRepository(name, base)
	if err != nil {
		helper.Logger.Error(fmt.Sprintf("Error: %s", err))
	} else {
		helper.Logger.Println("repository created")
	}

	err = generateValidation(name, base, fields)
	if err != nil {
		helper.Logger.Error(fmt.Sprintf("Error: %s", err))
	} else {
		helper.Logger.Println("validation created")
	}

	err = generateService(name, base, fields)
	if err != nil {
		helper.Logger.Error(fmt.Sprintf("Error: %s", err))
	} else {
		helper.Logger.Println("service created")
	}

	err = generateRoutes(name, base)
	if err != nil {
		helper.Logger.Error(fmt.Sprintf("Error: %s", err))
	} else {
		helper.Logger.Println("routes created")
	}

	err = generatePostgresRepository(name, base)
	if err != nil {
		helper.Logger.Error(fmt.Sprintf("Error: %s", err))
	} else {
		helper.Logger.Println("postgres repository created")
	}

	err = generateHandler(name, base, fields)
	if err != nil {
		helper.Logger.Error(fmt.Sprintf("Error: %s", err))
	} else {
		helper.Logger.Println("handler created")
	}
}

func generateEntity(pkg, dest string, fields map[string]string) error {
	file := jen.NewFilePathName(dest, "entity")
	upperPkg := strcase.UpperCamelCase(pkg)

	file.Add(jen.Id("import").Parens(jen.Id("\n").Id(`"github.com/satori/uuid"`).Id("\n").Id(`"time"`).Id("\n")))

	file.Comment(upperPkg + " " + pkg + " entity")

	generatedEntityFields := []jen.Code{
		jen.Id("ID").Id("string").Tag(map[string]string{"json": "id", "xorm": "id"}),
	}

	createStruct := []jen.Code{
		jen.Id("ID").Op(":").Id("uuid.NewV4().String(),"),
	}

	createParams := make([]jen.Code, 0)

	for field, fieldType := range fields {
		generatedEntityFields = append(
			generatedEntityFields,
			jen.Id(strcase.UpperCamelCase(field)).Id(fieldType).Tag(map[string]string{"json": field, "xorm": field}),
		)

		createStruct = append(
			createStruct,
			jen.Id(strcase.UpperCamelCase(field)).Op(":").Id(strcase.LowerCamelCase(field)).Id(","),
		)

		createParams = append(
			createParams,
			jen.Id(strcase.LowerCamelCase(field)).Id(fieldType),
		)
	}

	generatedEntityFields = append(
		generatedEntityFields,
		jen.Id("CreatedAt").Id("time.Time").Tag(map[string]string{"json": "created_at", "xorm": "created"}),
		jen.Id("UpdatedAt").Id("time.Time").Tag(map[string]string{"json": "updated_at", "xorm": "updated"}),
		jen.Id("DeletedAt").Id("*time.Time").Tag(map[string]string{"json": "deleted_at", "xorm": "deleted"}),
	)

	file.Type().Id(upperPkg).Struct(
		generatedEntityFields...,
	)

	file.Comment(upperPkg + "ChangeSet change set for" + pkg)
	file.Type().Id(upperPkg + "ChangeSet").Struct(
		generatedEntityFields...,
	)

	file.Comment("New" + upperPkg + " create new" + pkg)

	file.Func().Id("New"+upperPkg).Params(createParams...).
		Parens(jen.List(jen.Id(pkg).Id(upperPkg), jen.Id("err").Id("error"))).Block(
		jen.Id(pkg).Id("=").Id(upperPkg).Block(
			createStruct...,
		),
		jen.Return(),
	)

	err := file.Save(dest + "/" + pkg + ".go")
	return err
}

func generateRepository(pkg, dest string) error {
	file := jen.NewFilePathName(dest, pkg)

	upperPkg := strcase.UpperCamelCase(pkg)

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
			jen.Id("pagination").Id("entity.Pagination"),
		).Params(jen.Id(upperPkg+"s").Id("[]entity."+upperPkg), jen.Id("count").Id("int"), jen.Id("err").Id("error")),
	)

	err := file.Save(dest + pkg + "/repository.go")
	return err
}

func generateValidation(pkg, dest string, fields map[string]string) error {
	file := jen.NewFilePathName(dest, pkg)

	generatedValidationFields := make([]jen.Code, 0)
	for field, fieldType := range fields {
		generatedValidationFields = append(
			generatedValidationFields,
			jen.Id(strcase.UpperCamelCase(field)).Id(fieldType).Tag(map[string]string{"json": field}),
		)
	}

	file.Comment("CreateRequest request for create new " + pkg)
	file.Type().Id("CreateRequest").Struct(
		generatedValidationFields...,
	)

	file.Comment("UpdateRequest request for update " + pkg)
	file.Type().Id("UpdateRequest").Struct(
		generatedValidationFields...,
	)

	err := file.Save(dest + pkg + "/validation.go")
	return err
}

func generateService(pkg, dest string, fields map[string]string) error {
	file := jen.NewFilePathName(dest, pkg)

	upperPkg := strcase.UpperCamelCase(pkg)

	file.Add(jen.Id("import").Params(
		jen.Id(`
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

	createParams := make([]jen.Code, 0)
	createArgs := make([]jen.Code, 0)
	for field, fieldType := range fields {
		name := strcase.LowerCamelCase(field)
		createParams = append(createParams, jen.Id(name).Id(fieldType))
		createArgs = append(createArgs, jen.Id(name))
	}

	file.Comment("Create" + upperPkg + " create new " + pkg)
	file.Func().Params(jen.Id("service").Id("Service")).Id("Create"+upperPkg).Params(createParams...).Params(
		jen.Id(pkg).Id("entity."+upperPkg),
		jen.Id("err").Id("error"),
	).Block(
		jen.List(jen.Id(pkg), jen.Err()).Op("=").Id("entity.New"+upperPkg).Params(createArgs...),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(),
		),

		jen.Id("err").Op("=").Id("service.repository.Save("+pkg+")"),

		jen.Return(),
	)

	file.Comment("GetList get list of " + pkg)
	file.Func().Params(jen.Id("service").Id("Service")).Id("GetList").Params(jen.Id("pagination").Id("entity.Pagination")).Params(
		jen.Id(pkg).Id("[]entity."+upperPkg),
		jen.Id("count").Id("int"),
		jen.Id("err").Id("error"),
	).Block(
		jen.List(jen.Id(pkg), jen.Id("count"), jen.Err()).Op("=").Id("service.repository.GetList(pagination)"),

		jen.Return(),
	)

	file.Comment("Update update " + pkg)
	file.Func().Params(jen.Id("service").Id("Service")).Id("Update").Params(jen.Id("id").Id("string"), jen.Id("changeset").Id("entity."+upperPkg+"ChangeSet")).Params(
		jen.Id(pkg).Id("entity."+upperPkg),
		jen.Id("err").Id("error"),
	).Block(
		jen.List(jen.Err()).Op("=").Id("service.repository.Update(id, changeset)"),
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
		jen.Id("prefix").Id("iris.Party"),
		jen.Id("adapters").Id("adapters.Adapters"),
	).Block(
		jen.Id("repository").Op(":=").Id("CreatePosgresRepository(adapters.Postgres)"),
		jen.Id("service").Op(":=").Id("CreateService(repository)"),
		jen.Id("handler").Op(":=").Id("handler{service, adapters}"),

		jen.Id(pkg).Op(":=").Id("prefix.Party(name)"),

		jen.Id(pkg+`.Get("/", middlewares.ValidatePaginationQuery, handler.GetList)`),
		jen.Id(pkg+`.Post("/", middlewares.ValidateBody(&CreateRequest{}), handler.Create)`),
		jen.Id(pkg+`.Get("/{id:string}", handler.GetByID)`),
		jen.Id(pkg+`.Delete("/{id:string}", handler.DeleteByID)`),
		jen.Id(pkg+`.Put("/{id:string}", middlewares.ValidateBody(&UpdateRequest{}), handler.Update)`),
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
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("GetList").Params(jen.Id("pagination").Id("entity.Pagination")).Params(jen.Id(pkgWithS).Id("[]entity."+upperPkg), jen.Id("count").Id("int"), jen.Id("err").Id("error")).Block(
		jen.List(jen.Id("count"), jen.Id("err")).Op("=").Id(`repo.db.
		Paginate("`+pkgWithS+`", &`+pkgWithS+`, pagination)`),
		jen.Return(),
	)

	file.Comment("Update update " + pkg)
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("Update").Params(jen.Id("id").Id("string"), jen.Id("changeset").Id("entity."+upperPkg+"ChangeSet")).Params(jen.Id("error")).Block(
		jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(`repo.db.Table("`+pkgWithS+`").Where("id = ?", id).Update(&changeset)`),
		jen.Return().Id("err"),
	)

	file.Comment("FindByID find " + pkg + " by id")
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("FindByID").Params(jen.Id("id").Id("string")).Params(jen.Id(pkg).Id("entity."+upperPkg), jen.Id("err").Id("error")).Block(
		jen.List(jen.Id("_"), jen.Id("err")).Op("=").Id(`repo.db.SQL("SELECT * FROM `+pkgWithS+` WHERE id = ? AND deleted_at IS null", id).Get(&`+pkg+`)`),
		jen.Return(),
	)

	file.Comment("DeleteByID delete " + pkg + " by id")
	file.Func().Params(jen.Id("repo").Id("PostgresRepository")).Id("DeleteByID").Params(jen.Id("id").Id("string")).Params(jen.Id("error")).Block(
		jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(`repo.db.Table("`+pkgWithS+`").Where("id = ?", id).Delete(&entity.`+upperPkg+`{})`),
		jen.Return().Id("err"),
	)

	err := file.Save(dest + pkg + "/postgres_repository.go")
	return err
}

func generateHandler(pkg, dest string, fields map[string]string) error {
	file := jen.NewFilePathName(dest, pkg)

	upperPkg := strcase.UpperCamelCase(pkg)
	pkgWithS := pkg + "s"
	errResponse := `helper.
	CreateErrorResponse(ctx, err).
	InternalServer().
	JSON()`
	successResponse := "helper.CreateResponse(ctx).Ok().WithData(" + pkg + ").JSON()"

	file.Add(jen.Id("import").Params(
		jen.Id(`"fmt"
		"go-boilerplate/adapters"
		"go-boilerplate/entity"
		"go-boilerplate/helper"
	
		"github.com/kataras/iris/v12"`),
	))

	updateFields := make([]jen.Code, 0)
	createFields := make([]jen.Code, 0)

	for field := range fields {
		name := strcase.UpperCamelCase(field)

		createFields = append(createFields, jen.Id("request."+name))
		updateFields = append(updateFields, jen.Id(name).Id(":").Id("request."+name).Id(","))
	}

	file.Type().Id("handler").Struct(
		jen.Id(pkgWithS).Id("Service"),
		jen.Id("adapters").Id("adapters.Adapters"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("GetList").Params(jen.Id("ctx").Id("iris.Context")).Block(
		jen.Id("request").Op(":=").Id(`ctx.Values().Get("pagination").(entity.Pagination)`),

		jen.List(jen.Id(pkgWithS), jen.Id("count"), jen.Id("err")).Op(":=").Id("h."+pkgWithS+".GetList(request)"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id("helper.CreatePaginationResponse(ctx, request,"+pkgWithS+", count).JSON()"),
		jen.Id("ctx.Next()"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("GetByID").Params(jen.Id("ctx").Id("iris.Context")).Block(
		jen.Id("id").Op(":=").Id(`ctx.Params().GetString("id")`),

		jen.List(jen.Id(pkg), jen.Id("err")).Op(":=").Id("h."+pkgWithS+".GetByID(id)"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id(successResponse),
		jen.Id("ctx.Next()"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("DeleteByID").Params(jen.Id("ctx").Id("iris.Context")).Block(
		jen.Id("id").Op(":=").Id(`ctx.Params().GetString("id")`),

		jen.Id("err").Op(":=").Id("h."+pkgWithS+".DeleteByID(id)"),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id(`helper.CreateResponse(ctx).Ok().WithMessage(fmt.Sprintf("%s deleted", id)).JSON()`),
		jen.Id("ctx.Next()"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("Update").Params(jen.Id("ctx").Id("iris.Context")).Block(
		jen.Id("request").Op(":=").Id(`ctx.Values().Get("body").(*UpdateRequest)`),
		jen.Id("id").Op(":=").Id(`ctx.Params().GetString("id")`),

		jen.List(jen.Id(pkg), jen.Id("err")).Op(":=").Id("h."+pkgWithS+".Update").Params(
			jen.Id("id"),
			jen.Id("entity."+upperPkg+"ChangeSet").Block(
				updateFields...,
			),
		),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id(errResponse),
			jen.Return(),
		),

		jen.Id(successResponse),
		jen.Id("ctx.Next()"),
	)

	file.Func().Params(jen.Id("h").Id("handler")).Id("Create").Params(jen.Id("ctx").Id("iris.Context")).Block(
		jen.Id("request").Op(":=").Id(`ctx.Values().Get("body").(*CreateRequest)`),

		jen.List(jen.Id(pkg), jen.Id("err")).Op(":=").Id("h."+pkgWithS+".Create"+upperPkg).Params(createFields...),

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
