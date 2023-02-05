<script>


export default {

    props: {
        posts: {
            type: Object // Photoposts to display
        }
    },

    methods: {

        async initialize() {

        }

    },
    mounted() {

        this.initialize();

    }
}
</script>

<template>

    <div v-for="post in posts" :key="post.id" class="m-1">
        <PhotoPost :post_data="post" />
    </div>

</template>

<style>

</style>