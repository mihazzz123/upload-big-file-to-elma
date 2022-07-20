package util

import (
	"context"
	"fmt"
	"strings"

	"git.elewise.com/elma365/common/pkg/collection"
	"git.elewise.com/elma365/common/pkg/connection/sqldb"
	"git.elewise.com/elma365/common/pkg/md"

	"github.com/Masterminds/squirrel"
)

// BodyJSON шорткат для работы с подполями элемента коллекции
type BodyJSON string

// Column - шорткат обращения к подполяем необходимой колонки таблицы
func Column(column string, codes ...string) BodyJSON {
	return createBuilder(column, false, codes...)
}

// RawColumn - шорткат для обращения к полям необходимой колонки в виде jsonb
func RawColumn(column string, codes ...string) BodyJSON {
	return createBuilder(column, true, codes...)
}

// Body шорткат для обращения к подполям элемента коллекции
func Body(codes ...string) BodyJSON {
	return createBuilder(collection.CollectionBodyCol, false, codes...)
}

// RawBody шорткат для обращения к подполям элемента коллекции в виде jsonb
func RawBody(codes ...string) BodyJSON {
	return createBuilder(collection.CollectionBodyCol, true, codes...)
}

func createBuilder(column string, isRaw bool, codes ...string) BodyJSON {
	delimeter := ">>"
	if isRaw {
		delimeter = ">"
	}

	format := collection.CollectionBodyCol

	if len(codes) > 0 {
		format = fmt.Sprintf("%s#%s'{%s}'", column, delimeter, strings.Join(codes, ", "))
	}

	return BodyJSON(format)
}

// Cast привести значение поля к определённому типу
func (bj BodyJSON) Cast(t string) BodyJSON {
	return BodyJSON(fmt.Sprintf("%s::%s", bj, t))
}

// Table добавляет имя таблицы
func (bj BodyJSON) Table(t string) BodyJSON {
	return BodyJSON(fmt.Sprintf("%s.%s", t, string(bj)))
}

// CastTS привести значение поля к формату временной метки и добавить компанию из контекста
func (bj BodyJSON) CastTS(ctx context.Context) BodyJSON {
	company := md.CompanyFromContext(ctx)
	return BodyJSON(fmt.Sprintf("%q.f_cast_isots%s", company, bj))
}

// String implements fmt.Stringer
func (bj BodyJSON) String() string {
	return fmt.Sprintf("(%s)", string(bj))
}

// ToSql implements sqldb.QueryBuilder
//
//nolint:golint
func (bj BodyJSON) ToSql() (query string, args []interface{}, err error) {
	return string(bj), nil, nil
}

// As формирует выражения для алиаса к полю
func (bj BodyJSON) As(alias string) string {
	return fmt.Sprintf("%s AS %s", bj, alias)
}

// Eq возвращает выражение для выражения WHERE на равенство
//
// В качестве входного параметра должен передавать экземпляр sqldb.QueryBuilder или string.
// Если параметр имеет тип отличный от указанных, то следует сперва воспользоваться методом Cast.
func (bj BodyJSON) Eq(to interface{}) sqldb.QueryBuilder {
	return squirrel.Eq{bj.String(): to}
}

// Neq возвращает выражение для выражения WHERE на равенство
//
// В качестве входного параметра должен передавать экземпляр sqldb.QueryBuilder или string.
// Если параметр имеет тип отличный от указанных, то следует сперва воспользоваться методом Cast.
func (bj BodyJSON) Neq(to interface{}) sqldb.QueryBuilder {
	return squirrel.NotEq{bj.String(): to}
}

// Gt возвращает выражение для выражения WHERE на равенство
//
// В качестве входного параметра должен передавать экземпляр sqldb.QueryBuilder или string.
// Если параметр имеет тип отличный от указанных, то следует сперва воспользоваться методом Cast.
func (bj BodyJSON) Gt(to interface{}) sqldb.QueryBuilder {
	return squirrel.Gt{bj.String(): to}
}

// Gte возвращает выражение для выражения WHERE на равенство
//
// В качестве входного параметра должен передавать экземпляр sqldb.QueryBuilder или string.
// Если параметр имеет тип отличный от указанных, то следует сперва воспользоваться методом Cast.
func (bj BodyJSON) Gte(to interface{}) sqldb.QueryBuilder {
	return squirrel.GtOrEq{bj.String(): to}
}

