/*
 Navicat Premium Dump SQL

 Source Server         : LearnPostgreSQL
 Source Server Type    : PostgreSQL
 Source Server Version : 170004 (170004)
 Source Host           : localhost:5432
 Source Catalog        : bookings
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 170004 (170004)
 File Encoding         : 65001

 Date: 14/04/2025 22:51:49
*/


-- ----------------------------
-- Sequence structure for reservations_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."reservations_id_seq";
CREATE SEQUENCE "public"."reservations_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for restrictions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."restrictions_id_seq";
CREATE SEQUENCE "public"."restrictions_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for room_restrictions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."room_restrictions_id_seq";
CREATE SEQUENCE "public"."room_restrictions_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for rooms_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."rooms_id_seq";
CREATE SEQUENCE "public"."rooms_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for reservations
-- ----------------------------
DROP TABLE IF EXISTS "public"."reservations";
CREATE TABLE "public"."reservations" (
  "id" int4 NOT NULL DEFAULT nextval('reservations_id_seq'::regclass),
  "first_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "last_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "phone" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "room_id" int4 NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "processed" int4 NOT NULL DEFAULT 0
)
;

-- ----------------------------
-- Records of reservations
-- ----------------------------
INSERT INTO "public"."reservations" VALUES (2, 'john', 'go', 'test@g.com', '555-555-1234', '2025-05-01', '2025-06-01', 1, '2025-04-10 22:40:09.055871', '2025-04-10 22:40:09.055871', 0);
INSERT INTO "public"."reservations" VALUES (4, 'test', 'test', 'test@gmai.com', '555-55-555', '2025-04-11', '2025-04-12', 1, '2025-04-11 21:44:45.235173', '2025-04-11 21:44:45.235173', 0);
INSERT INTO "public"."reservations" VALUES (5, 'test', 'test', 'test@gmail.com', '555-555-555', '2025-04-26', '2025-04-27', 1, '2025-04-12 10:19:48.29721', '2025-04-12 10:19:48.29721', 0);
INSERT INTO "public"."reservations" VALUES (6, 'test', 'test', 'test@gmail.com', '555-555-555', '2025-04-30', '2025-05-01', 1, '2025-04-12 10:39:14.837786', '2025-04-12 10:39:14.837786', 0);
INSERT INTO "public"."reservations" VALUES (3, 'Test001', 'Test002', 'test@mgil.com', '555-555-555', '2025-04-10', '2025-04-10', 1, '2025-04-11 17:54:45.40366', '2025-04-14 22:03:45.06831', 0);

-- ----------------------------
-- Table structure for restrictions
-- ----------------------------
DROP TABLE IF EXISTS "public"."restrictions";
CREATE TABLE "public"."restrictions" (
  "id" int4 NOT NULL DEFAULT nextval('restrictions_id_seq'::regclass),
  "restriction_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
)
;

-- ----------------------------
-- Records of restrictions
-- ----------------------------
INSERT INTO "public"."restrictions" VALUES (1, 'Reservation', '2025-04-10 22:34:46', '2025-04-10 22:34:49');
INSERT INTO "public"."restrictions" VALUES (2, 'Owner Block', '2025-04-11 15:39:46', '2025-04-11 15:39:48');

-- ----------------------------
-- Table structure for room_restrictions
-- ----------------------------
DROP TABLE IF EXISTS "public"."room_restrictions";
CREATE TABLE "public"."room_restrictions" (
  "id" int4 NOT NULL DEFAULT nextval('room_restrictions_id_seq'::regclass),
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "room_id" int4 NOT NULL,
  "reservation_id" int4,
  "restriction_id" int4 NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
)
;

-- ----------------------------
-- Records of room_restrictions
-- ----------------------------
INSERT INTO "public"."room_restrictions" VALUES (1, '2021-02-01', '2021-02-04', 1, 2, 1, '2025-04-10 22:40:09.058295', '2025-04-10 22:40:09.058295');
INSERT INTO "public"."room_restrictions" VALUES (2, '2021-02-01', '2021-02-04', 2, NULL, 2, '2025-04-11 16:43:29', '2025-04-11 16:43:31');
INSERT INTO "public"."room_restrictions" VALUES (3, '2025-04-10', '2025-04-10', 1, 3, 1, '2025-04-11 17:54:45.408183', '2025-04-11 17:54:45.408183');
INSERT INTO "public"."room_restrictions" VALUES (4, '2025-04-11', '2025-04-12', 1, 4, 1, '2025-04-11 21:44:45.236469', '2025-04-11 21:44:45.236469');
INSERT INTO "public"."room_restrictions" VALUES (5, '2025-04-26', '2025-04-27', 1, 5, 1, '2025-04-12 10:19:48.30518', '2025-04-12 10:19:48.30518');
INSERT INTO "public"."room_restrictions" VALUES (6, '2025-04-30', '2025-05-01', 1, 6, 1, '2025-04-12 10:39:14.840105', '2025-04-12 10:39:14.840105');

