<script>
  import {createEventDispatcher} from 'svelte';
  import Sections from './Sections.svelte';
  export let resume;

  const dispatch = createEventDispatcher();

  function addSection() {
    dispatch('addSection');
  }
</script>

<h2 class="mb-2">Overview</h2>
<div class="mb-4">
  <label class="block text-gray-700 font-bold">Name</label>
  <input class="shadow w-full px-1 py-1" bind:value={resume.name} />
</div>
<div class="mb-4">
  <label class="block text-gray-700 font-bold">Description</label>
  <textarea class="block w-full" bind:value={resume.description}></textarea>
</div>
{#if resume.id}
  <h2 class="mt-2 mb-2">Sections</h2>
  <Sections sections={resume.sections} />
{/if}
<div>
  {#if resume.id}
    <button type="button" class="bg-blue-500 px-2 py-2 text-white" on:click={addSection}>Add Section</button>
  {/if}
  <button type="submit" class="bg-green-500 px-2 py-2 text-white">Save Resume</button>
</div>
