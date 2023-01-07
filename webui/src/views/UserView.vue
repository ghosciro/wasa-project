<script>
export default {
	data: function() {
		return {
			User:{
				Username: null,
				Follow: null,
				Follows: null,
				Nphotos: 0,
				NFollow:0,
				NFollows:0,
			},
			errormsg: null,
			loading: false,
			photos: null,
		}
	},
        methods: {
            async refresh() {		
				document.getElementById("user_data").style.display="initial";
				document.getElementById("followers").style.display="none";
				document.getElementById("follows").style.display="none";
                this.loading = true;
                this.errormsg = null;
                let response = await this.$axios.get("/users/"+this.$route.params.username,this.$config);

					this.User.Username= response.data.Username;
					this.User.Follow = response.data.Follower;
					this.User.Follows = response.data.Follows;
					if (response.data.Follower != null){
						this.User.NFollow = response.data.Follower.length;
					}
					else{
						this.User.NFollow = 0;
					}
					if (response.data.Follows != null){
						this.User.NFollows =response.data.Follows.length;
					}
					else{
						this.User.NFollows = 0;
					}
					this.User.Nphotos = response.data.Nphotos; 
					let response2   = await this.$axios.get("/users/"+this.$route.params.username+"/Photos",this.$config);
					this.photos = response2.data;
                this.loading = false;
				console.log(this.User)
		    },
			async showFollowers(){
				document.getElementById("user_data").style.display="none";
				document.getElementById("followers").style.display="initial";
				document.getElementById("follows").style.display="none";
			},
			async showFollows(){
				document.getElementById("user_data").style.display="none";
				document.getElementById("followers").style.display="none";
				document.getElementById("follows").style.display="initial";
			},
			async redirect(name){
				console.log(name)
				await this.$router.replace("/users/"+name)
				this.refresh()
			},
			async follow(){
				console.log(this.config.headers.token)
					let response = await this.$axios.post("users/dennis/following/"+this.$route.params.username, null,this.$config);
					console.log(response)
					this.refresh()
			}
        },
    mounted() {
				this.refresh();
	}
}
</script>


<template>
	<div>
		<div>
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
		<div class="row">
			<div class="col-1" >
			<h1	>{{User.Username}}</h1>
			</div>
			<div class="col-1">
			<button @click="follow()">Follow</button>
			</div>
		</div>
		<div v-if="User" id="user_data">
			<div class="row">
				<div class="col-4">
					<button @click="showFollowers()">Followers</button>
					<div>
						{{ User.NFollow }}
					</div>

				</div>
				<div class="col-4">
					<button @click="showFollows()">Follows</button>
					<div>
					{{ User.NFollows }}
					</div>
				</div>
				<div class="col-4">Photos:
					<div> {{ User.Nphotos }}</div>
				</div>
			</div>
			<div v-if="photos">
				<div v-for="photo in photos" :key="photo.ID">
					<img :src="photo.URL" alt="photo" width="200" height="200">
				</div>

			</div>
		</div>
		<div id="followers">
			<div v-if="User.Follow">
				<div v-for="follow in User.Follow">
					<button @click="redirect(follow)">{{follow}}</button>
				</div>
			</div>
		</div>
		<div id="follows">
			<div v-if="User.Follows">
				<div v-for="follows in User.Follows">
					<button @click="redirect(follows)">{{follows}}</button>
				</div>
			</div>	
		</div>
    </div>
</template>