const express = require('express');
const bodyParser = require('body-parser');
const { graphqlExpress } = require('apollo-server-express');
const { makeExecutableSchema } = require('graphql-tools');

const typeDefs = `
  type Query { echo(text: String!): Echo }
  type Echo { text: String }
`;

const resolvers = {
  Query: {
    echo: (_, { text }) => { return {text} }
  }
};

const schema = makeExecutableSchema({
  typeDefs,
  resolvers
});

const app = express();

app.use('/graphql', bodyParser.json(), graphqlExpress({ schema }));

app.get('/', (req, res) => {
  res.send('Hello World');
});

app.listen(3000, () => {
  console.log('Now server is running on port 3000');
});