<script>
  import {createEventDispatcher} from 'svelte';
  export let section;
  export let sectionIndex;

  const dispatch = createEventDispatcher();

  function deleteSection() {
    dispatch('delete', {sectionIndex});
  }
</script>

<div class="border-r border-b border-l border-t border-gray-400 p-5 mt-5 mb-5">
  <div class="mb-4">
    <label class="block text-gray-700 font-bold">Name</label>
    <input class="w-full px-2 py-2" bind:value={section.name} />
  </div>
  <div class="mb-4">
    <label class="block text-gray-700 font-bold">Description</label>
    <textarea class="block w-full px-2 py-2" bind:value={section.description}></textarea>
  </div>
  <button class="bg-red-500 text-white p-2" on:click={deleteSection}>
    Delete this section
  </button>
</div>
