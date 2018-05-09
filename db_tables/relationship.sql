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

 Date: 05/09/2018 18:16:25 PM
*/

-- ----------------------------
--  Table structure for relationship
-- ----------------------------
DROP TABLE IF EXISTS "public"."relationship";
CREATE TABLE "public"."relationship" (
	"id" int4 NOT NULL DEFAULT nextval('relationship_id_seq'::regclass),
	"user_id" int4 NOT NULL,
	"other_user_id" int4 NOT NULL,
	"state" varchar NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."relationship" OWNER TO "fmzhen";

COMMENT ON COLUMN "public"."relationship"."user_id" IS '主语';
COMMENT ON COLUMN "public"."relationship"."other_user_id" IS '谓语';
COMMENT ON COLUMN "public"."relationship"."state" IS 'liked | disliked | matched';

-- ----------------------------
--  Primary key structure for table relationship
-- ----------------------------
ALTER TABLE "public"."relationship" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Uniques structure for table relationship
-- ----------------------------
ALTER TABLE "public"."relationship" ADD CONSTRAINT "relation_uniq" UNIQUE ("user_id","other_user_id") NOT DEFERRABLE INITIALLY IMMEDIATE;

-- ----------------------------
--  Foreign keys structure for table relationship
-- ----------------------------
ALTER TABLE "public"."relationship" ADD CONSTRAINT "first_user_for" FOREIGN KEY ("user_id") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "public"."relationship" ADD CONSTRAINT "second_user_for" FOREIGN KEY ("other_user_id") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE;

