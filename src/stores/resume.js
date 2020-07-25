import {writable} from 'svelte/store';

const resume = writable({
  name: '',
  description: '',
  sections: [],
});

export {resume as default};
