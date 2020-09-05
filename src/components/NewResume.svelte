<!-- <script>
  import {navigate} from 'svelte-routing';
  import {resumeStore} from '../stores';

  import ResumeForm from './ResumeForm.svelte';

  let resume = {};

  resumeStore.subscribe((data) => {
    resume = data;
  });

  async function onSubmit() {
    try {
      const result = await window.fetch('http://localhost:8080/resumes', {
        method: 'POST',
        body: JSON.stringify(resume),
        header: {'Content-Type': 'application/json', 'Accept': 'application/json'},
      });
      const json = await result.json();

      navigate(`/resumes/${json.id}`);
    } catch (err) {
      console.error(err);
    }
  }
</script> -->

<script>
  import {getClient, mutate} from 'svelte-apollo';
  import {gql} from 'apollo-boost';
  import {navigate} from 'svelte-routing';
  import {resumeStore} from '../stores';

  import ResumeForm from './ResumeForm.svelte';

  let resume = {};

  resumeStore.subscribe((data) => {
    resume = data;
  });

  const CREATE_RESUME_MUTATION = gql`
    mutation CreateResumeMutation(
      $name: String!
      $description: String
    ) {
      createResume(input: {name: $name, description: $description}) {
        _id
        name
        description
      }
    }
  `;

  const client = getClient();
  
  async function onSubmit() {
    try {
      const {data} = await mutate(client, {
        mutation: CREATE_RESUME_MUTATION,
        variables: {...resume},
      });

      navigate(`/resumes/${data.createResume._id}`);
    } catch (err) {
      console.error(err);
    }
  };
</script>

<section class="bg-white rounded py-5 px-5 block border-2 border-gray">
  <form on:submit|preventDefault={onSubmit}>
    <ResumeForm />
  </form>
</section>
