#!/usr/bin/env node
/**
 * curl-based integration tests for the Kindle-Go bookstore API.
 * Usage: node scripts/curl-tests.js [baseURL]
 * Default baseURL: http://localhost:9010 — run the app first.
 */
const { execFileSync } = require("child_process");

const BASE = process.argv[2] || "http://localhost:9010";
let passed = 0;
let failed = 0;
let cookie = null; // stores created book id via /book/:id path

function curl({ method = "GET", path, body = null }) {
  const args = ["-s", "-X", method, `${BASE}${path}`, "-w", "\n%{http_code}"];
  if (body) {
    args.push("-H", "Content-Type: application/json", "-d", JSON.stringify(body));
  }
  const raw = execFileSync("curl", args).toString();
  const idx = raw.lastIndexOf("\n");
  const status = parseInt(raw.slice(idx + 1));
  let json = null;
  try {
    json = JSON.parse(raw.slice(0, idx));
  } catch (_) {
    /* non-JSON body */
  }
  return { status, json };
}

function test(name, { method = "GET", path, body, expectStatus = 200, check = () => true }) {
  try {
    const res = curl({ method, path, body });
    const statusOk = res.status === expectStatus;
    const bodyOk = check(res);
    if (statusOk && bodyOk) {
      console.log(`✅ ${name} (${method} ${path} → ${res.status})`);
      passed++;
    } else {
      console.log(`❌ ${name} (${method} ${path} → ${res.status}, want ${expectStatus}${bodyOk ? "" : ", body check failed"} ${res.json ? JSON.stringify(res.json) : ""})`);
      failed++;
    }
    return res;
  } catch (err) {
    console.log(`❌ ${name} (${method} ${path} → request failed: ${err.message})`);
    failed++;
    return { status: 0, json: null };
  }
}

console.log(`Running curl tests against ${BASE}\n`);

// 1. Create a book
const created = test("Create book returns 201 with generated id", {
  method: "POST",
  path: "/book/",
  body: { isbn: `CURL-${Date.now()}`, title: "Curl Test Book" },
  expectStatus: 201,
  check: (r) => r.json && Number.isInteger(r.json.id),
});
if (created.json && created.json.id) cookie = created.json.id;
const learnPort = cookie ? `/book/${cookie}` : "/book/0";

// 2. Get all books
test("List all books returns array", {
  path: "/book/",
  check: (r) => Array.isArray(r.json),
});

// 3. Get by id
test("Get book by id", {
  path: learnPort,
  check: (r) => r.json && r.json.id === cookie,
});

// 4. Update
test("Update book by id", {
  method: "PUT",
  path: learnPort,
  body: { isbn: created?.json?.isbn, title: "Curl Test Updated" },
  check: (r) => r.json && r.json.title === "Curl Test Updated",
});

// 5. Delete
test("Delete book by id", {
  method: "DELETE",
  path: learnPort,
  check: (r) => r.json && r.json.id === cookie,
});

// 6. Deleted book is gone
test("Deleted book returns 404", {
  path: learnPort,
  expectStatus: 404,
});

console.log(`\n${passed} passed, ${failed} failed`);
process.exit(failed ? 1 : 0);
