<template>
  <div class="homepage-container">
    <!-- Sezione di ricerca -->
    <div class="search-bar">
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Cerca utenti..."
      />
      <button @click="searchUsers">Cerca</button>
    </div>

    <!-- Sezione dei post -->
    <div class="posts-grid">
      <div
        v-for="(post, index) in posts"
        :key="index"
        class="post-item"
      >
        <div class="post-author">{{ post.author.username }}</div>
        <img :src="post.imageUrl" alt="Post Image" />
      </div>
    </div>

    <!-- Modale per la lista di utenti -->
    <div v-if="showUserModal" class="modal">
      <div class="modal-content">
        <!-- Bottone di chiusura in alto a sinistra -->
        <button class="close-left" @click="closeModal">X</button>
        <span class="close" @click="closeModal">&times;</span>
        <h2>Risultati della ricerca</h2>
        <ul>
          <li v-for="user in users" :key="user.id" @click="goToUserPage(user.id)">
            {{ user.username }}
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>


<script>
import api from '../services/axios'
import {readToken, readUser, writeUser} from '../services/session'

export default {
  data() {
    return {
      searchQuery: '', // Testo per la ricerca
      showUserModal: false,
      posts: [], // Array per i post
    };
  },

  methods: {
    async searchUsers() {
      
    },

    async loadStream() {
      await api.get("/users/" + readUser().uid + "/feeds").then((response) => {
        if (response.data) {
          console.log(response.data);
          response.data.forEach((post) => {
              this.posts.push(post);
              console.log(post);
          })
        }
      })
    }
  },

  mounted() {
    this.loadStream();
  },
};
</script>

<style scoped>
.homepage-container {
  padding: 20px;
  font-family: Arial, sans-serif;
}

.search-bar {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.search-bar input {
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.search-bar button {
  padding: 10px;
  margin-left: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.search-bar button:hover {
  background-color: #0056b3;
}

.posts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
}

.post-item {
  border: 1px solid #ddd;
  border-radius: 4px;
  overflow: hidden;
  text-align: center;
}

.post-author {
  background-color: #007bff;
  color: white;
  padding: 10px;
  font-weight: bold;
}

.post-item img {
  width: 100%;
  height: auto;
  display: block;
}
</style>
