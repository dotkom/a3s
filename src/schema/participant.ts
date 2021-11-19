import { objectType } from "nexus";

export const Participant = objectType({
  name: "Participant",
  definition(t) {
    t.nonNull.id("id");
    t.nonNull.string("name");
  },
});
