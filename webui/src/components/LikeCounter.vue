<script>

export default {
    props: ['likes_count', 'liked'],

    data: function () {
        return {
            is_liked: this.liked,
            count: this.likes_count
        }
    },

    methods: {
        async like() {
            this.is_liked = !this.is_liked
            if (this.is_liked) {
                this.count += 1

            } else {
                this.count -= 1
            }
        },

        async refresh() {
            this.loading = true;
            this.errormsg = null;
            /*
            try {
                let response = await this.$axios.get("/");
                this.some_data = response.data;
            } catch (e) {
                this.errormsg = e.toString();
            }
            */
            this.loading = false;

            console.log("refresh")
        }

    },

    mounted() {
        this.refresh()
    }

}

</script>

<template>
    <div class="d-flex align-items-center pb-2">

        <a type="button" class="btn btn-sm btn-outline-secondary likeButton" @click="like"
            :class="{'btn-outline-success': is_liked, 'btn-outline-danger': !is_liked}">
            <Transition name="slide-up" mode="out-in">
                <span v-if="is_liked" class="bi bi-heart-fill"> Like</span>
                <span v-else class="bi bi-heart"> Unlike</span>
            </Transition>
        </a>
        <span class="badge bg-secondary ms-2 me-2">
            <Transition name="slide-up" mode="out-in" :duration="{ 'enter': 300, 'leave': 100 }">
                <div :key="count">
                    {{ count }}
                </div>
            </Transition>
        </span>
    </div>
</template>


<style>
.slide-up-enter-active,
.slide-up-leave-active {
    transition: all .2s cubic-bezier(0.165, 0.84, 0.44, 1);
}

.slide-up-enter,
.slide-up-leave-to {
    transform: translateY(-10px);
    opacity: 0;
}

.slide-down-enter-active,
.slide-down-leave-active {
    transition: all .2s cubic-bezier(0.165, 0.84, 0.44, 1);
}

.slide-down-enter,
.slide-down-leave-to {
    transform: translateY(10px);
    opacity: 0;
}

.likeButton {
    padding: 0.2rem 0.5rem;
    font-size: 0.8rem;
    width: 80px;
}
</style>