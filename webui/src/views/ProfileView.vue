<script>
export default {
    data: function () {
        return {
            followers: 0,
            following: 0,
            posts: 0,
            is_me: false,
            is_banned: false,
            is_following: false,
            username: null,
            has_banned_you: false,
            photos: [] // list of IDs, pairs of ("hash", SHA256 hash of the photo)
        }
    },
    methods: {
        async refresh() {

            this.username = this.$route.params.username;

            // Redirect to login if not logged in
            if (this.$user_state.username == null) {
                this.$router.push("/");
                return
            }

            if (this.$route.params.username == this.$user_state.username) {
                this.is_me = true;
            }

            this.$user_state.current_view = this.$views.PROFILE;

            // Check if you are banned

            let response = await this.$axios.get("/users/" + this.username + "/bans", {
                headers: this.$user_state.headers
            });

            this.has_banned_you = response.data.users.map(x => x["username-string"]).includes(this.$user_state.username);

            // Check if banned

            response = await this.$axios.get("/users/" + this.$user_state.username + "/bans", {
                headers: this.$user_state.headers
            });

            this.is_banned = response.data.users.map(x => x["username-string"]).includes(this.username);

            if (!this.has_banned_you) {

                // Get following

                let response = await this.$axios.get("/users/" + this.username + "/following", {
                    headers: this.$user_state.headers
                });


                this.following = response.data["follow-list"].length;

                // Get followers 

                response = await this.$axios.get("/users/" + this.username + "/followers", {
                    headers: this.$user_state.headers
                });

                this.followers = response.data["follow-list"].length;

                if (!this.is_me) {

                    // check if I follow him

                    this.is_following = response.data["follow-list"].map(x => x["username-string"]).includes(this.$user_state.username);
                }

                // Get photos

                response = await this.$axios.get("/users/" + this.username + "/profile/photos", {
                    headers: this.$user_state.headers
                });

                this.photos = response.data["posts"];

                this.posts = this.photos.length;

                console.log(this.has_banned_you)

            }

        },

        async DeletePost(post_data) {

            this.refresh();
        },

        async ChangeName() {

            const new_name = prompt("Change name", "New name");

            if (new_name == null || new_name == "") {
                return
            }


            if (!new_name.match("^[a-zA-Z][a-zA-Z0-9_]{2,32}$")) {
                alert("Invalid username, must respect RegEx: ^[a-zA-Z][a-zA-Z0-9_]{2,32}$ (3 - 32 characters, first character must be a letter, only letters, numbers and underscores allowed)");
                return;
            }

            const req_body = {
                "username-string": new_name
            }

            const res = await this.$axios.put("/users/" + this.$user_state.username + "/profile", req_body, {
                headers: this.$user_state.headers
            });

            if (res.statusText != "No Content" && res.statusText != "OK") {

                alert("Error: " + res.statusText);
                console.table(res);
                return
            }

            this.$user_state.username = new_name;
            this.username = new_name;
            this.$user_state.headers["Authorization"] = res.data.token.hash;

            this.refresh();
        },

        async Follow() {

            const req_body = {
                "username-string": this.$user_state.username
            }

            const res = await this.$axios.put("/users/" + this.$user_state.username + "/following/" + this.username, req_body, {
                headers: this.$user_state.headers
            });

            if (res.statusText != "No Content") {

                alert("Error: " + res.statusText);
                console.table(res);
                return
            }

            this.is_following = true;
            this.followers += 1;
        },

        async Unfollow() {

            if (!this.is_following) {
                return
            }

            const req_body = {
                "username-string": this.$user_state.username
            }

            const res = await this.$axios.delete("/users/" + this.$user_state.username + "/following/" + this.username, {
                headers: this.$user_state.headers,
                data: req_body
            });

            if (res.statusText != "No Content") {

                alert("Error: " + res.statusText);
                console.table(res);
                return
            }

            this.is_following = false;
            this.followers -= 1;
        },

        async Ban() {

            const res = await this.$axios.put("/users/" + this.$user_state.username + "/bans/" + this.username, {}, {
                headers: this.$user_state.headers
            });

            if (res.statusText != "No Content") {

                alert("Error: " + res.statusText);
                console.table(res);
                return
            }

            this.is_banned = true;
        },

        async UnBan() {

            const res = await this.$axios.delete("/users/" + this.$user_state.username + "/bans/" + this.username, {
                headers: this.$user_state.headers
            });

            if (res.statusText != "No Content") {

                alert("Error: " + res.statusText);
                console.table(res);
                return
            }

            this.is_banned = false;

        },

    },

    mounted() {
        this.refresh()
    }
}
</script>

