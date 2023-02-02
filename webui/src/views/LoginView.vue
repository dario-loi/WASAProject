<script>
export default {
    data: function () {
        return {
            errormsg: null,
            loading: false,
            some_data: null,
        }
    },
    methods: {
        async refresh() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.get("/");
                this.some_data = response.data;
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;

            console.table(this.some_data)
        },
    },
    mounted() {
        this.refresh()
    }
}
</script>

<template>
    <div>
        <div class="container text-center ">
            <h2>Login</h2>
        </div>


        <div class="container">

        </div>
    </div>
</template>

<style>

</style>
