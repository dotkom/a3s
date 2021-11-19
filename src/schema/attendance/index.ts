import { objectType } from "nexus";

export const AttendanceInfo = objectType({
  name: "AttendanceInfo",
  definition(t) {
    t.nonNull.datetime("registrationStartDate");
    t.nonNull.datetime("registrationEndDate");
    t.nonNull.datetime("unattendDeadline");
    t.list.nonNull.string("allowedGroups");
    t.nonNull.boolean("canCauseMarks");
    t.nonNull.boolean("delayMarkedUsers");
  },
});