<template>
    <div class="container">
        <div class="align-items-center text-center h-100">
            <div class="container text-center pt-3 pb-2 border-bottom">
                <div class="row w-100 my-3">
                    <h2 class="col-3 text-break d-inline-block" style="vertical-align: middle;">
                        <i class="bi-person-circle mx-1"></i>{{ username }}'s profile.
                    </h2>
                    <div class="col-9" style="align-items: center; vertical-align: middle;">
                        <div class="row">
                            <div class="col-4">
                                <div class="row border p-1 pt-2 rounded me-1 shadow-sm">
                                    <div class="col-12">
                                        <h5>Posts</h5>
                                    </div>
                                    <div class="col-12">
                                        <h5> {{ posts }}</h5>
                                    </div>
                                </div>
                            </div>
                            <div class="col-4">
                                <div class="row border p-1 pt-2 rounded me-1 shadow-sm">
                                    <div class="col-12">
                                        <h5>Followers</h5>
                                    </div>
                                    <div class="col-12">
                                        <h5>{{ followers }}</h5>
                                    </div>
                                </div>
                            </div>
                            <div class="col-4">
                                <div class="row border p-1 pt-2 rounded me-1 shadow-sm">
                                    <div class="col-12">
                                        <h5>Following</h5>
                                    </div>
                                    <div class="col-12">
                                        <h5>{{ following }}</h5>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div v-if="is_me" class="row w-100">
                    <div class="row w-100">
                        <div class="col-3">
                            <button class="btn btn-primary btn-md" type="button" @click="ChangeName()">
                                <i class="bi-pencil-square"></i>
                                Change Name
                            </button>
                        </div>
                    </div>
                </div>
                <div v-else>
                    <div class="row w-100 align-content-between my-1">
                        <!-- Follow Button -->
                        <div class="col">
                            <Transition name="fade" mode="out-in">
                                <div v-if="is_following && !has_banned_you">
                                    <button class="btn btn-warning btn-lg" type="button" @click="Unfollow()">
                                        <i class="bi-person-dash-fill"></i>
                                        Unfollow
                                    </button>
                                </div>
                                <div v-else-if="!is_following && !has_banned_you">
                                    <button class="btn btn-primary btn-lg" type="button" @click="Follow()">
                                        <i class="bi-person-plus-fill"></i>
                                        Follow
                                    </button>
                                </div>
                            </Transition>
                        </div>
                        <!-- Ban Button -->
                        <div class="col">
                            <Transition name="fade" mode="out-in">
                                <div v-if="is_banned && !has_banned_you">
                                    <button class="btn btn-success btn-lg" type="button" @click="UnBan()">
                                        <i class="bi-person-check-fill"></i>
                                        Unban
                                    </button>
                                </div>
                                <div v-else-if="!is_banned && !has_banned_you">
                                    <button class=" btn btn-danger btn-lg" type="button" @click="Ban()">
                                        <i class="bi-person-x-fill"></i>
                                        Ban
                                    </button>
                                </div>
                            </Transition>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div v-if="has_banned_you" class="container">
        <div class="row">
            <div class="col-12">
                <div class="alert alert-danger" role="alert">
                    <h4 class="alert-heading">You have been banned by this user!</h4>
                    <p>Sorry, but you have been banned from this user's profile. You cannot view their posts or
                        interact with them.</p>
                    <hr>
                    <p class="mb-0">Try not to be so mean next time!</p>
                </div>
            </div>
        </div>
    </div>
    <div v-else class="container">
        <Stream :posts="photos" @delete-post="DeletePost" />
    </div>
</template>

<style>
.fade-enter-active,
.fade-leave-active {
    transition: opacity cubic-bezier(0.4, 0, 0.2, 1) 0.1s
}

.fade-enter,
.fade-leave-to {
    opacity: 0
}
</style>