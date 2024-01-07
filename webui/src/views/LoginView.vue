<script>
export default {
    data: function() {
        return {
            username: "",
            status_msg: "",
        }
    },
    methods: {
        async login() {
            try {
                const headers = {
                    'Content-Type': 'application/json'
                };
                const data = {
                    'username': this.username,
                };
                let request = await this.$axios.post("/session", data, {
                    headers: headers
                });
                if (request.status == 200 || request.status == 201) {
                    console.log(this.user);
                    localStorage.setItem('user', JSON.stringify(request.data));
                    window.location = '/#/users/' + request.data.uid;
                } else {
                    this.status_msg = "Login Failed";
                    console.log(request.data);
                }
            } catch (e) {
                console.log(e);
            }
        },
    },
    mounted() {
        console.log('Mounted');
    }
}
</script>

<template>
    <div class="login-container">
        <h1 class="page-title">Login</h1>
        <form @submit.prevent="login">
            <label for="username">Username:</label>
            <input type="text" id="username" v-model="username" required />
            <div>{{status_msg}}</div>
            <button type="submit">Login</button>
        </form>
  </div>
</template>

<style scoped>
.login-container {
    display: flex;
    flex-direction: column;
    align-items: center; /* Center horizontally */
    justify-content: center; /* Center vertically */
    height: 100vh; /* Full height of the viewport */
    max-width: 400px;
    margin: auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 8px;
}

.page-title {
    margin-top: 0;
}

form {
    display: flex;
    flex-direction: column;
}

label {
    margin-bottom: 8px;
}

input {
    padding: 8px;
    margin-bottom: 16px;
}

button {
    padding: 10px;
    background-color: #3498db;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

button:hover {
    background-color: #2980b9;
}
</style>