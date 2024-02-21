ALTER TABLE IF EXISTS "room_inventory" DROP CONSTRAINT IF EXISTS "room_inventory_room_id_fkey";
ALTER TABLE IF EXISTS "room_inventory" DROP CONSTRAINT IF EXISTS "room_inventory_hotel_id_fkey";

ALTER TABLE IF EXISTS "room" DROP CONSTRAINT IF EXISTS "room_room_type_id_fkey";
ALTER TABLE IF EXISTS "room" DROP CONSTRAINT IF EXISTS "room_hotel_id_fkey";

ALTER TABLE IF EXISTS "reservation" DROP CONSTRAINT IF EXISTS "reservation_hotel_id_fkey";
ALTER TABLE IF EXISTS "reservation" DROP CONSTRAINT IF EXISTS "reservation_user_id_fkey";
ALTER TABLE IF EXISTS "reservation" DROP CONSTRAINT IF EXISTS "reservation_room_id_fkey";

ALTER TABLE IF EXISTS "rate" DROP CONSTRAINT IF EXISTS "rate_room_id_fkey";
ALTER TABLE IF EXISTS "rate" DROP CONSTRAINT IF EXISTS "rate_hotel_id_fkey";

DROP TABLE IF EXISTS room;
DROP TABLE IF EXISTS room_inventory;
DROP TABLE IF EXISTS rate;
DROP TABLE IF EXISTS reservation;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS hotel;
DROP TABLE IF EXISTS room_type;