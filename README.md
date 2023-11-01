<div align="center">
<pre>
███████╗     ██████╗  ██████╗ ███╗   ███╗███╗   ███╗███████╗██████╗  ██████╗███████╗
██╔════╝    ██╔════╝ ██╔═══██╗████╗ ████║████╗ ████║██╔════╝██╔══██╗██╔════╝██╔════╝
█████╗█████╗██║  ███╗██║   ██║██╔████╔██║██╔████╔██║█████╗  ██████╔╝██║     █████╗  
██╔══╝╚════╝██║   ██║██║   ██║██║╚██╔╝██║██║╚██╔╝██║██╔══╝  ██╔══██╗██║     ██╔══╝  
███████╗    ╚██████╔╝╚██████╔╝██║ ╚═╝ ██║██║ ╚═╝ ██║███████╗██║  ██║╚██████╗███████╗
╚══════╝     ╚═════╝  ╚═════╝ ╚═╝     ╚═╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝╚══════╝
                                                                                    
</pre>
</div>

## Overview

This is personal project that will serve as a e-commerce service that I want to publish in the future as a real-world product :) It uses Go (Golang), and includes a variety of services, each serving a specific purpose, and a front-end application to interact with these services.

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

## Requirements/dependencies
- Docker
- docker-compose (more info: https://docs.docker.com/compose/)
- minikube (more info: https://github.com/kubernetes/minikube)
- kubectl (install instructions: https://kubernetes.io/docs/tasks/tools/)

## Virtual hosts configuration
As the project emulates the front and backend, it's necessary to virtualize one or both of the hosts to run locally. In the `project/ingress.yml` configuration file I've defined the domains: `front-end.info` and `broker-service.info` you can change it as you like. Before starting the application, you have to add your hosts in your configuration hosts files. 
On MacOS the file is in `/etc/hosts` and you can edit it including the following line:
```bash
127.0.0.1 front-end.info broker-service.info
```
For other OS, please follow this tutorial: https://www.manageengine.com/network-monitoring/how-to/how-to-add-static-entry.html

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/jhonipereira/eGommerce.git
    ```

2. If you use VS Code, open the file **workspace.code-workspace** directly from the file explorer to open the structure with workspaces. Follow the setup instructions in the `project` directory to configure Docker, Kubernetes, and other dependencies.

3. In the `project` folder, start the PostgreSQL database:

    ```bash
    docker-compose -f postgres.yml up -d
    ```

4. Create one or more instance with minikube:

```bash
minikube start --nodes=N
```
Where `N` is the number of nodes to start. The default value is `1`

5. Install all K8s: this will create all the resources in a cluster

```bash
kubectl apply -f k8s
```

6. Make a tunnel for local access: this will create a route to services deployed with LoadBalancer and set the Ingress to the ClusterIP. This command will be active until you cancel it with `CTRL+C`

```bash
minikube tunnel
```

## Contributing

Contributions to the project are welcome. Please follow the [contribution guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

For any inquiries or issues, please open an issue on the [GitHub repository](https://github.com/jhonipereira/eGommerce/issues).
