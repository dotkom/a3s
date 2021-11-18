import { ApolloServer, gql } from "apollo-server-lambda";
import { Context } from "./context";
import { baseSchema } from "./schema";

const isDev = process.env.NODE_ENV != "production";
const server = new ApolloServer({
  schema: baseSchema,
  introspection: isDev,
  context: async (): Promise<Context> => {
    return { userID: "2" };
  },
});

export const handler = server.createHandler();
