# Introduction to Backend and docker

This repository contains a series of exercises focused on key topics in software architecture and engineering, including package by layer design, REST API development, integration testing, and Docker containerization.

## Exercises

### 1. Create an Endpoint for Item Creation
**Objective:**  
Develop an endpoint that allows for the creation of an item using JSON data. Ensure that the endpoint returns the correct HTTP status code upon successful creation or the correct HTTP status on error.

**Instructions:**  
- Implement an endpoint that processes a JSON payload with the following structure:
  ```json
  {
    "name": "string",
    "id": "string",
    "price": "int"
  }
  ```
 - Validate the input and handle the request.
 - Return the appropriate HTTP status code based on the outcome.
### 2. Refactor Using Package by Layer Design
**Objective:**  
Improve the solution from Exercise 1 by refactoring the code to follow the package by layer design approach. This helps in better organization, maintainability, and separation of concerns.

**Instructions:**  
- Organize your code into layers such as controllers, services, repositories, and models.
- Refactor the existing endpoint to follow this structure.
- Ensure that the code remains functional after refactoring.

### 3. Integration Test for Item Creation Endpoint
**Objective:**  
Write an integration test for the item creation endpoint developed in Exercise 1. This test should validate the full flow of the endpoint, ensuring that the API works as expected with all components.

**Instructions:**  
- Set up the necessary testing framework.
- Write a test case that simulates an item creation request.
- Verify that the endpoint returns the expected HTTP status code and response.

### 4. Dockerize the Application
**Objective:**  
Create a Dockerfile to containerize the application, allowing it to run in any environment with ease. The Dockerfile should expose the necessary ports for the API to be accessible.

**Instructions:**  
- Write a Dockerfile that installs dependencies, builds the project, and runs the application.
- Ensure that the correct ports are exposed and mapped for API access.
- Test the Dockerized application locally to verify it works as expected.

