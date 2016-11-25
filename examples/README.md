# Example tasks

[This](example-task.yml) example task will publish metrics to a file
from the Elasticsearch collector plugin.  


### Requirements 
 * `docker` and `docker-compose` are **installed** and **configured** 

Running the sample is as *easy* as running the script `./run-example.sh`. 

## Files

- [run-example.sh](run-example.sh) 
    - The example is launched with this script     
- [example-task.yml](example-task.yml)
    - The example Snap task definition
- [docker-compose.yml](docker-compose.yml)
    - A docker compose file which defines two linked containers
        - "runner" is the container where snapteld is run from.  You will be dumped 
        into a shell in this container after running 
        [run-example.sh](run-example.sh).  Exiting the shell will 
        trigger cleaning up the containers used in the example.
        - "elasticsearch" is the container running Elasticsearch server. 
- [example.sh](example.sh)
    - Downloads `snapteld`, `snaptel`, `snap-plugin-collector-elasticsearch`,
    `snap-plugin-publisher-file` and starts the task 
    [example-task.yml](example-task.yml).
- [.setup.sh](.setup.sh)
    - Verifies dependencies and starts the containers.  It's called 
    by [run-example.sh](run-example.sh).
