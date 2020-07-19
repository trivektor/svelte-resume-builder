<script>
	import router from 'page';
	import Resumes from './components/Resumes.svelte';
	import NewResume from './components/NewResume.svelte';

	let page;

	router('/', () => page = Resumes);
	router('/resumes/new', () => page = NewResume);

	router.start();
</script>


<div class="px-4 py-4">
	<svelte:component this={page} />
</div>