-- ----------------------------
-- Table structure for rooms
-- ----------------------------
DROP TABLE IF EXISTS "public"."rooms";
CREATE TABLE "public"."rooms" (
  "id" int4 NOT NULL DEFAULT nextval('rooms_id_seq'::regclass),
  "room_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
)
;

-- ----------------------------
-- Records of rooms
-- ----------------------------
INSERT INTO "public"."rooms" VALUES (1, 'General', '2025-04-10 22:21:12', '2025-04-10 22:21:18');
INSERT INTO "public"."rooms" VALUES (2, 'Major', '2025-04-11 16:41:09', '2025-04-11 16:41:12');

-- ----------------------------
-- Table structure for schema_migration
-- ----------------------------
DROP TABLE IF EXISTS "public"."schema_migration";
CREATE TABLE "public"."schema_migration" (
  "version" varchar(14) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of schema_migration
-- ----------------------------
INSERT INTO "public"."schema_migration" VALUES ('20250410075836');
INSERT INTO "public"."schema_migration" VALUES ('20250410081207');
INSERT INTO "public"."schema_migration" VALUES ('20250410081407');
INSERT INTO "public"."schema_migration" VALUES ('20250410081547');
INSERT INTO "public"."schema_migration" VALUES ('20250410081643');
INSERT INTO "public"."schema_migration" VALUES ('20250410083137');
INSERT INTO "public"."schema_migration" VALUES ('20250410084451');
INSERT INTO "public"."schema_migration" VALUES ('20250410092653');
INSERT INTO "public"."schema_migration" VALUES ('20250410092951');
INSERT INTO "public"."schema_migration" VALUES ('20250410094013');
INSERT INTO "public"."schema_migration" VALUES ('20250411084135');
INSERT INTO "public"."schema_migration" VALUES ('20250414140430');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "first_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "last_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(60) COLLATE "pg_catalog"."default" NOT NULL,
  "access_level" int4 NOT NULL DEFAULT 1,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES (1, 'a', 'b', 'test@admin.com', '$2a$12$t1ok2B6FR9KFOD0iBuGLCO/keESK9kJScR8M3LUIbJNlmYn7R0rcO', 3, '2025-04-13 10:29:49', '2025-04-13 10:29:53');

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."reservations_id_seq"
OWNED BY "public"."reservations"."id";
SELECT setval('"public"."reservations_id_seq"', 6, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."restrictions_id_seq"
OWNED BY "public"."restrictions"."id";
SELECT setval('"public"."restrictions_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."room_restrictions_id_seq"
OWNED BY "public"."room_restrictions"."id";
SELECT setval('"public"."room_restrictions_id_seq"', 6, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."rooms_id_seq"
OWNED BY "public"."rooms"."id";
SELECT setval('"public"."rooms_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq"', 1, true);

-- ----------------------------
-- Indexes structure for table reservations
-- ----------------------------
CREATE INDEX "reservations_email_idx" ON "public"."reservations" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "reservations_last_name_idx" ON "public"."reservations" USING btree (
  "last_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table reservations
-- ----------------------------
ALTER TABLE "public"."reservations" ADD CONSTRAINT "reservations_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table restrictions
-- ----------------------------
ALTER TABLE "public"."restrictions" ADD CONSTRAINT "restrictions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table room_restrictions
-- ----------------------------
CREATE INDEX "room_restrictions_reservation_id_idx" ON "public"."room_restrictions" USING btree (
  "reservation_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "room_restrictions_room_id_idx" ON "public"."room_restrictions" USING btree (
  "room_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "room_restrictions_start_date_end_date_idx" ON "public"."room_restrictions" USING btree (
  "start_date" "pg_catalog"."date_ops" ASC NULLS LAST,
  "end_date" "pg_catalog"."date_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table room_restrictions
-- ----------------------------
ALTER TABLE "public"."room_restrictions" ADD CONSTRAINT "room_restrictions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table rooms
-- ----------------------------
ALTER TABLE "public"."rooms" ADD CONSTRAINT "rooms_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table schema_migration
-- ----------------------------
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "public"."schema_migration" USING btree (
  "version" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE UNIQUE INDEX "users_email_idx" ON "public"."users" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table reservations
-- ----------------------------
ALTER TABLE "public"."reservations" ADD CONSTRAINT "reservations_rooms_id_fk" FOREIGN KEY ("room_id") REFERENCES "public"."rooms" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- ----------------------------
-- Foreign Keys structure for table room_restrictions
-- ----------------------------
ALTER TABLE "public"."room_restrictions" ADD CONSTRAINT "room_restrictions_reservations_id_fk" FOREIGN KEY ("reservation_id") REFERENCES "public"."reservations" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "public"."room_restrictions" ADD CONSTRAINT "room_restrictions_restrictions_id_fk" FOREIGN KEY ("restriction_id") REFERENCES "public"."restrictions" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "public"."room_restrictions" ADD CONSTRAINT "room_restrictions_rooms_id_fk" FOREIGN KEY ("room_id") REFERENCES "public"."rooms" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
