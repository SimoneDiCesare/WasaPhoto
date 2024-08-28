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
        @click="selectPost(post)"
      >
        <div class="post-author" @click="goToUserPage(post.author.uid)">{{ post.author.username }}</div>
        <img :src="post.imageUrl" alt="Post Image" />
      </div>
    </div>
    <!-- View Dettagliata Post selezionato -->
    <PostDetail
        v-if="selectedPost"
        :post="selectedPost"
        @close="selectedPost = null"
      />

    <!-- Modale per la lista di utenti -->
    <div v-if="showUserModal" class="modal">
      <div class="modal-content">
        <!-- Bottone di chiusura in alto a sinistra -->
        <button class="close-left" @click="closeModal">X</button>
        <h2>Risultati della ricerca</h2>
        <ul>
          <li v-for="user in users" :key="user.id" @click="goToUserPage(user.uid)">
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
import router from "../router/index.js"

export default {
  data() {
    return {
      searchQuery: '', // Testo per la ricerca
      showUserModal: false,
      posts: [], // Array per i post
      users: [],  // Array per gli utenti
      selectedPost: null,
    };
  },

  methods: {

    async selectPost(post) {
      await api.get("/users/" + post.author.uid + "/posts/" + post.pid).then((response) => {
        if (response.data) {
          this.selectedPost = response.data;
        }
      });
    },

    async searchUsers() {
      if (!this.searchQuery.trim()) {
        // Se la query di ricerca è vuota, non fare nulla
        return;
      }

      await api.get("/users?username=" + this.searchQuery).then((response) => {
        if (response.data) {
          this.users = response.data;
          this.showUserModal = true;
        }
      }).catch((error) => {
        if (error.response) {
          console.log("Can't search users:", error.response);
        }
      });
    },

    async loadStream() {
      await api.get("/users/" + readUser().uid + "/feeds").then((response) => {
        if (response.data) {
          this.posts = response.data;
          console.log(this.posts);
        }
      }).catch((error) => {
        if (error.response) {
          console.log("Can't get feeds:", error.response);
        }
      });
    },

    async goToUserPage(uid) {
      router.push("/users/" + uid);
    },

    closeModal() {
      this.showUserModal = false;
      this.users = []; // Pulisci l'elenco degli utenti quando chiudi la modale
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

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 400px;
  max-width: 90%;
  position: relative;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.close-left {
  position: absolute;
  top: 10px;
  left: 10px;
  border: none;
  background: transparent;
  font-size: 20px;
  cursor: pointer;
  z-index: 1; /* Assicura che il bottone sia sopra il contenuto */
}

.modal-content h2 {
  margin: 0;
  padding-left: 40px; /* Spazio per il bottone di chiusura */
}

.modal-content ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.modal-content li {
  padding: 10px;
  cursor: pointer;
}

.modal-content li:hover {
  background-color: #f1f1f1;
}
</style>