const AUTH_KEY = "token"
const UID_KEY = "uid"
const USERNAME_KEY = "username"

function read(key) {
    return localStorage.getItem(key)
}

function write(key, item) {
    if (item) {
        localStorage.setItem(key, item)
    } else {
        console.log("Clearing", key);
        localStorage.removeItem(key);
    }
}

export function readUser() {
    const data = {'uid': read(UID_KEY), 'username': read(USERNAME_KEY), 'token': readToken()};
    return data;
}

export function writeUser(data) {
    if (data) {
        writeUid(data.uid)
        writeUsername(data.username)
        writeToken(data.token)
    } else {
        writeUid()
        writeUsername()
        writeToken()
    }
}

export function writeUid(uid) {
    write(UID_KEY, uid)
}

export function writeUsername(username) {
    write(USERNAME_KEY, username)
}

export function writeToken(token) {
    write(AUTH_KEY, token)
}

export function readToken() {
    return read(AUTH_KEY)
}