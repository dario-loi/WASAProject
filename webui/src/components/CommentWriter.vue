<script>
export default {
    props: ['photo_id', 'author_name'],

    emits: ['comment'],

    data: function () {
        return {
            author: this.author_name,
            photo_id_: this.photo_id,
            curr_time: null,
        }
    },

    methods: {
        async initialize() {

            this.format_date_now();
            setInterval(this.format_date_now, 60000);
        },

        async format_date_now() {
            // Get date in RFC3339 format
            let now = new Date().toISOString();

            // format to dd month yyyy at hh:mm

            let date_split = now.split("T");
            let date = date_split[0].split("-");
            let time = date_split[1].split(":");
            time = time[0] + ":" + time[1];
            date = date[2] + " " + this.$months[parseInt(date[1]) - 1] + " " + date[0] + " at " + time;

            console.log("Date: " + date);

            this.curr_time = date;
        },

        async add_comment() {
            let text = document.getElementById("comment").value;
            this.$emit('comment', text);

            // Clear the text area

            document.getElementById("comment").value = "";
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
                <h5 class="card-title col-3 mx-2">Leave a Comment!</h5>
                <!-- Align the text to the top right of the card -->
                <div class="col-3 text-end align-middle">
                    <span class="card-subtitle text-muted w-auto">{{ curr_time }}</span>
                </div>
            </div>

            <p class="card-text">

                <textarea class="form-control p-1" id="comment" rows="3"></textarea>

            </p>
            <div class="row justify-content-end">
                <button type="button" class="btn mx-2 btn-primary col-auto" @click="add_comment"
                    style="font-size: 0.8em, font-style: italic;">Comment</button>
            </div>

        </div>
    </div>
</template>

<style>
.WasaPhotoComment {
    margin-bottom: 1rem;
}
</style>