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
        <button type = "button" @click=get_users()>
        search:<input type="text" v-model = name ></button>
        <div>
            <ul>
                <li v-for="item in found_users">
                    <button type = "button" @click=get_user(item)>{{ item}}</button>
                    
                </li>
            </ul>
        </div>
    </div>
</template>

