import { objectType } from "nexus";

export const Company = objectType({
  name: "Company",
  definition(t) {
    t.nonNull.id("id");
    t.nonNull.string("name");
    t.nonNull.string("link");
  },
});
