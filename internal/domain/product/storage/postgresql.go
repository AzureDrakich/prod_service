package storage

import (
	"app/internal/domain/product/model"
	db "app/pkg/client/postgresql/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"log"
)

type ProductStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
}

func NewProductStorage(client PostgreSQLClient) ProductStorage {
	return ProductStorage{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

const (
	scheme = "public"
	tabl   = "category"
)

//func (s *ProductStorage) queryLogger(sql, table string, args []interface{}) string {
//q := fmt.Sprintf("sql: %s\ntable: %s, args:%s", sql, table, args)
//return q
//}

func (s *ProductStorage) All(ctx context.Context) ([]model.Product, error) {
	query := s.queryBuilder.Select("*").
		From("prod_service")

	sql, args, err := query.ToSql()
	if err != nil {
		err := db.ErrCreateQuery(err)
		log.Print(err)
		return nil, err
	}

	log.Print("do query")
	rows, err := s.client.Query(ctx, sql, args...)
	if err != nil {
		err = db.ErrDoQuery(err)
		log.Print(err)
		return nil, err
	}

	defer rows.Close()
	list := make([]model.Product, 0)
	for rows.Next() {
		p := model.Product{}
		if err = rows.Scan(
			&p.Id, &p.Name, &p.Description, &p.Image_id, &p.Price, &p.Currency_id, &p.Rating, &p.Category_id, &p.Specification, &p.Created_at, &p.Updated_at); err != nil {
			err = db.ErrScan(err)
			log.Printf("Error: %s", err)
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}
