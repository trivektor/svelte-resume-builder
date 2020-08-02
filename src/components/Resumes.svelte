<script>
  import {Link} from 'svelte-routing';
  import {onMount} from 'svelte';

  let resumes = [];

  onMount(async () => {
    const response = await window.fetch('http://localhost:8080/resumes', {
      header: {'Accept': 'application/json'}
    });
    const json = await response.json();

    resumes = json;
  });
</script>

<Link to="/resumes/new">
  <button class="bg-green-500 px-2 py-2 text-white inline-block">
    New Resume
  </button>
</Link>

<table class="table-auto">
  <thead>
    <tr>
      <th class="px-4 py-2"></th>
      <th class="px-4 py-2">Name</th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    {#each resumes as resume, index}
      <tr>
        <td class="border px-4 py-2">{index + 1}</td>
        <td class="border px-4 py-2">{resume.name}</td>
        <td class="border px-4 py-2">
          <Link to={`/resumes/${resume.id}`}>
            <button class="bg-blue-500 px-2 py-2 text-white">View Resume</button>
          </Link>
        </td>
      </tr>
    {/each}
  </tbody>
</table>
