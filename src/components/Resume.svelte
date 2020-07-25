<script>
  export let id;
  import {onMount} from 'svelte';
  import ResumeForm from './ResumeForm.svelte';
  import {resumeStore} from '../stores';

  onMount(async () => {
    const result = await fetch(`http://localhost:8080/resumes/${id}`, {
      header: {'Accept': 'application/json'},
    });
    const json = await result.json();

    resumeStore.update(() => ({
      ...json,
      sections: json.sections || [],
    }));
  });

  async function onSubmit(event) {
    event.preventDefault();

    try {
      console.log({resume});
    } catch (err) {

    }
  }
</script>

<section class="bg-white rounded py-5 px-5 block border-2 border-gray">
  <form on:submit={onSubmit}>
    <ResumeForm />
  </form>
</section>
