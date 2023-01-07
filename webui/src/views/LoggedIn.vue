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
		},
		async dologin(){
			console.log(this.username)
			this.loading = true;
			this.errormsg = null;
			try{
				let response = await this.$axios.post("/session?username="+this.username);
				console.log(response.data)
				this.$config.header.token = response.data;

			}
			catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.$router.back()
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Logged in mate</h1>
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
			</div>
		</div>
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
