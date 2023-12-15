package queries

var TablesCheck = []struct {
	TableName   string
	CreateQuery string
}{
	{"users", CreateUserTable},
	{"posts", CreatePostsTable},
	{"comments", CreateCommentsTable},
	{"follows", CreateFollowsTable},
	{"likes", CreateLikesTable},
	{"bans", CreateBansTable},
}

const (
	// User Queries
	CreateUserTable = "CREATE TABLE IF NOT EXISTS users (" +
		"uid TEXT CHECK(LENGTH(uid) = 12)," +
		"username TEXT CHECK(LENGTH(username) >= 3 AND LENGTH(username) <= 20)," +
		"biography TEXT CHECK(LENGTH(biography) <= 100)," +
		"token TEXT," +
		"PRIMARY KEY (uid)" +
		");"
	CreateNewUser       = "INSERT INTO users (uid, username, biography, token) VALUES ($1, $2, \"\", $3);"
	ChangeUsername      = "UPDATE users SET username = $1 WHERE uid = $2;"
	GetUserFromUsername = "SELECT * FROM users WHERE username = $1;"
	GetUserFromUid      = "SELECT * FROM users WHERE uid = $1;"
	GetUserToken        = "SELECT token FROM users WHERE token = $1;"
	GetUseridFromToken  = "SELECT uid FROM users WHERE token = $1;"
	GetUserPosts        = "SELECT * FROM posts WHERE posts.uid = $1;"
	SearchUser          = "SELECT u.uid, u.username FROM users u LEFT JOIN bans b ON" +
		" (u.uid = b.uid2 AND b.uid1 = $1) OR (u.uid = b.uid1 AND b.uid2 = $1)" +
		" WHERE u.username LIKE $2 AND b.uid2 IS NULL LIMIT 50;"
	// TODO: delete profile picture from external source
	//	Implement db.Begin(), tx.Rollback/Commit() logic
	DeleteUser = "DELETE FROM users WHERE uid = $1;" +
		"DELETE FROM posts WHERE uid = $1;" +
		"DELETE FROM comments WHERE uid = $1;" +
		"DELETE FROM follows WHERE uid1 = $1 OR uid2 = $1;" +
		"DELETE FROM likes WHERE uid = $1;" +
		"DELETE FROM bans WHERE uid1 = $1 OR uid2 = $2;"
	// Post Queries
	CreatePostsTable = "CREATE TABLE IF NOT EXISTS posts (" +
		"pid TEXT CHECK(LENGTH(pid) = 15)," +
		"uid TEXT CHECK(LENGTH(uid) = 12)," +
		"caption TEXT CHECK(LENGTH(caption) <= 200)," +
		"createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (pid)," +
		"FOREIGN KEY (uid) REFERENCES users(uid)" +
		");"
	CreateNewPost = "INSERT INTO posts (pid, uid, caption) VALUES ($1, $2, $3);"
	GetBasePost   = "SELECT * FROM posts WHERE pid = $1;"
	// TODO: delete post image from external source
	//	Implement db.Begin(), tx.Rollback/Commit() logic
	DeletePost = "DELETE FROM posts WHERE pid = $1 AND uid = $2;" +
		"DELETE FROM comments WHERE pid = $1;" +
		"DELETE FROM likes WHERE pid = $1;"
	// Comment Queries
	CreateCommentsTable = "CREATE TABLE IF NOT EXISTS comments (" +
		"cid TEXT CHECK(LENGTH(cid) = 15)," +
		"pid TEXT CHECK(LENGTH(pid) = 15)," +
		"uid TEXT CHECK(LENGTH(uid) = 12)," +
		"content TEXT CHECK(LENGTH(content) <= 200 and LENGTH(content) >= 1)," +
		"PRIMARY KEY (cid, pid)," +
		"FOREIGN KEY (uid) REFERENCES users(uid)," +
		"FOREIGN KEY (pid) REFERENCES posts(pid)" +
		");"
	CreateNewComment = "INSERT INTO comments (cid, pid, uid, content) VALUES ($1, $2, $3, $4);"
	GetPostComments  = "SELECT * FROM comments WHERE pid = $1;"
	GetPostComment   = "SELECT * FROM comments WHERE cid = $1 AND pid = $2;"
	ModifyComment    = "UPDATE comments SET content = $1 WHERE cid = $2 AND pid = $3 AND uid = $4;"
	RemoveComment    = "DELETE FROM comments WHERE cid = $1 AND pid = $2 AND uid = $3;"
	// Follow Queries
	CreateFollowsTable = "CREATE TABLE IF NOT EXISTS follows (" +
		"uid1 TEXT CHECK(LENGTH(uid1) = 12)," +
		"uid2 TEXT CHECK(LENGTH(uid2) = 12)," +
		"PRIMARY KEY (uid1, uid2)," +
		"FOREIGN KEY (uid1) REFERENCES users(uid)," +
		"FOREIGN KEY (uid2) REFERENCES users(uid)" +
		");"
	GetFollowers     = "SELECT users.uid, users.username FROM users JOIN follows ON users.uid = follows.uid1 WHERE follows.uid2 = ?;"
	GetFollows       = "SELECT users.uid, users.username FROM users JOIN follows ON users.uid = follows.uid2 WHERE follows.uid1 = ?;"
	IsFollowing      = "SELECT * FROM follows WHERE uid1 = $1 AND uid2 = $2;"
	IsFollowedBy     = "SELECT * FROM follows WHERE uid1 = $2 AND uid2 = $1;"
	FollowUser       = "INSERT INTO follows (uid1, uid2) VALUES ($1, $2);"
	GetFollowerCount = "SELECT COUNT(*) FROM follows WHERE uid1 = $1;"
	GetFollowsCount  = "SELECT COUNT(*) FROM follows WHERE uid2 = $1;"
	UnfollowUser     = "DELETE FROM follows WHERE uid1 = $1 AND uid2 = $2;"
	// Likes Queries
	CreateLikesTable = "CREATE TABLE IF NOT EXISTS likes (" +
		"pid TEXT CHECK(LENGTH(pid) = 15)," +
		"uid TEXT CHECK(LENGTH(uid) = 12)," +
		"PRIMARY KEY (pid, uid)," +
		"FOREIGN KEY (pid) REFERENCES posts(pid)," +
		"FOREIGN KEY (uid) REFERENCES users(uid)" +
		");"
	AddLikeToPost    = "INSERT INTO likes (pid, uid) VALUES ($1, $2);"
	GetNumberOfLikes = "SELECT COUNT(*) FROM likes WHERE pid = $1;"
	GetPostLikes     = "SELECT * FROM likes WHERE pid = $1;"
	RemoveLike       = "DELETE FROM likes WHERE pid = $1 AND uid = $2;"
	// Bans Queries
	CreateBansTable = "CREATE TABLE IF NOT EXISTS bans (" +
		"uid1 TEXT CHECK(LENGTH(uid1) = 12)," +
		"uid2 TEXT CHECK(LENGTH(uid2) = 12)," +
		"PRIMARY KEY (uid1, uid2)," +
		"FOREIGN KEY (uid1) REFERENCES users(uid)," +
		"FOREIGN KEY (uid2) REFERENCES users(uid)" +
		");"
	GetBans = "SELECT users.* FROM users " +
		"JOIN bans ON users.uid = bans.uid2 " +
		"WHERE follows.uid1 = ?;"
	// TODO: $1 can still follow $2 if it bans $2?
	BanUser = "INSERT INTO bans (bans.uid1, bans.uid2) VALUES ($1, $2);" +
		"DELETE FROM follows WHERE follows.uid1 = $2 AND follows.uid2 = $1;" +
		"DELETE FROM follows WHERE follows.uid1 = $1 AND follows.uid2 = $2;"
	UnbanUser = "DELETE FROM bans WHERE uid1 = $1 AND uid2 = $2;"
)
