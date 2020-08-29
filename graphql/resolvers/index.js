const resumes = require('./resumes');
const resume = require('./resume');

const resolvers = {
  Query: {
    resumes,
    resume,
  },
};

module.exports = resolvers;
