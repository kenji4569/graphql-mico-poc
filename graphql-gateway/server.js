const express = require('express');
const bodyParser = require('body-parser');
const { graphqlExpress } = require('apollo-server-express');
const { makeExecutableSchema } = require('graphql-tools');

const { HttpLink } = require('apollo-link-http');
const fetch = require('node-fetch');
const {
  introspectSchema,
  mergeSchemas,
  makeRemoteExecutableSchema,
} = require('graphql-tools');

const app = express();

new Promise(async (resolve, reject) => {
  const link = new HttpLink({ uri: process.env.SERVICE1 + '/graphql', fetch });
  let schema;
  try {
    schema = await introspectSchema(link);
  } catch (e) {
    reject(e)
    return
  }

  const executableSchema = makeRemoteExecutableSchema({
    schema,
    link,
  });

  resolve(executableSchema);

}).then((executableSchema) => {
  app.use('/graphql', bodyParser.json(), graphqlExpress({ schema: executableSchema }));

  app.get('/', async (req, res) => {
    result = await fetch(process.env.SERVICE1)
    body = await result.text()
    res.send(body);
  });

  app.listen(3000, () => {
    console.log('Now server is running on port 3000');
  });
}).catch((e) => { console.error(e) });

