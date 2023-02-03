<script setup>
import { RouterLink, RouterView } from "vue-router";
</script>
<script>
export default {

	data: function () {
		return {
			search_results: null,
		}
	},

	methods: {
		async Logout() {

			this.$user_state.username = null;
			this.$user_state.headers.Authorization = null;
			this.$router.push("/");

		},

		async PerformSearch() {

			let search = document.querySelector("input").value;

			search = search.trim();

			if (search.length > 0) {
				// query /users/ for results

				const searcher_id = this.$user_state.headers.Authorization;

				if (searcher_id == null) {
					return
				}

				//TODO: standardize to Authorization header
				const header = {
					"searcher_id": searcher_id
				}

				let response = await this.$axios.get("/users", {
					params: {
						"search_term": search,
					},
					headers: header
				});

				if (response.status == 200) {
					this.search_results = response.data;
				} else {
					this.search_results = null;
				}
			}
			else {
				this.search_results = null;
			}
		},

		async ToProfile() {

			if (this.$user_state.username == null) {
				return
			}

			this.$router.push("/profile/" + this.$user_state.username);
		},

		async refresh() {
			if (this.$user_state.username == null) {
				this.$router.push("/");
			}
		}

	},

	mounted() {
		this.refresh()
	}
}

</script>

<template>
	<nav class="navbar navbar-expand navbar-dark bg-dark bg-gradient shadow-lg fixed-top z-depth-5" role="navigation">
		<div class="container">
			<span class="navbar-brand mb-0 bg-dark bg-gradient h1">WasaPhoto!</span>
			<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#Links"
				aria-controls="Links" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>

			<div class="collapse navbar-collapse row" id="Links">
				<ul class="navbar-nav mr-auto col-md-9 col-sm-6">
					<li class="nav-item active">
						<a class="nav-link" :class="{
							disabled: $user_state.username == null, 'd-none': $user_state.username == null,
							active: $current_view == $views.STREAM
						}" to="/">Stream
						</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" :class="{
							disabled: $user_state.username == null, 'd-none': $user_state.username == null,
							active: $current_view == $views.PROFILE
						}" href="#" @click="ToProfile()">Profile</a>
					</li>
					<li class="nav-item">
						<a class="nav-link"
							:class="{ disabled: $user_state.username == null, 'd-none': $user_state.username == null }"
							href="#" @click="Logout()">Logout</a>
					</li>
				</ul>

				<div class="col-md-3 col-sm-6">
					<form class="nav form-inline my-2 my-md-0" :class="{
						disabled: $user_state.username == null, 'd-none': $user_state.username == null
					}">
						<input class="form-control" ref="SearchBox" type="text" placeholder="Search" aria-label="Search"
							@input="PerformSearch()">
						<!-- Results -->
						<datalist class="list-group w-75 dropdown mt-5 position-absolute" :style="{
							'left': $refs.SearchBox.offsetLeft + 'px',
							'width': $refs.SearchBox.offsetWidth + 'px',
							'z-index': 1000
						}">

							<option class=" list-group-item w-75" v-for="user in search_results"
								:key="user['username-string']">
								<RouterLink :to="'/users/' + user['username-string']">
									{{ user['username-string'] }}
								</RouterLink>
							</option>

						</datalist>
					</form>
				</div>
			</div>
		</div>
	</nav>

	<div class="h-100 w-100 ">
		<main class="h-100">
			<div class="container-fluid h-100">
				<div class="row h-100 p-4">
					<div class="col-md-3 col-sm-1"></div>
					<div class="col-md-6 col-sm-10 shadow-lg bg-light opacity-75 rounded">
						<RouterView></RouterView>
					</div>
					<div class="col-md-3 col-sm-1"></div>
				</div>
			</div>
		</main>
	</div>

</template>

<style>

</style>
