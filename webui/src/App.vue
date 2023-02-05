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
			console.log("Logging out")
			this.$router.push("/");

		},

		async PerformSearch() {

			let search = document.querySelector("input").value;

			search = search.trim();

			if (search.length > 0) {
				// query ./users for results

				const searcher_id = this.$user_state.headers.Authorization;

				if (searcher_id == null) {
					return
				}

				const header = {
					"Authorization": searcher_id,
					"user_name": this.$user_state.username
				}

				let response = await this.$axios.get("/users", {
					params: {
						"search_term": search
					},
					headers: header
				});

				if (response.status == 200) {
					this.search_results = response.data;
				} else {
					console.log
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

		async ToStream() {

			if (this.$user_state.username == null) {
				return
			}

			this.$router.push("/stream/" + this.$user_state.username);
		},

		async refresh() {
			if (this.$user_state.username == null) {
				console.log("Empty username, redirecting to login")
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

			<span class="navbar-brand mb-0 bg-dark bg-gradient h1" style="vertical-align: middle">
				<i class="bi bi-camera-reels text-white m-1 mb-1 mt-1 bg-dark bg-gradient 
				"></i>WasaPhoto!</span>
			<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#Links"
				aria-controls="Links" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>

			<div class="collapse navbar-collapse row p-2 pb-0 pt-1" id="Links">
				<ul class="navbar-nav mr-auto col-md-4 col-sm-6 p-5 pt-1 pb-0" style="font-size: large;">
					<li class="nav-item active">
						<div class="flex-row align-content-center">
							<a class="nav-link" role="button" :class="{
								disabled: $user_state.username == null, 'd-none': $user_state.username == null,
								active: $user_state.current_view == $views.STREAM
							}" @click="ToStream()"><i class="bi m-1 mt-0 mb-0 bi-film text-white" style="vertical-align: middle"></i>Stream
							</a>
						</div>

					</li>
					<li class="nav-item">

						<a class="nav-link" role="button" :class="{
							disabled: $user_state.username == null, 'd-none': $user_state.username == null,
							active: $user_state.current_view == $views.PROFILE
						}" @click="ToProfile()"><i class="bi m-1 mt-0 mb-0 bi-person-circle text-white m-1 mb-1 mt-1"
								style="vertical-align: middle;">
							</i>Profile</a>
					</li>
					<li class="nav-item">

						<a class="nav-link" role="button"
							:class="{ disabled: $user_state.username == null, 'd-none': $user_state.username == null }"
							@click="Logout()"><i class="bi m-1 mt-0 mb-0 bi-door-open text-white m-1 mb-1 mt-1"
								style="vertical-align: middle;">
							</i>Logout</a>
					</li>
				</ul>

				<div class="col-md-4 col-sm-0 text-light">

					<h5>
						{{
							$user_state.username == null ? "Not Logged In" : "Logged in as " + $user_state.username
						}}
					</h5>

				</div>


				<div class="col-md-4 col-sm-6">
					<form class="nav form-inline my-2 my-md-0" :class="{
						disabled: $user_state.username == null, 'd-none': $user_state.username == null
					}">
						<input class="form-control" id="SearchBox" type="text" placeholder="Search" aria-label="Search"
							@input="PerformSearch()">
						<!-- Results -->
						<datalist class="list-group custom-select w-25 dropdown mt-5 position-absolute">

							<option class=" list-group-item align-middle" v-for="user in search_results"
								:key="user['username-string']">

								<i class="bi-person-circle m-2 fa-lg" style="font-size: 1.5rem;"></i>
								<RouterLink class="text-dark text-decoration-none m-0" style="font-size: 1.0rem;"
									:to="'/profile/' + user['username-string']">
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
					<div class="col-md-6 col-sm-10 mt-2 mb-4 shadow-lg bg-light rounded">
						<RouterView :key="$route.fullPath"></RouterView>
					</div>
					<div class="col-md-3 col-sm-1"></div>
				</div>
			</div>
		</main>
	</div>

</template>

<style>

</style>
