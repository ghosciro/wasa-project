<script>

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
            photo:null,
            comments:null,
            likes:[],
            mycomment:null,
		}
	},
	methods: {
		async refresh() {  
            document.getElementById("likebutton").style.display="initial";
            document.getElementById("unlikebutton").style.display="none";
            this.photo= await this.$axios.get("users/"+this.$route.params.username+"/Photos/"+this.$route.params.id,this.$config);
            this.photo=this.photo.data;
            this.comments= await this.$axios.get("users/"+this.$route.params.username+"/Photos/"+this.$route.params.id+"/comments",this.$config);
            this.comments=this.comments.data;
            this.likes= await this.$axios.get("users/"+this.$route.params.username+"/Photos/"+this.$route.params.id+"/likes",this.$config);
            this.likes=this.likes.data;
            if (this.likes!= null) {
            if(this.likes.includes(this.$username.username)){
                document.getElementById("likebutton").style.display="none";
                document.getElementById("unlikebutton").style.display="initial";
            }
            else{
                document.getElementById("likebutton").style.display="initial";
                document.getElementById("unlikebutton").style.display="none";
            }
        }
		},
        async postcomment(){
            await this.$axios.post("users/"+this.$route.params.username+"/Photos/"+this.$route.params.id+"/comments",this.mycomment,this.$config);
            this.mycomment=null;
            this.refresh();
        },
        async like(){
            await this.$axios.put("users/"+this.$route.params.username+"/Photos/"+this.$route.params.id+"/likes","",this.$config);
            this.refresh();
        },
        async unlike(){
            await this.$axios.delete("users/"+this.$route.params.username+"/Photos/"+this.$route.params.id+"/likes",this.$config);
            this.refresh();
        }
	},
	mounted() {
		this.refresh()
	}
}
</script>
<template>
	<div>
		<h1>Photo</h1>
        <div v-if=photo>
            <p>{{photo.Date}}</p>
            <p>{{photo.Username}}</p>
            <img :src="photo.Photo"  class="Bordered" alt="photo">
        </div>
        <div>
            <button @click="like()" id="likebutton">Like</button>
            <button @click="unlike()" id="unlikebutton">unLike</button>
        </div>
        <br>
        <div>
            <input type="text" v-model="mycomment">
            <button @click="postcomment()">Post comment</button>
        </div>
        <div v-if=comments>
            <ul>
                <li v-for="comment in comments">
                    <p>{{comment.Username}}:    {{comment.Comment}}</p>
                    <br>
                </li>
            </ul>
        </div>
	</div>
</template>
