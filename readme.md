
# GoLang GraphQL API - Boilerplate

This repository created for GraphQL API with GoLang. Used **gqlgen** package. 

It's a schema-based project, if you want add some new models, firstly you should add new .graphqls files under to schemas folder.

And, you must generate your folders with this command;

```
go run github.com/99designs/gqlgen generate

```

After this, gpqlgen will create **resolver** and **model** file according to your schemas. You can make changes on resolver functions.

Now, you are ready for **run**.