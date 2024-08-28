<template>
	<header class="navbar navbar-dark sticky-top bg-dark shadow p-0">
		<div class="container-fluid d-flex justify-content-between align-items-center">
			<!-- Pulsante di logout a sinistra -->
			<button v-if="isLoggedIn" class="btn btn-outline-light me-2" @click="logout">
				Logout
			</button>

			<!-- Titolo centrato -->
			<a class="navbar-brand mx-auto fs-6 text-center" @click="goHome">WasaPhoto</a>

			<!-- Pulsante Home a destra -->
			<button v-if="isLoggedIn" class="btn btn-outline-light ms-2" @click="goProfile">
				Profile
			</button>
		</div>
	</header>

	<main>
		<RouterView />
	</main>
</template>

<script>
import { readToken, writeUser, readUser } from './services/session';
import router from './router/index.js';

export default {
	data() {
		return {
			isUserLoggedIn: false,
		};
	},

	computed: {
		isLoggedIn() {
			return this.isUserLoggedIn;
		},
	},

	created() {
		this.isUserLoggedIn = !!readToken();
	},

	methods: {
		logout() {
			if (this.isLoggedIn) {
				writeUser();
				this.isUserLoggedIn = false;
				router.push('/login');
			}
		},
		goHome() {
			if (this.isLoggedIn) {
				const uid = readUser().uid;
				router.push('/users/' + uid + '/feeds');
			} else {
				writeUser();
				router.push('/login');
			}
		},
		goProfile() {
			if (this.isLoggedIn) {
				const uid = readUser().uid;
				router.push('/users/' + uid);
			}
		},
	},

	watch: {
		'$route'() {
			this.isUserLoggedIn = !!readToken();
		},
	},
};
</script>

<style scoped>
.navbar-brand {
    flex-grow: 1;
    text-align: center;
    margin-left: -2.5rem; /* Adjust the value as needed */
}
</style>
