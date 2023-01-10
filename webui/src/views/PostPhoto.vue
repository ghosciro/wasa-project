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
        async postPhoto(e){
            this.loading=true
            this.errormsg=null
            const image = e.target.files[0]
            const reader = new FileReader()
            reader.readAsDataURL(image)
            reader.onload = e=>{
                this.previewImage=e.target.result
                console.log(this.previewImage)
            }
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
    <input type="file" accept="image/jpg/png" @change="postPhoto">
    <img: src="previewImage" class="uploading-image"/>
	</div>
</template>

<style>
.uploading-image{
    display: flex;
}
</style>