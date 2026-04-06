# Deployment / Physical view

Illustrates the physical organization of the aplication, its about "what code runs in what hardware"

> In this case, our entire system will be hosted on __AWS__ services

```mermaid
flowchart TD

    subgraph Device ["User Device"]
        Angular["Angular App
        SPA"]
    end

    subgraph AWS ["AWS Cloud"]

        subgraph APIGW ["AWS API Gateway"]
            GW["API Gateway
            Routing + Rate limiting"]
        end

        subgraph ECS1 ["ECS Container"]
            GO["GO Backend
            REST API"]
        end

        subgraph ECS2 ["ECS Container"]
            KC["Keycloak
            Auth Server"]
        end

        subgraph ECS3 ["ECS Container (GPU)"]
            ML["Python ML Service
            Docker"]
        end

        subgraph MSK ["Amazon MSK"]
            Kafka["Kafka Broker
            Event Broker"]
        end

        subgraph RDS ["AWS Aurora"]
            Aurora[("Aurora PostgreSQL
            Metadata + Results")]
        end

        subgraph Storage ["AWS S3"]
            S3[("Image Storage
            Bucket")]
        end

    end

    Angular -- "HTTPS REST / JWT" --> GW
    GW -- "Rate limit check" --> GW
    GW -- "Route / HTTPS" --> GO
    GO -- "HTTPS / JWT introspect" --> KC
    GO -- "TCP / publish event" --> Kafka
    Kafka -- "TCP / consume event" --> GO
    GO -- "HTTPS / image + metadata" --> ML
    ML -- "HTTPS / classification result" --> GO
    GO -- "TCP 5432 / SQL" --> Aurora
    GO -- "HTTPS / image upload" --> S3
```