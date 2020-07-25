<script>
  export let id;
  import {onMount} from 'svelte';
  import ResumeForm from './ResumeForm.svelte';
  import {resumeStore} from '../stores';

  let resume;

  resumeStore.subscribe((data) => resume = data);

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

  async function onSubmit() {
    try {
      await window.fetch(`http://localhost:8080/resumes/${id}`, {
        header: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        method: 'PUT',
        body: JSON.stringify(resume),
      });
    } catch (err) {
      console.error(err);
    }
  }
</script>

<section class="bg-white rounded py-5 px-5 block border-2 border-gray">
  <form on:submit|preventDefault={onSubmit}>
    <ResumeForm />
  </form>
</section>
