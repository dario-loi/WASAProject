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

    emits: ["delete-post"],

    methods: {

        async initialize() {

        },

        async RemovePost(post_data) {

            // Bubble up the event, let the parent handle it
            // Not optimal design but its what we are stuck with

            this.$emit("delete-post", post_data);
        }

    },
    mounted() {

        this.initialize();

    }
}
</script>

<template>

    <div v-for="post in posts_" :key="post.id" class="m-1">
        <PhotoPost :post_data="post" @delete-post="RemovePost" />
    </div>

</template>

<style>

</style>