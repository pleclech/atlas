function "pgp_pub_decrypt" {
  schema     = schema.foo
  args       = "bytea, bytea"
  returns    = "text"
  definition = "CREATE OR REPLACE FUNCTION foo.pgp_pub_decrypt(bytea, bytea)\n RETURNS text\n LANGUAGE c\n IMMUTABLE PARALLEL SAFE STRICT\nAS '$libdir/pgcrypto', $function$pgp_pub_decrypt_text$function$\n"
}
schema "foo" {
}
