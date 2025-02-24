# Generate Resource Components

This document provides commands to generate various components like Resource Controller, Service, Repository, and Model in your project.

## Generate Resource

- `make admincmd.gen name="todo"`: Generate all components in a resource.
- `make admincmd.gen name="todo" option=MD`: Generate only model in a resource.
- `make admincmd.gen name="todo" option=P`: Generate only permission in a resource.
- `make admincmd.gen name="todo" option=M`: Generate only module in a resource.
- `make admincmd.gen name="todo" option=C`: Generate only controller in a resource.
- `make admincmd.gen name="todo" option=-S`: Generate only service in a resource.
- `make admincmd.gen name="todo" option=-S,-C,-S,-D`: Generate module, controller, service and dto in a resource.
