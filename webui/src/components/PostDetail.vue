<script>
import api from '../services/axios'
import {readToken, readUser, writeUser} from '../services/session'
import router from "../router/index.js"
export default {
  props: {
    post: Object,  // Il post che viene passato alla view
  },
  data() {
    return {
      userLiked: false,  // Stato del like da parte dell'utente
      isMyPost: false,
      newComment: '',
    };
  },
  mounted() {
    const me = readUser();
    if (me.uid == this.post.author.uid) {
      this.isMyPost = true;
    }
    if (!this.post.likes) {
      return;
    }
    this.post.likes.forEach((user) => {
      if (user.uid == me.uid) {
        this.userLiked = true;
      }
    })
  },
  methods: {
    async toggleLike() {
      const me = readUser();
      if (this.userLiked) {
        await api.delete("/posts/" + this.post.pid + "/likes/" + me.uid).then((response) => {
          this.userLiked = false;
          this.post.likesCount -= 1;
          console.log(response);
        }).catch((error) => {
          if (error.response) {
            console.log("Error unliking:", error);
          }
        });
      } else {
        await api.put("/posts/" + this.post.pid + "/likes/" + me.uid).then((response) => {
          this.userLiked = true;
          this.post.likesCount += 1;
          console.log(response);
        }).catch((error) => {
          if (error.response) {
            console.log("Error liking:", error);
          }
        });
      }
    },
    closeView() {
      this.$emit('close');
    },
    async commentPost() {
      if (!this.newComment.trim()) {
        return;
      }
      console.log("Commenting:", this.newComment);
      await api.post("/posts/" + this.post.pid + "/comments", this.newComment).then((response) => {
        console.log(response);
        if (response.data) {
          if (!!this.post.comments) {
            this.post.comments = [];
          }
          this.post.comments.push(response.data);
          this.newComment = '';
        }
      }).catch((error) => {
        if (error.response) {
          console.log("Error commenting:", error);
        }
      });
    },
    async removeComment(comment) {
      console.log(comment.author.uid, "==", readUser().uid, readUser().uid == comment.author.uid);
      await api.delete("/posts/" + this.post.pid + "/comments/" + comment.cid).then((response) => {
        console.log(response);
        // Remove from list
        this.post.comments = this.post.comments.filter((item) => {
          return item.cid != comment.cid;
        });
      }).catch((error) => {
        if (error.response) {
          console.log("Error removing comment:", error);
        }
      });
    },
    async goToUserPage(uid) {
      router.push("/users/" + uid);
    },
    isMyComment(comment) {
      console.log(comment.author.uid, "==", readUser().uid, readUser().uid == comment.author.uid);
      return comment.author.uid == readUser().uid;
    },
    getFormattedDate(time) {
      const date = new Date(time);
      const options = {
        day: 'numeric',
        month: 'short',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        hour12: false,
      };
      return new Intl.DateTimeFormat('en-GB', options).format(date);
    },
    // TODO: Implement delete logic
    async deletePost() {
      await api.delete("/posts/" + this.post.pid).then((response) => {
        console.log(response);
        this.closeView();
        this.$router.go();
      })
    },
  },
};
</script>

<template>
  <div class="modal-overlay">
    <div class="post-detail-container">
      <!-- Header con pulsante di chiusura e nome utente -->
      <div class="post-header">
        <button class="close-button" @click="closeView">X</button>
        <div class="post-author" @click="goToUserPage(post.author.uid)">
          {{ post.author.username }} <br> {{ getFormattedDate(this.post.uploadTime) }}
        </div>
        <button v-if="isMyPost" class="delete-post" @click="deletePost()">
          Delete
        </button>
      </div>

      <!-- Contenitore del post e dei commenti -->
      <div class="post-body">
        <!-- Immagine del post -->
        <div class="post-content">
          <img :src="post.imageUrl" alt="Post Image" />
          <div class="post-footer">
            <div class="like-count">{{ post.likesCount }} Likes</div>
            <button v-if="!isMyPost" @click="toggleLike">{{ userLiked ? 'Unlike' : 'Like' }}</button>
          </div>
        </div>

        <!-- Lista dei commenti -->
        <div class="comments-section">
          <div class ="comments-list">
            <div v-for="comment in post.comments" :key="comment.id" class="comment-item">
              <div class="comment-top">
                <div class="comment-author">{{ comment.author.username }}</div>
                <div v-if="isMyComment(comment)" class="comment-button" @click="removeComment(comment)">X</div>
              </div>
              <div class="comment-text">{{ comment.text }}</div>
            </div>
          </div>
          <div class="new-comment">
            <textarea v-model="newComment" placeholder="Aggiungi un commento..."></textarea>
            <button @click="commentPost">Comment</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
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
  z-index: 1000;
}

.post-detail-container {
  background-color: white;
  padding: 20px;
  max-width: 900px;
  width: 90%;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  position: relative;
  z-index: 1001;
}

.post-header {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  background-color: #007bff;
  padding: 10px;
  margin-bottom: 10px;
  color: white;
}

.close-button {
  background-color: transparent;
  border: none;
  font-size: 20px;
  cursor: pointer;
  margin-right: 10px;
  color: white;
}

.post-author {
  font-weight: bold;
  text-align: center;
  flex-grow: 1;
  font-size: 24px;
}

.post-body {
  display: flex;
  flex-direction: row;
  gap: 20px;
}

.post-content {
  flex: 1.5;
}

.post-content img {
  width: 100%;
  height: 50vh;
  object-fit: contain;
  display: block;
}

.comments-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  border-left: 2px solid #ddd;
  padding-left: 20px;
  max-height: 50vh;
  overflow-y: auto;
}

.comments-list {
  flex: 1;
  display: flex;
  flex-direction: column;
  max-height: 60vh;
  overflow-y: auto;
}

.new-comment {
  padding: 10px;
  display: flex;
  align-items: center; /* Allinea verticalmente gli elementi */
  gap: 10px; /* Aggiungi spazio tra gli elementi */
}

.new-comment textarea {
  flex: 1; /* Occupa tutto lo spazio disponibile */
  resize: none; /* Rimuovi la possibilità di ridimensionare la textarea */
  height: 40px; /* Imposta un'altezza specifica */
  padding: 5px;
}

.new-comment button {
  padding: 5px 10px;
  margin-left: 5px;
  white-space: nowrap; /* Impedisce che il testo vada su una nuova linea */
  height: 40px; /* Imposta la stessa altezza della textarea */
}

.comment-top {
  display: flex;
  justify-content: space-between;
  align-items: top;
  padding: 0px;
}

.comment-item {
  background-color: #f0f0f0; /* Sfondo grigio chiaro per i commenti */
  padding: 10px;
  margin-bottom: 15px;
  border-radius: 8px;
}

.comment-author {
  font-weight: bold;
  margin-bottom: 5px;
}

.comment-button {
  color: #ff0000;
  margin-left: 10px;
}

.comment-text {
  font-size: 14px;
}

.post-footer {
  border-top: 1px solid #ddd;
  padding-top: 5px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.like-count {
  font-weight: bold;
}

button {
  background-color: #007bff;
  color: white;
  border: none;
  padding: 5px 10px;
  cursor: pointer;
  border-radius: 4px;
}

button:hover {
  background-color: #0056b3;
}
</style>
