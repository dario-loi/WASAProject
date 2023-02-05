<script>
export default {
    props: ['comment'],


    data: function () {
        return {
            author: this.comment.author,
            date: null,
            text: this.comment.text
        }
    },

    methods: {
        async initialize() {
            this.date = this.comment.creation_time;

            // format to dd month yyyy at hh:mm

            this.date = this.date.split("T");
            let date = this.date[0].split("-");
            let time = this.date[1].split(":");
            time = time[0] + ":" + time[1];
            date = date[2] + " " + this.$months[parseInt(date[1]) - 1] + " " + date[0] + " at " + time;

            this.date = date;
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
            <div class="row justify-content-between">
                <h5 class="card-title col-3">{{ author }}</h5>
                <!-- Align the text to the top right of the card -->
                <div class="col-3 text-end align-middle">
                    <h6 class="card-subtitle mb-text-muted w-auto">{{ date }}</h6>
                </div>
            </div>

            <p class="card-text">{{ text }}</p>
        </div>
    </div>
</template>

<style>
.WasaPhotoComment {
    margin-bottom: 1rem;
}
</style>