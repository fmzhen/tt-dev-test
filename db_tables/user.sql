/*
 Navicat Premium Data Transfer

 Source Server         : local_pg
 Source Server Type    : PostgreSQL
 Source Server Version : 100300
 Source Host           : localhost
 Source Database       : tt-dev-test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100300
 File Encoding         : utf-8

 Date: 05/09/2018 18:16:57 PM
*/

-- ----------------------------
--  Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
	"id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
	"name" varchar(64) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."user" OWNER TO "fmzhen";

-- ----------------------------
--  Primary key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Indexes structure for table user
-- ----------------------------
CREATE UNIQUE INDEX  "user_id_key" ON "public"."user" USING btree("id" "pg_catalog"."int4_ops" ASC NULLS LAST);

