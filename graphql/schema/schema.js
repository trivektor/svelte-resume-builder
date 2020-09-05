const {gql} = require('apollo-server');
const Query = require('./query');
const Mutation = require('./mutation');
const Resume = require('./resume');

const typeDefs = gql`
  ${Resume}
  ${Query}
  ${Mutation}
`;

module.exports = typeDefs;
