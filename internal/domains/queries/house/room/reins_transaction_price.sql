-- name: CreateReinsTransactionPrice :one
INSERT INTO reins_transaction_price (room_id, price, start_month, end_month)
SELECT $1, $2, $3, $4
WHERE NOT EXISTS (
    SELECT room_id
    FROM reins_transaction_price
    WHERE room_id = $1 AND start_month = $3 AND end_month = $4
)
    RETURNING room_id, price, start_month, end_month;
