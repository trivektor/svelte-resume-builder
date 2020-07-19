<script>
  import ResumeForm from './ResumeForm.svelte';

  let resume = {
    name: '',
    description: '',
  };

  function onSubmit(event) {
    event.preventDefault();

    console.log(resume);
  }
</script>

<section class="bg-white rounded py-5 px-5 block border-2 border-gray">
  <form on:submit={onSubmit}>
    <ResumeForm resume={resume} />
  </form>
</section>
