import { objectType } from "nexus";

export const Organizer = objectType({
  name: "Organizer",
  definition(t) {
    t.nonNull.id("id");
    t.nonNull.string("name");
    t.nonNull.string("description");
    t.nonNull.string("contactMail");
  },
});
