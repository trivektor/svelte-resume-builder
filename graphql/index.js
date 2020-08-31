const express = require('express');
const {ApolloServer} = require('apollo-server-express');
const {MongoClient} = require('mongodb');
const cors = require('cors');

const client = new MongoClient(process.env.CHATTYY_DB_CONNECTION_STRING, {
  useUnifiedTopology: true,
});

client.connect().then(() => {
  const typeDefs = require('./schema');
  const resolvers = require('./resolvers');
  const server = new ApolloServer({
    typeDefs,
    resolvers,
    context: () => ({
      db: client.db('svelte_resumes_builder'),
    })
  });
  const app = express();

  app.use(cors());

  server.applyMiddleware({app});

  app.listen({port: 4000}, () => {
    console.log(`🚀 Server ready at http://localhost:4000${server.graphqlPath}`);
  });
}).catch((err) => {
  console.error(err);
});
