graph TB
    Client[Client/Browser]
    
    subgraph "HTTP Layer"
        Router[Gin Router]
        ErrorMW[Error Middleware]
        ProductHandler[Product Handler]
        HealthHandler[Health Handler]
    end
    
    subgraph "Business Layer"
        ListUC[List Products UseCase]
        GetUC[Get Product UseCase]
    end
    
    subgraph "Data Layer"
        RepoInterface[Repository Interface]
        ProductRepo[Product Repository]
        SQLite[(SQLite In-Memory)]
    end
    
    Client -->|HTTP Request| Router
    Router --> ErrorMW
    ErrorMW --> ProductHandler
    ErrorMW --> HealthHandler
    
    ProductHandler -->|List| ListUC
    ProductHandler -->|Get by ID| GetUC
    
    ListUC --> RepoInterface
    GetUC --> RepoInterface
    RepoInterface -.->|implements| ProductRepo
    
    ProductRepo -->|SQL Query| SQLite
    
    style Client fill:#e1f5ff
    style Router fill:#fff4e1
    style ErrorMW fill:#fff4e1
    style ListUC fill:#e8f5e9
    style GetUC fill:#e8f5e9
    style ProductRepo fill:#fce4ec
    style SQLite fill:#f3e5f5