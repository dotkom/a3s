import { objectType } from "nexus";

export const Event = objectType({
  name: "Event",
  definition(t) {
    t.nonNull.string("title");
    t.nonNull.datetime("startDate");
    t.nonNull.datetime("endDate");
    t.nonNull.string("location");
    t.nonNull.string("image");
    t.nonNull.string("description");
    t.nonNull.field("eventType", { type: "EventType" });
    t.nonNull.field("organizer", { type: "Organizer" });
    t.nonNull.list.nonNull.string("tags");
    t.field("attendance", { type: "AttendanceInfo" });
    t.field("company", { type: "Company" });
  },
});
