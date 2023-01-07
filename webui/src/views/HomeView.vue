<script>
import router from '../router';

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
			
			if(this.$config.header.token !=null ){
				document.getElementById("button_login").style.display= "none"
				document.getElementById("logged in").style.display="initial"
			}
			else{
				console.log("no token")
				document.getElementById("logged in").style.display="none"
				document.getElementById("button_login").style.display="initial"
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
		async login(){
			this.$router.push("/session");
		},
		async search(){
			this.$router.push("/users");
		},
		async do_logout(){
			console.log("logging out");
			this.$axios.delete("/session?",this.$config);
			localStorage.clear();
			this.$config.header.token=null;
			this.refresh();
		}
	},
	mounted() {
		console.log(this.$config.header)
		this.refresh();
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
				<div>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click=do_logout()>
						Log out
					</button>
				</div>
			</div>
		</div>
		<div id ="button_login">
			<button
  				@click="login()"
  				color="primary">do log-in
			</button>
		</div>
		<div id ="logged in">
			<button @click="search()">search users</button>
			<button @click="mystream()">my stream</button>
			<button @click="postphoto()">post photo</button>
			<button @click="change_username()">change username</button>
		</div>

	</div>
	<div>
		
	</div>
	
</template>


<style>
</style>
