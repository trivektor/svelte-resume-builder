<script>
  export let sections;
</script>

{#each sections as section}
  <div class="mb-4">
    <label class="block text-gray-700 font-bold">Name</label>
    <input class="shadow w-full px-1 py-1" bind:value={section.name} />
  </div>
  <div class="mb-4">
    <label class="block text-gray-700 font-bold">Description</label>
    <textarea class="block w-full" bind:value={section.description}></textarea>
  </div>
{/each}
