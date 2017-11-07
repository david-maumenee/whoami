whoami
======

Fork of jwilder/whoami an HTTP docker service that prints it's container ID

Display the hostname of the current container and display the result of a backend URL defined by the environment variable named BACK_END_URL

If you configure BACK_END_URL to jwilder/whoami you can verify the internal load balancing of our CaaS


    $ cd docker-compose
    $ docker-compose up -d
    
    OR with swarm
    $ docker swarm init
    $ docker stack deploy --compose-file docker-compose.yml  whoami

    $ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
    
