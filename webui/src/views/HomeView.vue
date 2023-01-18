<script>
import router from '../router';

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			photos: null,
		}
	},
	methods: {	
		async refresh() {
			console.log(this.$config)
			if(this.$config.headers.token !=null ){
				//get my stream and show it
				let response = await this.$axios.get("/home",this.$config );
				this.photos=response.data;

			}
			else{
				console.log("no token")
			}

			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");

				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async go(username,id){
			router.push("/users/"+username+"/photos/"+id);
		},
	},
	mounted() {
		this.refresh();
	}
}
</script>

<template>
<div v-if="this.$username.username">
	<div v-if=photos>
		<div v-for="photo in photos" :key="photo.Id">
			{{photo.Username}}
			<button @click="go(photo.Username,photo.Id)">
				<img :src="photo.Photo"  class="Bordered" alt="photo" width="200" height="200">
			</button>
			{{photo.Date}}
		</div>
	</div>
</div>
</template>


<style>
</style>
