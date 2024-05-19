<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">IMAGE_SLIDESHOW</h1>
</p>
<p align="center">
    <em>Picture perfect moments, beautifully displayed.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/commit-activity/m/sharanreddy99/dbms_algos" alt="last-commit">
	<img src="https://img.shields.io/github/created-at/sharanreddy99/dbms_algos" alt="created_at">
   <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/sharanreddy99/dbms_algos">
   <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/sharanreddy99/dbms_algos">
   <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/sharanreddy99/dbms_algos">

</p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<br><!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary><br>

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
</details>
<hr>


##  Overview

The image_slideshow project is a multifaceted application combining Golang, PHP, and React components to create a comprehensive image slideshow feature with user management capabilities. It facilitates user authentication, image upload, deletion, and browsing functionalities, ensuring seamless interaction between the frontend and backend services. With robust API handling, responsive frontend design, and secure backend operations, image_slideshow provides a user-friendly platform for managing and showcasing images effectively.

---

##  Features

|    |   Feature         | Description |
|----|-------------------|---------------------------------------------------------------|
| ⚙️  | **Architecture**  | The project consists of frontend (React), backend (Golang with Gorilla Mux and PHP with Firebase JWT), and Nginx for proxying requests. Docker is used for containerization, ensuring reliable deployment and scalability. MySQL is the database of choice for data storage. |
| 🔩 | **Code Quality**  | The codebase maintains good quality standards with structured code organization, clear naming conventions, and adherence to best practices like component-based architecture in React. The backend and frontend code are well-segmented, promoting maintainability and readability. |
| 📄 | **Documentation** | The repository contains comprehensive documentation across its components, outlining setup instructions, API endpoints, and code structure. Each section has detailed descriptions, aiding developers in understanding and extending the project. |
| 🔌 | **Integrations**  | Key dependencies include jwt-go for Golang JWT authentication, mux for Golang routing, and various frontend libraries like React, Bootstrap, and Axios. Nginx is integrated for request routing, and MySQL for data storage. |
| 🧩 | **Modularity**    | The codebase is modular and allows for easy extension and reuse. Components in React are well-isolated, facilitating component reusability and ensuring a clear separation of concerns between frontend and backend functionalities. |
| 🧪 | **Testing**       | Testing frameworks such as @testing-library/jest-dom for React and unit testing for backend services are employed. This supports the project's reliability and aids in maintaining code integrity during development. |
| ⚡️  | **Performance**   | The project demonstrates good performance by efficiently handling backend requests, optimized data retrieval, and fast rendering of frontend components. The use of Docker containers ensures quick deployment and scalable resource allocation. |
| 🛡️ | **Security**      | Security measures are implemented with JWT token authentication for user sessions, user authorization checks, secure API endpoints, and CORS handling. The backend services follow best practices for data protection and access control. |
| 📦 | **Dependencies**  | Key dependencies include jwt-go, Gorilla Mux, React, Bootstrap, Axios, and MySQL. These libraries enhance functionalities, promote efficient data handling, and streamline frontend development with popular frameworks. |

---

##  Repository Structure

```sh
└── image_slideshow/
    ├── README.md
    ├── backend_golang
    │   ├── .dockerignore
    │   ├── .gitignore
    │   ├── Dockerfile
    │   ├── constants
    │   ├── go.mod
    │   ├── go.sum
    │   ├── main.go
    │   ├── router
    │   ├── start.sh
    │   └── utils
    ├── backend_php
    │   ├── .dockerignore
    │   ├── Dockerfile
    │   ├── api
    │   ├── composer.json
    │   ├── composer.lock
    │   ├── index.php
    │   └── start.sh
    ├── docker-compose.yml
    ├── frontend
    │   ├── .gitignore
    │   ├── Dockerfile
    │   ├── README.md
    │   ├── generate-react-cli.json
    │   ├── package-lock.json
    │   ├── package.json
    │   ├── public
    │   ├── src
    │   └── start.sh
    └── nginx
        ├── Dockerfile
        └── default.conf
```

---

##  Modules

<details closed><summary>.</summary>

