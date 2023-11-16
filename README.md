# Shopify-Integrator-docs

API Documentation for Shopify Integrator created with Swagger

1. Description of project
2. Installation

## Description of project

This project serves as the API documentation of the Shopify-Integrator application and is built using [Swagger](https://swagger.io).

## Installation

Note that this is installed automatically in the shell `install.sh` script of the Shopify Integrator Application.

### To manually run this in your browser

1. Install the `swagger` go module. A guide can be found on the [go-swagger repository](https://goswagger.io/install.html).
2. Ensure you have the latest version of the repository (`git pull` or you may `clone` it)
3. Run the command below:

```bash
swagger serve ./swagger.yml --port=8080 --base-path=/api/ --no-open
```
