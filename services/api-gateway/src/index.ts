import { GraphQLFileLoader } from "@graphql-tools/graphql-file-loader";
import { loadSchemaSync } from "@graphql-tools/load";
import { join } from "path";
import { ApolloServer } from "apollo-server";
import { addResolversToSchema } from "@graphql-tools/schema";
import resolvers from "./graphql/resolvers";

const schema = loadSchemaSync(join(__dirname, "../schema.graphql"), {
  loaders: [new GraphQLFileLoader()],
});
const schemaWithResolvers = addResolversToSchema({ schema, resolvers });

const server = new ApolloServer({
  schema: schemaWithResolvers,
  csrfPrevention: true, // see below for more about this
  cache: "bounded",
  cors: {
    origin: ["http://localhost:3000", "https://studio.apollographql.com"],
  },
});
server.listen().then((url) => console.log(`🚀  Server ready at ${url}`));
