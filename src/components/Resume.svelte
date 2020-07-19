<script>
  export let id;
  import {onMount} from 'svelte';

  onMount(async () => {
    const result = await fetch(`http://localhost:8080/resumes/${id}`, {
      header: {'Accept': 'application/json'},
    });
    const json = await result.json();

    console.log(json);
  });
</script>
