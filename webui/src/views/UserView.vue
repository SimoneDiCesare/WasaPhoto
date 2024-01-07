<script>
export default {
	data: function() {
		return {
			user: JSON.parse(localStorage.getItem('user')) || null,
			uid: null,
			profile_pic: null,
			is_owner: false,
			loagind: false,
			some_data: null,
			errormsg: null,
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

		loadPage() {
			this.uid = this.$route.params.uid;
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
	<div class="user-profile">
		<img :src="profile_pic" alt="Profile picture" width="50" height="50">
    	<h2>{{ user.username }}</h2>
 	</div>
</template>

<style scoped>
.user-profile {
 display: flex;
 flex-direction: column;
 align-items: center;
 padding: 20px;
 background-color: #f0f0f0;
 border-radius: 5px;
}

.profile-pic {
 width: 100px;
 height: 100px;
 border-radius: 50%;
 object-fit: cover;
 margin-bottom: 10px;
}
</style>
