<script>
  export let resume;
</script>

<div class="mb-4">
  <label class="block text-gray-700 font-bold">Name</label>
  <input class="shadow w-full px-1 py-1" bind:value={resume.name} />
</div>
<div class="mb-4">
  <label class="block text-gray-700 font-bold">Description</label>
  <textarea class="block w-full" bind:value={resume.description}></textarea>
</div>
<div>
  <button type="submit" class="bg-green-500 px-2 py-2 text-white">Save Resume</button>
</div>
