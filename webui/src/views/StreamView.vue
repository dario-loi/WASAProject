<script>
import { Axios } from 'axios';

export default {
    data: function () {
        return {
            modal: null,
            uploaded_image: null,
        }
    },
    methods: {
        async initialize() {

            this.$user_state.current_view = this.$views.STREAM;

            const mod = bootstrap.Modal.getOrCreateInstance(document.getElementById('exampleModal'))
            document.body.appendChild(mod._element)
            this.modal = mod

            // Get the default image's base64 data from ./assets/default.png

            const default_image = await this.$axios.get("/resources/photos/default", {
                responseType: "image/jpeg"
            });

            this.uploaded_image = default_image.data;
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

            const textAsBuffer = new TextEncoder().encode(filename);
            const hashBuffer = await window.crypto.subtle.digest('SHA-256', textAsBuffer);
            const hashArray = Array.from(new Uint8Array(hashBuffer))
            const img_id = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');

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
        </div>
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
                                <input class="form-control" type="file" id="fileInput">

                                <label for="formFile" class="form-label">Caption</label>
                                <input class="form-control" type="text" id="captionInput">
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