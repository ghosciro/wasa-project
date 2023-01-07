<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			username: null,
            name : "",
            found_users : null,
		}
	},
        methods: {
            async refresh() {
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
            async get_users(){
                let response = await this.$axios.get("/users?username="+this.name,this.$config)
                console.log(response.data)
                this.found_users= response.data
            },
            async get_user(name){
                this.$router.push("/users/"+name)
            },
           
    },
    mounted() {
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
                search:<input type="text" v-model = name >
                <button type = "button" @click=get_users()></button>
			</div>
		</div>
        <div>
            <ul>
                <li v-for="item in found_users">
                    <button type = "button" @click=get_user(item)>{{ item}}</button>
                    
                </li>
            </ul>
        </div>
    </div>
</template>

