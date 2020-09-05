const resumes = require('./resumes');
const resume = require('./resume');
const createResume = require('./create-resume');

const resolvers = {
  Query: {
    resumes,
    resume,
  },
  Mutation: {
    createResume,
  },
};

module.exports = resolvers;
