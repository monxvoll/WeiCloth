# Logical / Structural View

> **Nota:** These diagrams are for illustrative purposes only and are subject to change.

Concerns itself with the functionality that is provided by the system and how the code is designed to provide such functionality

> In this case, we only create domain models, since it is designed for object-oriented decomposition and our system is service-oriented / microservice-based

This object diagram is about the entities of domain. Esential entities for the data flow in this system. 

### 1. Domain model (pure 4+1)

```mermaid
classDiagram
    direction TB

    class User {
        +UUID id
        +String SubKeyCloak
        +String FirstName
        +String LastName
        +String NickName
        +String email
        +DateTime DateBirth
        +String Gender
        +DateTime createdAt
        +DateTime updatedAt
    }

    class Session {
        +UUID id
        +UUID userId
        +String jwtToken
        +DateTime issuedAt
        +DateTime expiresAt
        +Boolean revoked
    }

    class GarmentUpload {
        +UUID id
        +UUID userId
        +String imageS3Path
        +String status
        +DateTime uploadedAt
        +DateTime processedAt
    }

    class ClassificationResult {
        +UUID id
        +UUID uploadId
        +String category
        +String subcategory
        +Float confidenceScore
        +DateTime classifiedAt
    }

    class Garment {
        +UUID id
        +UUID userId
        +UUID classificationId
        +String name
        +String category
        +String color
        +String season
        +DateTime registeredAt
    }

    class Recommendation {
        +UUID id
        +UUID userId
        +List~UUID~ garmentIds
        +String context
        +DateTime generatedAt
    }

    class AuditLog {
        +UUID id
        +UUID userId
        +String action
        +String service
        +JSON payload
        +DateTime timestamp
    }

    User "1" --> "0..*" Session : has
    User "1" --> "0..*" GarmentUpload : owns
    User "1" --> "0..*" Garment : wardrobe
    User "1" --> "0..*" Recommendation : receives
    GarmentUpload "1" --> "1" ClassificationResult : produces
    ClassificationResult "1" --> "1" Garment : registers
    Recommendation "1" --> "0..*" Garment : suggests
    User "1" --> "0..*" AuditLog : generates
```

---

This object diagram refers to the domain model (DTO mapping). It helps clarify what the client sends and what is stored, as well as identify the boundaries of transformation between layers.

### 2. Domain model + DTO mapping

```mermaid
classDiagram
    direction TB

    %% ── Domain Entities ──
    class User {
        +UUID id
        +String email
        +String keycloakUID
        +DateTime createdAt
        +DateTime updatedAt
    }

    class Session {
        +UUID id
        +UUID userId
        +String jwtToken
        +DateTime issuedAt
        +DateTime expiresAt
        +Boolean revoked
    }

    class GarmentUpload {
        +UUID id
        +UUID userId
        +String imageS3Path
        +String status
        +DateTime uploadedAt
        +DateTime processedAt
    }

    class ClassificationResult {
        +UUID id
        +UUID uploadId
        +String category
        +String subcategory
        +Float confidenceScore
        +DateTime classifiedAt
    }

    class Garment {
        +UUID id
        +UUID userId
        +UUID classificationId
        +String name
        +String category
        +String color
        +String season
        +DateTime registeredAt
    }

    class Recommendation {
        +UUID id
        +UUID userId
        +List~UUID~ garmentIds
        +String context
        +DateTime generatedAt
    }

    class AuditLog {
        +UUID id
        +UUID userId
        +String action
        +String service
        +JSON payload
        +DateTime timestamp
    }

    %% ── DTOs ──
    class RegisterUserInput {
        +String email
        +String password
        +String fullName
    }

    class LoginRequestDTO {
        +String email
        +String password
    }

    class LoginResponseDTO {
        +String jwtToken
        +DateTime expiresAt
        +String userId
    }

    class UploadRequestDTO {
        +File image
        +String metadata
        +String jwt
    }

    class UploadResponseDTO {
        +String uploadId
        +String status
        +String message
    }

    class ClassificationResultDTO {
        +String category
        +String subcategory
        +Float confidenceScore
    }

    class RecommendationRequestDTO {
        +String userId
        +String context
        +String jwt
    }

    class RecommendationResponseDTO {
        +List~GarmentDTO~ garments
        +String context
        +DateTime generatedAt
    }

    class GarmentDTO {
        +String id
        +String name
        +String category
        +String color
        +String season
    }

    %% ── Domain relationships ──
    User "1" --> "0..*" Session : has
    User "1" --> "0..*" GarmentUpload : owns
    User "1" --> "0..*" Garment : wardrobe
    User "1" --> "0..*" Recommendation : receives
    GarmentUpload "1" --> "1" ClassificationResult : produces
    ClassificationResult "1" --> "1" Garment : registers
    Recommendation "1" --> "0..*" Garment : suggests
    User "1" --> "0..*" AuditLog : generates

    %% ── DTO mappings ──
    RegisterUserInput ..> User : maps to
    LoginRequestDTO ..> Session : maps to
    LoginResponseDTO ..> Session : built from
    UploadRequestDTO ..> GarmentUpload : maps to
    UploadResponseDTO ..> GarmentUpload : built from
    ClassificationResultDTO ..> ClassificationResult : maps to
    RecommendationRequestDTO ..> Recommendation : maps to
    RecommendationResponseDTO ..> Recommendation : built from
    GarmentDTO ..> Garment : built from
```