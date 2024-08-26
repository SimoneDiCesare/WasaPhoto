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
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      searchQuery: '', // Testo per la ricerca
      posts: [
		{
			imageUrl: 'http://www.cacciaepescatognini.it/fotonotizie/19_ThinkstockPhotos-608516088.jpg',
			author: {
				username: 'Tizio'
			}
		}
	  ], // Array per i post
    };
  },

  methods: {
    async searchUsers() {
      try {
        // Effettua la ricerca degli utenti
        console.log('Ricerca utenti:', this.searchQuery);
        // Implementa la logica di ricerca qui (ad esempio, chiamando un endpoint API)
        // Poi aggiorna i post basati sui risultati di ricerca
        await this.fetchPosts();
      } catch (error) {
        console.error('Errore nella ricerca utenti:', error);
      }
    },

    async fetchPosts() {
      try {
        // URL del tuo endpoint API per ottenere i post
        const url = 'https://api.example.com/posts'; 
        const response = await axios.get(url);
        this.posts = response.data; // Aggiorna i post con i dati ottenuti
      } catch (error) {
        console.error('Errore nel recupero dei post:', error);
      }
    },
  },

  mounted() {
    // Recupera i post all'inizio
    this.fetchPosts();
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