| File                                                                                                      | Summary                                                                                                                                                                                                                                                                                      |
| ---                                                                                                       | ---                                                                                                                                                                                                                                                                                          |
| [docker-compose.yml](https://github.com/sharanreddy99/image_slideshow.git/blob/master/docker-compose.yml) | Orchestrates Docker services for backend Golang, PHP, and frontend React with MySQL and PHPMyAdmin. Aligns services, manages dependencies, sets environment variables, and maps volumes. Ensures seamless integration and communication within the image_slideshow repositorys architecture. |

</details>

<details closed><summary>frontend</summary>

| File                                                                                                                         | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---                                                                                                                          | ---                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [package-lock.json](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/package-lock.json)             | This code file, `main.go`, serves as the core backend component for the image slideshow application within the `image_slideshow` repository. It defines the primary functionality and routing logic needed to handle image data requests and responses. By leveraging this file, the backend utilizes Golang to efficiently process and serve image-related content, contributing to the seamless operation of the overall image slideshow feature in the application. |
| [Dockerfile](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/Dockerfile)                           | Enables frontend setup by installing dependencies, setting working directory, exposing port, and initiating the application start script. Complements the architecture of the parent repositorys frontend section.                                                                                                                                                                                                                                                     |
| [start.sh](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/start.sh)                               | Sets environment variable for backend host in React app and initiates development server.                                                                                                                                                                                                                                                                                                                                                                              |
| [package.json](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/package.json)                       | Manages dependencies, scripts, and configurations for a dynamic image slideshow frontend. Enables React development with testing, routing, and UI components like Bootstrap. Integrates Axios for API interactions and Web Vitals for performance monitoring.                                                                                                                                                                                                          |
| [generate-react-cli.json](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/generate-react-cli.json) | CSS, TypeScript preferences, and component structure details. Contributes to the architecture by standardizing component creation for consistent UI development within the parent repository.                                                                                                                                                                                                                                                                          |

</details>

<details closed><summary>frontend.public</summary>

| File                                                                                                            | Summary                                                                                                                                                                                 |
| ---                                                                                                             | ---                                                                                                                                                                                     |
| [manifest.json](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/public/manifest.json) | Sets display mode, icons, theme colors & background. Crucial for PWA functionality and user experience in the frontend of the image_slideshow repository.                               |
| [robots.txt](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/public/robots.txt)       | Defines crawl settings for web robots. Allows universal access and no specified disallow paths. Crucial for search engine indexing in the frontend section of the repository.           |
| [index.html](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/public/index.html)       | Defines the main webpage structure for the React frontend, including metadata and essential script links. Ensures proper rendering and functionality of the web app on various devices. |

</details>

<details closed><summary>frontend.src</summary>

| File                                                                                                           | Summary                                                                                                                                                                                                                                         |
| ---                                                                                                            | ---                                                                                                                                                                                                                                             |
| [index.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/index.js)             | Renders the main React component in the frontend, setting up the initial view for the image slideshow application. Integrates Bootstrap styling and links the App component to the root HTML element.                                           |
| [App.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/App.css)               | Resets default styles. Prevents text selection. Enhances frontend styling consistency within the image_slideshow repository.                                                                                                                    |
| [App.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/App.js)                 | Routes requests based on chosen backend, enabling seamless communication between frontend and backend APIs in the image slideshow repository architecture. Displays UI components through React, integrating routing via RouterSetup component. |
| [index.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/index.css)           | Enhances frontend styling for the image_slideshow project. Centralizes CSS rules to ensure consistent design across components. This file contributes to the overall visual cohesion of the web application.                                    |
| [RouterSetup.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/RouterSetup.js) | Redirects to login, displays Dashboard and EditUser under a common template. Shows Error component for unknown routes. Incorporated within frontend to manage navigation in the web app.                                                        |

</details>

<details closed><summary>frontend.src.components.Error</summary>

| File                                                                                                                  | Summary                                                                                                                                                                                                                                                                             |
| ---                                                                                                                   | ---                                                                                                                                                                                                                                                                                 |
| [Error.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Error/Error.css) | Defines styling for error component, ensuring full viewport coverage, distinctive background colors, and centered display. Implements visual effects like shadows and border radius for an aesthetically pleasing error display within the frontend architecture of the repository. |
| [Error.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Error/Error.js)   | Defines a React component to display a Page Not Found error in the frontend of the repository's image slideshow project.                                                                                                                                                            |

</details>

<details closed><summary>frontend.src.components.Dashboard</summary>

| File                                                                                                                              | Summary                                                                                                                                                                                                                                               |
| ---                                                                                                                               | ---                                                                                                                                                                                                                                                   |
| [Dashboard.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Dashboard/Dashboard.js)   | Manages images for a dashboard, enabling upload, display, and deletion functionalities. Integrates authorization and error handling through modals. Supports navigation between images and ensures proper image management.                           |
| [Dashboard.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Dashboard/Dashboard.css) | Styles Dashboard component for responsive image slideshow layout with interactive elements and customizable buttons. Aesthetic design with dynamic resizing for different screen sizes. Enhances user experience in frontend display of slide images. |

</details>

<details closed><summary>frontend.src.components.EditUser</summary>

| File                                                                                                                           | Summary                                                                                                                                                                                                                                                                                     |
| ---                                                                                                                            | ---                                                                                                                                                                                                                                                                                         |
| [EditUser.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/EditUser/EditUser.css) | Defines styles for user editing interface in the frontend component EditUser. Manages form aesthetics, error UI, and interactions like button states and hover effects for enhanced user experience.                                                                                        |
| [EditUser.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/EditUser/EditUser.js)   | Enables user data editing in a React component. Validates input fields and triggers updates via an API call, displaying feedback through modal notifications. Supports dynamic user details population and user-friendly form interactions within the frontend structure of the repository. |

</details>

<details closed><summary>frontend.src.components.Login</summary>

| File                                                                                                                  | Summary                                                                                                                                                                                                                                                |
| ---                                                                                                                   | ---                                                                                                                                                                                                                                                    |
| [Login.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Login/Login.js)   | Implements a login component in the frontend, displaying signin and signup forms. Organizes the login layout into subcontainers, enhancing user experience. Complements the repositorys architecture by enhancing user authentication functionalities. |
| [Login.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Login/Login.css) | Defines responsive styling for the login component in the frontend section. Ensures proper layout and alignment for login forms across different screen sizes. Maintains consistency with the overall design principles and enhances user experience.  |

</details>

<details closed><summary>frontend.src.components.Signup</summary>

| File                                                                                                                     | Summary                                                                                                                                                                                                                                                               |
| ---                                                                                                                      | ---                                                                                                                                                                                                                                                                   |
| [Signup.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Signup/Signup.js)   | Enables user signup validation and submission in the frontend of the repositorys image slideshow project. Validates user input fields and sends data to the backend for signup. Displays modal messages based on the signup outcome.                                  |
| [Signup.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Signup/Signup.css) | Defines CSS styles for the Signup component in the frontend section. It specifies the padding, background, and styling for form elements like inputs and buttons, enhancing the visual appeal and user experience of the signup functionality within the application. |

</details>

<details closed><summary>frontend.src.components.PageTemplate</summary>

| File                                                                                                                                       | Summary                                                                                                                                                                                                                                     |
| ---                                                                                                                                        | ---                                                                                                                                                                                                                                         |
| [PageTemplate.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/PageTemplate/PageTemplate.css) | Responsive navbar, body layout, and icons. Enhances user experience through visually appealing design.                                                                                                                                      |
| [PageTemplate.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/PageTemplate/PageTemplate.js)   | Manages user authorization, profile editing, and logout functionalities with error handling. Renders a customizable page template for a slideshow app. Integrates modals for notifications. Key use of React hooks and Axios for API calls. |

</details>

<details closed><summary>frontend.src.components.Signin</summary>

| File                                                                                                                     | Summary                                                                                                                                                                                                                                           |
| ---                                                                                                                      | ---                                                                                                                                                                                                                                               |
| [Signin.css](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Signin/Signin.css) | Defines styles for the Signin component, enhancing user experience with custom forms, text, and buttons. Maintains visual consistency and interactivity, crucial for frontend functionality within the architecture.                              |
| [Signin.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Signin/Signin.js)   | Enables user sign-in process with form data handling in React. Utilizes axios for API interactions. Redirects to the dashboard with token and image data upon successful login. Displays error messages using modals for failed sign-in attempts. |

</details>

<details closed><summary>frontend.src.components.Modals</summary>

| File                                                                                                                                           | Summary                                                                                                                                                                                                                                                                      |
| ---                                                                                                                                            | ---                                                                                                                                                                                                                                                                          |
| [TemplateModal.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Modals/TemplateModal.js)           | Enables rendering of stylized modals in the React frontend for handling Template-related content. Influences user interaction by displaying defined ModalTitle and modal body.                                                                                               |
| [AuthorizationModal.js](https://github.com/sharanreddy99/image_slideshow.git/blob/master/frontend/src/components/Modals/AuthorizationModal.js) | Enables user authorization handling in React frontend. Displays a modal with customizable title and content. Provides a button to close the modal and redirect users to the login page. Integrates with the repositorys frontend architecture for cohesive user interaction. |

</details>

<details closed><summary>nginx</summary>

| File                                                                                                | Summary                                                                                                                                                                                                      |
| ---                                                                                                 | ---                                                                                                                                                                                                          |
| [Dockerfile](https://github.com/sharanreddy99/image_slideshow.git/blob/master/nginx/Dockerfile)     | Configures Nginx container with custom default settings.                                                                                                                                                     |
| [default.conf](https://github.com/sharanreddy99/image_slideshow.git/blob/master/nginx/default.conf) | Defines upstream servers and proxies requests for different backend services in Docker container orchestration environment; routes traffic based on specified endpoints to Go, PHP, or React app containers. |

</details>

<details closed><summary>backend_php</summary>

| File                                                                                                        | Summary                                                                                                                                                                                                                                      |
| ---                                                                                                         | ---                                                                                                                                                                                                                                          |
| [index.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/index.php)         | Implements JWT token generation and verification for user authentication in the PHP backend. Utilizes Firebase JWT library to securely manage user sessions, enabling secure access to the Image Slideshow application.                      |
| [Dockerfile](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/Dockerfile)       | Facilitates the setup and configuration for a PHP backend with Apache server. Installs necessary dependencies, sets working directory, copies files, and exposes the service on port 80. Management is handled by the provided start script. |
| [start.sh](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/start.sh)           | Initiates PHP server upon successful MySQL connection, displaying accessible APIs and PHPMyAdmin URLs. Monitors database status before launching server for seamless integration within the image slideshow backend architecture.            |
| [composer.json](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/composer.json) | Manages PHP dependencies for the backend. Requires Firebase JWT and PHP dotenv packages. Orchestrates essential functionalities for the PHP backend architecture of the image slideshow project.                                             |

</details>

<details closed><summary>backend_php.api</summary>

| File                                                                                                                                    | Summary                                                                                                                                                                                                                                                                            |
| ---                                                                                                                                     | ---                                                                                                                                                                                                                                                                                |
| [getsingleimage.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/getsingleimage.php)               | Retrieves a single image data after validating user, via PHP backend. Parses input, queries database, and sends JSON response. Core to the image slideshow backend architecture.                                                                                                   |
| [updateuser.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/updateuser.php)                       | Handles user updates securely, leveraging PHP and SQL. Verifies user data and performs updates, ensuring unique email registration and linking images to updated emails for proper user management. Maintains user experience through error handling and response messages.        |
| [checkvaliduserhandler.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/checkvaliduserhandler.php) | Implements user validation in PHP backend. Verifies user token, checks authorization status, and sets appropriate HTTP response codes and messages. Essential for secure user access within the PHP backend of the image_slideshow repository.                                     |
| [getallimages.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/getallimages.php)                   | Retrieves user-specific images data in JSON format based on email, ensuring secure access via checkvaliduserhandler.php. Implements image querying and formatting for frontend presentation in the PHP backend.                                                                    |
| [deletesingleimage.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/deletesingleimage.php)         | Implements image deletion logic based on user verification. Handles database operation, returning success or failure messages in an API response. Maintains security by checking user authentication before deleting an image.                                                     |
| [fetchuser.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/fetchuser.php)                         | Retrieves user data based on email, handling validation and database queries. Implements error handling to ensure smooth user experience.                                                                                                                                          |
| [logout.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/logout.php)                               | Implements token invalidation logic for user logout, ensuring secure user session management. Verifies user, updates token to null in the database, and confirms successful logout. Enhances user privacy and system security within the repositorys PHP backend architecture.     |
| [signin.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/signin.php)                               | Implements backend API for user signin in PHP. Validates credentials, generates tokens, and handles error responses. Contributes to the PHP backend of the image slideshow project, enhancing user authentication capabilities.                                                    |
| [imageupload.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/imageupload.php)                     | Handles image upload, verifies user, and stores image info in database. Checks for valid file formats and size before uploading. Responds with success/failure message and file name.                                                                                              |
| [isvaliduser.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/isvaliduser.php)                     | Handles user validation, invoking checkvaliduserhandler.php for authorization. Parses email data and catches exceptions, generating JSON response.                                                                                                                                 |
| [signup.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/signup.php)                               | Implements user signup functionality by handling POST requests, checking for existing email, and inserting new user details into the database. Handles errors, responds with JSON, and ensures a smooth registration process in the PHP backend of the image slideshow repository. |

</details>

<details closed><summary>backend_php.api.config</summary>

| File                                                                                                                     | Summary                                                                                                                                                                                                               |
| ---                                                                                                                      | ---                                                                                                                                                                                                                   |
| [connection.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/config/connection.php) | Establishes database connection parameters and initializes necessary tables for user and image data storage in the backend PHP API of the image slideshow app.                                                        |
| [validuser.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/config/validuser.php)   | Validates user credentials by verifying the token against email in the database. Returns the email if the user is valid; otherwise, returns null. contributes to the PHP backend API for the image slideshow project. |
| [php_cors.php](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_php/api/config/php_cors.php)     | Enables cross-origin resource sharing (CORS) for PHP backend, allowing unrestricted resource access. Essential for frontend-backend interaction in the image_slideshow repositorys architecture.                      |

</details>

<details closed><summary>backend_golang</summary>

| File                                                                                                     | Summary                                                                                                                                                                                                                                 |
| ---                                                                                                      | ---                                                                                                                                                                                                                                     |
| [go.sum](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/go.sum)         | Improve dependency management for the Golang backend by specifying exact versions, including jwt-go v3.2.0+incompatible and gorilla/mux v1.8.0. This ensures stability and predictability in the projects environment.                  |
| [Dockerfile](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/Dockerfile) | Defines Golang runtime environment and sets up project dependencies. Establishes working directory, installs required packages, and configures proxy settings. Executes the startup script for the backend service to run on port 8080. |
| [start.sh](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/start.sh)     | Facilitates waiting for MySQL connection before starting Golang APIs, exposes API and PHPMyAdmin URLs, and initiates the server. Key to orchestrating backend services in the repository structure.                                     |
| [main.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/main.go)       | Initiates request handling using the Go backend router within the image_slideshow repository architecture.                                                                                                                              |
| [go.mod](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/go.mod)         | Defines module dependencies for the Golang backend in the repository, ensuring integration with JWT authentication, MySQL database, and routing functionalities. Also specifies the required Go version.                                |

</details>

<details closed><summary>backend_golang.constants</summary>

| File                                                                                                                   | Summary                                                                                                                                                                       |
| ---                                                                                                                    | ---                                                                                                                                                                           |
| [constants.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/constants/constants.go) | Defines environmental constants for the Golang backend server, facilitating database connection and port configuration within the cross-language Image Slideshow application. |

</details>

<details closed><summary>backend_golang.router</summary>

| File                                                                                                          | Summary                                                                                                                                                                                                                                  |
| ---                                                                                                           | ---                                                                                                                                                                                                                                      |
| [router.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/router.go) | Handles API routes for DynamicImageSlideshow by connecting to SQL database and setting up endpoints for user authentication, image manipulation, and user management. Enables HTTP server to listen on port 8080 via Gorilla Mux router. |

</details>

<details closed><summary>backend_golang.router.routes</summary>

| File                                                                                                                                       | Summary                                                                                                                                                                                                                                                                                     |
| ---                                                                                                                                        | ---                                                                                                                                                                                                                                                                                         |
| [logout.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/logout.go)                       | Implements a logout API handler in the backend_golang router. Handles user logout by updating the token status in the database. Ensures service availability and returns logged out status message.                                                                                         |
| [imageupload.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/imageupload.go)             | Image validation, database interaction, error handling.                                                                                                                                                                                                                                     |
| [deletesingleimage.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/deletesingleimage.go) | Implements deletion of a single image in the backend Golang API. Validates user authorization, handles database queries, and responds with success or error messages accordingly. Maintains the integrity of the image storage system.                                                      |
| [signup.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/signup.go)                       | Handles user sign-up requests, checks existing email registrations, and processes new user registrations. Utilizes MySQL queries for user validation and insertion. Provides appropriate response messages on success or failure.                                                           |
| [signin.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/signin.go)                       | Implements the user authentication logic for the backend Go service. Handles user sign-in requests securely by validating credentials against the database and issuing access tokens. Ensures service availability and responds with appropriate messages for success or failure scenarios. |
| [getsingleimage.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/getsingleimage.go)       | Handles GET requests for retrieving a single images data based on user authentication, querying the database, and sending a JSON response containing the image data if found.                                                                                                               |
| [isvaliduser.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/isvaliduser.go)             | Validates user authorization for backend APIs, handling POST requests. Checks if the database server is reachable and verifies user credentials based on a token. Provides appropriate response statuses and messages for either success or failure.                                        |
| [fetchuser.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/fetchuser.go)                 | Manages user authentication, fetches user details from a MySQL database, and sends appropriate responses based on authorization status. Contributes to the backend Golang service in the repositorys architecture.                                                                          |
| [updateuser.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/updateuser.go)               | Handles user updates, checking authentication with tokens and MySQL database. Verifies email availability, updates user details, and associated images. Manages responses and error handling. Influences backend functionality within the repositorys architecture.                         |
| [getallimages.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/routes/getallimages.go)           | Retrieves images and details based on user authorization.-Handles DB connectivity with error responses.-Parses request data and constructs image responses in JSON format for frontend consumption.                                                                                         |

</details>

<details closed><summary>backend_golang.router.sqldb</summary>

| File                                                                                                        | Summary                                                                                                                                                                                                                                           |
| ---                                                                                                         | ---                                                                                                                                                                                                                                               |
| [db.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/router/sqldb/db.go) | Establishes database connection, creates tables in MySQL for users and images, and stores the connection globally. This enables the backend Golang service to interact with the MySQL database seamlessly in the parent repositorys architecture. |

</details>

<details closed><summary>backend_golang.utils</summary>

| File                                                                                                       | Summary                                                                                                                                                                                                                                    |
| ---                                                                                                        | ---                                                                                                                                                                                                                                        |
| [utils.go](https://github.com/sharanreddy99/image_slideshow.git/blob/master/backend_golang/utils/utils.go) | Defines utility functions for handling CORS headers, generating random strings and JWT tokens, and validating user tokens in the backend Golang service. Enables secure authentication and user validation within the Image Slideshow app. |

</details>

---

##  Getting Started

###  Installation

<h4>From <code>source</code></h4>

> 1. Clone the image_slideshow repository:
>
> ```console
> $ git clone https://github.com/sharanreddy99/image_slideshow.git
> ```
>
> 2. Change to the project directory:
> ```console
> $ cd image_slideshow
> ```
>
> 3. Run the project using docker compose
> ```console
> $ docker compose up
> ```

###  Usage


> Access the application in the browser at http://localhost

---

##  Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/sharanreddy99/image_slideshow.git/issues)**: Submit bugs found or log feature requests for the `image_slideshow` project.
- **[Submit Pull Requests](https://github.com/sharanreddy99/image_slideshow.git/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/sharanreddy99/image_slideshow.git/discussions)**: Share your insights, provide feedback, or ask questions.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/sharanreddy99/image_slideshow.git
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="center">
   <a href="https://github.com/sharanreddy99/image_slideshow.git/graphs/contributors">
      <img src="https://contrib.rocks/image?repo=sharanreddy99/image_slideshow">
   </a>
</p>
</details>

---
