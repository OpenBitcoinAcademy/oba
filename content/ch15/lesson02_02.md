## The Type System

Every Miniscript fragment has a type that describes its behavior on the Script stack. The four base types are B, V, K, and W.

A B-type (base) expression pushes a single nonzero value on success and zero on failure. A V-type (verify) expression either succeeds with no stack output or aborts the script. A K-type (key) expression pushes a public key on success. A W-type (wrapped) expression behaves like B but consumes the top stack element first, used for combining sub-expressions.

The compiler checks types at every composition point. An `and_v` fragment requires its first argument to be V-type and its second to be B-type. If you pass arguments with wrong types, the compiler rejects the expression. This catches errors that, in raw Script, would produce a script that "works" but enforces the wrong conditions.

The type system also tracks properties like dissipatibility (can a dissatisfied branch leave the stack clean?) and non-malleability (is there exactly one valid witness for each spending path?). These properties propagate through composition. A Miniscript expression either has them or does not, and the compiler can report which.
