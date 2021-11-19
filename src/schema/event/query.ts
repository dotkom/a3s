import { list, nonNull, queryField } from "nexus";

export const EventsQuery = queryField("events", {
  type: nonNull(list(nonNull("Event"))),
  resolve: async (_, __, ctx) => {
    return [];
  },
});
