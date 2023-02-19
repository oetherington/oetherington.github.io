# Introduction

I work on [ForumMagnum](https://github.com/ForumMagnum/ForumMagnum) which is
the open-source project powering the [EA Forum](https://forum.effectivealtruism.org/)
and [LessWrong](https://www.lesswrong.com/).

 - Background
   - LessWrong 1 was a Reddit fork
   - LessWrong 2 is build on Vulcan
   - Migrated from Javascript to Typescript a while ago
   - Static types are an ongoing project
 - Motivation for the migration
   - Performance
   - Maximum number of indexes
   - Developer ergonomics
   - JOIN

# Schemas

Mongo collections are modelled in the codebase as instances of the
`MongoCollection` class, each of which includes a schema and definitions of
required indexes, as well as other information not needed for our purposes such
as permissions and form layout information.

Each field in the schema has a type. Sometimes these map nicely onto basic
Postgres types like `TEXT` for strings or `DOUBLE PRECISION` for numbers, but
they can also be arrays or arbitrary objects (requiring the Postgres `JSONB`
type) which are both considerably less idiomatic in Postgres, even if they are
supported.

One early idea was that we could just store all the data in a single JSONB
column, but initial testing showed the performance of this approach to be
untenable. It's also a lot less aesthetically pleasing, and whilst we said that
we'd eventually migrate to proper statically-typed tables, it's arguable whether
or not we'd have ever found time for these migrations. The idea that our tables
should be statically-typed from the outset then became a fundamental design
decision.

 - We want to avoid catch-all JSON blobs
 - Reddit style tables?
 - Nope - fully type safe (matches our move to Typescript)
 - legacyData
 - Schema already exists (and caught some preexisting bugs)
 - MongoCollection vs PgCollection

# Indexes

Our index definitions look something like
```
ensureIndex(ReviewVotes, {year: 1, userId: 1, dummy: 1});
```
which will create an index on the `ReviewVotes` collection on the fields
`year`, `userId` and `dummy`. The can also take an `options` object which can
specify a name for the index, mark it as being unique, and/or provide a Mongo
expression to turn it into a partial index.

It turns our this is relatively simple to convert to Postgres. Our stategy is
to use btree indexes by default, but to switch to GIN if any of the fields are
JSON or an array.

One gotcha to be aware of when using GIN indexes for searching inside array
fields is that these indexes cannot be used by `ANY`, and you must use the `@>`
operator instead. If you don't do this correctly it can cause some terrible
performance regressions (guess how we found this out...).

# The Query Builder

 - How we build queries
 - Reflection: should have been AST based instead of concatenative

# Aggregation Pipelines

 - Mongo pipelines can often be modelled as nested simple queries
 - Caveat: ForumMagnum doesn't to anything super complicated here
 - For more complex pipelines we still just rewrite in raw SQL

# Migrations

 - Nothing too fancy here
 - Distinguish between migrations and scripts

# Testing Infrastructure

 - ???

# Zero-Downtime

 - Switching collection
 - Data cleanup
 - Multiple rounds of copying
 - Done in a transaction

# The Switch-Over Process

 - Very smooth
 - Client IDs

# Conclusion
