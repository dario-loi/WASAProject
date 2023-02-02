<script>
export default {
    props: ['author', 'date', 'text']
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