// Lt возвращает выражение для выражения WHERE на равенство
//
// В качестве входного параметра должен передавать экземпляр sqldb.QueryBuilder или string.
// Если параметр имеет тип отличный от указанных, то следует сперва воспользоваться методом Cast.
func (bj BodyJSON) Lt(to interface{}) sqldb.QueryBuilder {
	return squirrel.Lt{bj.String(): to}
}

// Lte возвращает выражение для выражения WHERE на равенство
//
// В качестве входного параметра должен передавать экземпляр sqldb.QueryBuilder или string.
// Если параметр имеет тип отличный от указанных, то следует сперва воспользоваться методом Cast.
func (bj BodyJSON) Lte(to interface{}) sqldb.QueryBuilder {
	return squirrel.LtOrEq{bj.String(): to}
}

// IsNull возвращает выражение для выражения WHERE на равенство
func (bj BodyJSON) IsNull() sqldb.QueryBuilder {
	return squirrel.Eq{bj.String(): nil}
}

// IsNotNull возвращает выражение для выражения WHERE на равенство
func (bj BodyJSON) IsNotNull() sqldb.QueryBuilder {
	return squirrel.NotEq{bj.String(): nil}
}

type bbField struct {
	Alias, Query string
}

// BodyBuilder конструктор объекта body для сложных запросов
type BodyBuilder []bbField

// Add добавляет к строящемуся объекту поле с указанным кодом из указанной таблицы
func (bb *BodyBuilder) Add(code, from string, codes ...string) *BodyBuilder {
	*bb = append(*bb, bbField{
		Alias: code,
		Query: RawBody(codes...).Table(from).String(),
	})
	return bb
}

// AddRaw - добавляем query к текущему запросу
func (bb *BodyBuilder) AddRaw(code, query string) *BodyBuilder {
	*bb = append(*bb, bbField{Alias: code, Query: query})

	return bb
}

// AddSubquery добавляет к строящемуся объекту поле с указанным кодом как подзапрос
//
// **Внимание!** Подзапрос должен быть SELECT без аргументов.
// По сути он вычисляется с откидыванием аргументов, оставляя лишь query.
func (bb *BodyBuilder) AddSubquery(code, aggregate string, stmt sqldb.QueryBuilder) *BodyBuilder {
	subquery, _, _ := stmt.ToSql()
	*bb = append(*bb, bbField{
		Alias: code,
		Query: fmt.Sprintf("%s(%s)", aggregate, subquery),
	})
	return bb
}

// BuildExpr - возврат запроса, как выражение для squirrel
// используется для подстановки выражений как подзапросов
func (bb *BodyBuilder) BuildExpr() sqldb.QueryBuilder {
	return squirrel.Expr(fmt.Sprintf("(%s)", bb.Build()))
}

// BuildAs построить запрос с aliase
func (bb *BodyBuilder) BuildAs(aliase string) string {
	return fmt.Sprintf("%s as %s", bb.Build(), aliase)
}

// Build построить запрос
func (bb *BodyBuilder) Build() string {
	pairs := make([]string, 0, 2*len(*bb))
	for _, field := range *bb {
		pairs = append(pairs, "'"+field.Alias+"'", field.Query)
	}
	return fmt.Sprintf("jsonb_build_object(%s)", strings.Join(pairs, ","))
}

// RemapJSONQuery - построение запроса, для селекта из таблицы с определенным мэпингом json
// используется для запросов insert from select, update from select
// Пример:
// мэпим namespace, __id приложения к namespace, id другого приложения
//
// bb := util.BodyBuilder{}
// bb.Add("namespace", "some_namespace:some_app", "namespace").
// Add("id", "some_namespace:some_app", "__id")
// query := bb.RemapJSONQuery("someTable", []string{}, []string{})
func (bb *BodyBuilder) RemapJSONQuery(from string, prependFields, appendFields []string) string {
	codes := make([]string, 0, len(prependFields)+len(appendFields)+1)
	codes = append(codes, prependFields...)
	codes = append(codes, bb.BuildAs("body"))
	codes = append(codes, appendFields...)

	query, _, _ := squirrel.Select(codes...).From(from).ToSql()

	return query
}
