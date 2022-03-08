This simulator simulates changes in the nonce algorithm, based on options set in a JSON file. For running the simulation we need to provide json file storing account information about various stakers. The simulation can be run for virtual day,week or month (passing the corresponding parameter). Various blockchain parameters can be defined in `energi/consensus/difficultysim/params/params.go` file. The output/result csv file is written in `energi/consensus/difficultysim/csv` directory

### Run the simulator
Show staker mining output, if we need to run DiffV1 functionality there is a parameter `AsgardIsActive` in `params/params.go` to be set correspondingly, simulator can be run for specific period - for a day, week or month
```bash
go run main.go -stakeConfig=./config/testnet.config.json -simulationPeriod=day
```
#### To get up-to-date help on the parameters:
```bash
go run . --help
```

### JSON File Description
The example JSON file `config.sample.json` represents a `SimConfig` object and defines a very basic simulation, similar to this example:

```json
{
   "Stakers":[
      {
         "name":"Jack",
         "accounts":[
            {
               "address":"eXAYraZVrH",
               "balance":2000000,
               "nonceCap":0
            }
         ]
      },
      {
         "name":"Rose",
         "accounts":[
            {
               "address":"SAaOknvphY",
               "balance":22737,
               "nonceCap":0
            },
            {
               "address":"lkKLml2j312",
               "balance":22737,
               "nonceCap":0
            }
         ]
      }
   ]
}
```
