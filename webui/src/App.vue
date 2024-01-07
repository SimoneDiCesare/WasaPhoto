<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data: function() {
		return {
			user: null,
		}
	}
}
</script>

<template>
	<RouterView />
</template>

<style>
</style>
