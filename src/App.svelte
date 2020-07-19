<script>
	import {Router, Route} from 'svelte-routing';
	import Resumes from './components/Resumes.svelte';
	import NewResume from './components/NewResume.svelte';
	import Resume from './components/Resume.svelte';

	export let url = '';
</script>

<div class="px-4 py-4">
	<Router url="{url}">
		<Route path='/' component="{Resumes}" />
		<Route path='/resumes/new' component="{NewResume}" />
		<Route path='/resumes/:id' let:params>
			<Resume id="{params.id}" />
		</Route>
	</Router>
</div>
