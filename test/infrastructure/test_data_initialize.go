package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var INSERT_PRODUCTS = `INSERT INTO products (name, price, discount,store) 
VALUES('AirFryer',3000.0, 22.0, 'ABC TECH'),
('Ütü',1500.0, 10.0, 'ABC TECH'),
('Çamaşır Makinesi',10000.0, 15.0, 'ABC TECH'),
('Lambader',2000.0, 0.0, 'Dekorasyon Sarayı');
`

func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertProductsResult, insertProductsErr := dbPool.Exec(ctx, INSERT_PRODUCTS)
	if insertProductsErr != nil {
		log.Error(insertProductsErr)
	} else {
		log.Info(fmt.Sprintf("Products data created with %d rows", insertProductsResult.RowsAffected()))
	}
}
