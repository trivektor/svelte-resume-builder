<script>
	import {Router, Route} from 'svelte-routing';
	import ApolloClient from 'apollo-client';
	import {createHttpLink} from 'apollo-link-http';
	import {InMemoryCache} from 'apollo-cache-inmemory';
	import {setClient} from 'svelte-apollo';

	import Resumes from './components/Resumes.svelte';
	import NewResume from './components/NewResume.svelte';
	import Resume from './components/Resume.svelte';

	export let url = '';
	const link = createHttpLink({uri: 'http://localhost:4000/graphql'});
	const cache = new InMemoryCache();
	const client = new ApolloClient({link, cache});

	setClient(client);
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
