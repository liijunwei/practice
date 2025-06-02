// deno add jsr:@std/assert
// deno test --unstable-kv ./demo1.ts
import { assertEquals, assert } from "@std/assert";

async function setDisplayName(
  kv: Deno.Kv,
  username: string,
  displayname: string,
) {
  await kv.set(["preferences", username, "displayname"], displayname);
}

async function getDisplayName(
  kv: Deno.Kv,
  username: string,
): Promise<string | null> {
  return (await kv.get(["preferences", username, "displayname"])).value as string;
}

Deno.test("Preferences", async (t) => {
  const kv = await Deno.openKv(":memory:");

  await t.step("can set displayname", async () => {
    const displayName1 = await getDisplayName(kv, "example");
    assertEquals(displayName1, null);

    await setDisplayName(kv, "example", "Exemplary User");

    const displayName2 = await getDisplayName(kv, "example");
    assertEquals(displayName2, "Exemplary User");
    assert(true)
  });
});

assert(true)
console.log("done")
