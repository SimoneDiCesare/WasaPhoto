<script>
export default {
	data: function() {
		return {
			user: JSON.parse(localStorage.getItem('user')) || null,
			profile_user: null,
			uid: null,
			profile_pic: null,
			is_owner: false,
			loagind: false,
			some_data: null,
			errormsg: null,
			postCount: 0,
			posts: null,
			uploadVisible: false,
			selectedFile: null,
		}
	},
	methods: {
		async loadProfilePic() {
            const headers = {
                'Token': this.user.token,
            };
			let request = await this.$axios.get('/users/' + this.uid + '/image',
			{
				responseType: 'arraybuffer',
				headers: headers
			});
			let base64 = btoa(
				new Uint8Array(request.data).reduce(
					(data, byte) => data + String.fromCharCode(byte), '',
				),
			);
    		this.profile_pic = `data:image/png;base64,${base64}`;
		},

		async loadPostPic(post) {
            const headers = {
                'Token': this.user.token,
            };
			let request = await this.$axios.get('/posts/' + post.pid + '/image',
			{
				responseType: 'arraybuffer',
				headers: headers
			});
			console.log(request);
			let base64 = btoa(
				new Uint8Array(request.data).reduce(
					(data, byte) => data + String.fromCharCode(byte), '',
				),
			);
    		post.image = `data:image` + post.pid + `/png;base64,${base64}`;
		},

		async loadProfile() {
			const headers = {
                'Token': this.user.token,
            };
			let request = await this.$axios.get('/users/' + this.uid + '/profile',
			{
				headers: headers,
			});
			this.profile_user = request.data;
			if (this.profile_user.posts) {
				this.posts = this.profile_user.posts;
				this.postCount = this.posts.length;
				this.posts.forEach(post => {
					this.loadPostPic(post);
				});
			}
			console.log(this.profile_user);
		},

		loadPage() {
			this.uid = this.$route.params.uid;
			this.loadProfile();
			this.loadProfilePic();
			// Load Profile pic
			let profileUid = this.$route.params.uid;
			if (profileUid == this.user.uid) {
				console.log('Personal Page');
				this.is_owner = true;
			} else {
				console.log(profileUid + ' Page');
			}
		},

		changeUsername() {
			console.log('Username');
		},

		changeProfilePic() {
			console.log('Profile Pic');
		},

		showUploadLayout() {
			this.uploadVisible = true;
		},

		closeUploadLayout() {
			this.uploadVisible = false;
		},

		async uploadPost() {
			if (this.selectedFile) {
				const formData = new FormData();
				formData.append('uid', this.user.uid);
				formData.append('caption', 'A sample Caption');
				formData.append('image', this.selectedFile);
				console.log(formData);
				const headers = {
                	'Token': this.user.token,
            	};
				let response = await this.$axios.post('/posts', formData, {
					headers: headers,
				});
        		console.log(response.data);
			}
			this.closeUploadLayout();
		},

		onFileChange(event) {
      		this.selectedFile = event.target.files[0];
    	},

	},
	mounted() {
		if (this.user) {
			this.loadPage();
			console.log('Mounted');
		} else {
			window.location = '/#/session';
		}
	}
}
</script>

<template>
<div v-if="profile_user" class="profile-page">
	<div class="profile-container">
		<img :src="profile_pic" alt="Profile picture" class="profile-pic">
		<div class="info-container">
			<div class="name-container">
				<h2>{{ profile_user.user.username }}</h2>
				<div v-if="is_owner" class="button-info">
					<button @click="changeUsername()">Change Username</button>
					<button @click="changeProfilePic()">Change Profile Pic</button>
				</div>
			</div>
			<div class="counts-container">
				<p class="count">{{ postCount }} Post</p>
				<p class="count">{{ profile_user.follower }} Followers</p>
				<p class="count">{{ profile_user.follows }} Follows</p>
			</div>
			<div class="bio-container">
				<p>{{ profile_user.user.bio }}</p>
			</div>
		</div>
	</div>
	<button @click="showUploadLayout()">Upload Post</button>
	<div class="post-grid">
		<div v-for="post in posts" :key="post.pid">
			<img :src="post.image" alt="Post picture" class="post-pic">
		</div>
	</div>
</div>
<div v-else>
	<h1> Loading Profile </h1>
</div>
<div class="username-modal" v-show="false">
	<p contenteditable="true">{{ user.username }}</p>
</div>
<div class="post-modal" v-show="uploadVisible">
  	<label for="file">File</label>
  	<input id="file" name="file" type="file" @change="onFileChange"/>
	<button @click="uploadPost()">Upload</button>
</div>
</template>


<style scoped>
.profile-page {
	display: flex;
	flex-direction: column;
	align-items: center;
}

.profile-container {
	display: grid;
	grid-template-columns: 1fr 2fr;
	grid-template-rows: 200px;
}

.profile-pic {
	width: 150px;
	height: 150px;
	border-radius: 50%;
	object-fit: cover;
	margin-right: 50px;
}

.info-container {
}

.name-container {
	display: flex;
	flex-direction: row;
	align-items: center;
	gap: 20px;
}

.button-info {
	display: flex;
	flex-direction: row;
	align-items: center;
	gap: 20px;
}

.button-info button {
	padding: 5px;
	border: none;
	text-decoration: none;
	color: #000000;
	border-radius: 5px;
}

.counts-container {
	display: flex;
	flex-direction: row;
	align-items: center;
	gap: 20px;
}

.count {
	color: #333333;
}

.bio-container {
	color: #000000
}

.post-grid {
	display: grid;
	grid-template-columns: auto auto auto;
	padding: 5px;
	gap: 2px;
	background: #cccccc;
	margin-top: 20px;
}

.post-pic {
	width: 250px;
	height: 250px;
	object-fit: cover;
}

.post-modal {
	position: fixed;
	top: 150px;
	left: 50%;
	width: 300px;
	height: 500px;
	transform: translate(-50%, 0px);
	border-style: solid;
	background: #ffffff;
	border-color: #222222;
	padding: 20px;
	z-index: 1000; /* Set a high z-index to make it appear on top of other elements */
}
</style>
