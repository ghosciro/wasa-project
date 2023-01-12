<script>

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
		}
	},
	methods: {
		async refresh() {
			await this.$axios.delete("/session?",this.$config)
			
			this.$config.headers.token=null
			this.$username=null
			console.log(this.$config.headers)
			this.$router.push("/session")
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>
<template>
	<div>
		<h1>Log Out </h1>
	</div>
</template>
