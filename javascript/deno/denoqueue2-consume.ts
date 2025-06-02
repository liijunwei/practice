// deno --allow-all --unstable-kv denoqueue2-consume.ts
// https://docs.deno.com/deploy/kv/manual/queue_overview/
import { format } from "https://deno.land/std@0.91.0/datetime/mod.ts";

function now(): string {
  return format(new Date(), "yyyy-MM-dd HH:mm:ss")
}

interface Notification {
  time: number;
  forUser: string;
  body: string;
}

const kv = await Deno.openKv("./tmp/denoqueue.db");

// why the long latency?
kv.listenQueue(async (msg: Notification) => {
  await console.log(`${now()} Message received:`, msg);
});
