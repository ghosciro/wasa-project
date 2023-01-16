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
				document.getElementById("follow").style.display="initial";
				document.getElementById("unfollow").style.display="none";
				document.getElementById("user_data").style.display="initial";
				document.getElementById("followers").style.display="none";
				document.getElementById("follows").style.display="none";
				try {
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
				//if this.username in followers turn button to unfollow
				if(this.User.Follow){
				if (this.User.Follow.includes(this.$username.username)){
					document.getElementById("follow").style.display="none";
					document.getElementById("unfollow").style.display="initial";
				}
				else{
					document.getElementById("follow").style.display="initial";
					document.getElementById("unfollow").style.display="none";
				}
				}

				//if this.username in banned turn button to unban
				response= await this.$axios.get("/users/"+this.$username.username+"/banned",this.$config);
				var banned=response.data;
				document.getElementById("ban").style.display="initial";
				document.getElementById("unban").style.display="none";
				if (banned!=null && banned.includes(this.$route.params.username)){
					document.getElementById("ban").style.display="none";
					document.getElementById("unban").style.display="initial";
				}
			}
			catch (e) {
				console.log(e.toString())
				this.errormsg = e.toString();
			}
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
				console.log(this.$config.headers.token)
					let response = await this.$axios.post("users/"+this.$username+"/following/"+this.$route.params.username, null,this.$config);
					console.log(response)
					this.refresh()
			},
			async unfollow(){
				console.log("unfollowing")
				await this.$axios.delete("users/"+this.$username.username+"/following/"+this.$route.params.username,this.$config);
				this.refresh()
			},
			async gophoto(id){
				await this.$router.push(this.User.Username+"/photos/"+id)
			},
			async ban(){
				console.log("banning")
				await this.$axios.post("users/"+this.$username.username+"/banned/"+this.$route.params.username, null,this.$config);
				this.refresh()
			},
			async unban(){
				console.log("unbanning")
				await this.$axios.delete("users/"+this.$username.username+"/banned/"+this.$route.params.username,this.$config);
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
		<div class="row">
			<h1	>{{User.Username}} </h1>
			<div class="col-1" >
				<button id="follow" @click="follow()">Follow</button>
				<button id="unfollow"  @click="unfollow()">Unfollow</button>
			</div>
			<div class="col-1" >
				<button id="ban" @click="ban()">ban</button>
				<button id="unban"  @click="unban()">unban</button>
			</div>
		</div>
		<br>
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
				<div v-for="photo in photos" :key="photo.Id">
					<br>
					<button @click="gophoto(photo.Id)">
					<img :src="photo.Photo"  class="Bordered" alt="photo" width="200" height="200">
					</button>	
				</div>
			</div>
		</div>
		<div id="followers">
			<div v-if="User.Follow">
				<div v-for="follow in User.Follow" :key="follow">
					<button @click="redirect(follow)">{{follow}}</button>
				</div>
			</div>
		</div>
		<div id="follows">
			<div v-if="User.Follows">
				<div v-for="follows in User.Follows" :key="follows">
					<button @click="redirect(follows)">{{follows}}</button>
				</div>
			</div>	
		</div>
    </div>
</template>