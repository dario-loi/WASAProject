<script>
export default {
    data: function () {
        return {
        }
    },
    methods: {
        async refresh() {

            this.$user_state.current_view = this.$views.STREAM;

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
            <span>
                {{ this.$user_state.username }}
            </span>
        </div>

    </div>

</template>