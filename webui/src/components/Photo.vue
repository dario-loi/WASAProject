<script>


export default {

    props: {
        alt: {
            type: String
        },
        style: {
            type: Object
        }
    },
    data: function () {
        return {
            src: null,
            alt_: null,
            style_: null,
        }
    },

    methods: {
    },
    mounted() {

        // Here we need to perform a sequence of async operations through axios 
        // to get the image data from the server. We need to get the image data
        // likes, comments, etc...



    }

}



</script>

<template>

    <img :src="src" :alt="(alt != null ? alt : 'Wasaphoto Image')" class="img-fluid"
        :style="(style_ != null ? style_ : '')" />



</template>

<style>

</style>