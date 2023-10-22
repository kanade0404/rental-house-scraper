-- name: CreateChangeRentRoomPriceEvent :one
INSERT INTO change_rent_room_price_event (monthly_rent, security_deposit, key_money, brokerage_fee, renewal_fee, guarantee_fee, common_service_fee, room_id, recorded_at)
SELECT $1, $2, $3, $4, $5, $6, $7, $8, $9
WHERE NOT EXISTS (
    SELECT room_id
    FROM change_rent_room_price_event
    WHERE room_id = $8
)
    RETURNING id, monthly_rent, security_deposit, key_money, brokerage_fee, renewal_fee, guarantee_fee, common_service_fee, room_id, recorded_at;
