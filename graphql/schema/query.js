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

module.exports = Query;