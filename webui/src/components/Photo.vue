<script>


export default {

    props: {
        src: {
            type: String
        },
        alt: {
            type: String
        },
        style: {
            type: Object
        }
    },
    data: function () {
        return {
            is_loading: false,
            src_: null,
            alt_: null,
            style_: null,
        }
    },

    methods: {

        async initialize() {

            this.is_loading = true
            this.alt_ = this.alt;
            this.style_ = this.style;

            // Get the image data from the server

            const response = await this.$axios.get("/resources/photos/" + this.src, {
                responseType: "image/jpeg"
            });

            this.src_ = response.data;
            this.is_loading = false;
        }

    },
    mounted() {

        // Here we need to perform a sequence of async operations through axios 
        // to get the image data from the server. We need to get the image data
        // likes, comments, etc...

        this.initialize();

    }
}
</script>

<template>

    <div v-if="!is_loading">
        <img :src="src_" :alt="(alt != null ? alt : 'Wasaphoto Image')" class="img-fluid opacity-100"
            :style="(style_ != null ? style_ : '')" />
    </div>
    <div v-else>
        <LoadingSpinner></LoadingSpinner>
    </div>

</template>

<style>

</style>