<script>
export default{
    data: function() {
        return {
            errormsg: null,
            loading: false,
			photos: null,
			User:{
				Username: null,
				Follow: null,
				Follows: null,
				Nphotos: 0,
				NFollow:0,
				NFollows:0,
			},
            my_banned: null,
            new_username: null,
        }
    }, 
    methods: {
        async refresh() {
            this.loading = true;
            this.errormsg = null;
            try {
				document.getElementById("user_data").style.display="initial";
				document.getElementById("followers").style.display="none";
				document.getElementById("follows").style.display="none";
                document.getElementById("banned").style.display="none";
                let response = await this.$axios.get("/users/"+this.$username.username,this.$config);
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
                response = await this.$axios.get("/users/"+this.$username.username+"/Photos",this.$config);
                this.photos = response.data;
                response = await this.$axios.get("/users/"+this.$username.username+"/banned",this.$config);
                this.my_banned = response.data;
                if (this.my_banned==null){
                    this.my_banned = [];
                }

            } catch (e) {
                console.log(e.toString())
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

        async showFollowers(){
				document.getElementById("user_data").style.display="none";
                document.getElementById("banned").style.display="none";
				document.getElementById("followers").style.display="initial";
				document.getElementById("follows").style.display="none";
			},
			async showFollows(){
				document.getElementById("user_data").style.display="none";
				document.getElementById("followers").style.display="none";
                document.getElementById("banned").style.display="none";
				document.getElementById("follows").style.display="initial";
			},
			async redirect(name){
				console.log(name)
				await this.$router.replace("/users/"+name)
				this.refresh()
			},
            async showBanned(){
                document.getElementById("user_data").style.display="none";
                document.getElementById("followers").style.display="none";
                document.getElementById("follows").style.display="none";
                document.getElementById("banned").style.display="initial";
            },
            async gophoto(id){
				await this.$router.push("users/"+this.User.Username+"/photos/"+id)
			},
            async deletePhoto(id){
                try {
                    await this.$axios.delete("/users/"+this.$username.username+"/Photos/"+id,this.$config);
                    this.refresh();
                } catch (e) {
                    console.log(e.toString())
                    this.errormsg = e.toString();
                }
            },
            async Change_username(){
                try {
                    if (this.new_username == null){
                        this.errormsg = "Username cannot be empty";
                        return;
                    }
                    console.log(this.new_username)
                    await this.$axios.put("/users/"+this.$username.username+"/options?username="+this.new_username,null,this.$config);
                    this.$username.username = this.new_username;
                    this.refresh();
                } catch (e) {
                    console.log(e.toString())
                    this.errormsg = e.toString();
                }
            },
    },
    mounted() {
        this.refresh();
    }
}
</script>
<template>
	<div>
        <div class="row">
            <div class="col-4">
                <div class="row">
                    <div class="col-6">
                        <input type="text" v-model="new_username" placeholder="New Username">
                    </div>
                    <div class="col-6">
                        <button @click="Change_username()">Change Username</button>
                    </div>
                </div>
            </div>
        <h1>{{User.Username}} </h1>
		<br>
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
				<div class="col-4" v-if="my_banned">
                    <button @click="showBanned()">Banned</button>
					<div>
					{{ my_banned.length }}
					</div>
				</div>
			</div>
			<div v-if="photos">
				<div v-for="photo in photos" :key="photo.Id">
					<br>
					<button @click="gophoto(photo.Id)">
					<img :src="photo.Photo"  class="Bordered" alt="photo" width="200" height="200">
					</button>
                    <button @click="deletePhoto(photo.Id)">Delete</button>"	
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
        <div id="banned">
			<div v-if="my_banned">
				<div v-for="banned in my_banned" :key="banned">
					<button @click="redirect(banned)">{{banned}}</button>
				</div>
			</div>	
		</div>
    </div>
</template>