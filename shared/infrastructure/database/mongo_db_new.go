package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reflect"
	"regexp"
	"strings"
)

type GetAllParam struct {
	page   int64
	size   int64
	sort   map[string]any
	filter map[string]any
}

func (g GetAllParam) Page(page int64) GetAllParam {
	g.page = page
	return g
}

func (g GetAllParam) Size(size int64) GetAllParam {
	g.size = size
	return g
}

func (g GetAllParam) Sort(field string, sort any) GetAllParam {
	g.sort[field] = sort
	return g
}

func (g GetAllParam) Filter(key string, value any) GetAllParam {
	g.filter[key] = value
	return g
}

func NewDefaultParam() GetAllParam {
	return GetAllParam{
		page:   1,
		size:   20,
		sort:   map[string]any{},
		filter: map[string]any{},
	}
}

type InsertOrUpdateRepo[T any] interface {
	InsertOrUpdate(obj T) error
}

type InsertManyRepo[T any] interface {
	InsertMany(objs ...T) error
}

type GetOneRepo[T any] interface {
	GetOne(filter map[string]any, result *T) error
}

type GetAllRepo[T any] interface {
	GetAll(param GetAllParam, results *[]T) (int64, error)
}

type GetAllEachItemRepo[T any] interface {
	GetAllEachItem(param GetAllParam, resultEachItem func(result T)) (int64, error)
}

type Repository[T any] interface {
	InsertOrUpdateRepo[T]
	InsertManyRepo[T]
	GetOneRepo[T]
	GetAllRepo[T]
	GetAllEachItemRepo[T]
	getTypeName() string
}

type Basic[T any] struct{}

func (b Basic[T]) getTypeName() string {
	var x T
	return snakeCase(reflect.TypeOf(x).Name())
}

var matchFirstCapSnakeCase = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCapSnakeCase = regexp.MustCompile("([a-z\\d])([A-Z])")

// SnakeCase is
func snakeCase(str string) string {
	snake := matchFirstCapSnakeCase.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCapSnakeCase.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func toSliceAny[T any](objs []T) []any {
	var results []any
	for _, obj := range objs {
		results = append(results, obj)
	}
	return results
}

type mongoGateway[T any] struct {
	Basic[T]
	Database *mongo.Database
}

func NewMongoGateway[T any](uri, databaseName string) Repository[T] {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	return &mongoGateway[T]{
		Database: client.Database(databaseName),
	}
}

func (g *mongoGateway[T]) InsertOrUpdate(obj T) error {

	name := g.getTypeName()

	sf, exist := reflect.TypeOf(obj).FieldByName("ID")
	if !exist {
		return fmt.Errorf("field ID as primary key is not found in %s", reflect.TypeOf(obj).Name())
	}

	tagValue, exist := sf.Tag.Lookup("bson")
	if !exist || tagValue != "_id" {
		return fmt.Errorf("field ID must have tag `bson:\"_id\"`")
	}

	filter := bson.D{{"_id", reflect.ValueOf(obj).FieldByName("ID").Interface()}}
	update := bson.D{{"$set", obj}}
	opts := options.Update().SetUpsert(true)

	coll := g.Database.Collection(name)
	_, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (g *mongoGateway[T]) InsertMany(objs ...T) error {

	if len(objs) == 0 {
		return fmt.Errorf("objs must > 0")
	}

	name := g.getTypeName()

	opts := options.InsertMany().SetOrdered(false)

	coll := g.Database.Collection(name)
	_, err := coll.InsertMany(context.TODO(), toSliceAny(objs), opts)
	if err != nil {
		return err
	}

	return nil
}

func (g *mongoGateway[T]) GetOne(filter map[string]any, result *T) error {

	name := g.getTypeName()

	coll := g.Database.Collection(name)

	singleResult := coll.FindOne(context.TODO(), filter)

	err := singleResult.Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (g *mongoGateway[T]) GetAll(param GetAllParam, results *[]T) (int64, error) {

	name := g.getTypeName()

	coll := g.Database.Collection(name)

	skip := param.size * (param.page - 1)
	limit := param.size

	findOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  param.sort,
	}

	ctx := context.TODO()

	count, err := coll.CountDocuments(ctx, param.filter)
	if err != nil {
		return 0, err
	}

	cursor, err := coll.Find(ctx, param.filter, &findOpts)
	if err != nil {
		return 0, err
	}

	err = cursor.All(ctx, results)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (g *mongoGateway[T]) GetAllEachItem(param GetAllParam, resultEachItem func(result T)) (int64, error) {

	name := g.getTypeName()

	coll := g.Database.Collection(name)

	skip := param.size * (param.page - 1)
	limit := param.size

	findOpts := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  param.sort,
	}

	ctx := context.TODO()

	count, err := coll.CountDocuments(ctx, param.filter)
	if err != nil {
		return 0, err
	}

	cursor, err := coll.Find(ctx, param.filter, &findOpts)
	if err != nil {
		return 0, err
	}

	for cursor.Next(ctx) {

		var result T
		err := cursor.Decode(&result)
		if err != nil {
			return 0, err
		}

		resultEachItem(result)

	}

	err = cursor.Err()
	if err != nil {
		return 0, err
	}

	return count, nil

}