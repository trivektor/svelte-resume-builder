<script>
  import {navigate} from 'svelte-routing';

  import Section from './Section.svelte';
  import {resumeStore} from '../stores';

  let resume = {};

  resumeStore.subscribe((data) => {
    resume = data;
  });

  function addSection() {
    resumeStore.update((state) => ({
      ...state,
      sections: [
        ...(state.sections || []),
        {},
      ],
    }));
  }

  function deleteSection(event) {
    const {detail: {sectionIndex}} = event;
    const updatedSections = (resume.sections || []).concat([]);

    updatedSections.splice(sectionIndex, 1);

    resumeStore.update((state) => ({
      ...state,
      sections: updatedSections,
    }));
  }

  async function deleteResume() {
    try {
      await window.fetch(`http://localhost:8080/resumes/${resume.id}`, {
        method: 'DELETE',
      });

      navigate("/", {replace: true});
    } catch (err) {
      console.error(err);
    }
  }
</script>

<h2 class="mb-2">Overview</h2>
<div class="mb-4">
  <label class="block text-gray-700 font-bold">Name</label>
  <input class="w-full px-2 py-2" bind:value={resume.name} />
</div>
<div class="mb-4">
  <label class="block text-gray-700 font-bold">Description</label>
  <textarea class="block w-full px-2 py-2" bind:value={resume.description}></textarea>
</div>
{#if resume.id}
  <h2 class="mt-2 mb-2">Sections</h2>
  {#each resume.sections as section, sectionIndex}
    <Section section={section} {sectionIndex} on:delete={deleteSection} />
  {/each}
{/if}
<div>
  {#if resume.id}
    <button type="button" class="bg-blue-500 px-2 py-2 text-white" on:click={addSection}>Add Section</button>
  {/if}
  <button type="submit" class="bg-green-500 px-2 py-2 text-white">Save Resume</button>
  {#if resume.id}
  <button type="button" class="bg-red-500 px-2 py-2 text-white" on:click={deleteResume}>Delete Resume</button>
  {/if}
</div>
