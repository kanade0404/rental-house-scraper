-- スクレイピング対応サイト
CREATE TABLE
  web_site (id SERIAL PRIMARY KEY, host VARCHAR(255) NOT NULL);

CREATE UNIQUE INDEX unique_web_site_host ON web_site (host);

-- スクレイピング対象ページ
CREATE TABLE
  page (
    id SERIAL PRIMARY KEY,
    -- ページURLのsearch部
    search VARCHAR(255) NOT NULL,
    web_site_id INTEGER NOT NULL,
    CONSTRAINT fk_page_web_site_id FOREIGN KEY (web_site_id) REFERENCES web_site (id)
  );

CREATE UNIQUE INDEX unique_page_web_site_id_search ON page (web_site_id, search);

-- 路線
CREATE TABLE
  train (name VARCHAR(255) PRIMARY KEY);

-- 駅
CREATE TABLE
  station (name VARCHAR(255) PRIMARY KEY);

-- 路線と駅の関連
CREATE TABLE
  station_train (
    station VARCHAR(255) NOT NULL,
    train VARCHAR(255) NOT NULL,
    CONSTRAINT fk_station_train_station FOREIGN KEY (station) REFERENCES station (name),
    CONSTRAINT fk_station_train_train FOREIGN KEY (train) REFERENCES train (name)
  );

CREATE UNIQUE INDEX unique_station_train_station_train ON station_train (station, train);

-- 移動手段
CREATE TABLE
  movement_method (name VARCHAR(255) PRIMARY KEY);

-- 都道府県
CREATE TABLE
  prefecture (name VARCHAR(255) PRIMARY KEY);

-- 市区町村
CREATE TABLE
  city (
    name VARCHAR(255) PRIMARY KEY,
    prefecture VARCHAR(255) NOT NULL,
    CONSTRAINT fk_city_prefecture FOREIGN KEY (prefecture) REFERENCES prefecture (name)
  );

CREATE UNIQUE INDEX unique_city_name_prefecture ON city (name, prefecture);

-- 住所
CREATE TABLE
  address (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    CONSTRAINT fk_address_city FOREIGN KEY (city) REFERENCES city (name)
  );

CREATE UNIQUE INDEX unique_address_name_city ON address (name, city);

-- 物件所有者
CREATE TABLE
  owner (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    -- 電話番号
    tel VARCHAR(255) NOT NULL,
    -- 住所
    address_id INTEGER NOT NULL,
    CONSTRAINT fk_owner_address_id FOREIGN KEY (address_id) REFERENCES address (id)
  );

CREATE UNIQUE INDEX unique_owner_name_tel_address_id ON owner (name, tel, address_id);

-- 構造
CREATE TABLE
  structure (name VARCHAR(255) PRIMARY KEY);

-- 物件
CREATE TABLE
  house (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    -- 建築年(西暦)
    construction_year INTEGER NOT NULL,
    address_id INTEGER NOT NULL,
    owner_id INTEGER NOT NULL,
    structure VARCHAR(255) NOT NULL,
    CONSTRAINT fk_house_address_id FOREIGN KEY (address_id) REFERENCES address (id),
    CONSTRAINT fk_house_owner_id FOREIGN KEY (owner_id) REFERENCES owner (id),
    CONSTRAINT fk_room_metadata_structure FOREIGN KEY (structure) REFERENCES structure (name)
  );

CREATE UNIQUE INDEX unique_house_name_owner_id_address_id ON house (name, owner_id, address_id);

-- 物件の賃貸メタデータ

CREATE TABLE house_rent_metadata (
    house_id INTEGER PRIMARY KEY,
    -- 戸数
    num_rooms INTEGER NOT NULL,
    CONSTRAINT fk_house_rent_metadata_house_id FOREIGN KEY (house_id) REFERENCES house (id)
);

-- 地上の階建て
CREATE TABLE
  story (
    house_id INTEGER PRIMARY KEY,
    num INTEGER NOT NULL,
    CONSTRAINT fk_story_house_id FOREIGN KEY (house_id) REFERENCES house (id)
  );

-- 地下の階建て
CREATE TABLE
  basement_story (
    house_id INTEGER PRIMARY KEY,
    num INTEGER NOT NULL
  );

