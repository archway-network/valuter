# valuter
This is a testnet evaluator tool that is linked with [cosmologger](https://github.com/archway-network/cosmologger/) database and extracts the winners information.

# Install & development
The best way to install it is to do it through a bundle called [testnet-evaluator](https://github.com/archway-network/testnet-evaluator/). 

## Configuration

### ENV Variables


* **SERVING_ADDR**: The serving address, where the Valuter APIs are accessible.

* **POSTGRES_DB**: The name of `postgres` database
* **POSTGRES_USER**: Database username
* **POSTGRES_PASSWORD**: Password for the database user
* **POSTGRES_PORT**: Port number of the database
* **POSTGRES_HOST**: Host address of the server running postgres


### Config file
There is a `config.json` file, which has to be mapped into the app directory of the container. i.e. be in the same path of the executable.

Here is what is inside the conf file:

```json
{
    "tasks":{
        
        "validators-genesis": {
            "max-winners" : 10,
            "reward": 480
        },
        "validators-joined": {
            "max-winners" : 10,
            "reward": 720
        },
       
        "jail-unjail": {
            "max-winners" : 10,
            "reward": 720
        },
        "staking": {
            "max-winners" : 10,
            "reward": 480
        },
        "gov": {
            "max-winners" : 10,
            "proposals": [1,2,3],
            "reward": 480          
        },
        "node-upgrade": {
            "max-winners" : 10,
            "reward": 480,
            "condition": {
                "upgrade-hight": 1000
            }
        },
        "uptime": {
            "max-winners" : 10,
            "reward": 480,
            "conditions": [{
                "start-hight": 100,
                "end-hight": 200,
                "uptime-percent": 0.80
            }]
        }
    },

    "api":{
        "rows-per-page" : 200
    }

}
```

`tasks` keeps the configs related to the tasks:

* `max-winners`: The maximum number of winners for the task.
* `reward`: The reward amount for the task

```json
"gov": {
    "max-winners" : 10,
    "proposals": [1,2,3],
    "reward": 480          
},
```

* `proposals`: is a list of proposal Ids that participants need to vote for. 
The reward amount is calculated for each proposal vote. e.g. if there are 3 proposals and a user votes for two, they will get `2 x reward`


```json
"node-upgrade": {
    "max-winners" : 10,
    "reward": 480,
    "condition": {
        "upgrade-hight": 1000
    }
},
```

* `condition.upgrade-hight`: refers to a specific block height that the upgrade has to happen. We check if a validator has signed this particular block to pick the winners.


```json
"uptime": {
    "max-winners" : 10,
    "reward": 480,
    "conditions": [{
        "start-hight": 100,
        "end-hight": 200,
        "uptime-percent": 0.80
    }]
}
```
* `conditions` holds the configs for multiple rounds of load bursts. Each round has three parameters:

* * `start-hight`: The beginning of the load burst.
* * `end-hight`: The end of the load burst.
* * `uptime-percent`: minimum uptime requirement.