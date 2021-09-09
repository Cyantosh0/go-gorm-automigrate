## GORM AutoMigrate
Doc Link: https://gorm.io/docs/migration.html

**This is a repository on GO project to work with database migration using AutoMigrate**

#### Pros:
- No need to have knowledge or write raw SQL queries
- Models get migrated with:
> db.Automigrate(Model{})

This looks simpler and easier to do.
- New column/field gets automatically migrated if added to the model.

#### Cons:
- The major disadvantage of using GORM AutoMigrate is that it does not support version control. There is no tracking of changes in the database. Also, to know the status of the database, we have to manually check all the tables in the database.
- When the model is migrated and then the field’s properties are changed in the model then nothing gets updated into the database. We can change the type of column using ModifyColumn(); However, to add further properties like not null, default value, etc. Firstly, we have to run Drop command then add the column with the necessary properties and run AutoMigrate().
- The same query gets executed every time when running the application.
