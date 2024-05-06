package postgres

import (
	"ShoesShop/internal/models/entity"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddProduct(ctx context.Context, prod entity.Product) (id int32, err error) {
	q := `
	INSERT INTO products
		(category, name, color, price, size, quantity, images)
	VALUES 
	    ($1, $2, $3, $4, $5, $6, $7)
	RETURNING 
		id
	`

	err = r.db.QueryRowContext(
		ctx,
		q,
		prod.Category,
		prod.Name,
		prod.Color,
		prod.Price,
		prod.Size,
		prod.Quantity,
		prod.Images,
	).Scan(&id) //TODO
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) DeleteProduct(ctx context.Context, id int32) (err error) {
	q := `
	DELETE FROM
        products 
    WHERE 
        id = $1
`
	_, err = r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListProducts(ctx context.Context, filters map[string]any) ([]entity.Product, error) {
	var products []entity.Product
	params := squirrel.And{}
	for key, value := range filters {
		params = append(params, squirrel.Eq{key: value})
	}
	qf := squirrel.Select("id, category, name, color, price, size, quantity, images").
		From("products").
		Where(params).
		PlaceholderFormat(squirrel.Dollar)
	q, args, err := qf.ToSql()
	if err != nil {
		return products, err
	}
	if len(args) == 0 {
		err = r.db.SelectContext(ctx, &products, q)
	} else {
		err = r.db.SelectContext(ctx, &products, q, args...)
	}
	if err != nil {
		return nil, err
	}
	return products, nil
}
