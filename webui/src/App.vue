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
		async logout() {

			this.$user_state.username = null;
			this.$user_state.headers.Authorization = null;
			this.$router.push("/login");

		},

		async PerformSearch() {

			let search = document.querySelector("input").value;

			search = search.trim();

			console.log("searching for: " + search);

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

				let response = await this.$axios.get("/users/", {
					params: {
						"search-term": search,
					},
					headers: header
				});

				if (response.status == 200) {
					this.search_results = response.data;

					console.table(this.search_results);
				} else {
					console.log(response);
				}
			}
		}
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

			<div class="collapse navbar-collapse" id="Links">
				<ul class="navbar-nav mr-auto col-9">
					<li class="nav-item active">
						<RouterLink class="nav-link" to="/">Stream</RouterLink>
					</li>
					<li class="nav-item">
						<RouterLink class="nav-link" to="/about">Profile</RouterLink>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#" @click="Logout()">Logout</a>
					</li>
				</ul>

				<div class="col-3">
					<form class="nav form-inline my-2 my-md-0">
						<input class="form-control" type="text" placeholder="Search" aria-label="Search"
							@input="PerformSearch()">
						<!-- Results -->
						<ul class="list-group">
							<li class="list-group-item" v-for="user in search_results" :key="user.id">
								<RouterLink :to="'/users/' + user.id">
									{{ user.username }}
								</RouterLink>
							</li>

						</ul>
					</form>
				</div>
			</div>
		</div>
	</nav>

	<div class="h-100 w-100 ">
		<main class="h-100">
			<div class="container-fluid h-100">
				<div class="row h-100 p-4">
					<div class="col-3"></div>
					<div class="col-6 shadow-lg bg-light opacity-75 rounded">
						<RouterView></RouterView>
					</div>
					<div class="col-3"></div>
				</div>
			</div>
		</main>
	</div>

</template>

<style>

</style>
