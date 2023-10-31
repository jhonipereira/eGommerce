# eGommerce

## Overview

This is personal project that will serve as a ecommerce service that I want to publish in the future as a real-world product :) It uses Go (Golang), and includes a variety of services, each serving a specific purpose, and a front-end application to interact with these services.

## Project Structure

- `front-end`: The front-end application for interacting with the microservices.
- `broker-service`: Acts as an API Gateway, routing requests to the appropriate services.
- `project`: Contains configuration files and scripts for Docker, Makefile, Caddy, and Kubernetes (k8s).
- `authentication-service`: Manages user authentication and authorization.
- `logger-service`: Handles logging of application data and events.
- `mail-service`: Manages sending emails and notifications.
- `listener-service`: Listens to events and messages from other services, performing actions accordingly.
- `product-service`: Manages product-related operations and data storage.

## Technologies Used

- **Go (Golang)**: The primary programming language used to write the services.
- **Docker**: For containerizing the application and services.
- **Kubernetes (k8s)**: For orchestrating and managing the containers.
- **RabbitMQ**: For message queuing and event-driven architecture.
- **MongoDB**: Used by the logger service for storing logs.
- **PostgreSQL**: For database operations in the authentication and product services.
- **Caddy**: As a web server and for handling HTTPS.

## Getting Started

1. Clone the repository:

    ```bash
    git clone [<repository-url>](https://github.com/jhonipereira/eGommerce)
    ```

2. If you use VS Code, open the file **workspace.code-workspace** directly from the file explorer to open the structure with workspaces. Follow the setup instructions in the `project` directory to configure Docker, Kubernetes, and other dependencies.

3. Start the services:

    ```bash
    make start
    ```

4. Access the front-end application at [http://localhost](http://localhost) (or the configured URL). I configured the domain **front-end.info** in my etc/hosts.

## Contributing

Contributions to the project are welcome. Please follow the [contribution guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

For any inquiries or issues, please open an issue on the [GitHub repository](https://github.com/jhonipereira/eGommerce/issues).
