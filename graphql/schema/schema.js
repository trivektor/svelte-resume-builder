const {gql} = require('apollo-server');

const Resume = require('./resume');

const Query = `
  type Query {
    resume(id: ID!): Resume
    resumes: [Resume]!
  }
`;

const typeDefs = gql`
  ${Resume}
  ${Query}
`;

module.exports = typeDefs;
