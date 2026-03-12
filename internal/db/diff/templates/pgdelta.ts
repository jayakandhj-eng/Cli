import { createPlan } from "npm:@Indobase/pg-delta@1.0.0-alpha.3";
import { Indobase } from "npm:@Indobase/pg-delta@1.0.0-alpha.3/integrations/Indobase";

const source = Deno.env.get("SOURCE");
const target = Deno.env.get("TARGET");

const opts = { ...Indobase, role: "postgres" };
const includedSchemas = Deno.env.get("INCLUDED_SCHEMAS");
if (includedSchemas) {
  opts.filter = { schema: includedSchemas.split(",") };
}

const result = await createPlan(source, target, opts);
const statements = result?.plan.statements ?? [];
for (const sql of statements) {
  console.log(`${sql};`);
}

