<script>

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			username: null,
		}
	},
	methods: {
		async refresh() {
			if(this.$config.headers.Authorization != null){
				this.$router.push("/")
			}
		},
		async dologin(){

			this.loading = true;
			this.errormsg = null;
			try{
				let response = await this.$axios.post("/session",{"username":this.username},this.$config);

				
				this.$config.headers.Authorization = response.data;
				console.log("username:"+this.$username.username)
				
			}
			catch (e) {
				this.errormsg = e.toString();
			}
			if (this.$config.headers.Authorization != null){
				this.$username.username=this.username
				this.$router.push("/")
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div id="login">
		username : <input type="text" v-model = username>
		<button
			type="Submit"
			class = "e-success"
			@click="dologin()"
			color="primary">
			Log-in
		</button>
	</div>
</template>


<style>
</style>
