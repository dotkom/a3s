import { makeSchema } from "nexus";
import { join } from "path";
import { GQLDateTime } from "./scalars";
import * as EventSchema from "./event";
import * as AttendanceSchema from "./attendance";
import * as HostSchema from "./hosts";
import * as ParticipantSchema from "./participant";

export const baseSchema = makeSchema({
  types: [
    GQLDateTime,
    ParticipantSchema,
    HostSchema,
    AttendanceSchema,
    EventSchema,
  ], // 1
  outputs: {
    typegen: join(__dirname, "../__generated__", "nexus-typegen.ts"), // 2
    schema: join(__dirname, "../__generated__", "schema.graphql"), // 3
  },
  contextType: {
    module: require.resolve("../context"),
    export: "Context",
  },
  features: {
    abstractTypeStrategies: {
      __typename: true,
    },
  },
});
