## Development Flow

- Create Model
    - model/domain/category.go

- Create Repository
    - category_repository.go
    - category_repository_impl.go (impl = implementations)

- Create Helper
    - helper/*

- Create Web (value for request and response)
    - model/web/*

- Create Service (bussiness logic)
    - service/category_service.go
    - service/category_service_impl.go (impl = implementations)

- Create App
    - app/*

- Create Exception
    - exception/*
    
- Create Middleware
    - middleware/auth_middleware.go


## What is?
- Context in category_repository.go

## Note
- domain (category domain.Category) just for repository

## Term
- tx = transactional
- impl = implementations