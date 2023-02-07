<script>
export default {
    props: ['comment'],


    data: function () {
        return {
            author: this.comment.author["username-string"],
            created_at: null,
            text: this.comment.body
        }
    },

    emits: ['delete'],

    methods: {
        async initialize() {

            // format to dd month yyyy at hh:mm

            console.log(this.comment)

            let creation_time = this.comment["creation-time"];
            let date_split = creation_time.split("T");
            let date = date_split[0].split("-");
            let time = date_split[1].split(":");
            time = time[0] + ":" + time[1];
            date = date[2] + " " + this.$months[parseInt(date[1]) - 1] + " " + date[0] + " at " + time;

            this.created_at = date;
        }
    },

    mounted() {
        this.initialize();
    }
}
</script>

<template>
    <div class="card WasaPhotoComment shadow-lg">
        <div class="card-body">
            <div class="row justify-content-between" style="vertical-align: middle;">
                <h5 class="card-title col-5" style="font-size: 1.2em;">
                    <i class="bi bi-person-circle mx-1" style="font-size: 1.5em;"></i>{{ author }}
                </h5>

                <!-- Delete button-->

                <div class="col-3 text-end align-middle">
                    <button type="button" class="btn btn-danger btn-sm" v-on:click="$emit('delete', comment)">
                        <i class="bi bi-trash mx-1"></i>
                    </button>

                </div>

            </div>

            <hr class="my-2">

            <p class="card-text p-3">{{ text }}</p>

            <div class="row justify-content-end">
                <!-- Align the text to the top right of the card -->
                <div class="col-3 text-end align-middle">
                    <span class="card-subtitle v-center text-muted w-auto"
                        style="font-size: 0.8em, font-style: italic;">
                        {{ created_at }}</span>
                </div>

            </div>
        </div>
    </div>
</template>

<style>
.WasaPhotoComment {
    margin-bottom: 1rem;
}
</style>