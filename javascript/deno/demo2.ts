// deno repl --unstable-kv
// deno --unstable-kv demo2.ts
const kv = await Deno.openKv();
const prefs = {
  username: "ada",
  theme: "dark",
  language: "en-US",
};

await kv.set(["preferences", "ada"], prefs);
const entry = await kv.get(["preferences", "ada"]);
console.log(entry.key);
console.log(entry.value);
console.log(entry.versionstamp);
