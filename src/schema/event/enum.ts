import { enumType } from "nexus";

export const EventType = enumType({
  name: "EventType",
  members: ["SOCIAL", "COMPANY_PRESENTATION", "COMPANY_COURSE", "COURSE"],
});
