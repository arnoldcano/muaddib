Muad'dib - Web UI for Usul (http://github.com/arnoldcano/usul)

Supported Languages:

- Ruby
- Python
- Javascript

Requires:

- Install go
- Install virtualbox
- Install docker toolbox

Create docker machine in virtualbox:

- Run 'docker-machine create --driver=virtualbox dune'

Source the environment variables:

- Run 'eval "$(docker-machine env dune)"'

Build the docker image locally:

- Run 'docker build -t muaddib .'

Run muaddib:

- Run 'docker run --rm -p 8081:8081 --name muaddib --link usul muaddib

Get docker ip address:

- Run 'docker-machine ip dune'

Try web editor:

- http://$(docker-machine ip dune):8081
