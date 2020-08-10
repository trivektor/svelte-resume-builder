<script>
  import {Link} from 'svelte-routing';
  import {paginate, LightPaginationNav} from 'svelte-paginate'
  import {onMount} from 'svelte';

  let resumes = [];
  let page = 1;
  let pageSize = 5;
  let totalItems = 0;

  async function fetchResumes(params) {
    const response = await window.fetch('http://localhost:8080/resumes?' + new URLSearchParams(params), {
      header: {'Accept': 'application/json'},
    });
    const json = await response.json();

    resumes = json.resumes || [];
    totalItems = json.pagination.total;
  }

  onMount(async () => {
    await fetchResumes({page, pageSize});
  });

  async function setPage(event) {
    page = event.detail.page;

    await fetchResumes({page, pageSize});
  }
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

<div class="mt-4">
  <LightPaginationNav
    totalItems={totalItems}
    pageSize={pageSize}
    currentPage={page}
    limit="{1}"
    showStepOptions="{true}"
    on:setPage={setPage}
  />
</div>
