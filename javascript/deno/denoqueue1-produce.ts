// deno --allow-all --unstable-kv denoqueue1-produce.ts
// https://docs.deno.com/deploy/kv/manual/queue_overview/
import { format } from "https://deno.land/std@0.91.0/datetime/mod.ts";

function now(): string {
  return format(new Date(), "yyyy-MM-dd HH:mm:ss")
}

interface Notification {
  time: string;
  forUser: string;
  body: string;
}

const kv = await Deno.openKv("./tmp/denoqueue.db");

const message: Notification = {
  time: now(),
  forUser: "alovelace",
  body: "You've got mail!",
};

await kv.enqueue(message);
console.log(now(), "Message enqueued:", message);
