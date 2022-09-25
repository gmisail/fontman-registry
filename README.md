# fontman-registry

TBD.

## Creating Schemas 

To manage database interactions, the registry uses the [ent framework](https://entgo.io). To generate
new schemas, run `go run -mod=mod entgo.io/ent/cmd/ent init <SchemaName>`. This will generate a new
schema file in the `ent/schema/` folder. 

Once you are done making your changes, you will need to run the code generation step. To do this, simply
run `task generate`.
