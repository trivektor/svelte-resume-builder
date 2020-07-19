<script>
  import {navigate} from 'svelte-routing';
  import ResumeForm from './ResumeForm.svelte';

  let resume = {
    name: '',
    description: '',
  };

  async function onSubmit(event) {
    event.preventDefault();

    try {
      const result = await window.fetch('http://localhost:8080/resumes', {
        method: 'POST',
        body: JSON.stringify(resume),
        header: {'Content-Type': 'application/json', 'Accept': 'application/json'},
      });
      const json = await result.json();

      navigate(`/resumes/${json.id}`);
    } catch (err) {

    }
  }
</script>

<section class="bg-white rounded py-5 px-5 block border-2 border-gray">
  <form on:submit={onSubmit}>
    <ResumeForm resume={resume} />
  </form>
</section>
