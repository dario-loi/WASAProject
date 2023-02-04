<script>
export default {
    data: function () {
        return {
            followers: 0,
            following: 0,
            posts: 0,

        }
    },
    methods: {
        async initialize() {

            // Redirect to login if not logged in
            if (this.$user_state.username == null) {
                this.$router.push("/");
                return
            }

            this.$user_state.current_view = this.$views.PROFILE;

        },
    },
    mounted() {
        this.initialize()
    }
}
</script>

<template>

    <div class="container">

        <div class="align-items-center text-center h-100">

            <div class="container text-center pt-3 pb-2 border-bottom">
                <div class="row w-100">
                    <h2 class="col-3">
                        <i class="bi-person-circle mx-1"></i>{{ this.$user_state.username }}'s profile.
                    </h2>

                    <div class="col-9" style="align-items: center; vertical-align: middle;">
                        <div class="row">
                            <div class="col-4">
                                <div class="row">
                                    <div class="col-12">
                                        <h5>Followers</h5>
                                    </div>
                                    <div class="col-12">
                                        <h5>{{ followers }}</h5>
                                    </div>
                                </div>
                            </div>
                            <div class="col-4">
                                <div class="row">
                                    <div class="col-12">
                                        <h5>Following</h5>
                                    </div>
                                    <div class="col-12">
                                        <h5>{{ following }}</h5>
                                    </div>
                                </div>
                            </div>
                            <div class="col-4">
                                <div class="row">
                                    <div class="col-12">
                                        <h5>Posts</h5>
                                    </div>
                                    <div class="col-12">
                                        <h5> {{ posts }}</h5>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>
                <div class="row w-100 align-content-between my-1">

                    <!-- Follow Button -->

                    <div class="col">
                        <button class="btn btn-primary btn-lg" type="button">
                            <i class="bi-person-plus-fill"></i>
                            Follow
                        </button>
                    </div>

                    <!-- Ban Button -->

                    <div class="col">
                        <button class="btn btn-danger btn-lg" type="button">
                            <i class="bi-person-x-fill"></i>
                            Ban
                        </button>
                    </div>

                </div>
            </div>

        </div>

    </div>

</template>