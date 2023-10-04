### ****Steps for Starting :****
1. Migrating the model to DB, Command (go run cmd/migrate/main.go)
2. Starting the Server, Command (go run cmd/local/main.go)

### **Folder Descriptions :**

**cmd** -> Contains local, db migrations files which has to started before testing.

**conf** -> This folder contains the config of db.

**internal/blog** -> This contains APIs functions and models.

**internal/config** -> This folder contains functions related to reading of config.

**pkg/logger** -> This folder contains logger related files.

**pkg/sql** -> This folder contains file related to DB connection setup.
