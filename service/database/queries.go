package database

var TablesCheck = []struct {
	TableName   string
	CreateQuery string
}{
	{"users", CreateUsersTalbe},
	{"posts", CreatePostsTalbe},
	{"comments", CreateCommentsTable},
	{"likes", CreateLikesTable},
	{"follows", CreateFollowsTable},
	{"bans", CreateBansTable},
}

const (
	// DB Creations Queries
	CreateUsersTalbe = "CREATE TABLE IF NOT EXISTS users (" +
		"id TEXT CHECK(LENGTH(id) >= 1 AND LENGTH(id) <= 16)," +
		"username TEXT CHECK(LENGTH(username) >= 3 AND LENGTH(username) <= 16)," +
		"token TEXT," +
		"PRIMARY KEY (id)" +
		");"
	CreatePostsTalbe = "CREATE TABLE IF NOT EXISTS posts (" +
		"id TEXT CHECK(LENGTH(id) >= 1 AND LENGTH(id) <= 16)," +
		"uid TEXT CHECK(LENGTH(uid) >= 1 AND LENGTH(uid) <= 16)," +
		"createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (id)," +
		"FOREIGN KEY (uid) REFERENCES users(id)" +
		");"
	CreateCommentsTable = "CREATE TABLE IF NOT EXISTS comments (" +
		"id TEXT CHECK(LENGTH(id) >= 1 AND LENGTH(id) <= 16)," +
		"uid TEXT CHECK(LENGTH(uid) >= 1 AND LENGTH(uid) <= 16)," +
		"pid TEXT CHECK(LENGTH(pid) >= 1 AND LENGTH(pid) <= 16)," +
		"content TEXT CHECK(LENGTH(content) >= 1 AND LENGTH(content) <= 256)," +
		"PRIMARY KEY (id)," +
		"FOREIGN KEY (uid) REFERENCES users(id)" +
		"FOREIGN KEY (pid) REFERENCES posts(id)" +
		");"
	CreateLikesTable = "CREATE TABLE IF NOT EXISTS likes (" +
		"uid TEXT CHECK(LENGTH(uid) >= 1 AND LENGTH(uid) <= 16)," +
		"pid TEXT CHECK(LENGTH(pid) >= 1 AND LENGTH(pid) <= 16)," +
		"PRIMARY KEY (uid, pid)," +
		"FOREIGN KEY (uid) REFERENCES users(id)" +
		"FOREIGN KEY (pid) REFERENCES posts(id)" +
		");"
	CreateFollowsTable = "CREATE TABLE IF NOT EXISTS follows (" +
		"follower TEXT CHECK(LENGTH(follower) >= 1 AND LENGTH(follower) <= 16)," +
		"followed TEXT CHECK(LENGTH(followed) >= 1 AND LENGTH(followed) <= 16)," +
		"CHECK(follower != followed)," +
		"PRIMARY KEY (follower, followed)," +
		"FOREIGN KEY (follower) REFERENCES users(follower)" +
		"FOREIGN KEY (followed) REFERENCES posts(followed)" +
		");"
	CreateBansTable = "CREATE TABLE IF NOT EXISTS bans (" +
		"banner TEXT CHECK(LENGTH(banner) >= 1 AND LENGTH(banner) <= 16)," +
		"banned TEXT CHECK(LENGTH(banned) >= 1 AND LENGTH(banned) <= 16)," +
		"CHECK(banner != banned)," +
		"PRIMARY KEY (banner, banned)," +
		"FOREIGN KEY (banner) REFERENCES users(banner)" +
		"FOREIGN KEY (banned) REFERENCES posts(banned)" +
		");"
	// users Table Queries
	GetUsers            = "SELECT users.id, users.username, users.token FROM users WHERE 1;"
	CreateUser          = "INSERT INTO users (id, username, token) VALUES ($1, $2, $3);"
	GetSimpleUserFromId = "SELECT users.id, users.username FROM users WHERE users.id = $1;"
	GetUserById         = "SELECT users.id, users.username, users.token FROM users WHERE users.id = $1;"
	GetUserByName       = "SELECT users.id, users.username, users.token FROM users WHERE users.username = $1;"
	GetUidByToken       = "SELECT users.id FROM users WHERE users.token = $1;"
	UpdateUserToken     = "UPDATE users SET token = $1 WHERE users.id = $2;"
	UpdateUserName      = "UPDATE users SET username = $1 WHERE users.id = $2;"
	SearchUsersByName   = "SELECT users.id, users.username FROM users WHERE users.id NOT IN (SELECT bans.banner FROM bans WHERE bans.banned = $1) AND users.username LIKE $2 || '%' LIMIT 20;"
	GetFeeds            = "SELECT posts.id, posts.uid, users.username, posts.createdAt FROM posts INNER JOIN users ON posts.uid = users.id WHERE users.id IN (SELECT follows.followed FROM follows WHERE follows.follower = $1) ORDER BY posts.createdAt DESC;"
	// follow table query
	GetFollows   = "SELECT users.id, users.username FROM users INNER JOIN follows ON users.id = follows.followed WHERE follows.follower = $1;"
	GetFollowers = "SELECT users.id, users.username FROM users INNER JOIN follows ON users.id = follows.follower WHERE follows.followed = $1;"
	FollowUser   = "INSERT INTO follows (follower, followed) VALUES ($1, $2);"
	UnfollowUser = "DELETE FROM follows WHERE follows.follower = $1 AND follows.followed = $2;"
	// ban table query
	GetBans   = "SELECT users.id, users.username FROM users INNER JOIN bans ON users.id = bans.banned WHERE bans.banner = $1;"
	BanUser   = "INSERT INTO bans (banner, banned) VALUES ($1, $2);"
	UnbanUser = "DELETE FROM bans WHERE bans.banner = $1 AND bans.banned = $2;"
	// post table query
	CreatePost        = "INSERT INTO posts (id, uid) VALUES ($1, $2);"
	GetPostIdFromId   = "SELECT posts.id FROM posts WHERE posts.id = $1;"
	GetSimplePost     = "SELECT posts.id, posts.uid, users.username, posts.createdAt FROM posts INNER JOIN users ON posts.uid = users.id WHERE posts.id = $1;"
	GetUserPosts      = "SELECT posts.id, posts.uid, users.username, posts.createdAt FROM posts INNER JOIN users ON posts.uid = users.id WHERE posts.uid = $1 ORDER BY posts.createdAt DESC;"
	DeletePhoto       = "DELETE FROM posts WHERE posts.id = $1 AND posts.uid = $2;"
	LikePost          = "INSERT INTO likes (uid, pid) VALUES ($1, $2);"
	UnlikePost        = "DELETE FROM likes WHERE likes.uid = $1 AND likes.pid = $1;"
	CheckLike         = "SELECT likes.pid FROM likes WHERE likes.uid = $1 AND likes.pid = $1;"
	GetPostLikesCount = "SELECT COUNT(*) AS likes_count FROM likes WHERE likes.pid = $1;"
	// comment table query
	GetPostComments    = "SELECT comments.id, comments.pid, users.id, users.username, comments.content FROM comments INNER JOIN users ON comments.uid = users.id WHERE comments.pid = $1;"
	GetCommentIdFromId = "SELECT comments.id FROM comments WHERE comments.id = $1;"
	CommentPhoto       = "INSET INTO comments (id, pid, uid, content) VALUES ($1, $2, $3, $4);"
	UncommentPhoto     = "DELETE FROM comments WHERE comments.id = $1;"
)
