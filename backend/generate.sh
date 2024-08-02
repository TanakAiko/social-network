#!/bin/bash

# Function to create a service file
create_service() {
  service_name=$1
  service_file="${service_name}.go"

  cat <<EOL > $service_file
package services

import "net/http"

// ${service_name}Service struct
type ${service_name}Service struct {}

// ConfigureEndpoints configures the endpoints for the service
func (s *${service_name}Service) ConfigureEndpoints() {
  // Implement endpoint configuration
}

// InitService initializes the service
func (s *${service_name}Service) InitService() {
  // Implement service initialization
}
EOL

  echo "Service file '$service_file' created."
}

# Function to create a controller file
create_controller() {
  controller_name=$1
  controller_file="${controller_name}_controller.go"

  cat <<EOL > $controller_file
package services

import "net/http"

// ${controller_name}Controller struct
type ${controller_name}Controller struct {}

// HTTPServe serves HTTP requests
func (c *${controller_name}Controller) HTTPServe() http.Handler {
  // Implement HTTP serving
  return nil
}

// SetMethods sets the HTTP methods for the controller
func (c *${controller_name}Controller) SetMethods() []string {
  // Implement setting methods
  return nil
}

// EndPoint returns the endpoint for the controller
func (c *${controller_name}Controller) EndPoint() string {
  // Implement getting endpoint
  return ""
}
EOL

  echo "Controller file '$controller_file' created."
}

# Function to display the usage of the script
usage() {
  echo "Usage: $0 -t <type> -n <name>"
  echo "  -t <type>    Specify the type to create (service or controller)"
  echo "  -n <name>    Specify the name for the service or controller"
  exit 1
}

# Parse options
while getopts ":t:n:" opt; do
  case $opt in
    t)
      type=$OPTARG
      ;;
    n)
      name=$OPTARG
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      usage
      ;;
    :)
      echo "Option -$OPTARG requires an argument." >&2
      usage
      ;;
  esac
done

# Check if both type and name are provided
if [ -z "$type" ] || [ -z "$name" ]; then
  usage
fi

# Create service or controller based on the type provided
case $type in
  service)
    create_service $name
    ;;
  controller)
    create_controller $name
    ;;
  *)
    echo "Invalid type: $type" >&2
    usage
    ;;
esac
