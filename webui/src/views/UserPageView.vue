<template>
  <div class="profile-container">
    <!-- Nome Utente al Centro -->
    <div class="profile-info">
      <h2>{{ username }}</h2>
      <!-- Aggiunta dei parametri sotto il nome utente -->
      <div class="profile-stats">
        <span><strong>{{ followers.length }}</strong> Followers</span>
        <span><strong>{{ follows.length }}</strong> Follows</span>
        <span><strong>{{ posts.length }}</strong> Posts</span>
      </div>
    </div>

    <!-- Linea di divisione sopra i pulsanti -->
    <hr class="divider">

    <!-- Pulsanti a seconda della condizione -->
    <div class="action-buttons" v-if="isMyPage">
      <button @click="changeUsername">Cambia Username</button>
      <button @click="publishPost">Carica Foto</button>
    </div>
    <div class="action-buttons" v-else-if="!isFollowing && !isBanned">
      <button @click="followUser">Segui Utente</button>
      <button @click="banUser">Banna Utente</button>
    </div>
    <div class="action-buttons" v-else-if="isFollowing && !isBanned">
      <button @click="unfollowUser">Smetti di Seguire</button>
      <button @click="banUser">Banna Utente</button>
    </div>
    <div class="action-buttons" v-else-if="isBanned">
      <button @click="unbanUser">Rimuovi Ban</button>
    </div>

    <!-- Linea di divisione sotto i pulsanti -->
    <hr class="divider">

    <!-- Griglia dei post -->
    <div v-if="!isBanned" class="posts-grid">
      <div 
        v-for="(post, index) in posts" 
        :key="index" 
        class="post-item"
        @click="selectPost(post)"
      >
        <img :src="`${post.imageUrl}`" alt="image" />
      </div>
    </div>
    <!-- View Dettagliata Post selezionato -->
    <PostDetail
        v-if="selectedPost"
        :post="selectedPost"
        @close="selectedPost = null"
      />
  </div>

  <!-- Modale per cambiare username -->
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
import {readToken, readUser, writeUser} from '../services/session'
import api from '../services/axios'
import router from "../router/index.js"
export default {
  data() {
    return {
      username: '',
      posts: [],
      isMyPage: false,
      isFollowing: false,
      isBanned: false,
      followers: [],
      follows: [],
      showModal: false,
      newUsername: '',
      selectedPost: null,
    };
  },

  async beforeMount() {
    await this.checkToken();
  },

  mounted() {
  },

  methods: {
    async changeUsername() {
      this.newUsername = '';
      this.showModal = true;
    },

    async handleCancel() {
      this.showModal = false;
    },

    async handleSubmit() {
      await api.put(this.$route.fullPath, this.newUsername).then((response) => {
        this.username = response.data.username;
      }).catch((error) => {
        if (error.response) {
          console.log("Can't change username:", error.response);
        }
      });
      this.showModal = false;
    },

    async publishPost() {
        this.$refs.fileInput.click();
    },

    async handleFileUpload() {
        const file = event.target.files[0];
        if (file) {
            const formData = new FormData();
            formData.append("image", file);
            await api.post("/posts", formData).then((response) => {
              this.posts.unshift(response.data);
            }).catch((error) => {
              if (error.response) {
                console.log("Can't upload Post:", error.response);
              }
            });
        }
    },

    async followUser() {
      const me = readUser();
      await api.put(this.$route.fullPath + "/followers/" + me.uid).then((response) => {
        this.isFollowing = true;
        this.followers.push(me);
      }).catch((error) => {
        if (error.response) {
          console.log("Can't follow user:", error.response);
        }
      });
    },

    async unfollowUser() {
      const me = readUser();
      await api.delete(this.$route.fullPath + "/followers/" + me.uid).then((response) => {
        this.isFollowing = false;
        this.followers.pop(me);
      }).catch((error) => {
        if (error.response) {
          console.log("Can't unfollow user:", error.response);
        }
      });
    },

    async banUser() {
      const me = readUser();
      const bid = this.$route.fullPath.substring(7);
      await api.put("/users/" + me.uid + "/bans/" + bid).then((response) => {
        this.isFollowing = false;
        this.isBanned = true;
      }).catch((error) => {
        if (error.response) {
          console.log("Can't ban user:", error.response);
        }
      });
    },

    async unbanUser() {
      const me = readUser();
      const bid = this.$route.fullPath.substring(7);
      await api.delete("/users/" + me.uid + "/bans/" + bid).then((response) => {
        this.isFollowing = false;
        this.isBanned = false;
      }).catch((error) => {
        if (error.response) {
          console.log("Can't unban user:", error.response);
        }
      });
    },

    async checkToken() {
        if (readToken()) {
            const currentUser = readUser();
            // Get generic infos
            await api.get(this.$route.fullPath).then((response) => {
                this.posts = response.data.posts || [];
                this.isMyPage = response.data.user.uid === currentUser.uid;
                this.username = response.data.user.username;
            }).catch((error) => {
                if (error.response) {
                  console.log("Error reading infos:", error.response);
                }
            });
            // Get follows and followers infos
            await api.get(this.$route.fullPath + "/followers").then((response) => {
              if (response.data) {
                this.followers = response.data;
                if (!this.isMyPage) {
                  this.followers.forEach((follower) => {
                    if (follower.uid == currentUser.uid) {
                      this.isFollowing = true;
                    }
                  });
                }
              }
            }).catch((error) => {
              if (error.response) {
                console.log("Error reading followers:", error.response);
              }
            });
            await api.get(this.$route.fullPath + "/follows").then((response) => {
              if (response.data) {
                this.follows = response.data;
              }
            }).catch((error) => {
              if (error.response) {
                console.log("Error reading follows:", error.response);
              }
            });
            // Get My Ban
            if (!this.isMyPage) {
              await api.get("/users/"+currentUser.uid+"/bans").then((response) => {
                if (response.data) {
                  response.data.forEach((ban) => {
                    if (ban.username == this.username) {
                      this.isBanned = true;
                      this.isFollowing = false;
                    }
                  })
                }
              }).catch((error) => {
                if (error.response) {
                  console.log("Error reading my bans:", error.response);
                }
              });
            }
        } else { // No credentials
          router.push('/login');
        }
    },

    async loadProfile() {

    },

    async selectPost(post) {
      await api.get("/users/" + post.author.uid + "/posts/" + post.pid).then((response) => {
        if (response.data) {
          this.selectedPost = response.data;
        }
      });
    },
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

.profile-info h2 {
  margin: 0;
  font-size: 24px;
  font-weight: bold;
  text-align: center;
}

.profile-stats {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-top: 10px;
}

.profile-stats span {
  font-size: 18px;
}

.action-buttons {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin: 10px 0;
}

.divider {
  width: 100%;
  max-width: 800px;
  border-top: 2px solid #ddd;
  margin: 20px 0;
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
