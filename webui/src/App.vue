<script setup>
import { RouterLink, RouterView } from "vue-router";
</script>
<script>
export default {};
</script>

<template>
	<nav class="navbar navbar-expand-lg navbar-dark bg-dark bg-gradient shadow-lg fixed-top z-depth-5">
		<div class="container">
			<span class="navbar-brand mb-0 bg-dark bg-gradient h1">WasaPhoto!</span>
			<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExample07"
				aria-controls="navbarsExample07" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>

			<div class="collapse navbar-collapse" id="navbarsExample07">
				<ul class="navbar-nav mr-auto col-9">
					<li class="nav-item active">
						<RouterLink class="nav-link" to="/">Home</RouterLink>
					</li>
					<li class="nav-item">
						<RouterLink class="nav-link" to="/about">About</RouterLink>
					</li>

				</ul>

				<div class="col-3">
					<form class="form-inline my-2 my-md-0">
						<input class="form-control" type="text" placeholder="Search" aria-label="Search">
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
					<div class="col-6 shadow-lg bg-light opacity-75">
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
