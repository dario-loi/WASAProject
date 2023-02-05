<script>


export default {

    props: {
        post_data: {
            type: Object
        }
    },

    data: function () {
        return {
            datetime: null,
            is_your_post: false,
            photo_id: null,
            have_i_liked_this: false,
            username: null,
            likes: 0,
            comments: []
        }
    },

    methods: {

        async initialize() {
            this.datetime = this.post_data.created_at;
            this.photo_id = this.post_data.photo_id.hash;
            this.username = this.post_data.author_name["username-string"];

            console.log("ID is " + this.photo_id);

            // format to dd month yyyy at hh:mm

            this.datetime = this.datetime.split("T");
            let date = this.datetime[0].split("-");
            let time = this.datetime[1].split(":");
            time = time[0] + ":" + time[1];
            date = date[2] + " " + this.$months[parseInt(date[1]) - 1] + " " + date[0] + " at " + time;

            this.datetime = date;
            this.is_your_post = this.post_data.author_name["username-string"] == this.$user_state.username;

            // Fetch likes

            let response = await this.$axios.get("/users/" + this.post_data.author_name["username-string"] + "/profile/photos/" + this.photo_id + "/likes");

            this.likes = response.data.users.length;
            this.have_i_liked_this = response.data.users.map((user) => user["username-string"]).includes(this.$user_state.username);

            // Fetch comments

            response = await this.$axios.get("/users/" + this.post_data.author_name["username-string"] + "/profile/photos/" + this.photo_id + "/comments");

            console.log(this.likes)
        },

        async WriteComment() {

        },

        async DeletePost() {

        },

        async Like() {

            // Update the state on the server

            console.log("Request Path: " + "/users/" + this.post_data.author_name["username-string"] + "/profile/photos/" + this.photo_id + "/likes/" + this.$user_state.headers.Authorization);

            console.log(this.$user_state.headers.Authorization)

            let response = await this.$axios.put("/users/" + this.post_data.author_name["username-string"] + "/profile/photos/" + this.photo_id + "/likes/" + this.$user_state.headers.Authorization, {}, {
                headers: {
                    "Authorization": this.$user_state.headers.Authorization
                }
            });

            if (response.statusText != "No Content") {
                alert("Error: " + response.statusText);
                return;
            }

            // Only for consistency, the component does this internally.
            this.have_i_liked_this = true;
            this.likes++;
        },

        async Unlike() {

            // Update the state on the server

            console.log("Request Path: " + "/users/" + this.post_data.author_name["username-string"] + "/profile/photos/" + this.photo_id + "/likes/" + this.$user_state.headers.Authorization);

            console.log(this.$user_state.headers.Authorization)

            let response = await this.$axios.delete("/users/" + this.post_data.author_name["username-string"] + "/profile/photos/" + this.photo_id + "/likes/" + this.$user_state.headers.Authorization, {
                headers: {
                    "Authorization": this.$user_state.headers.Authorization
                }
            });

            if (response.statusText != "No Content") {
                alert("Error: " + response.statusText);
                return;
            }

            this.have_i_liked_this = false;
            this.likes--;
        }

    },

    mounted() {

        this.initialize();

    }
}
</script>

<template>


    <!-- Bordered Wrapper -->

    <div class="rounded p-2 m-4 border shadow-lg">


        <div class="row align-content-between my-2">

            <div class="col">
                <i class="bi-person-circle mx-1" style="font-size: 2em"></i>
                <span class="col font-weight-bold h4">
                    {{ post_data.author_name["username-string"] }}
                </span>
            </div>

            <!-- Right-aligned datetime -->
            <div class="col-auto ">
                <span class="text-muted v-center" style="font-size: 0.8em, font-style: italic;">
                    {{ datetime }}
                </span>

            </div>

            <!-- Delete Button -->

            <div class="col-auto" v-if="is_your_post">
                <button class="btn btn-danger v-center" @click="DeletePost">
                    <i class="bi-trash"></i>
                </button>
            </div>
        </div>

        <div class="row">

            <Photo :src="post_data.photo_id.hash" :alt="post_data.description" :style="{ width: '100%' }">
            </Photo>

        </div>

        <div class="row mt-3 align-content-start">
            <div class="col-auto">
                <LikeCounter class="v-center" :likes_count="this.likes" :liked="this.have_i_liked_this" @like="Like"
                    @unlike="Unlike"></LikeCounter>
            </div>
            <div class="col-auto d-flex align-items-center pb-2">
                <button class="btn bt-sm comment-button btn-outline-primary v-center">
                    <i class="bi-chat" @click="WriteComment()">Comment</i>
                </button>
            </div>

        </div>

        <!-- Divider -->

        <hr class="mt-1 mb-4">

        <!-- Caption -->

        <div class="row">

            <div class="col-12">

                <i class="bi-person-circle mx-1" style="font-size: 1.5em"></i>
                <span class="font-weight-bold h6" style="margin-right: 10px;">{{ username }}</span>
                <span> {{ post_data.description }}</span>

            </div>
        </div>

        <!-- Divider -->

        <hr class="my-4">

        <!-- Comments -->

        <div class="row">
            <div v-if="comments.length == 0" class="col-12 align-content-center w-100"><!-- Center the text -->
                <span class="h5 mx-1 font-weight-bold align-middle text-muted text-center">No comments yet.</span>
            </div>
            <div v-else class="col-12">
                <span class="h5 mx-1 font-weight-bold align-middle text-muted text-start">Comments: </span>
            </div>
            <div class="col-12">
                <Comment v-for="comment in comments" :comment="comment"></Comment>
            </div>

        </div>
    </div>


</template>

<style>
.v-center {

    display: inline-block;
    vertical-align: middle;
    line-height: normal;

}

.comment-button {

    padding: 0.35rem 0.5rem;
    font-size: 0.8rem;
    width: 120px;

}
</style>