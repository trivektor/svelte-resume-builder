const {gql} = require('apollo-server');

const Resume = require('./resume');

const Query = `
  type ResumeEdge {
    cursor: String!
    node: Resume
  }
  type PageInfo {
    endCursor: String
    hasNextPage: Boolean!
    hasPreviousPage: Boolean!
    startCursor: String
  }
  type ResumeConnection {
    edges: [ResumeEdge]!
    pageInfo: PageInfo!
    totalCount: Int!
  }
  type Query {
    resume(_id: ID!): Resume
    resumes(first: Int, after: String): ResumeConnection!
  }
`;

const typeDefs = gql`
  ${Resume}
  ${Query}
`;

module.exports = typeDefs;
