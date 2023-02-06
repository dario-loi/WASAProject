<script>


export default {

    props: {
        posts: {
            type: Object // Photoposts to display
        }
    },

    watch: {
        posts: function (new_posts) {
            this.posts_ = new_posts;
        }
    },

    data: function () {
        return {
            posts_: this.posts
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

    <div v-for="post in posts_" :key="post.id" class="m-1">
        <PhotoPost :post_data="post" />
    </div>

</template>

<style>

</style>