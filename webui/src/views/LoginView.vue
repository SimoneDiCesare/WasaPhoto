<template>
  <div class="login-container">
    <div class="login-box">
      <h2>Login</h2>
      <form @submit.prevent="login">
        <div class="input-group">
          <label for="username">Username</label>
          <input type="username" id="username" v-model="username" required />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  </div>
</template>

<script>
import {readUser, writeUser} from '../services/session'
import api from '../services/axios'
import router from "../router";
export default {
  data() {
    return {
      username: ''
    };
  },
  methods: {
    async login() {
      console.log('Username:', this.username);
	  await api.post('/login', {
		username: this.username
	  }).then((response) => {
		console.log(response)
		if (response.status < 300) {
			writeUser(response.data)
			router.push(`/users/${readUser().uid}`);
		} else {
			console.log(response.status)
		}
	  }).catch((error) => {
		console.log(error)
	  });
    }
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
}

.login-box {
  background-color: white;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 300px;
}

.input-group {
  margin-bottom: 15px;
  text-align: left;
}

.input-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

.input-group input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
}

button:hover {
  background-color: #38a174;
}
</style>
