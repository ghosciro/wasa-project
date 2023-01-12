<script>
import router from '../router';

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			username: null,
			photos: null,
		}
	},
	methods: {
		async refresh() {
			console.log(this.$config)
			if(this.$config.headers.token !=null ){
				//get my stream and show it
				document.getElementById("showphotos").style.display="initial"
				let response = await this.$axios.get("/home",this.$config );
				this.photos=response.data;

			}
			else{
				console.log("no token")
				document.getElementById("showphotos").style.display="none"
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
		async search(){
			this.$router.push("/users");
		},
		async postphoto(){
			this.$router.push("/postPoto")
		},
		async Gocomment(){
			this.$routher.push("/comments")
		}
	},
	mounted() {
		this.refresh();
	}
}
</script>

<template>
<div>
	<div id ="logged in">
		<button @click="search()">search users</button>
		<button @click="postphoto()">post photo</button>
		<button @click="change_username()">change username</button>
	</div>
	<div id="showphotos">
		<div v-if=photos  v-for="photo in photos" :key="photo.Id">
			<div>
				<br>
				{{photo.Date}}
				<button>
					<img :src="photo.Photo"  class="Bordered" alt="photo" width="200" height="200">
				</button>
			</div>
		</div>
	</div>
</div>

	
</template>


<style>
</style>
