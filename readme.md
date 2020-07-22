# Order App

This is a sample application  built with  Go backend and vuejs frontend 

The applications displays a list of orders and also provides a search functionality.


## Getting Started

These instructions will get you a copy of the project up and running on your local machine.

### Prerequisites
 
### Installing 

Ensure that you have Docker and docker compose installed on your system

## Built With

* [Docker for Mac](https://docs.docker.com/docker-for-mac/install/) - Install docker for mac 
* [Docker for windows](https://docs.docker.com/docker-for-windows/install/) - Install docker for windows
* [Docker compose](https://docs.docker.com/compose/install/) - Install docker compose

Clone The project 

```
git clone git@github.com:ysfkel/order-app.git
```

Change directory to the root of the cloned repository (where you have docker-compose.yaml) and run:

```
docker-compose up
```

Wait for the application to build

When the application build is done, open a web browser and navigate to:

```
http://localhost:8080/orders
```
 
## Backend folder structure

Directory name: order-app

The backend project follows the domain driven design project structure according 
to the following directories

* Api
   This directory contains a single file in which the rest api routes are defined

* Boot 
   This direcory contains a single file in which code that bootstrap the application is written

* Controllers:
   This directory contains a controller that handles the Api routes

* Domain:
   This directory contains code that define the models and repository interfaces

* Infratrsuture
   This directory contains the following sub directories
     * data_reader: Contains code that reads the csv files in test_data  
     * persistence: contains code that implement database repository interfaces 
        * Migration: Contains database migration code
     * service Contains code that processes user request and returns the result to the controller

## Application Flow
   
* User Makes a http request to fetch orders list 
* The request hits one of the rest api routes defines in the routers.go file  in the api directory 
* The controller validates the user request and instantiates the service and passes the request to the service located in the service directory
* the service retrieves the requested data from the database by making a call to the database repository
* service returns the result to the controller 
* controller returns response to the client 


