-- CREATE TABLE "posts" ----------------------------------------
CREATE TABLE `posts` ( 
	`id` Int( 255 ) AUTO_INCREMENT NOT NULL,
	`name` VarChar( 255 ) NOT NULL,
	`slug` VarChar( 255 ) NOT NULL,
	`content` VarChar( 255 ) NOT NULL,
	PRIMARY KEY ( `id` ) )
ENGINE = InnoDB;
-- -------------------------------------------------------------
