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

		async loadProfile() {
			const headers = {
                'Token': this.user.token,
            };
			let request = await this.$axios.get('/users/' + this.uid + '/profile',
			{
				headers: headers,
			});
			console.log(request.data);
			this.profile_user = request.data;
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
				<button>Change Settings</button>
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
	<div class="post-grid">
	</div>
</div>
<div v-else>
	<h1> Loading Profile </h1>
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

.name-container button {
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
</style>
