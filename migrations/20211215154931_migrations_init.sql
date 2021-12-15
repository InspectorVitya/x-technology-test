-- +goose Up
CREATE TABLE IF NOT EXISTS  stock_quotes (
            symbol varchar NOT NULL,
            price numeric NOT NULL,
            volume numeric NOT NULL,
            last_trade numeric NULL
);


-- +goose Down
DROP TABLE stock_quotes;

