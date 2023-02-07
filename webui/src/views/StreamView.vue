<script>
import { Axios } from 'axios';

export default {
    data: function () {
        return {
            modal: null,
            stream_posts: [],
            stream_top: 0,
            there_are_more_posts: true,
        }
    },
    methods: {
        async initialize() {

            this.$user_state.current_view = this.$views.STREAM;

            const mod = bootstrap.Modal.getOrCreateInstance(document.getElementById('exampleModal'))
            document.body.appendChild(mod._element)
            this.modal = mod

            // Load stream from user's 

            const loading_factor = 8;

            let batch = await this.LoadStream(this.stream_top, this.stream_top + loading_factor)

            if (batch.length == 0) {
                this.there_are_more_posts = false;
            }

            this.stream_posts.push(...batch);
            this.stream_top += batch.length;

            // listen when the user scrolls to the bottom of the page

            window.onscroll = async () => {
                if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {

                    let batch = await this.LoadStream(this.stream_top, this.stream_top + loading_factor)

                    if (this.there_are_more_posts == false) {
                        return;
                    }

                    if (batch.length == 0) {
                        this.there_are_more_posts = false;
                    }

                    this.stream_posts.push(...batch);
                    this.stream_top += batch.length;

                }
            };

        },

        async DeletePost(post_data) {

            this.refresh();
        },

        async LoadStream(start, end) {
            let ret = await this.$axios.get("/users/" + this.$user_state.username + "/stream?from=" + start + "&offset=" + end, {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": this.$user_state.headers.Authorization
                }
            }
            ).catch((error) => {
                console.log(error);
                alert("Error loading stream");
                return [];
            }).then((response) => {
                return response;
            });

            return ret.data["posts"];
        },


        async UploadPhoto() {

            // Manually set the submit button to be waiting

            document.getElementById("submit-button").innerHTML = "Uploading...";
            document.getElementById("submit-button").classList.add("disabled");

            const image = document.getElementById("fileInput").files[0];

            let reader = new FileReader();
            let data = null;

            reader.onload = async () => {
                data = reader.result;
            }

            reader.onerror = function (error) {
                console.log('Error: ', error);
                alert("Error uploading photo")
                return
            };

            reader.readAsDataURL(image);

            // Wait for the reader to finish reading the file

            while (data == null) {
                await new Promise(r => setTimeout(r, 1000));
                console.log("waiting for reader to finish")
            }

            const filename = image.name;

            // Hash the filename, username, and current time to get a unique id for the image
            // This avoids hash collisions by the "Trust me bro" theorem. (The converse of 
            // the law of large numbers)
            const to_hash = filename + this.$user_state.username + Date.now().toString();

            const img_id = this.$hasher(to_hash);
            const author = this.$user_state.username;

            const caption = document.getElementById("captionInput").value;
            // strip data:image/png;base64, from the beginning of the string

            data = data.substring(22);

            const req_body = {
                "photo_data": data,
                "photo_desc": caption
            }

            let response = await this.$axios.put("/users/" + author + "/profile/photos/" + img_id
                , req_body, {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": this.$user_state.headers.Authorization
                }
            });

            if (response.status == 204) {

                // manually restyle and rename the submit button

                const submit_button = document.getElementById("submit-button");

                submit_button.classList.remove("btn-primary");
                submit_button.classList.remove("disabled");
                submit_button.classList.add("btn-success");
                submit_button.innerHTML = "Success!";

                setTimeout(() => {

                    // reset the button

                    submit_button.classList.remove("btn-success");
                    submit_button.classList.add("btn-primary");
                    submit_button.innerHTML = "Submit";
                }, 3000);

                // Clear the form

                document.getElementById("fileInput").value = "";
                document.getElementById("captionInput").value = "";

            }
            else {
                alert("Error uploading photo")
            }


        }
    },
    mounted() {
        this.initialize()
    },

    updated() {
        const mod = bootstrap.Modal.getOrCreateInstance(document.getElementById('exampleModal'))

        document.body.appendChild(mod._element)
        document.body.removeChild(this.modal._element)

        this.modal = mod
    }
}

</script>

<template>

    <div class="container">

        <div class="align-items-center text-center h-100">

            <div class="container text-center pt-3 pb-2 border-bottom">
                <h2>
                    <i class="bi-film mx-1"></i>Your Stream.
                </h2>

                <!-- Post Photo Button -->
                <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#exampleModal"
                    data-backdrop="false">
                    Post Photo
                </button>
            </div>

            <!-- Stream -->

        </div>


        <Stream :posts="stream_posts" :key="stream_posts.length" @delete-post="DeletePost"></Stream>
    </div>

    <!-- Modal -->
    <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="exampleModalLabel">Create a Post</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <!-- Image Input -->
                    <div class="mb-3">
                        <div class="row g-3 align-items-center">
                            <form id="formFile">
                                <label for="formFile" class="form-label">Upload Image</label>
                                <input class="form-control" type="file" id="fileInput" accept="image/png">

                                <label for="formFile" class="form-label">Caption</label>
                                <textarea class="form-control" type="text" id="captionInput" rows="6"></textarea>
                            </form>
                        </div>
                    </div>

                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                    <button type="button" class="btn btn-primary" id="submit-button" @click="UploadPhoto()">
                        Submit
                    </button>
                </div>
            </div>
        </div>
    </div>


</template>