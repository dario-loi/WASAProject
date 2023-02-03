<script>
export default {
    data: function () {
        return {
            error: false,
            loading: false,
            some_data: null,
        }
    },
    methods: {
        async refresh() {

        },
        async login() {

            let username = document.getElementById("login-form").value;

            let response = await this.$axios.put("/session", {
                "username-string": username
            });

            console.log({
                "username-string": username
            })

            //check if the response is 201

            if (response.status == 201) {
                this.error = false;
                this.$user_state.username = username
                this.$user_state.headers.Authorization = response.data["token"]["hash"]
                this.$router.push("/stream/" + username + "/");
            } else {
                console.table(response);
                this.error = true;
            }
        }
    },
    mounted() {
        this.refresh()
    }
}
</script>

<template>
    <div class="container text-center pt-3 pb-2 border-bottom">
        <h2>Login</h2>
    </div>


    <div class="h-75 d-flex align-items-center justify-content-center">
        <form class="border border-dark p-5 rounded shadow-lg">
            <!-- Username input -->
            <div class="form-outline mb-4">
                <input type="text" id="login-form" class="form-control" pattern="^[a-zA-Z][a-zA-Z0-9_]{2,32}$" />
                <label class="form-label" for="login-form">Username</label>
            </div>

            <!-- Submit button -->
            <button type="button" class="btn btn-primary btn-block mb-4" @click="login()">Sign in</button>

        </form>

    </div>


</template>

<style>

</style>
