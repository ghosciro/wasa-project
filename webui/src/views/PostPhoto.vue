<script>

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			username: null,
            previewImage:null
		}
	},
	methods: {
		async refresh() {
		},
        async selectPhoto(e){
            this.loading=true
            this.errormsg=null
            const image = e.target.files[0]
            const reader = new FileReader()
            reader.readAsDataURL(image)
            reader.onload = e=>{
                this.previewImage=e.target.result
                console.log(this.previewImage)
            }
        },
        async postPhoto(){
            this.loading=true
            this.errormsg=null
            try{
                console.log("users/"+this.$username.username+"/Photos")
                let response = await this.$axios.post("users/"+this.$username.username+"/Photos",this.previewImage,this.$config)
                console.log(response)
                this.$router.push("/")
            }catch(e){
                this.errormsg=e.toString()
            }
            this.loading=false

        }
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
    <h1>Post Your Photo Here</h1>
    <div class="uploading-image">
        <input type="file" accept="image/jpg/png" @change="selectPhoto">
        <button @click="postPhoto">Post</button>
    </div>
    <div>
        <img :src="previewImage" alt="preview"  width="200" height="200" >
    </div>
	</div>
</template>

<style>
.uploading-image{
    display: flex;
}
</style>