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
	CreateUserTable = "CREATE TABLE IF NOT EXISTS users (" +
		"uid TEXT CHECK(LENGTH(uid) = 12)," +
		"username TEXT CHECK(LENGTH(username) >= 3 AND LENGTH(username) <= 20)," +
		"biography TEXT CHECK(LENGTH(biography) <= 100)," +
		"token TEXT," +
		"PRIMARY KEY (uid)" +
		");"
	CreateNewUser    = "INSERT INTO users (uid, username, biography, token) VALUES ($1, $2, \"\", $3);"
	CreatePostsTable = "CREATE TABLE IF NOT EXISTS posts (" +
		"pid TEXT CHECK(LENGTH(pid) = 15)," +
		"uid TEXT CHECK(LENGTH(uid) = 12)," +
		"caption TEXT CHECK(LENGTH(caption) <= 200)," +
		"createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (pid)," +
		"FOREIGN KEY (uid) REFERENCES users(uid)" +
		");"
	CreateCommentsTable = "CREATE TABLE IF NOT EXISTS comments (" +
		"cid TEXT CHECK(LENGTH(cid) = 15)," +
		"pid TEXT CHECK(LENGTH(pid) = 15)," +
		"uid TEXT CHECK(LENGTH(uid) = 12)," +
		"content TEXT CHECK(LENGTH(content) <= 200 and LENGTH(content) >= 1)," +
		"PRIMARY KEY (cid, pid)," +
		"FOREIGN KEY (uid) REFERENCES users(uid)," +
		"FOREIGN KEY (pid) REFERENCES posts(pid)" +
		");"
	CreateFollowsTable = "CREATE TABLE IF NOT EXISTS follows (" +
		"uid1 TEXT CHECK(LENGTH(uid1) = 12)," +
		"uid2 TEXT CHECK(LENGTH(uid2) = 12)," +
		"PRIMARY KEY (uid1, uid2)," +
		"FOREIGN KEY (uid1) REFERENCES users(uid)," +
		"FOREIGN KEY (uid2) REFERENCES users(uid)" +
		");"
	CreateLikesTable = "CREATE TABLE IF NOT EXISTS likes (" +
		"pid TEXT CHECK(LENGTH(pid) = 15)," +
		"uid TEXT CHECK(LENGTH(uid) = 12)," +
		"PRIMARY KEY (pid, uid)," +
		"FOREIGN KEY (pid) REFERENCES posts(pid)," +
		"FOREIGN KEY (uid) REFERENCES users(uid)" +
		");"
	CreateBansTable = "CREATE TABLE IF NOT EXISTS bans (" +
		"uid1 TEXT CHECK(LENGTH(uid1) = 12)," +
		"uid2 TEXT CHECK(LENGTH(uid2) = 12)," +
		"PRIMARY KEY (uid1, uid2)," +
		"FOREIGN KEY (uid1) REFERENCES users(uid)," +
		"FOREIGN KEY (uid2) REFERENCES users(uid)" +
		");"
	GetFollower = "SELECT users.* FROM users " +
		"JOIN follows ON users.uid = follows.uid1 " +
		"WHERE follows.uid2 = ?;"
	GetFollowing = "SELECT users.* FROM users " +
		"JOIN follows ON users.uid = follows.uid2 " +
		"WHERE follows.uid1 = ?;"
	GetBans = "SELECT users.* FROM users " +
		"JOIN bans ON users.uid = bans.uid2 " +
		"WHERE follows.uid1 = ?;"
)
