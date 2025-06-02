import { format } from "https://deno.land/std@0.91.0/datetime/mod.ts";

const db = await Deno.openKv();

function now(): string {
  return format(new Date(), "yyyy-MM-dd HH:mm:ss")
}

db.listenQueue(async (msg) => {
  await console.log(now(), "got message:", msg)
});
await db.enqueue({ channel: "C123456", text: "Slack message" }, {
  delay: 100,
});
