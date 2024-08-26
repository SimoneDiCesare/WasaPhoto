<template>
  <div class="profile-container">
    <div class="profile-info">
      <h2>{{ username }}</h2>
      <div class="stats">
        <button @click="homePage">Homepage</button>
        <span><strong>{{ posts.length }}</strong> Posts</span>
      </div>
      <div v-if="isMyPage" class="settings">
        <button @click="changeUsername">Cambia Username</button>
        <button @click="publishPost">Pubblica Post</button>
      </div>
    </div>
    
    <div class="posts-grid">
      <div 
        v-for="(post, index) in posts" 
        :key="index" 
        class="post-item"
      >
        <img :src="`http://localhost:3000${post.imageUrl}`" alt="image" />
      </div>
    </div>
  </div>

  <!-- Modale -->
    <div v-if="showModal" class="modal-overlay">
      <div class="modal-content">
        <textarea v-model="newUsername" placeholder="Inserisci Nuovo Username"></textarea>
        <div class="modal-buttons">
          <button @click="handleCancel">Cancella</button>
          <button @click="handleSubmit">Invia</button>
        </div>
      </div>
    </div>

  <input type="file" ref="fileInput" @change="handleFileUpload" accept="image/*" style="display: none" />
</template>

<script>
import {readToken, readUser} from '../services/session'
import api from '../services/axios'
export default {
  data() {
    return {
      username: '',
      posts: [],
      isMyPage: false,
      showModal: false,
      newUsername: '',
    };
  },

  mounted() {
    this.checkToken();
  },

  methods: {

    async homePage() {
      this.$router.push("/users/" + readUser().uid + "/feeds");
    },

    async changeUsername() {
      this.newUsername = '';
      this.showModal = true;
    },

    async handleCancel() {
      this.showModal = false;
    },

    async handleSubmit() {
      await api.put(this.$route.fullPath, this.newUsername).then((response) => {
        console.log(response.data);
        if (response.data < 300) {
          this.username = response.data.username;
        }
      });
      this.showModal = false;
    },

    async publishPost() {
        this.$refs.fileInput.click();
    },

    async handleFileUpload() {
        const file = event.target.files[0]; // Ottieni il file selezionato
        if (file) {
            const formData = new FormData();
            formData.append("image", file);
            await api.post("/posts", formData).then((response) => {
                console.log(response);
            });
        }
    },

    async checkToken() {
        if (readToken()) {
            const currentUser = readUser();
            await api.get(this.$route.fullPath).then((response) => {
                if (response.status >= 400) {
                    // writeUser();
                    // this.$router.push('/users/' + data.uid);
                    this.$router.push('/login');
                    return;
                }
                if (response.data.items) {
                    response.data.items.forEach((post) => {
                        this.posts.push(post)
                    })
                }
                if (response.data.user.uid == currentUser.uid) {
                    this.isMyPage = true;
                }
                this.username = response.data.user.username;
                console.log(response.data);
            });
            return;
        }
        this.$router.push('/login');
    }
  }
};
</script>

<style scoped>
.profile-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}

.profile-info {
  text-align: center;
  margin-bottom: 30px;
}

.profile-info h2 {
  margin: 0;
  font-size: 24px;
  font-weight: bold;
}

.stats {
  display: flex;
  gap: 20px;
  margin-top: 10px;
}

.stats span {
  font-size: 18px;
}

.posts-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 10px;
  width: 100%;
  max-width: 800px;
}

.post-item img {
  width: 100%;
  height: auto;
  display: block;
  border-radius: 5px;
  object-fit: cover;
}

.modal-overlay {
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
}

.modal-buttons {
  margin-top: 10px;
  display: flex;
  justify-content: flex-end;
}

.modal-buttons button {
  margin-left: 10px;
}
</style>