-- 物件と路線の関連
CREATE TABLE
  access (
    station VARCHAR(255) NOT NULL,
    movement_method VARCHAR(255) NOT NULL,
    time_in_minutes INTEGER NOT NULL,
    house_id INTEGER NOT NULL,
    CONSTRAINT fk_access_station FOREIGN KEY (station) REFERENCES station (name),
    CONSTRAINT fk_access_movement_method FOREIGN KEY (movement_method) REFERENCES movement_method (name),
    CONSTRAINT fk_access_house_id FOREIGN KEY (house_id) REFERENCES house (id)
  );

CREATE UNIQUE INDEX unique_access_station_house_id_movement_method_id ON access (station, house_id, movement_method);

-- 間取りの種類
CREATE TABLE
  floor_plan_type (name VARCHAR(255) PRIMARY KEY);

-- 間取り
CREATE TABLE
  floor_plan (
    name VARCHAR(255) PRIMARY KEY,
    floor_plan_type VARCHAR(255) NOT NULL,
    CONSTRAINT fk_floor_plan_floor_plan_type FOREIGN KEY (floor_plan_type) REFERENCES floor_plan_type (name)
  );

-- 設備
CREATE TABLE
  equipment (name VARCHAR(255) PRIMARY KEY);

-- 部屋
CREATE TABLE
  room (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    house_id INTEGER NOT NULL,
    CONSTRAINT fk_room_house_id FOREIGN KEY (house_id) REFERENCES house (id)
  );

CREATE UNIQUE INDEX unique_room_name_house_id_page_id ON room (name, house_id);

-- 部屋のメタデータ
CREATE TABLE
  room_metadata (
    room_id INTEGER PRIMARY KEY,
    -- 方角
    direction VARCHAR(255) NOT NULL,
    -- 平米数(平方メートル)
    area_square_meter NUMERIC NOT NULL,
    -- 階
    floor INTEGER NOT NULL,
    -- 間取り
    floor_plan VARCHAR(255) NOT NULL,
    CONSTRAINT fk_room_metadata_room_id FOREIGN KEY (room_id) REFERENCES room (id),
    CONSTRAINT fk_room_metadata_floor_plan FOREIGN KEY (floor_plan) REFERENCES floor_plan (name)
  );

-- 部屋と設備の関連
CREATE TABLE
  room_equipment (
    room_id INTEGER NOT NULL,
    equipment VARCHAR(255) NOT NULL,
    CONSTRAINT fk_room_equipment_room_id FOREIGN KEY (room_id) REFERENCES room (id),
    CONSTRAINT fk_room_equipment_equipment FOREIGN KEY (equipment) REFERENCES equipment (name)
  );

CREATE UNIQUE INDEX unique_room_equipment_room_id_equipment ON room_equipment (room_id, equipment);

-- 部屋の価格
CREATE TABLE
  change_rent_room_price_event (
    id SERIAL PRIMARY KEY,
    -- 家賃
    monthly_rent INTEGER NOT NULL,
    -- 敷金
    security_deposit INTEGER NOT NULL,
    -- 礼金
    key_money INTEGER NOT NULL,
    -- 仲介手数料
    brokerage_fee INTEGER NOT NULL,
    -- 更新料
    renewal_fee INTEGER NOT NULL,
    -- 保証金
    guarantee_fee INTEGER NOT NULL,
    -- 共益費
    common_service_fee INTEGER NOT NULL,
    room_id INTEGER NOT NULL,
    -- 記録日時
    recorded_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_rent_room_price_room_id FOREIGN KEY (room_id) REFERENCES room (id)
  );

-- 部屋が空いた日時
CREATE TABLE
  open_room_event (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL,
    -- 記録日時
    recorded_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_open_room_event_room_id FOREIGN KEY (room_id) REFERENCES room (id)
  );

-- 部屋が埋まった日時
CREATE TABLE
  close_room_event (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL,
    -- 記録日時
    recorded_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_close_room_event_room_id FOREIGN KEY (room_id) REFERENCES room (id)
  );

-- REINSの不動産取引価格
CREATE TABLE
  reins_transaction_price (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL,
    -- 取引価格
    price INTEGER NOT NULL,
    -- 制約開始年月(YYYY-MM-01)
    start_month DATE NOT NULL,
    -- 制約終了年月(YYYY-MM-01)
    end_month DATE NOT NULL,
    CONSTRAINT fk_transaction_price_room_id FOREIGN KEY (room_id) REFERENCES room (id)
  );
