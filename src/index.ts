import { PrismaClient } from "@prisma/client";
import { ApolloServer } from "apollo-server-lambda";
import { baseSchema } from "./schema";
import * as nexus from "./__generated__/nexus-typegen";

const isDev = process.env.NODE_ENV != "production";
const server = new ApolloServer({
  schema: baseSchema,
  introspection: isDev,
  context: async (): Promise<nexus.NexusGenTypes["context"]> => {
    return { prisma: new PrismaClient() };
  },
});

export const handler = server.createHandler();
