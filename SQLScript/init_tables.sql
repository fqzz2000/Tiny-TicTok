USE tictok;
-- ALTER TABLE relations DROP CONSTRAINT `fk_relations_user1`;
-- ALTER TABLE relations DROP CONSTRAINT `fk_relations_user2`;
-- ALTER TABLE comments DROP CONSTRAINT `fk_comments_users`;
-- ALTER TABLE comments DROP CONSTRAINT `fk_comments_videos`;
-- ALTER TABLE likes DROP CONSTRAINT `fk_likes_users`;
-- ALTER TABLE likes DROP CONSTRAINT `fk_likes_videos`;
-- ALTER TABLE videos DROP CONSTRAINT `fk_videos_users`;
-- DROP TABLE IF EXISTS relations;
-- DROP TABLE IF EXISTS comments;
-- DROP TABLE IF EXISTS likes;
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS videos;


-- create users
CREATE TABLE users (
user_id 		BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
user_name 		char(255) 		NOT NULL,
user_pswd		char(40) 		NOT NULL,
PRIMARY KEY (user_id)
) ENGINE = innoDB;


-- create videos
CREATE TABLE videos (
video_id 		BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
video_title 	CHAR(255) 		NOT NULL,
video_desc		TEXT			NULL,
video_owner		BIGINT UNSIGNED NOT NULL,
video_crt_time	DATETIME 		NOT NULL,
video_file		TEXT			NOT NULL,
cover_file 		TEXT 			NOT NULL,
PRIMARY KEY (video_id),
CONSTRAINT `fk_videos_users`
FOREIGN KEY (video_owner) 
REFERENCES users(user_id)
ON DELETE CASCADE
) ENGINE = innoDB;

-- create relations
CREATE TABLE relations (
follower_id 	BIGINT UNSIGNED NOT NULL, 
fans_id 		BIGINT UNSIGNED NOT NULL,
follow_date		DATETIME 		NOT NULL,
PRIMARY KEY (follower_id, fans_id),
CONSTRAINT `fk_relations_users1`
FOREIGN KEY (follower_id) 
REFERENCES users(user_id) 
ON DELETE CASCADE,
CONSTRAINT `fk_relations_users2`
FOREIGN KEY (fans_id)
REFERENCES users(user_id)
ON DELETE CASCADE
) ENGINE = innoDB;

-- create comments
-- use comment_crt_time and comment_video_id as joint index
CREATE TABLE comments (
comment_id 			BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
comment_user_id 	BIGINT UNSIGNED NOT NULL,
comment_video_id	BIGINT UNSIGNED NOT NULL,
comment_content		TEXT 			NOT NULL,
comment_crt_time	DATETIME		NOT NULL, 
PRIMARY KEY (comment_id),
CONSTRAINT `fk_comments_users`
FOREIGN KEY (comment_user_id)
REFERENCES users(user_id) 
ON DELETE CASCADE,
CONSTRAINT `fk_comments_videos`
FOREIGN KEY (comment_video_id) 
REFERENCES videos(video_id) 
ON DELETE CASCADE
) ENGINE = innoDB;

-- create likes
CREATE TABLE likes (
user_id 		BIGINT UNSIGNED NOT NULL,
video_id 		BIGINT UNSIGNED NOT NULL,
like_time 		DATETIME 		NOT NULL,
PRIMARY KEY (user_id, video_id),
CONSTRAINT `fk_likes_users`
FOREIGN KEY (user_id)
REFERENCES users(user_id)
ON DELETE CASCADE, 
CONSTRAINT `fk_likes_videos`
FOREIGN KEY (video_id)
REFERENCES videos(video_id)
ON DELETE CASCADE
) ENGINE = innoDB;
