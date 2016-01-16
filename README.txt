Muad'dib - Web UI for Usul

Supported Languages:

- Ruby
- Python
- Javascript

Requires:

- Install go
- Install virtualbox
- Install docker toolbox

Create docker machine in virtualbox:

- Run 'docker-machine create --driver=virtualbox muaddib'

Source the environment variables:

- Run 'eval "$(docker-machine env muaddib)"'

Build the docker image locally:

- Run 'docker build -t muaddib .'

Build Usul:

- Run 'go build'

Run Usul:

- Run './muaddib'

Try web editor:

- http://localhost:8080
