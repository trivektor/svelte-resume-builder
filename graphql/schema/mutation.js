const Mutation = `
    input ResumeInput {
        name: String!
        description: String
    }
    type Mutation {
        createResume(input: ResumeInput): Resume
    }
`;

module.exports = Mutation;